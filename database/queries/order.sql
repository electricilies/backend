-- name: CreateOrder :one
INSERT INTO orders (
  user_id,
  address,
  total_amount,
  provider_id,
  status_id
) VALUES (
  sqlc.arg('user_id'),
  sqlc.arg('address'),
  sqlc.arg('total_amount'),
  sqlc.arg('provider_id'),
  sqlc.arg('status_id')
)
RETURNING
  *;

-- name: CreateOrderItems :many
WITH inserts AS (
  SELECT
    sqlc.arg('quantity') AS quantity,
    sqlc.arg('order_id') AS order_id,
    sqlc.arg('price') AS price,
    sqlc.arg('product_variant_id') AS product_variant_id
)
INSERT INTO order_items (
  quantity,
  order_id,
  price,
  product_variant_id
) VALUES (
  inserts.quantity,
  inserts.order_id,
  inserts.price,
  inserts.product_variant_id
)
RETURNING
  *;

-- name: ListOrders :many
SELECT
  *
FROM
  orders
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('user_ids')::uuid[] IS NULL THEN TRUE
    ELSE user_id = ANY (sqlc.narg('user_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('status_ids')::uuid[] IS NULL THEN TRUE
    ELSE status_id = ANY (sqlc.narg('status_ids')::uuid[])
  END
ORDER BY
  id ASC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: CountOrders :one
SELECT
  COUNT(*) AS count
FROM
  orders
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('user_ids')::uuid[] IS NULL THEN TRUE
    ELSE user_id = ANY (sqlc.narg('user_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('status_ids')::uuid[] IS NULL THEN TRUE
    ELSE status_id = ANY (sqlc.narg('status_ids')::uuid[])
  END;

-- name: GetOrder :one
SELECT
  *
FROM
  orders
WHERE
  id = sqlc.arg('id');

-- name: ListOrderItems :many
SELECT
  *
FROM
  order_items
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('order_ids')::uuid[] IS NULL THEN TRUE
    ELSE order_id = ANY (sqlc.narg('order_ids')::uuid[])
  END
ORDER BY
  id;

-- name: GetOrderItem :one
SELECT
  *
FROM
  order_items
WHERE
  id = sqlc.arg('id');

-- name: ListOrderStatuses :many
SELECT
  *
FROM
  order_statuses
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
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
    WHEN sqlc.narg('id')::uuid IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::uuid
  END
  AND CASE
    WHEN sqlc.narg('name')::text IS NULL THEN TRUE
    ELSE name = sqlc.narg('name')::text
  END;

-- name: GetOrderProvider :one
SELECT
  *
FROM
  order_providers
WHERE
  CASE
    WHEN sqlc.narg('id')::uuid IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::uuid
  END
  AND CASE
    WHEN sqlc.narg('name')::text IS NULL THEN TRUE
    ELSE name = sqlc.narg('name')::text
  END;

-- name: UpdateOrder :one
UPDATE orders
SET
  user_id = COALESCE(sqlc.arg('user_id')::uuid, user_id),
  status_id = COALESCE(sqlc.arg('status_id')::uuid, status_id),
  updated_at = COALESCE(sqlc.narg('updated_at')::timestamp, NOW())
WHERE
  id = sqlc.arg('id')
RETURNING
  *;
