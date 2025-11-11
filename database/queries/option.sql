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

-- name: ListOptions :many
SELECT
  *
FROM
  options
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE options.id = ANY(sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('product_id')::integer IS NULL THEN TRUE
    ELSE options.product_id = sqlc.narg('product_id')::integer
  END
  AND CASE
    WHEN sqlc.narg('include_deleted_only')::bool THEN deleted_at IS NOT NULL
    WHEN sqlc.narg('include_deleted_only')::bool = FALSE THEN deleted_at IS NULL
    ELSE TRUE
  END
ORDER BY
  options.id;

-- name: GetOptionByID :one
SELECT
  *
FROM
  options
WHERE
  id = @id::integer
  AND CASE
    WHEN sqlc.narg('include_deleted_only')::bool THEN deleted_at IS NOT NULL
    WHEN sqlc.narg('include_deleted_only')::bool = FALSE THEN deleted_at IS NULL
    ELSE TRUE
  END;

-- name: UpdateOptions :execrows
WITH updated_options AS (
  SELECT
    UNNEST(@ids::integer[]) AS id,
    UNNEST(@names::text[]) AS name
)
UPDATE options
SET
  name = updated_options.name
FROM
  updated_options
WHERE
  options.id = updated_options.id
  AND options.deleted_at IS NULL;

-- name: UpdateOptionValues :execrows
WITH updated_option_values AS (
  SELECT
    UNNEST(@ids::integer[]) AS id,
    UNNEST(@values::text[]) AS value
)
UPDATE option_values
SET
  value = updated_option_values.value
FROM
  updated_option_values
WHERE
  option_values.id = updated_option_values.id;

-- name: DeleteOptions :execrows
UPDATE
  options
SET
  deleted_at = NOW()
WHERE
  id = ANY(@ids::integer[])
  AND deleted_at IS NULL;

-- name: DeleteOptionValues :execrows
DELETE FROM
  option_values
WHERE
  id = ANY(@ids::integer[]);
