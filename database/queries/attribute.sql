-- name: GetAllAttributes :many
SELECT
  *
FROM
  attributes
ORDER BY
  id ASC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

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
DELETE FROM
  attributes
WHERE
  id = @id;
