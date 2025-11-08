package product

import (
	"backend/internal/domain/product"
	"backend/internal/infrastructure/presistence/postgres"
)

func ToGetProductImageByIDParams(id int) *postgres.GetProductImageByIDParams {
	return &postgres.GetProductImageByIDParams{
		ID: int32(id),
	}
}

type UploadURLImage struct {
	URL string
	Key string
}

func (p *UploadURLImage) ToDomain() *product.UploadImageURLModel {
	return &product.UploadImageURLModel{
		URL: p.URL,
		Key: p.Key,
	}
}
