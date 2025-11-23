# Repository Layer (`internal/infrastructure/repository`)

## Purpose

Implement domain repositories using PostgreSQL + sqlc

## Implementation

```go
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

## CRUD Methods

### Save (Upsert)

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
    return mapError(err)
}
```

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

```go
func mapError(err error) error {
    if err == nil {
        return nil
    }
    
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
    
    if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
        return domain.ErrTimeout
    }
    
    return domain.ErrInternal
}
```

## Type Mapping

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

**Location:** `database/queries/*.sql`

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

## Transactions

```go
func (r *ProductRepository) SaveWithRelated(
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
    err = queries.UpsertProduct(ctx, ...)
    if err != nil {
        return mapError(err)
    }
    
    // Save variants
    for _, v := range variants {
        err = queries.UpsertVariant(ctx, ...)
        if err != nil {
            return mapError(err)
        }
    }
    
    return tx.Commit(ctx)
}
```

## Quick Rules

1. ✅ Use sqlc for type-safe queries
2. ✅ Save = upsert (`ON CONFLICT DO UPDATE`)
3. ✅ Map `pgx.ErrNoRows` → `domain.ErrNotFound`
4. ✅ Map all PG errors to domain errors
5. ✅ Use `sqlc.arg()` for required params
6. ✅ Use `sqlc.narg()` for optional params
7. ✅ Filter pattern: `(param IS NULL OR condition)`
8. ✅ Use `int32` for limit/offset
9. ✅ Transactions for multi-entity ops
10. ✅ Mapper functions for type conversion
