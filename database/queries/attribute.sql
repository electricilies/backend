-- name: UpsertAttribute :exec
INSERT INTO attributes (
  id,
  code,
  name,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('code'),
  sqlc.arg('name'),
  sqlc.narg('deleted_at')
)
ON CONFLICT (id) DO UPDATE SET
  code = EXCLUDED.code,
  name = EXCLUDED.name,
  deleted_at = EXCLUDED.deleted_at;

-- name: ListAttributes :many
SELECT
  *
FROM
  attributes
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE
      code ||| (sqlc.narg('search')::text)::pdb.fuzzy(2)
      OR name ||| (sqlc.narg('search')::text)::pdb.fuzzy(2)
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  CASE WHEN sqlc.narg('search')::text IS NOT NULL THEN pdb.score(id) END DESC,
  id ASC
OFFSET sqlc.arg('offset')::integer
LIMIT sqlc.arg('limit')::integer;

-- name: CountAttributes :one
SELECT
  COUNT(*) AS count
FROM
  attributes
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END;

-- name: GetAttribute :one
SELECT
  *
FROM
  attributes
WHERE
  id = sqlc.arg('id')
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END;

-- name: ListProductsAttributeValues :many
SELECT
  *
FROM
  products_attribute_values
WHERE
  CASE
    WHEN sqlc.narg('product_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('product_ids')::uuid[]) = 0 THEN TRUE
    ELSE product_id = ANY (sqlc.narg('product_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('attribute_value_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('attribute_value_ids')::uuid[]) = 0 THEN TRUE
    ELSE attribute_value_id = ANY (sqlc.narg('attribute_value_ids')::uuid[])
  END
ORDER BY
  product_id ASC,
  attribute_value_id ASC
OFFSET sqlc.arg('offset')::integer
LIMIT sqlc.arg('limit')::integer;

-- name: ListAttributeValues :many
SELECT
  *
FROM
  attribute_values
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('attribute_id')::uuid IS NULL THEN TRUE
    ELSE attribute_id = sqlc.narg('attribute_id')::uuid
  END
  AND CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE value ||| (sqlc.narg('search')::text)::pdb.fuzzy(2)
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  CASE WHEN sqlc.narg('search')::text IS NOT NULL THEN pdb.score(id) END DESC,
  id ASC;

-- name: MergeAttributeValuesFromTemp :exec
MERGE INTO attribute_values AS target
USING temp_attribute_values AS source
  ON target.id = source.id
WHEN MATCHED THEN
  UPDATE SET
    attribute_id = source.attribute_id,
    value = source.value,
    deleted_at = source.deleted_at
WHEN NOT MATCHED THEN
  INSERT (
    id,
    attribute_id,
    value,
    deleted_at
  )
  VALUES (
    source.id,
    source.attribute_id,
    source.value,
    source.deleted_at
  )
WHEN NOT MATCHED BY SOURCE AND target.attribute_id = sqlc.arg('attribute_id')::uuid THEN
  DELETE;
