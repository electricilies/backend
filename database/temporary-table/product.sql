-- name: CreateTempTableProductVariants :exec
CREATE TEMPORARY TABLE temp_product_variants (
  id UUID PRIMARY KEY,
  sku TEXT NOT NULL,
  price DECIMAL(12, 0) NOT NULL,
  quantity INTEGER NOT NULL,
  purchase_count INTEGER NOT NULL,
  product_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
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

-- name: CreateTempTableProductImages :exec
CREATE TEMPORARY TABLE temp_product_images (
  id UUID PRIMARY KEY,
  url TEXT NOT NULL,
  "order" INTEGER NOT NULL,
  product_id UUID,
  product_variant_id UUID,
  created_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
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

