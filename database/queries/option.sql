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
  NULLIF(sqlc.arg('deleted_at')::timestamptz, '0001-01-01T00:00:00Z'::timestamptz)
)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  product_id = EXCLUDED.product_id,
  deleted_at = COALESCE(EXCLUDED.deleted_at, options.deleted_at);

-- name: ListOptions :many
SELECT
  *
FROM
  options
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE options.id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('product_id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE options.product_id = sqlc.arg('product_id')::uuid
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
    WHEN sqlc.arg('option_value_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('option_value_ids')::uuid[]) = 0 THEN TRUE
    ELSE option_values_product_variants.option_value_id = ANY (sqlc.arg('option_value_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('product_variant_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('product_variant_ids')::uuid[]) = 0 THEN TRUE
    ELSE option_values_product_variants.product_variant_id = ANY (sqlc.arg('product_variant_ids')::uuid[])
  END;

-- name: ListOptionValues :many
SELECT
  *
FROM
  option_values
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE option_values.id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('option_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('option_ids')::uuid[]) = 0 THEN TRUE
    ELSE option_values.option_id = ANY (sqlc.arg('option_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  option_values.id;

-- name: CreateTempTableOptions :exec
CREATE TEMPORARY TABLE temp_options (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  product_id UUID NOT NULL,
  deleted_at TIMESTAMPTZ
) ON COMMIT DROP;

-- name: InsertTempTableOptions :copyfrom
INSERT INTO temp_options (
  id,
  name,
  product_id,
  deleted_at
) VALUES (
  @id,
  @name,
  @product_id,
  @deleted_at
);

-- name: MergeOptionsFromTemp :exec
MERGE INTO options AS target
USING temp_options AS source
  ON target.id = source.id
WHEN MATCHED THEN
  UPDATE SET
    name = source.name,
    product_id = source.product_id,
    deleted_at = COALESCE(NULLIF(source.deleted_at, '0001-01-01T00:00:00Z'::timestamptz), target.deleted_at)
WHEN NOT MATCHED THEN
  INSERT (
    id,
    name,
    product_id,
    deleted_at
  )
  VALUES (
    source.id,
    source.name,
    source.product_id,
    NULLIF(source.deleted_at, '0001-01-01T00:00:00Z'::timestamptz)
  )
WHEN NOT MATCHED BY SOURCE
  AND target.product_id = (SELECT DISTINCT product_id FROM temp_options) THEN
  DELETE;

-- name: CreateTempTableOptionValues :exec
CREATE TEMPORARY TABLE temp_option_values (
  id UUID PRIMARY KEY,
  value TEXT NOT NULL,
  option_id UUID NOT NULL,
  deleted_at TIMESTAMPTZ
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
    deleted_at = COALESCE(NULLIF(source.deleted_at, '0001-01-01T00:00:00Z'::timestamptz), target.deleted_at)
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
    NULLIF(source.deleted_at, '0001-01-01T00:00:00Z'::timestamptz)
  )
WHEN NOT MATCHED BY SOURCE
  AND target.option_id = (SELECT DISTINCT option_id FROM temp_option_values) THEN
  DELETE;
