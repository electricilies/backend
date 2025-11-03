-- name: CreateCart :one
INSERT INTO carts (
  user_id
) VALUES (
  @user_id
)
RETURNING
  *;

-- name: CreateCartItem :one
INSERT INTO cart_items (
  quantity,
  cart_id,
  product_variant_id
) VALUES (
  @quantity,
  @cart_id,
  @product_variant_id
)
RETURNING
  *;

-- name: GetCartByUserID :many
SELECT
  sqlc.embed(carts),
  sqlc.embed(cart_items),
  sqlc.embed(products),
  sqlc.embed(product_variants)
FROM
  carts
INNER JOIN cart_items
  ON carts.id = cart_items.cart_id
INNER JOIN product_variants
  ON cart_items.product_variant_id = product_variants.id
INNER JOIN products
  ON product_variants.product_id = products.id
WHERE
  carts.user_id = @user_id
ORDER BY
  cart_items.id ASC;

-- name: UpdateCartItemByID :one
UPDATE cart_items
SET
  quantity = @quantity
WHERE
  id = @id
RETURNING
  *;

-- name: DeleteCartItemByID :execrows
DELETE FROM cart_items
WHERE
  id = @id;
