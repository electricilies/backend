package cart

import "backend/internal/domain/pagination"

type Repository interface {
	GetCartByUser(userID string, pagination *pagination.Params) (*Model, error)
}
