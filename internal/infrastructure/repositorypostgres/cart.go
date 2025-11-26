package repositorypostgres

import (
	"context"

	"backend/internal/domain"
	"backend/internal/helper/ptr"
	"backend/internal/infrastructure/repositorypostgres/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Cart struct {
	queries *sqlc.Queries
}

var _ domain.CartRepository = (*Cart)(nil)

func ProvideCart(q *sqlc.Queries) *Cart {
	return &Cart{queries: q}
}

func (r *Cart) Get(
	ctx context.Context,
	id *uuid.UUID,
	userID *uuid.UUID,
) (*domain.Cart, error) {
	cartPgID := pgtype.UUID{
		Bytes: ptr.Deref(id, uuid.UUID{}),
		Valid: id != nil,
	}
	userPgID := pgtype.UUID{
		Bytes: ptr.Deref(userID, uuid.UUID{}),
		Valid: userID != nil,
	}
	cartEntity, err := r.queries.GetCart(ctx, sqlc.GetCartParams{
		ID:     cartPgID,
		UserID: userPgID,
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	cart := &domain.Cart{
		ID:     cartEntity.ID,
		UserID: cartEntity.UserID,
	}
	cartItems, err := r.queries.ListCartItems(ctx, sqlc.ListCartItemsParams{
		CartID: cartPgID,
	})
	if err != nil {
		return nil, ToDomainErrorFromPostgres(err)
	}
	for _, item := range cartItems {
		cart.Items = append(cart.Items, domain.CartItem{
			ID:               item.ID,
			ProductVariantID: item.ProductVariantID,
			Quantity:         int(item.Quantity),
		})
	}
	return cart, nil
}

func (r *Cart) Save(ctx context.Context, cart domain.Cart) error {
	err := r.queries.UpsertCart(ctx, sqlc.UpsertCartParams{
		ID:     cart.ID,
		UserID: cart.UserID,
		UpdatedAt: pgtype.Timestamptz{
			Time:  cart.UpdatedAt,
			Valid: true,
		},
	})
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	err = r.queries.CreateTempTableCartItems(ctx)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	itemParams := make([]sqlc.InsertTempTableCartItemsParams, 0, len(cart.Items))
	for i, item := range cart.Items {
		itemParams[i] = sqlc.InsertTempTableCartItemsParams{
			ID:               item.ID,
			CartID:           cart.ID,
			ProductVariantID: item.ProductVariantID,
			Quantity:         int32(item.Quantity),
		}
	}
	_, err = r.queries.InsertTempTableCartItems(ctx, itemParams)
	if err != nil {
		return ToDomainErrorFromPostgres(err)
	}
	return nil
}
