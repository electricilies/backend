# Repository Layer Rules

## Overview

The repository layer implements domain repository interfaces using PostgreSQL and sqlc. Repositories are responsible for:

- Persisting domain entities
- Querying data with filters and pagination
- Mapping infrastructure errors to domain errors

## Structure

**Note:** Based on the current structure analysis, repository implementations appear to be missing from `internal/infrastructure/repository`. This document describes the expected patterns.

### Repository Organization

```
internal/infrastructure/repository/
├── postgres/              # sqlc generated code
│   ├── models.go
│   ├── product.sql.go
│   └── querier.go
└── product_repository.go  # Repository implementation
```

## Implementation Pattern

```go
package repository

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repository/postgres"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db      *pgxpool.Pool
	queries *postgres.Queries
}

func ProvideProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db:      db,
		queries: postgres.New(db),
	}
}

var _ domain.ProductRepository = &ProductRepository{}
```

## Repository Methods

### Save (Upsert)

The `Save` method uses PostgreSQL's `ON CONFLICT DO UPDATE` for upsert logic:

```go
func (r *ProductRepository) Save(ctx context.Context, entity domain.Product) error {
    err := r.queries.UpsertProduct(ctx, postgres.UpsertProductParams{
        ID:          entity.ID,
        Name:        entity.Name,
        Description: entity.Description,
        Price:       entity.Price,
        CreatedAt:   entity.CreatedAt,
        UpdatedAt:   entity.UpdatedAt,
        DeletedAt:   entity.DeletedAt,
    })

    if err != nil {
        return mapError(err)
    }

    return nil
}
```

**Conventions:**

- Use sqlc-generated `Upsert*` query
- Map domain entity to sqlc params
- Always map errors with `mapError()`
- Handle aggregate relationships (save children)

### Get

```go
func (r *ProductRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
    row, err := r.queries.GetProduct(ctx, id)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, domain.ErrNotFound
        }
        return nil, mapError(err)
    }

    entity := mapRowToProduct(row)
    return &entity, nil
}
```

**Conventions:**

- Map `pgx.ErrNoRows` to `domain.ErrNotFound`
- Use mapper function to convert sqlc types to domain entities
- Load related entities if needed (joins or separate queries)

### List

```go
func (r *ProductRepository) List(
    ctx context.Context,
    ids *[]uuid.UUID,
    search *string,
    deleted domain.DeletedParam,
    limit int,
    offset int,
) (*[]domain.Product, error) {
    rows, err := r.queries.ListProducts(ctx, postgres.ListProductsParams{
        IDs:     ids,
        Search:  search,
        Deleted: string(deleted),
        Limit:   int32(limit),
        Offset:  int32(offset),
    })

    if err != nil {
        return nil, mapError(err)
    }

    entities := make([]domain.Product, len(rows))
    for i, row := range rows {
        entities[i] = mapRowToProduct(row)
    }

    return &entities, nil
}
```

**Conventions:**

- Convert pointer parameters for sqlc (handle nullability)
- Use `int32` for limit/offset (PostgreSQL convention)
- Map all rows to domain entities
- Return pointer to slice

### Count

```go
func (r *ProductRepository) Count(
    ctx context.Context,
    ids *[]uuid.UUID,
    deleted domain.DeletedParam,
) (*int, error) {
    count, err := r.queries.CountProducts(ctx, postgres.CountProductsParams{
        IDs:     ids,
        Deleted: string(deleted),
    })

    if err != nil {
        return nil, mapError(err)
    }

    result := int(count)
    return &result, nil
}
```

## Error Mapping

Map PostgreSQL errors to domain errors:

```go
import (
    "github.com/jackc/pgerrcode"
    "github.com/jackc/pgx/v5/pgconn"
)

func mapError(err error) error {
    var pgErr *pgconn.PgError
    if errors.As(err, &pgErr) {
        switch pgErr.Code {
        case pgerrcode.UniqueViolation:
            return domain.ErrExists
        case pgerrcode.ForeignKeyViolation:
            return domain.ErrInvalid
        case pgerrcode.CheckViolation:
            return domain.ErrInvalid
        case pgerrcode.NotNullViolation:
            return domain.ErrInvalid
        default:
            return domain.ErrInternal
        }
    }

    if errors.Is(err, context.Canceled) {
        return domain.ErrTimeout
    }

    if errors.Is(err, context.DeadlineExceeded) {
        return domain.ErrTimeout
    }

    return domain.ErrInternal
}
```

## sqlc Configuration

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

## SQL Queries

Queries are defined in `database/queries/*.sql`:

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

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1;

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
LIMIT $1 OFFSET $2;

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

**Conventions:**

- Use `sqlc.arg()` for required parameters
- Use `sqlc.narg()` for nullable parameters
- Filter patterns: `(param IS NULL OR condition)`
- Use `ILIKE` for case-insensitive search with `%` wildcards
- Implement soft delete with `deleted_at` filtering
- Use `LIMIT` and `OFFSET` for pagination

## Type Mapping

Map sqlc types to domain types:

```go
func mapRowToProduct(row postgres.Product) domain.Product {
    return domain.Product{
        ID:          row.ID,
        Name:        row.Name,
        Description: row.Description,
        Price:       row.Price,
        CreatedAt:   row.CreatedAt,
        UpdatedAt:   row.UpdatedAt,
        DeletedAt:   row.DeletedAt,
    }
}

// For nullable fields
func mapNullableString(s pgtype.Text) *string {
    if !s.Valid {
        return nil
    }
    return &s.String
}

func mapNullableTime(t pgtype.Timestamp) *time.Time {
    if !t.Valid {
        return nil
    }
    return &t.Time
}
```

## Transactions

For complex operations spanning multiple repositories:

```go
import "github.com/Thiht/transactor/pgx"

func (r *ProductRepository) SaveWithVariants(
    ctx context.Context,
    product domain.Product,
    variants []domain.ProductVariant,
) error {
    tx, err := r.db.Begin(ctx)
    if err != nil {
        return mapError(err)
    }
    defer tx.Rollback(ctx)

    queries := r.queries.WithTx(tx)

    // Save product
    err = queries.UpsertProduct(ctx, /* params */)
    if err != nil {
        return mapError(err)
    }

    // Save variants
    for _, variant := range variants {
        err = queries.UpsertProductVariant(ctx, /* params */)
        if err != nil {
            return mapError(err)
        }
    }

    return tx.Commit(ctx)
}
```

## Rules Summary

1. ✅ Repository implements domain repository interface
2. ✅ Use sqlc for type-safe SQL queries
3. ✅ Implement Save as upsert (`ON CONFLICT DO UPDATE`)
4. ✅ Map `pgx.ErrNoRows` to `domain.ErrNotFound`
5. ✅ Map all PostgreSQL errors to domain errors
6. ✅ Use mapper functions for type conversion
7. ✅ Handle nullable fields with `sqlc.narg()`
8. ✅ Implement soft delete filtering
9. ✅ Use `LIMIT/OFFSET` for pagination
10. ✅ Use transactions for multi-entity operations
