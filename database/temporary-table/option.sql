-- name: CreateTempTableOptionValues :exec
CREATE TEMPORARY TABLE temp_option_values (
  id UUID PRIMARY KEY,
  value TEXT NOT NULL,
  option_id UUID NOT NULL,
  deleted_at TIMESTAMP
) ON COMMIT DROP;

-- name: InsertTempTableOptionValues :copyfrom
INSERT INTO temp_option_values (
  id,
  value,
  option_id,
  deleted_at
) VALUES (
  @id,
  @value,
  @option_id,
  @deleted_at
);

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
