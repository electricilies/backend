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
  *
FROM
  carts
WHERE
  carts.user_id = @user_id;

-- name: GetCartItemByCarID :many
SELECT
  *
FROM
  cart_items
WHERE
  cart_items.cart_id = @cart_id;

-- name: UpdateCartItemByID :execrows
UPDATE cart_items
SET
  quantity = @quantity
WHERE
  id = @id;

-- name: DeleteCartItemByIDs :execrows
DELETE FROM cart_items
WHERE
  id = ANY(@ids::UUID[]);
