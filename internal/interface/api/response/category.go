package response

import "time"

type Category struct {
	ID          int       `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" binding:"required"`
}

type CategoriesPagination struct {
	Meta Pagination `json:"meta" binding:"required"`
	Data []Category `json:"data" binding:"required"`
}
