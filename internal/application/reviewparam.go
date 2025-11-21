package application

import (
	"backend/internal/domain"

	"github.com/google/uuid"
)

type ListReviewsParam struct {
	PaginationParam
	OrderItemIDs     *[]uuid.UUID        `binding:"omitnil"`
	ProductVariantID *uuid.UUID          `binding:"omitnil"`
	UserIDs          *[]uuid.UUID        `binding:"omitnil"`
	Deleted          domain.DeletedParam `binding:"required,oneof=exclude only all"`
}

type CreateReviewParam struct {
	OrderItemID uuid.UUID        `binding:"required"`
	UserID      uuid.UUID        `binding:"required"`
	Data        CreateReviewData `binding:"required"`
}

type CreateReviewData struct {
	Rating   int    `json:"rating"   binding:"required,gte=1,lte=5"`
	Content  string `json:"content"  binding:"omitnil"`
	ImageURL string `json:"imageUrl" binding:"omitnil,url"`
}

type UpdateReviewParam struct {
	ReviewID uuid.UUID        `binding:"required"`
	Data     UpdateReviewData `binding:"required"`
}

type UpdateReviewData struct {
	Rating   *int    `json:"rating"   binding:"omitnil,gte=1,lte=5"`
	Content  *string `json:"content"  binding:"omitnil"`
	ImageURL *string `json:"imageUrl" binding:"omitnil,url"`
}

type GetReviewParam struct {
	ReviewID uuid.UUID `binding:"required"`
}

type DeleteReviewParam struct {
	ReviewID uuid.UUID `binding:"required"`
}
