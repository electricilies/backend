-- name: UpsertOrder :exec
INSERT INTO orders (
  id,
  user_id,
  address,
  total_amount,
  is_paid,
  provider_id,
  status_id,
  created_at,
  updated_at
) VALUES (
  sqlc.arg('id'),
  sqlc.arg('user_id'),
  sqlc.arg('address'),
  sqlc.arg('total_amount'),
  sqlc.arg('is_paid'),
  sqlc.arg('provider_id'),
  sqlc.arg('status_id'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at')
)
ON CONFLICT (id) DO UPDATE SET
  user_id = EXCLUDED.user_id,
  address = EXCLUDED.address,
  total_amount = EXCLUDED.total_amount,
  is_paid = EXCLUDED.is_paid,
  provider_id = EXCLUDED.provider_id,
  status_id = EXCLUDED.status_id,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at;

-- name: ListOrders :many
SELECT
  *
FROM
  orders
WHERE
  CASE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN cardinality(sqlc.arg('user_ids')::uuid[]) = 0 THEN TRUE
    ELSE user_id = ANY (sqlc.arg('user_ids')::uuid[])
  END
  AND CASE
    WHEN cardinality(sqlc.arg('status_ids')::uuid[]) = 0 THEN TRUE
    ELSE status_id = ANY (sqlc.arg('status_ids')::uuid[])
  END
ORDER BY
  id ASC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountOrders :one
SELECT
  COUNT(*) AS count
FROM
  orders
WHERE
  CASE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN cardinality(sqlc.arg('user_ids')::uuid[]) = 0 THEN TRUE
    ELSE user_id = ANY (sqlc.arg('user_ids')::uuid[])
  END
  AND CASE
    WHEN cardinality(sqlc.arg('status_ids')::uuid[]) = 0 THEN TRUE
    ELSE status_id = ANY (sqlc.arg('status_ids')::uuid[])
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
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN cardinality(sqlc.arg('order_ids')::uuid[]) = 0 THEN TRUE
    ELSE order_id = ANY (sqlc.arg('order_ids')::uuid[])
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
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
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
    WHEN sqlc.arg('id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE id = sqlc.arg('id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('name')::text = '' THEN TRUE
    ELSE name = sqlc.arg('name')::text
  END;

-- name: GetOrderProvider :one
SELECT
  *
FROM
  order_providers
WHERE
  CASE
    WHEN sqlc.arg('id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE id = sqlc.arg('id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('name')::text = '' THEN TRUE
    ELSE name = sqlc.arg('name')::text
  END;
