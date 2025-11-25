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
  attributes.*
FROM
  attributes
LEFT JOIN (
  SELECT
    id,
    attribute_id
  FROM
    attribute_values
  WHERE
    CASE
      WHEN sqlc.narg('attribute_value_ids')::uuid[] IS NULL THEN TRUE
      WHEN cardinality(sqlc.narg('attribute_value_ids')::uuid[]) = 0 THEN TRUE
      ELSE attribute_values.id = ANY (sqlc.narg('attribute_value_ids')::uuid[])
    END
) AS av ON attributes.id = av.attribute_id
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE attributes.id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE
      code ||| (sqlc.narg('search')::text)
      OR name ||| (sqlc.narg('search')::text)
  END
  AND CASE
    WHEN sqlc.narg('attribute_value_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('attribute_value_ids')::uuid[]) = 0 THEN TRUE
    ELSE av.id IS NOT NULL
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  CASE WHEN sqlc.narg('search')::text IS NOT NULL THEN pdb.score(attributes.id) END DESC,
  attributes.id ASC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

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
    ELSE deleted_at IS NULL
  END;

-- name: ListAttributeByAttributeValues :many
SELECT
  attributes.*
FROM
  attributes
JOIN
  attribute_values ON attribute_values.attribute_id = attributes.id
WHERE
  CASE
    WHEN sqlc.narg('attribute_value_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('attribute_value_ids')::uuid[]) = 0 THEN TRUE
    ELSE attribute_values.id = ANY (sqlc.narg('attribute_value_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN attributes.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN attributes.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE attributes.deleted_at IS NULL
  END
GROUP BY
  attributes.id
ORDER BY
  attributes.id ASC;

-- name: GetAttribute :one
SELECT
  *
FROM
  attributes
WHERE
  id = sqlc.arg('id')::uuid
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;

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
    WHEN sqlc.narg('attribute_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('attribute_ids')::uuid[]) = 0 THEN TRUE
     ELSE attribute_id = ANY (sqlc.narg('attribute_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE value ||| (sqlc.narg('search')::text)
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  CASE WHEN sqlc.narg('search')::text IS NOT NULL THEN pdb.score(id) END DESC,
  id ASC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountAttributeValues :one
SELECT
  COUNT(*) AS count
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
    WHEN sqlc.narg('attribute_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('attribute_ids')::uuid[]) = 0 THEN TRUE
     ELSE attribute_id = ANY (sqlc.narg('attribute_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;

-- name: CreateTempTableAttributeValues :exec
CREATE TEMPORARY TABLE temp_attribute_values (
  id UUID PRIMARY KEY,
  attribute_id UUID NOT NULL,
  value TEXT NOT NULL,
  deleted_at TIMESTAMPTZ
) ON COMMIT DROP;

-- name: InsertTempTableAttributeValues :copyfrom
INSERT INTO temp_attribute_values (
  id,
  attribute_id,
  value,
  deleted_at
) VALUES (
  @id,
  @attribute_id,
  @value,
  @deleted_at
);

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
WHEN NOT MATCHED BY SOURCE
  AND target.attribute_id IN (SELECT DISTINCT attribute_id FROM temp_attribute_values) THEN
  DELETE;
