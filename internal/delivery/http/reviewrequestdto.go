package http

import (
	"backend/internal/domain"

	"github.com/google/uuid"
)

type ListReviewsRequestDto struct {
	PaginationRequestDto
	OrderItemIDs     []uuid.UUID
	ProductVariantID uuid.UUID
	UserIDs          []uuid.UUID
	Deleted          domain.DeletedParam
}

type CreateReviewRequestDto struct {
	OrderItemID uuid.UUID
	UserID      uuid.UUID
	Data        CreateReviewData
}

type CreateReviewData struct {
	Rating   int    `json:"rating"   binding:"required,gte=1,lte=5"`
	Content  string `json:"content"  binding:"omitnil"`
	ImageURL string `json:"imageUrl" binding:"omitnil,url"`
}

type UpdateReviewRequestDto struct {
	ReviewID uuid.UUID
	UserID   uuid.UUID
	Data     UpdateReviewData
}

type UpdateReviewData struct {
	Rating   int    `json:"rating,omitempty"   binding:"omitempty,gte=1,lte=5"`
	Content  string `json:"content,omitempty"`
	ImageURL string `json:"imageUrl,omitempty" binding:"omitempty,url"`
}

type GetReviewRequestDto struct {
	ReviewID uuid.UUID
}

type DeleteReviewRequestDto struct {
	ReviewID uuid.UUID
}
