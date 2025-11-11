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
    sqlc.narg('search')::text IS NULL
    OR (
      categories.name ||| (sqlc.narg('search')::text)::pdb.fuzzy(2)
      AND categories.deleted_at IS NULL
    )
) AS category_scores
  ON products.id = category_scores.id
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE products.id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE products.name ||| (sqlc.narg('search')::text)::pdb.fuzzy(products.trending_score)
  END
  AND CASE
    WHEN sqlc.narg('min_price')::decimal IS NULL THEN TRUE
    ELSE products.price >= sqlc.narg('min_price')::decimal
  END
  AND CASE
    WHEN sqlc.narg('max_price')::decimal IS NULL THEN TRUE
    ELSE products.price <= sqlc.narg('max_price')::decimal
  END
  AND CASE
    WHEN sqlc.narg('rating')::real IS NULL THEN TRUE
    ELSE products.rating >= sqlc.narg('rating')::real
  END
  AND CASE
    WHEN sqlc.narg('category_ids')::integer[] IS NULL THEN TRUE
    ELSE products.category_id = ANY (sqlc.narg('category_ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('include_deleted_only')::bool THEN products.deleted_at IS NOT NULL
    WHEN sqlc.narg('include_deleted_only')::bool = FALSE THEN products.deleted_at IS NULL
    ELSE TRUE
  END
ORDER BY
  CASE WHEN
    sqlc.narg('search') IS NOT NULL THEN pdb.score(products.id) + category_scores
  END DESC,
  CASE WHEN
    sqlc.narg('sort_rating_asc')::bool = TRUE THEN products.rating
  END ASC,
  CASE WHEN
    sqlc.narg('sort_rating_asc')::bool = FALSE THEN products.rating
  END DESC,
  CASE WHEN
   sqlc.narg('sort_price_asc')::bool = TRUE THEN products.price
  END ASC,
  CASE WHEN
   sqlc.narg('sort_price_asc')::bool = FALSE THEN products.price
  END DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetProduct :one
SELECT
  *
FROM
  products
WHERE
  products.id = @id
  AND CASE
    WHEN sqlc.narg('include_deleted_only')::bool THEN deleted_at IS NOT NULL
    WHEN sqlc.narg('include_deleted_only')::bool = FALSE THEN deleted_at IS NULL
    ELSE TRUE
  END;

-- name: ListProductVariants :one
SELECT
  *
FROM
  product_variants
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('product_ids')::integer[] IS NULL THEN TRUE
    ELSE product_id = ANY (sqlc.narg('product_ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('include_deleted_only')::bool THEN deleted_at IS NOT NULL
    WHEN sqlc.narg('include_deleted_only')::bool = FALSE THEN deleted_at IS NULL
    ELSE TRUE
  END
ORDER BY
  id;

-- name: ListProductImages :many
SELECT
  *
FROM
  product_images
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('product_variant_ids')::integer[] IS NULL THEN TRUE
    ELSE product_variant_id = ANY (sqlc.narg('product_variant_ids'))
  END
  AND CASE
    WHEN sqlc.narg('product_ids')::integer[] IS NULL THEN TRUE
    ELSE product_id = ANY (sqlc.narg('product_ids')::integer[])
  END
ORDER BY
  id ASC;

-- name: UpdateProduct :execrows
UPDATE
  products
SET
  name = COALESCE(sqlc.narg('name')::text, name),
  description = COALESCE(sqlc.narg('description')::text, description),
  views_count = COALESCE(sqlc.narg('views_count')::integer, views_count),
  total_purchase = COALESCE(sqlc.narg('total_purchase')::integer, purchase_count),
  trending_score = COALESCE(sqlc.narg('trending_score')::float, trending_score), -- TODO: Do we ever update this manually?
  updated_at = NOW()
WHERE
  id = @id
  AND deleted_at IS NULL;

-- name: UpdateProductVariants :execrows
WITH updated_variants AS (
  SELECT
    UNNEST(@ids::integer[]) AS id,
    UNNEST(@skus::text[]) AS sku,
    UNNEST(@prices::decimal[]) AS price,
    UNNEST(@quantities::integer[]) AS quantity,
    UNNEST(@purchase_counts::integer[]) AS purchase_count
)
UPDATE
  product_variants
SET
  sku = COALESCE(updated_variants.sku, product_variants.sku),
  price = COALESCE(updated_variants.price, product_variants.price),
  quantity = COALESCE(updated_variants.quantity, product_variants.quantity),
  purchase_count = COALESCE(updated_variants.purchase_count, product_variants.purchase_count),
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
  id = ANY (@ids::integer[])
  AND deleted_at IS NULL;

-- name: DeleteProductVariants :execrows
UPDATE
  product_variants
SET
  deleted_at = NOW()
WHERE
  id = ANY (@ids::integer[])
  AND deleted_at IS NULL;

-- name: DeleteProductImages :execrows
DELETE FROM
  product_images
WHERE
  id = ANY (@ids::integer[]);

-- name: DeleteLinkedProductAttributeValues :execrows
WITH deleted_links AS (
  SELECT
    UNNEST(@product_id::integer[]) AS product_id,
    UNNEST(@attribute_value_ids::integer[]) AS attribute_value_id
)
DELETE FROM
  products_attribute_values
USING
  deleted_links
WHERE
  products_attribute_values.product_id = deleted_links.product_id
  AND products_attribute_values.attribute_value_id = deleted_links.attribute_value_id;

-- name: DeleteLinkedProductVariantsOptionValues :execrows
WITH deleted_links AS (
  SELECT
    UNNEST(@product_variant_ids::integer[]) AS product_variant_id,
    UNNEST(@option_value_ids::integer[]) AS option_value_id
)
DELETE FROM
  option_values_product_variants
USING
  deleted_links
WHERE
  option_values_product_variants.product_variant_id = deleted_links.product_variant_id
  AND option_values_product_variants.option_value_id = deleted_links.option_value_id;
