-- name: CreateReview :one
INSERT INTO reviews (
  rating,
  content,
  image_url,
  user_id,
  product_id
)
VALUES (
  @rating,
  @content,
  @image_url,
  @user_id,
  @product_id
)
RETURNING
  *;

-- name: ListReviews :many
SELECT
  *,
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  reviews
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('product_ids')::integer[] IS NULL THEN TRUE
    ELSE product_id = ANY (sqlc.narg('product_ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('include_deleted_only')::boolean IS TRUE THEN deleted_at IS NOT NULL
    WHEN sqlc.narg('include_deleted_only')::boolean IS FALSE THEN deleted_at IS NULL
    ELSE TRUE
  END
ORDER BY
  created_at DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 10);

-- name: DeleteReviews :execrows
UPDATE reviews
SET
  deleted_at = NOW()
WHERE
  id = ANY(@ids::integer[])
  AND deleted_at IS NULL;
