package product

import "backend/internal/infrastructure/presistence/postgres"

func ToGetProductImageByIDParams(id int) *postgres.GetProductImageByIDParams {
	return &postgres.GetProductImageByIDParams{
		ID: int32(id),
	}
}
