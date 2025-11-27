-- name: UpsertCategory :exec
INSERT INTO categories (
  id,
  name,
  created_at,
  updated_at,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('name'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at'),
  NULLIF(sqlc.arg('deleted_at'), '0001-01-01T00:00:00Z'::timestamptz)
)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at,
  deleted_at = COALESCE(EXCLUDED.deleted_at, categories.deleted_at);

-- name: ListCategories :many
SELECT
  *
FROM
  categories
WHERE
  CASE
    WHEN sqlc.arg('search')::text = '' THEN TRUE
    ELSE name ||| (sqlc.arg('search')::text)
  END
  AND CASE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  id DESC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountCategories :one
SELECT
  COUNT(*) AS count
FROM
  categories
WHERE
  CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;

-- name: GetCategory :one
SELECT
  *
FROM
  categories
WHERE
  id = sqlc.arg('id')
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;
