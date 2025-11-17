-- name: CreateCategory :one
INSERT INTO categories (
  name
)
VALUES (
  sqlc.arg('name')
)
RETURNING
  *;

-- name: ListCategories :many
SELECT
  *,
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  categories
WHERE
  CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE name ||| (sqlc.narg('search')::text)::pdb.fuzzy(2)
  END
  AND deleted_at IS NULL
ORDER BY
  id DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: UpdateCategory :one
UPDATE
  categories
SET
  name = sqlc.arg('name'),
  updated_at = COALESCE(sqlc.narg('updated_at')::timestamp, NOW())
WHERE
  deleted_at IS NULL
  AND id = sqlc.arg('id')
RETURNING
  *;

-- name: DeleteCategory :execrows
UPDATE
  categories
SET
  deleted_at = NOW()
WHERE
  deleted_at IS NULL
  AND id = sqlc.arg('id');
