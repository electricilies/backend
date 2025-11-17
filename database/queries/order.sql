-- name: CreateOrder :one
INSERT INTO orders (
  user_id,
  status_id
) VALUES (
  sql.arg('user_id'),
  sql.arg('status_id')
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
  sql.arg('quantity'),
  sql.arg('order_id'),
  sql.arg('price_at_order'),
  sql.arg('product_variant_id')
)
RETURNING
  *;

-- name: ListOrders :many
SELECT
  *,
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  orders
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('user_ids')::uuid[] IS NULL THEN TRUE
    ELSE user_id = ANY (sqlc.narg('user_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('status_ids')::integer[] IS NULL THEN TRUE
    ELSE status_id = ANY (sqlc.narg('status_ids')::integer[])
  END
ORDER BY
  id ASC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetOrder :one
SELECT
  *
FROM
  orders
WHERE
  id = sql.arg('id');

-- name: ListOrderItems :many
SELECT
  *
FROM
  order_items
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('order_ids')::integer[] IS NULL THEN TRUE
    ELSE order_id = ANY (sqlc.narg('order_ids')::integer[])
  END
ORDER BY
  id;

-- name: GetOrderItem :one
SELECT
  *
FROM
  order_items
WHERE
  id = sql.arg('id');

-- name: ListOrderStatuses :many
SELECT
  *
FROM
  order_statuses
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
ORDER BY
  id ASC;

-- name: GetOrderStatus :one
SELECT
  *
FROM
  order_statuses
WHERE
  CASE
    WHEN sqlc.narg('id')::integer IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::integer
  END
  AND CASE
    WHEN sqlc.narg('name')::text IS NULL THEN TRUE
    ELSE name = sqlc.narg('name')::text
  END;

-- name: UpdateOrder :one
UPDATE orders
SET
  user_id = COALESCE(sql.arg('user_id'), user_id),
  status_id = COALESCE(sql.arg('status_id'), status_id),
  updated_at = COALESCE(sqlc.narg('updated_at')::timestamp, NOW())
WHERE
  id = sql.arg('id')
RETURNING
  *;
