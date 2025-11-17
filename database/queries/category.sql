-- name: CreateCategory :one
INSERT INTO categories (
  name
)
VALUES (
  sql.arg('name')
)
RETURNING
  *;

-- TODO: Get category by name (paradedb)
-- name: ListCategories :many
SELECT
  *,
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  categories
WHERE
  deleted_at IS NULL
ORDER BY
  id DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetSuggestedCategories :many
SELECT
  *
FROM
  categories
WHERE
  deleted_at IS NULL
  AND name ||| sql.arg('name')
ORDER BY
  pdb.score(id) DESC
LIMIT COALESCE(sqlc.narg('limit')::integer, 10);

-- name: UpdateCategory :one
UPDATE
  categories
SET
  name = sql.arg('name'),
  updated_at = COALESCE(sqlc.narg('updated_at')::timestamp, NOW())
WHERE
  deleted_at IS NULL
  AND id = sql.arg('id')
RETURNING
  *;

-- name: DeleteCategory :execrows
UPDATE
  categories
SET
  deleted_at = NOW()
WHERE
  deleted_at IS NULL
  AND id = sql.arg('id');
