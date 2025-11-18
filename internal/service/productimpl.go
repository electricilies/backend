package service

import (
	"context"

	"backend/internal/domain"
)

type ProductImpl struct{}

func ProvideProduct() *ProductImpl {
	return &ProductImpl{}
}

var _ Product = &ProductImpl{}
