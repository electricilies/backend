-- name: UpsertUser :exec
INSERT INTO users (
  id
) VALUES (
  sqlc.arg('id')
)
ON CONFLICT (id) DO NOTHING;

-- name: ListUsers :many
SELECT
  *
FROM
  users
ORDER BY
  id ASC;

-- name: GetUserByID :one
SELECT
  *
FROM
  users
WHERE
  id = sqlc.arg('id')::uuid;
