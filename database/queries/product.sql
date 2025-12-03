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
  NULLIF(sqlc.arg('deleted_at')::timestamptz, '0001-01-01T00:00:00Z'::timestamptz)
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
  deleted_at = COALESCE(EXCLUDED.deleted_at, products.deleted_at);

-- This is used for list, search (with filter, order), suggest
-- name: ListProducts :many
SELECT
  products.*
FROM
  products
LEFT JOIN (
  SELECT
    categories.id AS category_id,
    pdb.score(categories.id) AS category_score
  FROM products
  INNER JOIN categories
    ON products.category_id = categories.id
  WHERE
    CASE
      WHEN sqlc.arg('search')::text = '' THEN FALSE
      ELSE (
        categories.name ||| sqlc.arg('search')::text
        AND categories.deleted_at IS NULL
      )
    END
) AS category_scores
  ON products.category_id = category_scores.category_id
WHERE
  CASE
    WHEN sqlc.arg('id')::uuid IS NULL THEN TRUE
    WHEN sqlc.arg('id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE products.id = sqlc.arg('id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE products.id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('search')::text = '' THEN TRUE
    ELSE products.name ||| sqlc.arg('search')::text
  END
  AND CASE
    WHEN sqlc.arg('min_price')::decimal = 0 THEN TRUE
    ELSE products.price >= sqlc.arg('min_price')::decimal
  END
  AND CASE
    WHEN sqlc.arg('max_price')::decimal = 0 THEN TRUE
    ELSE products.price <= sqlc.arg('max_price')::decimal
  END
  AND CASE
    WHEN sqlc.arg('rating')::real = 0 THEN TRUE
    ELSE products.rating >= sqlc.arg('rating')::real
  END
  AND CASE
    WHEN sqlc.arg('category_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('category_ids')::uuid[]) = 0 THEN TRUE
    ELSE products.category_id = ANY (sqlc.arg('category_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('variant_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('variant_ids')::uuid[]) = 0 THEN TRUE
    ELSE EXISTS (
      SELECT 1
      FROM product_variants
      WHERE product_variants.product_id = products.id
        AND product_variants.id = ANY (sqlc.arg('variant_ids')::uuid[])
    )
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN products.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN products.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE products.deleted_at IS NULL
  END
ORDER BY
  CASE WHEN
    sqlc.arg('search')::text <> '' THEN pdb.score(products.id) + category_scores.category_score + products.trending_score
  END DESC,
  CASE WHEN
    sqlc.arg('sort_rating')::text = 'asc' THEN products.rating
  END ASC,
  CASE WHEN
    sqlc.arg('sort_rating')::text = 'desc' THEN products.rating
  END DESC,
  CASE WHEN
   sqlc.arg('sort_price')::text = 'asc' THEN products.price
  END ASC,
  CASE WHEN
   sqlc.arg('sort_price')::text = 'desc' THEN products.price
  END DESC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountProducts :one
SELECT
  COUNT(*) AS count
FROM
  products
WHERE
  CASE
    WHEN sqlc.arg('id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE products.id = sqlc.arg('id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE products.id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('min_price')::decimal = 0 THEN TRUE
    ELSE products.price >= sqlc.arg('min_price')::decimal
  END
  AND CASE
    WHEN sqlc.arg('max_price')::decimal = 0 THEN TRUE
    ELSE products.price <= sqlc.arg('max_price')::decimal
  END
  AND CASE
    WHEN sqlc.arg('rating')::real = 0 THEN TRUE
    ELSE products.rating >= sqlc.arg('rating')::real
  END
  AND CASE
    WHEN sqlc.arg('category_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('category_ids')::uuid[]) = 0 THEN TRUE
    ELSE products.category_id = ANY (sqlc.arg('category_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN products.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN products.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE products.deleted_at IS NULL
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
    ELSE deleted_at IS NULL
  END;

-- name: ListProductVariants :many
SELECT
  *
FROM
  product_variants
WHERE
  CASE
    WHEN sqlc.arg('id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE id = sqlc.arg('id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('sku')::text = '' THEN TRUE
    ELSE sku = sqlc.arg('sku')::text
  END
  AND CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('product_id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE product_id = sqlc.arg('product_id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('product_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('product_ids')::uuid[]) = 0 THEN TRUE
    ELSE product_id = ANY (sqlc.arg('product_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  id
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

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
    ELSE deleted_at IS NULL
  END;

-- name: ListProductsAttributeValues :many
SELECT
  *
FROM
  products_attribute_values
WHERE
  CASE
    WHEN sqlc.arg('product_id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE product_id = sqlc.arg('product_id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('product_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('product_ids')::uuid[]) = 0 THEN TRUE
    ELSE product_id = ANY (sqlc.arg('product_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('attribute_value_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('attribute_value_ids')::uuid[]) = 0 THEN TRUE
    ELSE attribute_value_id = ANY (sqlc.arg('attribute_value_ids')::uuid[])
  END
ORDER BY
  product_id ASC,
  attribute_value_id ASC;

-- name: ListProductImages :many
SELECT
  *
FROM
  product_images
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('product_variant_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('product_variant_ids')::uuid[]) = 0 THEN TRUE
    ELSE product_variant_id = ANY (sqlc.arg('product_variant_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('product_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('product_ids')::uuid[]) = 0 THEN TRUE
    ELSE product_id = ANY (sqlc.arg('product_ids')::uuid[])
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

-- name: CreateTempTableProductVariants :exec
CREATE TEMPORARY TABLE temp_product_variants (
  id UUID PRIMARY KEY,
  sku TEXT NOT NULL,
  price DECIMAL(12, 0) NOT NULL,
  quantity INTEGER NOT NULL,
  purchase_count INTEGER NOT NULL,
  product_id UUID NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ
) ON COMMIT DROP;

-- name: InsertTempTableProductVariants :copyfrom
INSERT INTO temp_product_variants (
  id,
  sku,
  price,
  quantity,
  purchase_count,
  product_id,
  created_at,
  updated_at,
  deleted_at
) VALUES (
  @id,
  @sku,
  @price,
  @quantity,
  @purchase_count,
  @product_id,
  @created_at,
  @updated_at,
  @deleted_at
);

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
    deleted_at = COALESCE(NULLIF(source.deleted_at, '0001-01-01T00:00:00Z'::timestamptz), target.deleted_at)
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
    NULLIF(source.deleted_at, '0001-01-01T00:00:00Z'::timestamptz)
  )
WHEN NOT MATCHED BY SOURCE
  AND target.product_id = ANY (SELECT DISTINCT id FROM temp_product_variants) THEN
  DELETE;

-- name: CreateTempTableProductImages :exec
CREATE TEMPORARY TABLE temp_product_images (
  id UUID PRIMARY KEY,
  url TEXT NOT NULL,
  "order" INTEGER NOT NULL,
  product_id UUID NOT NULL,
  product_variant_id UUID,
  created_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ
) ON COMMIT DROP;

-- name: InsertTempTableProductImages :copyfrom
INSERT INTO temp_product_images (
  id,
  url,
  "order",
  product_id,
  product_variant_id,
  created_at,
  deleted_at
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
);

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
    deleted_at = COALESCE(NULLIF(source.deleted_at, '0001-01-01T00:00:00Z'::timestamptz), target.deleted_at)
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
    NULLIF(source.deleted_at, '0001-01-01T00:00:00Z'::timestamptz)
  )
WHEN NOT MATCHED BY SOURCE
  AND target.product_id = ANY (SELECT DISTINCT product_id FROM temp_product_images) THEN
  DELETE;

-- name: CreateTempTableProductsAttributeValues :exec
CREATE TEMPORARY TABLE temp_products_attribute_values (
  product_id UUID NOT NULL,
  attribute_value_id UUID NOT NULL,
  PRIMARY KEY (product_id, attribute_value_id)
) ON COMMIT DROP;

-- name: InsertTempTableProductsAttributeValues :copyfrom
INSERT INTO temp_products_attribute_values (
  product_id,
  attribute_value_id
) VALUES (
  @product_id,
  @attribute_value_id
);

-- name: MergeProductsAttributeValuesFromTemp :exec
MERGE INTO products_attribute_values AS target
USING temp_products_attribute_values AS source
  ON target.product_id = source.product_id
    AND target.attribute_value_id = source.attribute_value_id
WHEN NOT MATCHED THEN
  INSERT (
    product_id,
    attribute_value_id
  )
  VALUES (
    source.product_id,
    source.attribute_value_id
  )
WHEN NOT MATCHED BY SOURCE
  AND target.product_id = ANY (SELECT DISTINCT product_id FROM temp_products_attribute_values) THEN
  DELETE;

-- name: CreateTempTableOptionValuesProductVariants :exec
CREATE TEMPORARY TABLE temp_option_values_product_variants (
  product_variant_id UUID NOT NULL,
  option_value_id UUID NOT NULL,
  PRIMARY KEY (product_variant_id, option_value_id)
) ON COMMIT DROP;

-- name: InsertTempTableOptionValuesProductVariants :copyfrom
INSERT INTO temp_option_values_product_variants (
  product_variant_id,
  option_value_id
) VALUES (
  @product_variant_id,
  @option_value_id
);

-- name: MergeOptionValuesProductVariantsFromTemp :exec
MERGE INTO option_values_product_variants AS target
USING temp_option_values_product_variants AS source
  ON target.option_value_id = source.option_value_id
WHEN NOT MATCHED THEN
  INSERT (
    product_variant_id,
    option_value_id
  )
  VALUES (
    source.product_variant_id,
    source.option_value_id
  )
WHEN NOT MATCHED BY SOURCE
  AND target.option_value_id = ANY (SELECT id FROM temp_option_values) THEN
  DELETE;
