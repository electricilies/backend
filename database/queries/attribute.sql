-- name: CreateAttribute :one
INSERT INTO attributes (
  code,
  name
)
VALUES (
  sqlc.arg('code'),
  sqlc.arg('name')
)
RETURNING
  *;

-- name: CreateAttributeValue :one
INSERT INTO attribute_values (
  attribute_id,
  value
)
SELECT
  UNNEST(sqlc.arg('attribute_id')::integer) AS attribute_id,
  UNNEST(sqlc.arg('value')::text) AS value
RETURNING
  *;

-- name: ListAttributes :many
SELECT
  *
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
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  CASE WHEN sqlc.narg('search')::text IS NOT NULL THEN pdb.score(id) END DESC,
  id ASC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: CountAttributes :one
SELECT
  COUNT(*) AS count
FROM
  attributes
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END;

-- name: GetAttribute :one
SELECT
  *
FROM
  attributes
WHERE
  id = sqlc.arg('id');

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
  AND CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE value ||| (sqlc.narg('search')::text)::pdb.fuzzy(2)
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
  id = sqlc.arg('id')
RETURNING
  *;

-- name: UpdateAttributeValues :many
WITH updated_attribute_values AS (
  SELECT
    UNNEST(sqlc.arg('ids')::integer[]) AS id,
    UNNEST(sqlc.arg('values')::text[]) AS value
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
  id = ANY (sqlc.arg('ids')::integer[]);

-- name: DeleteAttributeValues :execrows
DELETE FROM
  attribute_values
WHERE
  id = ANY (sqlc.arg('ids')::integer[]);
