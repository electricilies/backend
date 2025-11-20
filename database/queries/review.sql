-- name: CreateReview :one
INSERT INTO reviews (
  rating,
  content,
  image_url,
  user_id,
  order_item_id
)
VALUES (
  sqlc.arg('rating'),
  sqlc.arg('content'),
  sqlc.arg('image_url'),
  sqlc.arg('user_id'),
  sqlc.arg('order_item_id')
)
RETURNING
  *;

-- name: ListReviews :many
SELECT
  *
FROM
  reviews
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('order_item_ids')::integer[] IS NULL THEN TRUE
    ELSE order_item_id = ANY (sqlc.narg('order_item_ids')::integer[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  created_at DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 10);

-- name: CountReviews :one
SELECT
  COUNT(*) AS count
FROM
  reviews
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('order_item_ids')::integer[] IS NULL THEN TRUE
    ELSE order_item_id = ANY (sqlc.narg('order_item_ids')::integer[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END;

-- name: UpdateReview :one
UPDATE reviews
SET
  rating = COALESCE(sqlc.narg('rating')::integer, rating),
  content = COALESCE(sqlc.narg('content')::text, content),
  image_url = COALESCE(sqlc.narg('image_url')::text, image_url),
  updated_at = COALESCE(sqlc.narg('updated_at')::timestamp, NOW())
WHERE
  id = sqlc.arg('id')::integer
  AND deleted_at IS NULL
RETURNING
  *;

-- name: DeleteReviews :execrows
UPDATE reviews
SET
  deleted_at = NOW()
WHERE
  id = ANY(sqlc.arg('ids')::integer[])
  AND deleted_at IS NULL;
