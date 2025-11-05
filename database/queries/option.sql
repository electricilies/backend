-- name: CreateOption :one
INSERT INTO options (
  name
)
VALUES (
  @name
)
RETURNING
  *;

-- name: CreateOptionValues :many
INSERT INTO option_values (
  option_id,
  value
)
SELECT
  @option_id,
  UNNEST(@values::text[]) AS value
RETURNING
  *;

-- name: GetOptions :many
SELECT
  sqlc.embed(options),
  sqlc.embed(option_values)
FROM
  options,
  option_values
WHERE
  options.id = option_values.option_id
  AND options.deleted_at IS NULL;

-- name: GetOptionByID :one
SELECT
  sqlc.embed(options),
  sqlc.embed(option_values)
FROM
  options,
  option_values
WHERE
  options.id = @id::integer
  AND option_values.option_id = options.id
  AND options.deleted_at IS NULL;

-- name: UpdateOptionValue :one
UPDATE option_values
SET
  value = @value
WHERE
  id = @id
  AND deleted_at IS NULL
RETURNING
  *;

-- name: DeleteOptionValue :execrows
WITH _ AS (
  DELETE FROM
    option_values
  WHERE
    id = @id
    AND deleted_at IS NULL
),
_ AS (
  DELETE FROM
    option_values_product_variants
  WHERE
    option_value_id = @id
)
SELECT 1;
