package domain

import (
	"time"
)

type Review struct {
	ID        int        `json:"id" binding:"required"`
	Rating    int        `json:"rating" binding:"required"`
	Content   *string    `json:"content" binding:"required"`
	OrderItem *OrderItem `json:"orderItem" binding:"omitnil"`
	ImageURL  *string    `json:"imageUrl" binding:"required,url"`
	User      *User      `json:"user" binding:"omitnil"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt" binding:"required"`
}
