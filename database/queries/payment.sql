-- name: GetPayments :many
SELECT
  sqlc.embed(payments),
  sqlc.embed(payment_statuses),
  sqlc.embed(payment_methods),
  sqlc.embed(payment_providers),
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  payments
INNER JOIN payment_statuses
  ON payments.status_id = payment_statuses.id
INNER JOIN payment_methods
  ON payments.method_id = payment_methods.id
INNER JOIN payment_providers
  ON payments.provider_id = payment_providers.id
ORDER BY
  payments.id DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetPaymentByOrderID :one
SELECT
  sqlc.embed(payments),
  sqlc.embed(payment_statuses),
  sqlc.embed(payment_methods),
  sqlc.embed(payment_providers)
FROM
  payments
INNER JOIN orders
  ON payments.id = orders.payment_id
INNER JOIN payment_statuses
  ON payments.status_id = payment_statuses.id
INNER JOIN payment_methods
  ON payments.method_id = payment_methods.id
INNER JOIN payment_providers
  ON payments.provider_id = payment_providers.id
WHERE
  orders.id = @order_id;

-- name: CreatePayment :one
WITH payments AS (
  INSERT INTO payments (
    amount,
    method_id,
    status_id,
    provider_id
  )
  VALUES (
    @amount,
    @method_id,
    @status_id,
    @provider_id
  )
  RETURNING
    *
)
SELECT
  sqlc.embed(payments),
  sqlc.embed(payment_statuses),
  sqlc.embed(payment_methods),
  sqlc.embed(payment_providers)
FROM payments
INNER JOIN payment_statuses
  ON payments.status_id = payment_statuses.id
INNER JOIN payment_methods
  ON payments.method_id = payment_methods.id
INNER JOIN payment_providers
  ON payments.provider_id = payment_providers.id;

-- name: UpdatePaymentStatus :one
WITH payments AS (
  UPDATE payments
  SET
    status_id = @status_id,
    updated_at = NOW()
  WHERE
    payments.id = @id -- # HACK: Wtf sqlc?
  RETURNING
    *
)
SELECT
  sqlc.embed(payments),
  sqlc.embed(payment_statuses),
  sqlc.embed(payment_methods),
  sqlc.embed(payment_providers)
FROM payments
INNER JOIN payment_statuses
  ON payments.status_id = payment_statuses.id
INNER JOIN payment_methods
  ON payments.method_id = payment_methods.id
INNER JOIN payment_providers
  ON payments.provider_id = payment_providers.id;
