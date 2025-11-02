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
  ON payments.payment_status_id = payment_statuses.id
INNER JOIN payment_methods
  ON payments.payment_method_id = payment_methods.id
INNER JOIN payment_providers
  ON payments.payment_provider_id = payment_providers.id
ORDER BY
  payments.id DESC -- FIXME: Or?? Updated at?
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
  ON payments.payment_status_id = payment_statuses.id
INNER JOIN payment_methods
  ON payments.payment_method_id = payment_methods.id
INNER JOIN payment_providers
  ON payments.payment_provider_id = payment_providers.id
WHERE
  orders.id = @order_id;

-- name: CreatePayment :one
WITH payments AS (
  INSERT INTO payments (
    amount,
    payment_method_id,
    payment_status_id,
    payment_provider_id
  )
  VALUES (
    @amount,
    @payment_method_id,
    @payment_status_id,
    @payment_provider_id
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
  ON payments.payment_status_id = payment_statuses.id
INNER JOIN payment_methods
  ON payments.payment_method_id = payment_methods.id
INNER JOIN payment_providers
  ON payments.payment_provider_id = payment_providers.id;

-- name: UpdatePaymentStatus :one
WITH payments AS (
  UPDATE payments
  SET
    payment_status_id = @payment_status_id,
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
  ON payments.payment_status_id = payment_statuses.id
INNER JOIN payment_methods
  ON payments.payment_method_id = payment_methods.id
INNER JOIN payment_providers
  ON payments.payment_provider_id = payment_providers.id;
