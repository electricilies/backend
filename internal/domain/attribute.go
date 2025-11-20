package domain

import "time"

type Attribute struct {
	ID        int               `json:"id"        binding:"required"      validate:"required"               example:"123"`
	Code      string            `json:"code"      binding:"required"      validate:"required,gte=2,lte=50"  example:"color"`
	Name      string            `json:"name"      binding:"required"      validate:"required,gte=2,lte=100" example:"Color"`
	Values    *[]AttributeValue `json:"values"    validate:"omitnil,dive"`
	DeletedAt *time.Time        `json:"deletedAt" validate:"omitempty"`
}

type AttributeValue struct {
	ID        int        `json:"id"        binding:"required" validate:"required"               example:"1"`
	Value     string     `json:"value"     binding:"required" validate:"required,gte=1,lte=100" example:"Red"`
	Attribute *Attribute `json:"attribute" validate:"omitnil"`
}
