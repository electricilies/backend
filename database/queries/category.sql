-- name: CreateCategory :one
INSERT INTO categories (
  name
)
VALUES (
  @name
)
RETURNING
  *;

-- name: GetCategories :many
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
  AND name ||| @name
ORDER BY
  pdb.score(id) DESC
LIMIT COALESCE(sqlc.narg('limit')::integer, 10);

-- name: UpdateCategory :one
UPDATE
  categories
SET
  name = @name
WHERE
  deleted_at IS NULL
  AND id = @id
RETURNING
  *;

-- name: DeleteCategory :execrows
UPDATE
  categories
SET
  deleted_at = NOW()
WHERE
  deleted_at IS NULL
  AND id = @id;
