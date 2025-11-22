-- name: CreateTempTableAttributeValues :exec
CREATE TEMPORARY TABLE temp_attribute_values (
  id UUID PRIMARY KEY,
  attribute_id UUID NOT NULL,
  value TEXT NOT NULL,
  deleted_at TIMESTAMP
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
