-- name: CreateTempTableCartItems :exec
CREATE TEMPORARY TABLE temp_cart_items (
  id UUID PRIMARY KEY,
  quantity INTEGER NOT NULL,
  cart_id UUID NOT NULL,
  product_variant_id UUID NOT NULL
) ON COMMIT DROP;

-- name: InsertTempTableCartItems :copyfrom
INSERT INTO temp_cart_items (
  id,
  quantity,
  cart_id,
  product_variant_id
) VALUES (
  @id,
  @quantity,
  @cart_id,
  @product_variant_id
);
