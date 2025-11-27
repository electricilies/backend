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
  NULLIF(sqlc.arg('deleted_at')::timestamptz, '0001-01-01T00:00:00Z'::timestamptz)
)
ON CONFLICT (id) DO UPDATE SET
  rating = EXCLUDED.rating,
  content = EXCLUDED.content,
  image_url = EXCLUDED.image_url,
  user_id = EXCLUDED.user_id,
  order_item_id = EXCLUDED.order_item_id,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at,
  deleted_at = COALESCE(EXCLUDED.deleted_at, reviews.deleted_at);

-- name: ListReviews :many
SELECT
  reviews.*
FROM
  reviews
LEFT JOIN order_items ON reviews.order_item_id = order_items.id
LEFT JOIN product_variants ON order_items.product_variant_id = product_variants.id
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE reviews.id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('order_item_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('order_item_ids')::uuid[]) = 0 THEN TRUE
    ELSE reviews.order_item_id = ANY (sqlc.arg('order_item_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('product_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('product_ids')::uuid[]) = 0 THEN TRUE
    ELSE product_variants.product_id = ANY (sqlc.arg('product_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN reviews.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN reviews.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE reviews.deleted_at IS NOT NULL
  END
ORDER BY
  reviews.created_at DESC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountReviews :one
SELECT
  COUNT(*) AS count
FROM
  reviews
LEFT JOIN order_items ON reviews.order_item_id = order_items.id
LEFT JOIN product_variants ON order_items.product_variant_id = product_variants.id
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE reviews.id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('order_item_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('order_item_ids')::uuid[]) = 0 THEN TRUE
    ELSE reviews.order_item_id = ANY (sqlc.arg('order_item_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('product_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('product_ids')::uuid[]) = 0 THEN TRUE
    ELSE product_variants.product_id = ANY (sqlc.arg('product_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN reviews.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN reviews.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE reviews.deleted_at IS NOT NULL
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
    ELSE deleted_at IS NULL
  END;
