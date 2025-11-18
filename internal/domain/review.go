package domain

import (
	"time"
)

type Review struct {
	ID        int       `json:"id" binding:"required"`
	Rating    int       `json:"rating" binding:"required"`
	Content   string    `json:"content,omitempty"`
	OrderItem OrderItem `json:"orderItem" binding:"required"`
	ImageURL  string    `json:"imageUrl,omitempty"`
	User      User      `json:"user" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}
