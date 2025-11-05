-- name: CreateAttribute :one
INSERT INTO attributes (
  code,
  name
)
VALUES (
  @code,
  @name
)
RETURNING
  *;

-- name: CreateAttributeValues :many
INSERT INTO attribute_values (
  attribute_id,
  value
)
SELECT
  UNNEST(@attribute_ids::integer[]) AS attribute_id,
  UNNEST(@values::text[]) AS value
RETURNING
  *;

-- name: GetAttributes :many
SELECT
  *,
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  attributes
ORDER BY
  id ASC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetAttributeByID :one
SELECT
  *
FROM
  attributes
WHERE
  id = @id;

-- name: UpdateAttribute :one
UPDATE
  attributes
SET
  code = @code,
  name = @name
WHERE
  id = @id
RETURNING
  *;

-- name: DeleteAttribute :execrows
UPDATE
  attributes
SET
  deleted_at = NOW()
WHERE
  id = @id;
