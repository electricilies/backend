-- name: UpsertProduct :exec
INSERT INTO products (
  id,
  name,
  description,
  price,
  views_count,
  total_purchase,
  rating,
  trending_score,
  category_id,
  created_at,
  updated_at,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('name'),
  sqlc.arg('description'),
  sqlc.arg('price'),
  sqlc.arg('views_count'),
  sqlc.arg('total_purchase'),
  sqlc.arg('rating'),
  sqlc.arg('trending_score'),
  sqlc.arg('category_id'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at'),
  sqlc.narg('deleted_at')
)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  price = EXCLUDED.price,
  views_count = EXCLUDED.views_count,
  total_purchase = EXCLUDED.total_purchase,
  rating = EXCLUDED.rating,
  trending_score = EXCLUDED.trending_score,
  category_id = EXCLUDED.category_id,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at,
  deleted_at = EXCLUDED.deleted_at;

-- This is used for list, search (with filter, order), suggest
-- name: ListProducts :many
SELECT
  products.*
FROM
  products
LEFT JOIN (
  SELECT
    products.id,
    pdb.score(products.id) AS category_score
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
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    ELSE products.id = ANY (sqlc.narg('ids')::uuid[])
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
    WHEN sqlc.narg('category_ids')::uuid[] IS NULL THEN TRUE
    ELSE products.category_id = ANY (sqlc.narg('category_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN products.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN products.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  CASE WHEN
    sqlc.narg('search') IS NOT NULL THEN pdb.score(products.id) + category_scores.category_score
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

-- name: CountProducts :one
SELECT
  COUNT(*) AS count
FROM
  products
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    ELSE products.id = ANY (sqlc.narg('ids')::uuid[])
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
    WHEN sqlc.narg('category_ids')::uuid[] IS NULL THEN TRUE
    ELSE products.category_id = ANY (sqlc.narg('category_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN products.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN products.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END;

-- name: GetProduct :one
SELECT
  *
FROM
  products
WHERE
  products.id = sqlc.arg('id')
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
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
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('product_ids')::uuid[] IS NULL THEN TRUE
    ELSE product_id = ANY (sqlc.narg('product_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  id;

-- name: GetProductVariant :one
SELECT
  *
FROM
  product_variants
WHERE
  id = sqlc.arg('id')
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END;

-- name: ListProductImages :many
SELECT
  *
FROM
  product_images
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('product_variant_ids')::uuid[] IS NULL THEN TRUE
    ELSE product_variant_id = ANY (sqlc.narg('product_variant_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('product_ids')::uuid[] IS NULL THEN TRUE
    ELSE product_id = ANY (sqlc.narg('product_ids')::uuid[])
  END
ORDER BY
  id ASC;

-- name: GetProductImage :one
SELECT
  *
FROM
  product_images
WHERE
  id = sqlc.arg('id');

-- name: MergeProductVariantsFromTemp :exec
MERGE INTO product_variants AS target
USING temp_product_variants AS source
  ON target.id = source.id
WHEN MATCHED THEN
  UPDATE SET
    sku = source.sku,
    price = source.price,
    quantity = source.quantity,
    purchase_count = source.purchase_count,
    product_id = source.product_id,
    created_at = source.created_at,
    updated_at = source.updated_at,
    deleted_at = source.deleted_at
WHEN NOT MATCHED THEN
  INSERT (
    id,
    sku,
    price,
    quantity,
    purchase_count,
    product_id,
    created_at,
    updated_at,
    deleted_at
  )
  VALUES (
    source.id,
    source.sku,
    source.price,
    source.quantity,
    source.purchase_count,
    source.product_id,
    source.created_at,
    source.updated_at,
    source.deleted_at
  )
WHEN NOT MATCHED BY SOURCE AND target.product_id = sqlc.arg('product_id')::uuid THEN
  DELETE;

-- name: MergeProductImagesFromTemp :exec
MERGE INTO product_images AS target
USING temp_product_images AS source
  ON target.id = source.id
WHEN MATCHED THEN
  UPDATE SET
    url = source.url,
    "order" = source."order",
    product_id = source.product_id,
    product_variant_id = source.product_variant_id,
    created_at = source.created_at,
    deleted_at = source.deleted_at
WHEN NOT MATCHED THEN
  INSERT (
    id,
    url,
    "order",
    product_id,
    product_variant_id,
    created_at,
    deleted_at
  )
  VALUES (
    source.id,
    source.url,
    source."order",
    source.product_id,
    source.product_variant_id,
    source.created_at,
    source.deleted_at
  )
WHEN NOT MATCHED BY SOURCE AND target.product_id = sqlc.arg('product_id')::uuid THEN
  DELETE;

-- name: MergeProductsAttributeValuesFromTemp :exec
MERGE INTO products_attribute_values AS target
USING temp_products_attribute_values AS source
  ON target.product_id = source.product_id AND target.attribute_value_id = source.attribute_value_id
WHEN NOT MATCHED THEN
  INSERT (
    product_id,
    attribute_value_id
  )
  VALUES (
    source.product_id,
    source.attribute_value_id
  )
WHEN NOT MATCHED BY SOURCE AND target.product_id = sqlc.arg('product_id')::uuid THEN
  DELETE;

