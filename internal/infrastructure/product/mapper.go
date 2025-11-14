package product

import (
	"backend/internal/domain/product"
)

func ToGetProductImageByIDParams(id int) any {
	return &struct {
		ID int32 `json:"id"`
	}{
		ID: int32(id),
	}
}

type UploadURLImage struct {
	URL *string
	Key *string
}

func (p *UploadURLImage) ToDomain() *product.UploadImageURLModel {
	return &product.UploadImageURLModel{
		URL: p.URL,
		Key: p.Key,
	}
}
