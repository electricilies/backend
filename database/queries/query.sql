-- noqa: disable=AM04

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
  AND deleted_at IS NULL
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: CreateUser :one
INSERT INTO users (
  avatar,
  birthday,
  phone_number
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET
  avatar = COALESCE($2, avatar),
  birthday = COALESCE($3, birthday),
  phone_number = COALESCE($4, phone_number)
WHERE id = $1
  AND deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
  AND deleted_at IS NULL;
