package domain

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID          uuid.UUID `json:"id"          binding:"required"                      validate:"required"`
	Rating      int       `json:"rating"      binding:"required"                      validate:"required,gte=1,lte=5"`
	Content     string    `json:"content"     validate:"omitempty,gte=10"`
	OrderID     uuid.UUID `json:"orderId"     binding:"required"                      validate:"required"`
	OrderItemID uuid.UUID `json:"orderItemId" binding:"required"                      validate:"required"`
	ImageURL    string    `json:"imageUrl"    validate:"omitempty,url"`
	CreatedAt   time.Time `json:"createdAt"   binding:"required"                      validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt"   binding:"required"                      validate:"required,gtefield=CreatedAt"`
	DeletedAt   time.Time `json:"deletedAt"   validate:"omitempty,gtefield=CreatedAt"`
}

func NewReview(
	rating int,
	content string,
	orderID uuid.UUID,
	orderItemID uuid.UUID,
	imageURL string,
) (Review, error) {
	now := time.Now()
	id, err := uuid.NewV7()
	if err != nil {
		return Review{}, err
	}
	return Review{
		ID:          id,
		Rating:      rating,
		Content:     content,
		OrderID:     orderID,
		OrderItemID: orderItemID,
		ImageURL:    imageURL,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (r *Review) Update(
	rating int,
	content string,
	imageURL string,
) {
	updated := false
	if rating != 0 && r.Rating != rating {
		r.Rating = rating
		updated = true
	}
	if content != "" && r.Content != content {
		r.Content = content
		updated = true
	}
	if imageURL != "" && r.ImageURL != imageURL {
		r.ImageURL = imageURL
		updated = true
	}
	if updated {
		r.UpdatedAt = time.Now()
	}
}
