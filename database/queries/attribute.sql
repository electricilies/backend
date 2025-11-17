-- name: CreateAttribute :one
INSERT INTO attributes (
  code,
  name
)
VALUES (
  sql.arg('code'),
  sql.arg('name')
)
RETURNING
  *;

-- name: CreateAttributeValues :many
INSERT INTO attribute_values (
  attribute_id,
  value
)
SELECT
  UNNEST(sql.arg('attribute_ids')::integer[]) AS attribute_id,
  UNNEST(sql.arg('values')::text[]) AS value
RETURNING
  *;

-- name: ListAttributes :many
SELECT
  *,
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  attributes
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE
      code ||| (sqlc.narg('search')::text)::pdb.fuzzy(2)
      OR name ||| (sqlc.narg('search')::text)::pdb.fuzzy(2)
  END
  AND CASE
    WHEN sql.arg('deleted')::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN sql.arg('deleted')::text = 'only' THEN deleted_at IS NULL
    WHEN sql.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  CASE WHEN sqlc.narg('search')::text IS NOT NULL THEN pdb.score(id) END DESC,
  id ASC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetAttribute :one
SELECT
  *
FROM
  attributes
WHERE
  id = sql.arg('id');

-- name: ListProductsAttributeValues :many
SELECT
  *
FROM
  products_attribute_values
WHERE
  CASE
    WHEN sqlc.narg('product_ids')::integer[] IS NULL THEN TRUE
    ELSE product_id = ANY (sqlc.narg('product_ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('attribute_value_ids')::integer[] IS NULL THEN TRUE
    ELSE attribute_value_id = ANY (sqlc.narg('attribute_value_ids')::integer[])
  END;

-- name: ListAttributeValues :many
SELECT
  *
FROM
  attribute_values
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('attribute_ids')::integer[] IS NULL THEN TRUE
    ELSE attribute_id = ANY (sqlc.narg('attribute_ids')::integer[])
  END
ORDER BY
  id ASC;

-- name: UpdateAttribute :one
UPDATE
  attributes
SET
  code = COALESCE(sqlc.narg('code')::varchar(100), code),
  name = COALESCE(sqlc.narg('name')::text, name)
WHERE
  id = sql.arg('id')
RETURNING
  *;

-- name: UpdateAttributeValues :many
WITH updated_attribute_values AS (
  SELECT
    UNNEST(sql.arg('ids')::integer[]) AS id,
    UNNEST(sql.arg('values')::text[]) AS value
)
UPDATE
  attribute_values
SET
  value = COALESCE(updated_attribute_values.value, attribute_values.value)
FROM
  updated_attribute_values
WHERE
  attribute_values.id = updated_attribute_values.id
RETURNING
  attribute_values.*;

-- name: DeleteAttributes :execrows
UPDATE
  attributes
SET
  deleted_at = NOW()
WHERE
  id = ANY (sql.arg('ids')::integer[]);

-- name: DeleteAttributeValues :execrows
DELETE FROM
  attribute_values
WHERE
  id = ANY (sql.arg('ids')::integer[]);
