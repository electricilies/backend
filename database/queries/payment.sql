-- name: CreatePayment :one
WITH payments AS (
  INSERT INTO payments (
    amount,
    status_id,
    provider_id
  )
  VALUES (
    sqlc.arg('amount'),
    sqlc.arg('status_id'),
    sqlc.arg('provider_id')
  )
  RETURNING
    *
)
SELECT
  sqlc.embed(payments),
  sqlc.embed(payment_statuses),
  sqlc.embed(payment_providers)
FROM payments
INNER JOIN payment_statuses
  ON payments.status_id = payment_statuses.id
INNER JOIN payment_providers
  ON payments.provider_id = payment_providers.id;

-- name: ListPayments :many
SELECT
  sqlc.embed(payments),
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  payments
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE payments.id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('status_ids')::integer[] IS NULL THEN TRUE
    ELSE payments.status_id = ANY (sqlc.narg('status_ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('provider_ids')::integer[] IS NULL THEN TRUE
    ELSE payments.provider_id = ANY (sqlc.narg('provider_ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('order_ids')::integer[] IS NULL THEN TRUE
    ELSE payments.order_id = ANY (sqlc.narg('order_ids')::integer[])
  END
ORDER BY
  id DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetPayment :one
SELECT
  sqlc.embed(payments)
FROM
  payments
WHERE
  CASE
    WHEN sqlc.narg('id')::integer IS NULL THEN TRUE
    ELSE payments.id = sqlc.narg('id')::integer
  END
  AND CASE
    WHEN sqlc.narg('order_id')::integer IS NULL THEN TRUE
    ELSE payments.order_id = sqlc.narg('order_id')::integer
  END;

-- name: ListPaymentProviders :many
SELECT
  *
FROM
  payment_providers
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
ORDER BY
  id;

-- name: GetPaymentProvider :one
SELECT
  *
FROM
  payment_providers
WHERE
  CASE
    WHEN sqlc.narg('id')::integer IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::integer
  END
  AND CASE
    WHEN sqlc.narg('name')::text IS NULL THEN TRUE
    ELSE name = sqlc.narg('name')::text
  END;

-- name: ListPaymentStatuses :many
SELECT
  *
FROM
  payment_statuses
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
ORDER BY
  id;

-- name: GetPaymentStatus :one
SELECT
  *
FROM
  payment_statuses
WHERE
  CASE
    WHEN sqlc.narg('id')::integer IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::integer
  END
  AND CASE
    WHEN sqlc.narg('name')::text IS NULL THEN TRUE
    ELSE name = sqlc.narg('name')::text
  END;

-- name: UpdatePayment :one
UPDATE payments
SET
  amount = COALESCE(sqlc.narg('amount')::decimal(12, 0), amount),
  provider_id = COALESCE(sqlc.narg('provider_id')::integer, provider_id),
  status_id = COALESCE(sqlc.narg('status_id')::integer, status_id),
  updated_at = COALESCE(sqlc.narg('updated_at')::timestamp, NOW())
WHERE
  id = sqlc.arg('id')::integer
RETURNING
  *;
