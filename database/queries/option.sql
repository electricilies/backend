-- name: UpsertOption :exec
INSERT INTO options (
  id,
  name,
  product_id,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('name'),
  sqlc.arg('product_id'),
  sqlc.narg('deleted_at')
)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  product_id = EXCLUDED.product_id,
  deleted_at = EXCLUDED.deleted_at;

-- name: ListOptions :many
SELECT
  *
FROM
  options
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE options.id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('product_id')::uuid IS NULL THEN TRUE
    ELSE options.product_id = sqlc.narg('product_id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  options.id;

-- name: GetOption :one
SELECT
  *
FROM
  options
WHERE
  id = sqlc.arg('id')::uuid
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;

-- name: ListOptionValuesProductVariants :many
SELECT
  *
FROM
  option_values_product_variants
WHERE
  CASE
    WHEN sqlc.narg('option_value_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('option_value_ids')::uuid[]) = 0 THEN TRUE
    ELSE option_values_product_variants.option_value_id = ANY (sqlc.narg('option_value_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('product_variant_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('product_variant_ids')::uuid[]) = 0 THEN TRUE
    ELSE option_values_product_variants.product_variant_id = ANY (sqlc.narg('product_variant_ids')::uuid[])
  END;

-- name: ListOptionValues :many
SELECT
  *
FROM
  option_values
WHERE
  CASE
    WHEN sqlc.narg('id')::uuid IS NULL THEN TRUE
    ELSE option_values.id = sqlc.narg('id')::uuid
  END
  AND CASE
    WHEN sqlc.narg('option_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('option_ids')::uuid[]) = 0 THEN TRUE
    ELSE option_values.option_id = ANY (sqlc.narg('option_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  option_values.id;

-- name: CreateTempTableOptionValues :exec
CREATE TEMPORARY TABLE temp_option_values (
  id UUID PRIMARY KEY,
  value TEXT NOT NULL,
  option_id UUID NOT NULL,
  deleted_at TIMESTAMP
) ON COMMIT DROP;

-- name: InsertTempTableOptionValues :copyfrom
INSERT INTO temp_option_values (
  id,
  value,
  option_id,
  deleted_at
) VALUES (
  @id,
  @value,
  @option_id,
  @deleted_at
);

-- name: MergeOptionValuesFromTemp :exec
MERGE INTO option_values AS target
USING temp_option_values AS source
  ON target.id = source.id
WHEN MATCHED THEN
  UPDATE SET
    value = source.value,
    option_id = source.option_id,
    deleted_at = source.deleted_at
WHEN NOT MATCHED THEN
  INSERT (
    id,
    value,
    option_id,
    deleted_at
  )
  VALUES (
    source.id,
    source.value,
    source.option_id,
    source.deleted_at
  )
WHEN NOT MATCHED BY SOURCE
  AND target.option_id = (SELECT DISTINCT option_id FROM source) THEN
  DELETE;
