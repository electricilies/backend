-- name: UpsertReview :exec
INSERT INTO reviews (
  id,
  rating,
  content,
  image_url,
  user_id,
  order_item_id,
  created_at,
  updated_at,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('rating'),
  sqlc.arg('content'),
  sqlc.arg('image_url'),
  sqlc.arg('user_id'),
  sqlc.arg('order_item_id'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at'),
  sqlc.narg('deleted_at')
)
ON CONFLICT (id) DO UPDATE SET
  rating = EXCLUDED.rating,
  content = EXCLUDED.content,
  image_url = EXCLUDED.image_url,
  user_id = EXCLUDED.user_id,
  order_item_id = EXCLUDED.order_item_id,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at,
  deleted_at = EXCLUDED.deleted_at;

-- name: ListReviews :many
SELECT
  *
FROM
  reviews
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('order_item_ids')::uuid[] IS NULL THEN TRUE
    ELSE order_item_id = ANY (sqlc.narg('order_item_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
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
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('order_item_ids')::uuid[] IS NULL THEN TRUE
    ELSE order_item_id = ANY (sqlc.narg('order_item_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END;

-- name: GetReview :one
SELECT
  *
FROM
  reviews
WHERE
  id = sqlc.arg('id')
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END;

