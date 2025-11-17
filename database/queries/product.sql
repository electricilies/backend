-- name: CreateProduct :one
INSERT INTO products (
  name,
  description
)
VALUES (
  sqlc.arg('name'),
  sqlc.arg('description')
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
  UNNEST(sqlc.arg('skus')::text[]) AS sku,
  UNNEST(sqlc.arg('prices')::decimal[]) AS price,
  UNNEST(sqlc.arg('quantities')::integer[]) AS quantity,
  sqlc.arg('product_id')
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
  UNNEST(sqlc.arg('urls')::text[]) AS url,
  UNNEST(sqlc.arg('orders')::integer[]) AS "order",
  UNNEST(sqlc.arg('product_variant_ids')::integer[]) AS product_variant_id,
  sqlc.arg('product_id')
RETURNING
  *;

-- name: LinkProductAttributeValues :execrows
INSERT INTO products_attribute_values (
  product_id,
  attribute_value_id
)
SELECT
  sqlc.arg('product_id'),
  UNNEST(sqlc.arg('attribute_value_ids')::integer[]) AS attribute_value_id;

-- name: LinkProductVariantsWithOptionValues :execrows
INSERT INTO option_values_product_variants (
  option_value_id,
  product_variant_id
)
SELECT
  UNNEST(sqlc.arg('option_value_ids')::integer[]) AS option_value_id,
  UNNEST(sqlc.arg('product_variant_ids')::integer[]) AS product_variant_id;

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
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN products.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN products.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  CASE WHEN
    sqlc.narg('search') IS NOT NULL THEN pdb.score(products.id) + category_scores
  END DESC,
  CASE WHEN
    sqlc.narg('sort_rating')::text = 'asc' THEN products.rating
  END ASC,
  CASE WHEN
    sqlc.narg('sort_rating')::text = 'desc' THEN products.rating
  END DESC,
  CASE WHEN
   sqlc.narg('sort_price')::text = 'asc' THEN products.price
  END ASC,
  CASE WHEN
   sqlc.narg('sort_price')::text = 'desc' THEN products.price
  END DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetProduct :one
SELECT
  *
FROM
  products
WHERE
  products.id = sqlc.arg('id')
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END;

-- name: ListProductVariants :many
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
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
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

-- name: UpdateProduct :one
UPDATE
  products
SET
  name = COALESCE(sqlc.narg('name')::text, name),
  description = COALESCE(sqlc.narg('description')::text, description),
  views_count = COALESCE(sqlc.narg('views_count')::integer, views_count),
  total_purchase = COALESCE(sqlc.narg('total_purchase')::integer, purchase_count),
  trending_score = COALESCE(sqlc.narg('trending_score')::float, trending_score), -- TODO: Do we ever update this manually?
  category_id = COALESCE(sqlc.narg('category_id')::integer, category_id),
  updated_at = COALESCE(sqlc.narg('updated_at')::timestamp, NOW())
WHERE
  id = sqlc.arg('id')
  AND deleted_at IS NULL
RETURNING
  *;

-- name: UpdateProductVariants :many
WITH updated_variants AS (
  SELECT
    UNNEST(sqlc.arg('ids')::integer[]) AS id,
    UNNEST(sqlc.arg('skus')::text[]) AS sku,
    UNNEST(sqlc.arg('prices')::decimal[]) AS price,
    UNNEST(sqlc.arg('quantities')::integer[]) AS quantity,
    UNNEST(sqlc.arg('purchase_counts')::integer[]) AS purchase_count
)
UPDATE
  product_variants
SET
  sku = COALESCE(updated_variants.sku, product_variants.sku),
  price = COALESCE(updated_variants.price, product_variants.price),
  quantity = COALESCE(updated_variants.quantity, product_variants.quantity),
  purchase_count = COALESCE(updated_variants.purchase_count, product_variants.purchase_count),
  updated_at = COALESCE(sqlc.narg('updated_at')::timestamp, NOW())
FROM
  updated_variants
WHERE
  product_variants.id = updated_variants.id
  AND product_variants.deleted_at IS NULL
RETURNING
  product_variants.*;

-- name: DeleteProducts :execrows
UPDATE
  products
SET
  deleted_at = NOW()
WHERE
  id = ANY (sqlc.arg('ids')::integer[])
  AND deleted_at IS NULL;

-- name: DeleteProductVariants :execrows
UPDATE
  product_variants
SET
  deleted_at = NOW()
WHERE
  id = ANY (sqlc.arg('ids')::integer[])
  AND deleted_at IS NULL;

-- name: DeleteProductImages :execrows
DELETE FROM
  product_images
WHERE
  id = ANY (sqlc.arg('ids')::integer[]);

-- name: GetProductImage :one
SELECT
  *
FROM
  product_images
WHERE
  id = sqlc.arg('id');
