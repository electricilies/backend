package cart

import (
	"backend/internal/domain/cart"
	"backend/internal/helper"
	"backend/internal/infrastructure/persistence/postgres"
)

func ToDomain(cartEntity *postgres.Cart) *cart.Model {
	return &cart.Model{
		ID: helper.ToPtr(int(cartEntity.ID)),
	}
}
