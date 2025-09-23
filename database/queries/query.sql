-- vim: ft=sql.postgresql
-- name: GetUser :one
SELECT * FROM users
WHERE id =$1
LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET name = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
