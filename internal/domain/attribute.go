package domain

import (
	"time"

	"github.com/google/uuid"
)

type Attribute struct {
	ID        uuid.UUID        `json:"id"        binding:"required"   validate:"required"               example:"123"`
	Code      string           `json:"code"      binding:"required"   validate:"required,gte=2,lte=50"  example:"color"`
	Name      string           `json:"name"      binding:"required"   validate:"required,gte=2,lte=100" example:"Color"`
	Values    []AttributeValue `json:"values"    binding:"required"   validate:"required,dive"`
	DeletedAt *time.Time       `json:"deletedAt" validate:"omitempty"`
}

func (a *Attribute) AddValues(values []AttributeValue) {
	a.Values = append(a.Values, values...)
}

type AttributeValue struct {
	ID        uuid.UUID  `json:"id"        binding:"required" validate:"required"               example:"1"`
	Value     string     `json:"value"     binding:"required" validate:"required,gte=1,lte=100" example:"Red"`
	Attribute *Attribute `json:"attribute" validate:"omitnil"`
}
