package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Cart struct {
	queries *sqlc.Queries
	conn    *pgxpool.Pool
}

var _ domain.CartRepository = (*Cart)(nil)

func ProvideCart(q *sqlc.Queries, conn *pgxpool.Pool) *Cart {
	return &Cart{
		queries: q,
		conn:    conn,
	}
}

func (r *Cart) Get(
	ctx context.Context,
	params domain.CartRepositoryGetParam,
) (*domain.Cart, error) {
	cartEntity, err := r.queries.GetCart(ctx, sqlc.GetCartParams{
		ID:     params.ID,
		UserID: params.UserID,
	})
	if err != nil {
		return nil, ToDomainError(err)
	}
	cart := &domain.Cart{
		ID:        cartEntity.ID,
		UserID:    cartEntity.UserID,
		UpdatedAt: cartEntity.UpdatedAt.Time,
	}
	cartItems, err := r.queries.ListCartItems(ctx, sqlc.ListCartItemsParams{
		CartID: params.ID,
	})
	if err != nil {
		return nil, ToDomainError(err)
	}
	productVariantIDs := make([]uuid.UUID, 0, len(cartItems))
	for _, item := range cartItems {
		productVariantIDs = append(productVariantIDs, item.ProductVariantID)
	}
	productVariantEntities, err := r.queries.ListProductVariants(ctx, sqlc.ListProductVariantsParams{
		IDs: productVariantIDs,
	})
	if err != nil {
		return nil, ToDomainError(err)
	}
	productVariantIDproductIDMap := make(map[uuid.UUID]uuid.UUID, len(productVariantEntities))
	for _, pv := range productVariantEntities {
		productVariantIDproductIDMap[pv.ID] = pv.ProductID
	}
	for _, item := range cartItems {
		cart.Items = append(cart.Items, domain.CartItem{
			ID:               item.ID,
			ProductID:        productVariantIDproductIDMap[item.ProductVariantID],
			ProductVariantID: item.ProductVariantID,
			Quantity:         int(item.Quantity),
		})
	}
	return cart, nil
}

func (r *Cart) Save(ctx context.Context, params domain.CartRepositorySaveParam) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return ToDomainError(err)
	}
	qtx := r.queries.WithTx(tx)
	defer func() { _ = tx.Rollback(ctx) }()
	err = qtx.UpsertCart(ctx, sqlc.UpsertCartParams{
		ID:     params.Cart.ID,
		UserID: params.Cart.UserID,
		UpdatedAt: pgtype.Timestamptz{
			Time:  params.Cart.UpdatedAt,
			Valid: true,
		},
	})
	if err != nil {
		return ToDomainError(err)
	}
	err = qtx.CreateTempTableCartItems(ctx)
	if err != nil {
		return ToDomainError(err)
	}
	itemParams := make([]sqlc.InsertTempTableCartItemsParams, len(params.Cart.Items))
	for i, item := range params.Cart.Items {
		itemParams[i] = sqlc.InsertTempTableCartItemsParams{
			ID:               item.ID,
			CartID:           params.Cart.ID,
			ProductVariantID: item.ProductVariantID,
			Quantity:         int32(item.Quantity),
		}
	}
	_, err = qtx.InsertTempTableCartItems(ctx, itemParams)
	if err != nil {
		return ToDomainError(err)
	}
	err = qtx.MergeCartItemsFromTemp(ctx)
	if err != nil {
		return ToDomainError(err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return ToDomainError(err)
	}
	return nil
}
