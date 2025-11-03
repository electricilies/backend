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

-- name: GetReviewsByProductID :many
SELECT
  *,
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  reviews
WHERE
  product_id = @product_id
  AND deleted_at IS NULL
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
