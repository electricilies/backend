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

-- name: GetCart :many
SELECT
  *
FROM
  carts
WHERE
  CASE
    WHEN sqlc.narg('id')::integer IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::integer
  END
  AND CASE
    WHEN sqlc.narg('user_id')::UUID IS NULL THEN TRUE
    ELSE user_id = sqlc.narg('user_id')::UUID
  END;

-- name: GetCartItems :many
SELECT
  *
FROM
  cart_items
WHERE
  CASE
    WHEN sqlc.narg('cart_id')::integer IS NULL THEN TRUE
    ELSE cart_id = sqlc.narg('cart_id')::integer
  END
ORDER BY
  id ASC;

-- name: UpdateCartItemByID :execrows
UPDATE cart_items
SET
  quantity = @quantity
WHERE
  id = @id;

-- name: DeleteCartItemByIDs :execrows
DELETE FROM cart_items
WHERE
  id = ANY (@ids::UUID[]);
