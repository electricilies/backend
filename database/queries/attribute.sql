-- name: CreateAttribute :one
INSERT INTO attributes (
  code,
  name
)
VALUES (
  @code,
  @name
)
RETURNING
  *;

-- name: CreateAttributeValues :many
INSERT INTO attribute_values (
  attribute_id,
  value
)
SELECT
  UNNEST(@attribute_ids::integer[]) AS attribute_id,
  UNNEST(@values::text[]) AS value
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
    WHEN sqlc.narg('include_deleted_only')::boolean IS TRUE THEN deleted_at IS NOT NULL
    WHEN sqlc.narg('include_deleted_only')::boolean IS FALSE THEN deleted_at IS NULL
    ELSE TRUE
  END
ORDER BY
  id ASC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetAttributeByID :one
SELECT
  *
FROM
  attributes
WHERE
  id = @id;

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

-- name: UpdateAttribute :execrows
UPDATE
  attributes
SET
  code = @code,
  name = @name
WHERE
  id = @id;

-- name: UpdateAttributeValues :execrows
WITH updated_attribute_values AS (
  SELECT
    UNNEST(@ids::integer[]) AS id,
    UNNEST(@values::text[]) AS value
)
UPDATE
  attribute_values
SET
  value = updated_attribute_values.value
FROM
  updated_attribute_values
WHERE
  attribute_values.id = updated_attribute_values.id;

-- name: DeleteAttributes :execrows
UPDATE
  attributes
SET
  deleted_at = NOW()
WHERE
  id = ANY (@ids::integer[]);

-- name: DeleteAttributeValues :execrows
DELETE FROM
  attribute_values
WHERE
  id = ANY (@ids::integer[]);
