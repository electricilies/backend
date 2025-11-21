-- name: CreateCart :one
INSERT INTO carts (
  user_id
) VALUES (
  sqlc.arg('user_id')
)
RETURNING
  *;

-- name: CreateCartItem :one
INSERT INTO cart_items (
  quantity,
  cart_id,
  product_variant_id
) VALUES (
  sqlc.arg('quantity'),
  sqlc.arg('cart_id'),
  sqlc.arg('product_variant_id')
)
RETURNING
  *;

-- name: GetCart :one
SELECT
  *
FROM
  carts
WHERE
  CASE
    WHEN sqlc.narg('id')::uuid IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::uuid
  END
  AND CASE
    WHEN sqlc.narg('user_id')::uuid IS NULL THEN TRUE
    ELSE user_id = sqlc.narg('user_id')::uuid
  END;

-- name: GetCartItems :many
SELECT
  *
FROM
  cart_items
WHERE
  CASE
    WHEN sqlc.narg('cart_id')::uuid IS NULL THEN TRUE
    ELSE cart_id = sqlc.narg('cart_id')::uuid
  END
ORDER BY
  id ASC;

-- name: UpdateCartItem :one
UPDATE cart_items
SET
  quantity = COALESCE(sqlc.narg('quantity')::integer, quantity)
WHERE
  id = sqlc.arg('id')
RETURNING
  *;

-- name: DeleteCartItemByIDs :execrows
DELETE FROM cart_items
WHERE
  id = ANY (sqlc.arg('ids')::uuid[]);
