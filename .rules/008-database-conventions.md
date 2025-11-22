# Database Conventions

## Overview

This project uses PostgreSQL with sqlc for type-safe SQL queries. Database schema, migrations, and queries follow consistent conventions.

## Schema Design

### Table Naming

- Use plural nouns: `products`, `categories`, `orders`
- Snake_case: `product_variants`, `option_values`
- Junction tables: `product_attributes` (singular_singular)

### Column Naming

- Snake_case: `created_at`, `total_purchase`, `category_id`
- Primary keys: `id` (UUID)
- Foreign keys: `<table>_id` (e.g., `product_id`, `category_id`)
- Timestamps: `created_at`, `updated_at`, `deleted_at`

### Data Types

**Standard types:**

- `UUID` - All primary keys and foreign keys
- `VARCHAR` - Text with known max length
- `TEXT` - Unlimited text (descriptions, content)
- `BIGINT` - Large numbers (prices in cents, counts)
- `INTEGER` - Regular numbers (quantities, small counts)
- `DECIMAL/NUMERIC` - Exact decimal values (avoid for money, use BIGINT cents)
- `BOOLEAN` - True/false flags
- `TIMESTAMPTZ` - Timestamps with timezone
- `JSONB` - Structured JSON data

**Example:**

```sql
CREATE TABLE products (
    id              UUID PRIMARY KEY,
    name            VARCHAR(200) NOT NULL,
    description     TEXT NOT NULL,
    price           BIGINT NOT NULL CHECK (price > 0),
    quantity        INTEGER NOT NULL CHECK (quantity >= 0),
    rating          DECIMAL(3,2) CHECK (rating >= 0 AND rating <= 5),
    views_count     INTEGER NOT NULL DEFAULT 0,
    total_purchase  INTEGER NOT NULL DEFAULT 0,
    trending_score  BIGINT NOT NULL DEFAULT 0,
    category_id     UUID NOT NULL REFERENCES categories(id),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);
```

### Constraints

**Primary Keys:**

```sql
id UUID PRIMARY KEY
```

**Foreign Keys:**

```sql
category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE
```

**Check Constraints:**

```sql
CHECK (price > 0)
CHECK (rating >= 0 AND rating <= 5)
CHECK (quantity >= 0)
```

**Unique Constraints:**

```sql
UNIQUE (email)
UNIQUE (category_id, attribute_id)
```

**Not Null:**

```sql
name VARCHAR(200) NOT NULL
```

### Indexes

Create indexes for:

- Foreign keys (automatically indexed)
- Frequently queried columns
- Sort columns
- Full-text search

```sql
-- Foreign key (automatic)
CREATE INDEX idx_products_category_id ON products(category_id);

-- Search optimization
CREATE INDEX idx_products_name ON products(name);

-- Soft delete filtering
CREATE INDEX idx_products_deleted_at ON products(deleted_at) WHERE deleted_at IS NULL;

-- Full-text search
CREATE INDEX idx_products_search ON products USING GIN(to_tsvector('english', name || ' ' || description));

-- Composite index for common queries
CREATE INDEX idx_products_category_price ON products(category_id, price) WHERE deleted_at IS NULL;
```

### Soft Deletes

Use `deleted_at` timestamp:

```sql
deleted_at TIMESTAMPTZ

-- Queries exclude deleted by default
WHERE deleted_at IS NULL

-- Include deleted
WHERE deleted_at IS NOT NULL

-- Show all
WHERE TRUE
```

## Migrations

### File Structure

**Directory:** `migration/`

**Naming:** `YYYYMMDDHHMMSS.sql` (timestamp-based)

**Example:** `20250923113651.sql`

### Migration Content

```sql
-- +migrate Up
CREATE TABLE products (
    id              UUID PRIMARY KEY,
    name            VARCHAR(200) NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_products_name ON products(name);

-- +migrate Down
DROP INDEX IF EXISTS idx_products_name;
DROP TABLE IF EXISTS products;
```

**Conventions:**

- Always include both Up and Down
- Create indexes after table creation
- Use `IF EXISTS` / `IF NOT EXISTS` for idempotency
- Group related changes together

### Schema Management

**Main schema:** `database/schema.sql` (complete current state)

**Migration tool:** Atlas

**Configuration:** `atlas.hcl`

```hcl
env "local" {
  src = "file://database/schema.sql"
  url = "postgres://user:pass@localhost:5432/dbname?sslmode=disable"
  dev = "docker://postgres/16/dev"
}
```

## sqlc Queries

### Query Files

**Location:** `database/queries/*.sql`

**Organization:**

- One file per aggregate: `product.sql`, `category.sql`, `user.sql`
- Related queries together

### Query Annotations

**Naming:** `-- name: <QueryName> :<type>`

**Types:**

- `:one` - Returns single row
- `:many` - Returns multiple rows
- `:exec` - No return value
- `:execrows` - Returns affected row count

### Query Patterns

**Insert/Upsert:**

```sql
-- name: UpsertProduct :exec
INSERT INTO products (
    id,
    name,
    description,
    price,
    category_id,
    created_at,
    updated_at,
    deleted_at
)
VALUES (
    sqlc.arg('id'),
    sqlc.arg('name'),
    sqlc.arg('description'),
    sqlc.arg('price'),
    sqlc.arg('category_id'),
    sqlc.arg('created_at'),
    sqlc.arg('updated_at'),
    sqlc.narg('deleted_at')
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    price = EXCLUDED.price,
    category_id = EXCLUDED.category_id,
    updated_at = EXCLUDED.updated_at,
    deleted_at = EXCLUDED.deleted_at;
```

**Get by ID:**

```sql
-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1;
```

**List with filters:**

```sql
-- name: ListProducts :many
SELECT * FROM products
WHERE
    -- Optional ID filter
    (sqlc.narg('ids')::uuid[] IS NULL OR id = ANY(sqlc.narg('ids')::uuid[]))
    -- Search filter
    AND (sqlc.narg('search') IS NULL OR name ILIKE '%' || sqlc.narg('search') || '%')
    -- Price range
    AND (sqlc.narg('min_price') IS NULL OR price >= sqlc.narg('min_price'))
    AND (sqlc.narg('max_price') IS NULL OR price <= sqlc.narg('max_price'))
    -- Soft delete filter
    AND (
        CASE sqlc.arg('deleted')
            WHEN 'exclude' THEN deleted_at IS NULL
            WHEN 'only' THEN deleted_at IS NOT NULL
            ELSE TRUE
        END
    )
ORDER BY
    CASE WHEN sqlc.narg('sort_price') = 'asc' THEN price END ASC,
    CASE WHEN sqlc.narg('sort_price') = 'desc' THEN price END DESC,
    created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
```

**Count:**

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

**Delete (soft):**

```sql
-- name: SoftDeleteProduct :exec
UPDATE products
SET deleted_at = NOW(), updated_at = NOW()
WHERE id = $1;
```

**Delete (hard):**

```sql
-- name: HardDeleteProduct :exec
DELETE FROM products WHERE id = $1;
```

### Parameter Conventions

**Required parameters:**

```sql
sqlc.arg('parameter_name')
```

**Optional/nullable parameters:**

```sql
sqlc.narg('parameter_name')
```

**Array parameters:**

```sql
sqlc.narg('ids')::uuid[]
```

**Filter pattern for optional params:**

```sql
(sqlc.narg('param') IS NULL OR column = sqlc.narg('param'))
```

### Joins

**Simple join:**

```sql
-- name: GetProductWithCategory :one
SELECT
    p.*,
    c.id as category_id,
    c.name as category_name
FROM products p
INNER JOIN categories c ON p.category_id = c.id
WHERE p.id = $1;
```

**With aggregates:**

```sql
-- name: ListProductsWithVariantCount :many
SELECT
    p.*,
    COUNT(v.id) as variant_count
FROM products p
LEFT JOIN product_variants v ON p.id = v.product_id
WHERE p.deleted_at IS NULL
GROUP BY p.id
ORDER BY p.created_at DESC;
```

## Triggers

**Location:** `database/trigger.sql`

**Common triggers:**

**Update timestamp:**

```sql
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

**Maintain aggregates:**

```sql
CREATE OR REPLACE FUNCTION update_product_rating()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE products
    SET rating = (
        SELECT COALESCE(AVG(rating), 0)
        FROM reviews
        WHERE product_id = NEW.product_id
    )
    WHERE id = NEW.product_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_product_rating
    AFTER INSERT OR UPDATE ON reviews
    FOR EACH ROW
    EXECUTE FUNCTION update_product_rating();
```

## Seeding

**Development seed:** `database/seed.sql`

**Fake data seed:** `database/seed-fake.sql`

```sql
-- Seed categories
INSERT INTO categories (id, code, name, created_at, updated_at)
VALUES
    (gen_random_uuid(), 'electronics', 'Electronics', NOW(), NOW()),
    (gen_random_uuid(), 'clothing', 'Clothing', NOW(), NOW());

-- Seed products
INSERT INTO products (id, name, description, price, category_id, created_at, updated_at)
SELECT
    gen_random_uuid(),
    'Product ' || i,
    'Description for product ' || i,
    (random() * 100000)::BIGINT,
    (SELECT id FROM categories ORDER BY random() LIMIT 1),
    NOW(),
    NOW()
FROM generate_series(1, 100) i;
```

## Temporary Tables

**Location:** `database/temporary-table/*.sql`

Used for sqlc to avoid complaints about `CREATE TEMP TABLE`:

```sql
-- database/temporary-table/product.sql
CREATE TEMP TABLE temp_product_ids (
    id UUID
);
```

Not part of actual schema, only for sqlc compilation.

## Transactions

Handle at repository/application layer using pgx:

```go
tx, err := db.Begin(ctx)
if err != nil {
    return err
}
defer tx.Rollback(ctx)

queries := queries.WithTx(tx)

// Execute queries
err = queries.UpsertProduct(ctx, params)
if err != nil {
    return err
}

return tx.Commit(ctx)
```

## Rules Summary

1. ✅ Use plural table names, snake_case columns
2. ✅ UUID for all IDs
3. ✅ BIGINT for money (cents), not DECIMAL
4. ✅ TIMESTAMPTZ for all timestamps
5. ✅ Soft delete with `deleted_at`
6. ✅ Index foreign keys, search columns, sort fields
7. ✅ Migrations include Up and Down
8. ✅ sqlc queries: `:one`, `:many`, `:exec`
9. ✅ Use `sqlc.arg()` for required, `sqlc.narg()` for optional
10. ✅ Filter pattern: `(param IS NULL OR condition)`
11. ✅ Always include `deleted_at` filtering
12. ✅ Use triggers for automated fields
13. ✅ Upsert with `ON CONFLICT DO UPDATE`
