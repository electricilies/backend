package domain

import (
	"time"
)

type Category struct {
	ID        int       `json:"id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
}
