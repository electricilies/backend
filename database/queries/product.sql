-- name: CreateProduct :exec
INSERT INTO products (
  name,
  description
)
VALUES (
  @name,
  @description
);

-- name: CreateProductVariants :execresult
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
  @product_id;

-- name: CreateProductImages :execresult
INSERT INTO product_images (
  url,
  "order",
  product_variant_id
)
SELECT
  UNNEST(@urls::text[]) AS url,
  UNNEST(@orders::integer[]) AS "order",
  @product_id,
  UNNEST(@product_variant_ids::integer[]) AS product_variant_id
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

-- name: LinkProductOptions :execrows
INSERT INTO option_values_product_variants (
  option_value_id,
  product_variant_id
)
SELECT
  UNNEST(@option_value_ids::integer[]) AS option_value_id,
  UNNEST(@product_variant_ids::integer[]) AS product_variant_id;

-- name: LinkProductCategories :execrows
INSERT INTO products_categories (
  product_id,
  category_id
)
SELECT
  @product_id,
  UNNEST(@category_ids::integer[]) AS category_id;

-- name: GetAllProducts :many
SELECT
  sqlc.embed(products),
  sqlc.embed(product_variants),
  sqlc.embed(product_images),
  sqlc.embed(products_attribute_values),
  sqlc.embed(attribute_values),
  sqlc.embed(attributes),
  sqlc.embed(option_values_product_variants),
  sqlc.embed(option_values),
  sqlc.embed(options),
  sqlc.embed(categories),
  sqlc.embed(products_categories),
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  products,
  product_variants,
  product_images,
  products_attribute_values,
  attribute_values,
  attributes,
  option_values_product_variants,
  option_values,
  options,
  categories,
  products_categories
INNER JOIN product_variants
  ON products.id = product_variants.product_id
INNER JOIN product_images
  ON product_variants.id = product_images.product_variant_id
INNER JOIN products_attribute_values
  ON products.id = products_attribute_values.product_id
INNER JOIN attribute_values
  ON products_attribute_values.attribute_value_id = attribute_values.id
INNER JOIN attributes
  ON attribute_values.attribute_id = attributes.id
INNER JOIN option_values_product_variants
  ON product_variants.id = option_values_product_variants.product_variant_id
INNER JOIN option_values
  ON option_values_product_variants.option_value_id = option_values.id
INNER JOIN options
  ON option_values.option_id = options.id
INNER JOIN products_categories
  ON products.id = products_categories.product_id
INNER JOIN categories
  ON products_categories.category_id = categories.id
WHERE
  products.deleted_at IS NULL
  AND categories.deleted_at IS NULL
  AND (
    sqlc.narg('category_id')::integer IS NULL
    OR categories.id = sqlc.narg('category_id')::integer
  )
  AND (
    sqlc.narg('search')::text IS NULL
    OR products.name ||| sqlc.narg('search')::pdb.fuzzy(products.trending_score) -- TODO: Have to check does paradedb support this
    OR categories.name ||| sqlc.narg('search')::pdb.fuzzy(2)
  )
  -- TODO: Do we support rating?
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetProductByID :one
SELECT
  sqlc.embed(products),
  sqlc.embed(product_variants),
  sqlc.embed(product_images),
  sqlc.embed(products_attribute_values),
  sqlc.embed(attribute_values),
  sqlc.embed(attributes),
  sqlc.embed(option_values_product_variants),
  sqlc.embed(option_values),
  sqlc.embed(options),
  sqlc.embed(categories),
  sqlc.embed(products_categories)
FROM
  products,
  product_variants,
  product_images,
  products_attribute_values,
  attribute_values,
  attributes,
  option_values_product_variants,
  option_values,
  options,
  categories,
  products_categories
INNER JOIN product_variants
  ON products.id = product_variants.product_id
INNER JOIN product_images
  ON product_variants.id = product_images.product_variant_id
INNER JOIN products_attribute_values
  ON products.id = products_attribute_values.product_id
INNER JOIN attribute_values
  ON products_attribute_values.attribute_value_id = attribute_values.id
INNER JOIN attributes
  ON attribute_values.attribute_id = attributes.id
INNER JOIN option_values_product_variants
  ON product_variants.id = option_values_product_variants.product_variant_id
INNER JOIN option_values
  ON option_values_product_variants.option_value_id = option_values.id
INNER JOIN options
  ON option_values.option_id = options.id
INNER JOIN products_categories
  ON products.id = products_categories.product_id
INNER JOIN categories
  ON products_categories.category_id = categories.id
WHERE
  products.id = @id::integer -- sqlc requires this
  AND products.deleted_at IS NULL;

-- name: GetSuggestedProducts :many
-- TODO: Will we implement this?

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
  purchase_count = updated_variants.purchase_count
  -- FIXME: Maybe missing updated_at field?
FROM
  updated_variants
WHERE
  product_variants.id = updated_variants.id
  AND product_variants.deleted_at IS NULL;

-- naee: UpdateProductImages :execrows
-- TODO: Will we implement this?

-- name: DeleteProducts :execrows
UPDATE
  products
SET
  deleted_at = NOW()
WHERE
  deleted_at IS NULL
  AND id = ANY(@ids::integer[]);

-- name: DeleteProductVariants :execrows
-- TODO: Soft delete or delete all
