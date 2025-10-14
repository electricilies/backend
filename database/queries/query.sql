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
  first_name,
  last_name,
  username,
  email,
  birthday,
  phone_number
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET
  avatar = COALESCE($2, avatar),
  first_name = COALESCE($3, first_name),
  last_name = COALESCE($4, last_name),
  username = COALESCE($5, username),
  email = COALESCE($6, email),
  birthday = COALESCE($7, birthday),
  phone_number = COALESCE($8, phone_number)
WHERE id = $1
  AND deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
  AND deleted_at IS NULL;
