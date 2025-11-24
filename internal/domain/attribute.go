package domain

import (
	"time"

	"github.com/google/uuid"
)

type Attribute struct {
	ID        uuid.UUID        `json:"id"        binding:"required"        validate:"required"               example:"123"`
	Code      string           `json:"code"      binding:"required"        validate:"required,gte=2,lte=50"  example:"color"`
	Name      string           `json:"name"      binding:"required"        validate:"required,gte=2,lte=100" example:"Color"`
	Values    []AttributeValue `json:"values"    validate:"omitempty,dive"`
	DeletedAt *time.Time       `json:"deletedAt"`
}

func (a *Attribute) GetValueByID(id uuid.UUID) *AttributeValue {
	if a.Values == nil {
		return nil
	}
	for _, value := range a.Values {
		if value.ID == id {
			return &value
		}
	}
	return nil
}

type AttributeValue struct {
	ID        uuid.UUID  `json:"id"        binding:"required"   validate:"required"               example:"1"`
	Value     string     `json:"value"     binding:"required"   validate:"required,gte=1,lte=100" example:"Red"`
	DeletedAt *time.Time `json:"deletedAt" validate:"omitempty"`
}
