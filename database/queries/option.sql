-- name: CreateOption :one
INSERT INTO options (
  name,
  product_id
)
VALUES (
   @name,
   @product_id
)
RETURNING
  *;

-- name: CreateOptionValues :many
INSERT INTO option_values (
  option_id,
  value
)
SELECT
  @option_id,
  UNNEST(@values::text[]) AS value
RETURNING
  *;

-- name: ListOptions :many
SELECT
  *
FROM
  options
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE options.id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('product_id')::integer IS NULL THEN TRUE
    ELSE options.product_id = sqlc.narg('product_id')::integer
  END
  AND CASE
    WHEN @deleted::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN @deleted::text = 'only' THEN deleted_at IS NULL
    WHEN @deleted::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  options.id;

-- name: GetOption :one
SELECT
  *
FROM
  options
WHERE
  id = @id::integer
  AND CASE
    WHEN @deleted::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN @deleted::text = 'only' THEN deleted_at IS NULL
    WHEN @deleted::text = 'all' THEN TRUE
    ELSE FALSE
  END;

-- name: ListOptionValuesProductVariants :many
SELECT
  *
FROM
  option_values_product_variants
WHERE
  CASE
    WHEN sqlc.narg('option_value_ids')::integer[] IS NULL THEN TRUE
    ELSE option_values_product_variants.option_value_id = ANY (sqlc.narg('option_value_ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('product_variant_ids')::integer[] IS NULL THEN TRUE
    ELSE option_values_product_variants.product_variant_id = ANY (sqlc.narg('product_variant_ids')::integer[])
  END;

-- name: ListOptionValues :many
SELECT
  *
FROM
  option_values
WHERE
  CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE option_values.id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('option_ids')::integer[] IS NULL THEN TRUE
    ELSE option_values.option_id = ANY (sqlc.narg('option_ids')::integer[])
  END
ORDER BY
  option_values.id;

-- name: UpdateOptions :many
WITH updated_options AS (
  SELECT
    UNNEST(@ids::integer[]) AS id,
    UNNEST(@names::text[]) AS name
)
UPDATE options
SET
  name = COALESCE(updated_options.name, options.name)
FROM
  updated_options
WHERE
  options.id = updated_options.id
  AND options.deleted_at IS NULL
RETURNING
  options.*;

-- name: UpdateOptionValues :many
WITH updated_option_values AS (
  SELECT
    UNNEST(@ids::integer[]) AS id,
    UNNEST(@values::text[]) AS value
)
UPDATE option_values
SET
  value = COALESCE(updated_option_values.value, option_values.value)
FROM
  updated_option_values
WHERE
  option_values.id = updated_option_values.id
RETURNING
  option_values.*;

-- name: DeleteOptions :execrows
UPDATE
  options
SET
  deleted_at = NOW()
WHERE
  id = ANY (@ids::integer[])
  AND deleted_at IS NULL;

-- name: DeleteOptionValues :execrows
DELETE FROM
  option_values
WHERE
  id = ANY (@ids::integer[]);
