package service

import (
	"backend/internal/domain"
)

type Category struct{}

func ProvideCategory() *Category {
	return &Category{}
}

var _ domain.CategoryService = &Category{}
