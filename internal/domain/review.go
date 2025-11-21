package domain

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID  `json:"id"        binding:"required"                      validate:"required"`
	Rating    int        `json:"rating"    binding:"required"                      validate:"required,gte=1,lte=5"`
	Content   *string    `json:"content"   binding:"required"                      validate:"required,gte=10"`
	OrderItem OrderItem  `json:"orderItem" binding:"required"                      validate:"required"`
	ImageURL  *string    `json:"imageUrl"  validate:"omitempty,url"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"                      validate:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"                      validate:"required,gtefield=CreatedAt"`
	DeletedAt *time.Time `json:"deletedAt" validate:"omitempty,gtefield=CreatedAt"`
}
