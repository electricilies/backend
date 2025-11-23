# Database Conventions (PostgreSQL + sqlc)

## Schema Design

**Tables:** plural, snake_case (`products`, `product_variants`)  
**Columns:** snake_case (`created_at`, `category_id`)  
**Primary keys:** `id UUID PRIMARY KEY`  
**Foreign keys:** `<table>_id` (`product_id`, `category_id`)  
**Timestamps:** `created_at`, `updated_at`, `deleted_at` (all `TIMESTAMPTZ`)

## Data Types

- `UUID` - All IDs
- `VARCHAR(n)` - Text with max length
- `TEXT` - Unlimited text
- `BIGINT` - Money (cents), large numbers
- `INTEGER` - Regular numbers
- `DECIMAL(p,s)` - Exact decimals (avoid for money)
- `BOOLEAN` - Flags
- `TIMESTAMPTZ` - Timestamps with timezone
- `JSONB` - Structured JSON

## Table Example

```sql
CREATE TABLE products (
    id              UUID PRIMARY KEY,
    name            VARCHAR(200) NOT NULL,
    description     TEXT NOT NULL,
    price           BIGINT NOT NULL CHECK (price > 0),
    quantity        INTEGER NOT NULL CHECK (quantity >= 0),
    rating          DECIMAL(3,2) CHECK (rating >= 0 AND rating <= 5),
    category_id     UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_deleted_at ON products(deleted_at) WHERE deleted_at IS NULL;
```

## Constraints

```sql
-- Primary key
id UUID PRIMARY KEY

-- Foreign key
category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE

-- Check
CHECK (price > 0)
CHECK (rating >= 0 AND rating <= 5)

-- Unique
UNIQUE (email)
UNIQUE (category_id, attribute_id)

-- Not null
name VARCHAR(200) NOT NULL
```

## Indexes

```sql
-- Foreign keys (auto-indexed)
CREATE INDEX idx_products_category_id ON products(category_id);

-- Search
CREATE INDEX idx_products_name ON products(name);

-- Soft delete filtering
CREATE INDEX idx_products_deleted_at ON products(deleted_at) WHERE deleted_at IS NULL;

-- Full-text search
CREATE INDEX idx_products_search ON products 
  USING GIN(to_tsvector('english', name || ' ' || description));

-- Composite
CREATE INDEX idx_products_category_price ON products(category_id, price) 
  WHERE deleted_at IS NULL;
```

## Soft Deletes

```sql
deleted_at TIMESTAMPTZ

-- Exclude deleted (default)
WHERE deleted_at IS NULL

-- Only deleted
WHERE deleted_at IS NOT NULL

-- All records
WHERE TRUE
```

## sqlc Queries

### Upsert

```sql
-- name: UpsertProduct :exec
INSERT INTO products (
  id, name, description, price, created_at, updated_at, deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('name'),
  sqlc.arg('description'),
  sqlc.arg('price'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at'),
  sqlc.narg('deleted_at')
)
ON CONFLICT (id) DO UPDATE SET
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  price = EXCLUDED.price,
  updated_at = EXCLUDED.updated_at,
  deleted_at = EXCLUDED.deleted_at;
```

### Get

```sql
-- name: GetProduct :one
SELECT * FROM products WHERE id = $1;
```

### List with Filters

```sql
-- name: ListProducts :many
SELECT * FROM products
WHERE
  (sqlc.narg('ids')::uuid[] IS NULL OR id = ANY(sqlc.narg('ids')::uuid[]))
  AND (sqlc.narg('search') IS NULL OR name ILIKE '%' || sqlc.narg('search') || '%')
  AND (
    CASE sqlc.arg('deleted')
      WHEN 'exclude' THEN deleted_at IS NULL
      WHEN 'only' THEN deleted_at IS NOT NULL
      ELSE TRUE
    END
  )
ORDER BY created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
```

### Count

```sql
-- name: CountProducts :one
SELECT COUNT(*) FROM products
WHERE
  (sqlc.narg('ids')::uuid[] IS NULL OR id = ANY(sqlc.narg('ids')::uuid[]))
  AND (
    CASE sqlc.arg('deleted')
      WHEN 'exclude' THEN deleted_at IS NULL
      WHEN 'only' THEN deleted_at IS NOT NULL
      ELSE TRUE
    END
  );
```

## Query Annotations

- `:one` - Single row
- `:many` - Multiple rows
- `:exec` - No return
- `:execrows` - Affected row count

## Parameters

- `sqlc.arg('name')` - Required
- `sqlc.narg('name')` - Nullable/optional
- `sqlc.narg('ids')::uuid[]` - Array parameter

**Filter pattern:**
```sql
(sqlc.narg('param') IS NULL OR column = sqlc.narg('param'))
```

## Migrations

**Directory:** `migration/`  
**Naming:** `YYYYMMDDHHMMSS.sql`

```sql
-- +migrate Up
CREATE TABLE products (
    id   UUID PRIMARY KEY,
    name VARCHAR(200) NOT NULL
);

CREATE INDEX idx_products_name ON products(name);

-- +migrate Down
DROP INDEX IF EXISTS idx_products_name;
DROP TABLE IF EXISTS products;
```

## Triggers

```sql
-- Auto-update updated_at
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_products_updated_at
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();
```

## sqlc Config

**File:** `sqlc.yaml`

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "database/queries/"
    schema: "database/schema.sql"
    gen:
      go:
        package: "postgres"
        out: "internal/infrastructure/repository/postgres"
        emit_json_tags: true
        emit_interface: true
        emit_empty_slices: true
```

## Quick Rules

1. ✅ Tables: plural, snake_case
2. ✅ UUID for all IDs
3. ✅ BIGINT for money (cents)
4. ✅ TIMESTAMPTZ for all timestamps
5. ✅ Soft delete with `deleted_at`
6. ✅ Index FKs, search, sort columns
7. ✅ `sqlc.arg()` required, `sqlc.narg()` optional
8. ✅ Filter: `(param IS NULL OR condition)`
9. ✅ Always filter `deleted_at`
10. ✅ Upsert: `ON CONFLICT DO UPDATE`
11. ✅ Migrations: Up + Down
12. ✅ Use triggers for auto-fields
