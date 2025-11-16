package cart

import (
	"backend/internal/domain/cart"
	"backend/internal/infrastructure/persistence/postgres"
)

func ToDomain(cartEntity postgres.Cart) *cart.Model {
	return &cart.Model{
		ID: int(cartEntity.ID),
	}
}
