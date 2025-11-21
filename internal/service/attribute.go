package service

import (
	"backend/internal/domain"
)

type Attribute struct{}

func ProvideAttribute() *Attribute {
	return &Attribute{}
}

var _ domain.AttributeService = &Attribute{}
