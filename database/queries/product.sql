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
  product_id,
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
  *
FROM
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
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetProductByID :one
SELECT
  *
FROM
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
