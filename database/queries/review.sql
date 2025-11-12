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
    WHEN @deleted::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN @deleted::text = 'only' THEN deleted_at IS NULL
    WHEN @deleted::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  created_at DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 10);

-- name: UpdateReview :one
UPDATE reviews
SET
  rating = COALESCE(sqlc.narg('rating')::integer, rating),
  content = COALESCE(sqlc.narg('content')::text, content),
  image_url = COALESCE(sqlc.narg('image_url')::text, image_url),
  updated_at = NOW()
WHERE
  id = @id::integer
  AND deleted_at IS NULL
RETURNING
  *;

-- name: DeleteReviews :execrows
UPDATE reviews
SET
  deleted_at = NOW()
WHERE
  id = ANY(@ids::integer[])
  AND deleted_at IS NULL;
