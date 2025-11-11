package response

import (
	"time"

	"backend/internal/domain/review"
)

type Review struct {
	ID        int       `json:"id" binding:"required"`
	Rate      int       `json:"rate" binding:"required"`
	Content   string    `json:"content" binding:"omitnil"`
	ImageURL  string    `json:"imageUrl" binding:"omitnil"`
	User      *User     `json:"user" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required"`
}

func ReviewFromDomain(r *review.Review) *Review {
	return &Review{
		ID:        r.ID,
		Rate:      r.Rate,
		Content:   r.Content,
		ImageURL:  r.ImageURL,
		User:      UserFromDomain(r.User),
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}
