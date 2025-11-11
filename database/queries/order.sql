-- name: CreateOrder :one
INSERT INTO orders (
  user_id,
  status_id
) VALUES (
  @user_id,
  @status_id
)
RETURNING
  *;

-- name: CreateOrderItem :one
INSERT INTO order_items (
  quantity,
  order_id,
  price_at_order,
  product_variant_id
) VALUES (
  @quantity,
  @order_id,
  @price_at_order,
  @product_variant_id
)
RETURNING
  *;

-- name: GetOrder :one
SELECT
  *
FROM
  orders
WHERE
  id = @id;
