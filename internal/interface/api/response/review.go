package response

import "time"

type Review struct {
	ID        int       `json:"id" binding:"required"`
	Rate      int       `json:"rate" binding:"required"`
	Content   string    `json:"content" binding:"omitnil"`
	ImageURL  string    `json:"image_url" binding:"omitnil"`
	User      User      `json:"user" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}
