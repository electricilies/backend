package response

import "time"

type Category struct {
	ID          int       `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
}
