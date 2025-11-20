package domain

import (
	"time"
)

type Category struct {
	ID        int        `json:"id"        binding:"required"                      validate:"required"`
	Name      string     `json:"name"      binding:"required"                      validate:"required,gte=2,lte=100"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"                      validate:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"                      validate:"required,gtefield=CreatedAt"`
	DeletedAt *time.Time `json:"deletedAt" validate:"omitempty,gtefield=CreatedAt"`
}
