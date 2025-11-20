package domain

import (
	"time"
)

type Review struct {
	ID        int        `json:"id"        binding:"required"                      validate:"required"`
	Rating    int        `json:"rating"    binding:"required"                      validate:"required,gte=1,lte=5"`
	Content   *string    `json:"content"   binding:"required"                      validate:"required,min=10"`
	OrderItem *OrderItem `json:"orderItem" binding:"omitnil"                       validate:"omitnil"`
	ImageURL  *string    `json:"imageUrl"  validate:"omitempty,url"`
	User      *User      `json:"user"      binding:"omitnil"                       validate:"omitnil"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"                      validate:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"                      validate:"required,gtefield=CreatedAt"`
	DeletedAt *time.Time `json:"deletedAt" validate:"omitempty,gtefield=CreatedAt"`
}
