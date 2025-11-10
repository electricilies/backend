-- name: CreateProduct :one
INSERT INTO products (
  name,
  description
)
VALUES (
  @name,
  @description
)
RETURNING
  *;

-- name: CreateProductVariants :many
INSERT INTO product_variants (
  sku,
  price,
  quantity,
  product_id
)
SELECT
  UNNEST(@skus::text[]) AS sku,
  UNNEST(@prices::decimal[]) AS price,
  UNNEST(@quantities::integer[]) AS quantity,
  @product_id
RETURNING
  *;

-- name: CreateProductImages :many
INSERT INTO product_images (
  url,
  "order",
  product_variant_id,
  product_id
)
SELECT
  UNNEST(@urls::text[]) AS url,
  UNNEST(@orders::integer[]) AS "order",
  UNNEST(@product_variant_ids::integer[]) AS product_variant_id,
  @product_id
RETURNING
  *;

-- name: LinkProductAttributeValues :execrows
INSERT INTO products_attribute_values (
  product_id,
  attribute_value_id
)
SELECT
  @product_id,
  UNNEST(@attribute_value_ids::integer[]) AS attribute_value_id;

-- name: LinkProductVariantsWithOptionValues :execrows
INSERT INTO option_values_product_variants (
  option_value_id,
  product_variant_id
)
SELECT
  UNNEST(@option_value_ids::integer[]) AS option_value_id,
  UNNEST(@product_variant_ids::integer[]) AS product_variant_id;

-- This is used for list, search (with filter, order), suggest
-- name: ListProducts :many
SELECT
  sqlc.embed(products),
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  products
LEFT JOIN (
  SELECT
    products.id,
    pdb.score(product.id) AS category_score
  FROM products
  INNER JOIN categories
    ON products.category_id = categories.id
  WHERE
    categories.deleted_at IS NULL
    AND CASE
      WHEN sqlc.narg('search')::text IS NULL THEN TRUE
      ELSE categories.name ||| sqlc.narg('search')::pdb.fuzzy(2)
    END
) AS category_scores
  ON products.id = category_scores.id
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE products.id = ANY(sqlc.narg('ids'))
  END
  AND CASE
    WHEN sqlc.narg('min_price')::decimal IS NULL THEN TRUE
    ELSE products.price >= sqlc.narg('min_price')
  END
  AND CASE
    WHEN sqlc.narg('max_price')::decimal IS NULL THEN TRUE
    ELSE products.price <= sqlc.narg('max_price')
  END
  AND CASE
    WHEN sqlc.narg('rating')::real IS NULL THEN TRUE
    ELSE products.rating >= sqlc.narg('rating')
  END
  AND CASE
    WHEN sqlc.narg('category_ids')::integer[] IS NULL THEN TRUE
    ELSE products.category_id = ANY(sqlc.narg('category_ids'))
  END
  AND (products.deleted_at IS NOT NULL) = @deleted::bool
  AND CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE
      products.name ||| sqlc.narg('search')::pdb.fuzzy(products.trending_score)
  END
ORDER BY
  CASE WHEN sqlc.narg('search') IS NOT NULL THEN pdb.score(products.id) + category_scores END DESC,
  CASE WHEN @sort_rating_asc::bool THEN products.rating END ASC,
  CASE WHEN NOT @sort_rating_asc::bool THEN products.rating END DESC,
  CASE WHEN @sort_price_asc::bool THEN products.price END ASC,
  CASE WHEN NOT @sort_price_asc::bool THEN products.price END DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetProductByID :one
SELECT
  *
FROM
  products
WHERE
  products.id = @id
  AND (products.deleted_at IS NOT NULL) = @deleted::bool;

-- name: ListProductVariants :one
SELECT
  *
FROM
  product_variants
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE product_variants.id = ANY(sqlc.narg('id'))
  END
  AND CASE
    WHEN sqlc.narg('product_ids')::integer[] IS NULL THEN TRUE
    ELSE product_variants.product_id = ANY(sqlc.narg('product_ids'))
  END
  AND (product_variants.deleted_at IS NOT NULL) = @deleted::bool
ORDER BY
  product_variants.id;

-- name: ListProductImages :many
SELECT
  *
FROM
  product_images
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE product_images.id = ANY(sqlc.narg('ids'))
  END
  AND CASE
    WHEN sqlc.narg('product_variant_ids')::integer[] IS NULL THEN TRUE
    ELSE product_images.product_variant_id = ANY(sqlc.narg('product_variant_ids'))
  END
  AND CASE
    WHEN sqlc.narg('product_ids')::integer[] IS NULL THEN TRUE
    ELSE product_images.product_id = ANY(sqlc.narg('product_ids'))
  END
ORDER BY
  product_images.id ASC;

-- name: UpdateProduct :one
UPDATE
  products
SET
  name = COALESCE(sql.narg('name')::text, name),
  description = COALESCE(sql.narg('description')::text, description),
  views_count = COALESCE(sql.narg('views_count')::integer, views_count),
  total_purchase = COALESCE(sql.narg('total_purchase')::integer, purchase_count),
  trending_score = COALESCE(sql.narg('trending_score')::float, trending_score) -- TODO: Do we ever update this manually?
WHERE
  deleted_at IS NULL
  AND id = @id
RETURNING
  *;

-- name: UpdateProductVariants :execrows
WITH updated_variants AS (
  SELECT
    UNNEST(@ids::integer[]) AS id,
    UNNEST(@skus::text[]) AS sku,
    UNNEST(@prices::decimal[]) AS price,
    UNNEST(@quantities::integer[]) AS quantity,
    UNNEST(@purchase_counts::integer[]) AS purchase_count,
    @updated_at::timestamp AS updated_at
)
UPDATE
  product_variants
SET
  sku = updated_variants.sku,
  price = updated_variants.price,
  quantity = updated_variants.quantity,
  purchase_count = updated_variants.purchase_count,
  updated_at = NOW()
FROM
  updated_variants
WHERE
  product_variants.id = updated_variants.id
  AND product_variants.deleted_at IS NULL;

-- name: DeleteProducts :execrows
UPDATE
  products
SET
  deleted_at = NOW()
WHERE
  deleted_at IS NULL
  AND id = ANY(@ids::integer[]);

-- name: DeleteProductVariants :execrows
UPDATE
  product_variants
SET
  deleted_at = NOW()
WHERE
  deleted_at IS NULL
  AND id = ANY(@ids::integer[]);

-- name: DeleteProductImages :execrows
DELETE FROM
  product_images
WHERE
  id = ANY(@ids::integer[]);
