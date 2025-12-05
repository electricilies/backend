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
  NULLIF(sqlc.arg('deleted_at')::timestamptz, '0001-01-01T00:00:00Z'::timestamptz)
)
ON CONFLICT (id) DO UPDATE SET
  code = EXCLUDED.code,
  name = EXCLUDED.name,
  deleted_at = COALESCE(EXCLUDED.deleted_at, attributes.deleted_at);

-- name: ListAttributes :many
SELECT
  attributes.*
FROM
  attributes
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE attributes.id::uuid = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('search')::text = '' THEN TRUE
    ELSE
      attributes.name ||| sqlc.arg('search')::text
      OR attributes.code ||| sqlc.arg('search')::text
  END
  AND CASE
    WHEN sqlc.arg('attribute_value_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('attribute_value_ids')::uuid[]) = 0 THEN TRUE
    ELSE EXISTS (
      SELECT 1
      FROM attribute_values
      WHERE
        attribute_values.attribute_id = attributes.id
        AND attribute_values.id = ANY (sqlc.arg('attribute_value_ids')::uuid[])
    )
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  CASE WHEN sqlc.arg('search')::text <> '' THEN pdb.score(attributes.id) END DESC,
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
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
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
    WHEN sqlc.arg('attribute_value_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('attribute_value_ids')::uuid[]) = 0 THEN TRUE
    ELSE attribute_values.id = ANY (sqlc.arg('attribute_value_ids')::uuid[])
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
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('attribute_id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE attribute_id = sqlc.arg('attribute_id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('attribute_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('attribute_ids')::uuid[]) = 0 THEN TRUE
     ELSE attribute_id = ANY (sqlc.arg('attribute_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('search')::text = '' THEN TRUE
    ELSE value ||| (sqlc.arg('search')::text)
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  CASE WHEN sqlc.arg('search')::text <> '' THEN pdb.score(id) END DESC,
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
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('attribute_id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE attribute_id = sqlc.arg('attribute_id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('attribute_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('attribute_ids')::uuid[]) = 0 THEN TRUE
     ELSE attribute_id = ANY (sqlc.arg('attribute_ids')::uuid[])
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
    deleted_at = COALESCE(NULLIF(source.deleted_at, '0001-01-01T00:00:00Z'::timestamptz), target.deleted_at)
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
    NULLIF(source.deleted_at, '0001-01-01T00:00:00Z'::timestamptz)
  )
WHEN NOT MATCHED BY SOURCE
  AND target.attribute_id IN (SELECT DISTINCT attribute_id FROM temp_attribute_values) THEN
  DELETE;
