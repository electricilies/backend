-- noqa: disable=AM04


-- name: CreateUser :one
INSERT INTO users (
  id
) VALUES (
  $1
)
RETURNING *;
