-- name: GetAllUsers :many
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
  id = @id;

-- name: CreateUser :one
INSERT INTO users (
  id
) VALUES (
  @id
)
RETURNING
  *;
