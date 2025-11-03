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

-- name: GetAllOptions :many
SELECT
  sqlc.embed(options),
  sqlc.embed(option_values)
FROM
  options,
  option_values
WHERE
  options.deleted_at IS NULL;

-- name: GetOptionByID :one
SELECT
  sqlc.embed(options),
  sqlc.embed(option_values)
FROM
  options,
  option_values
WHERE
  options.id = @id::integer
  AND options.deleted_at IS NULL;
