package service

import "backend/internal/domain"

type CreateReviewParam struct {
	OrderItemID int              `binding:"required"`
	UserID      int              `binding:"required"`
	Data        CreateReviewData `binding:"required"`
}

type CreateReviewData struct {
	Rating   int    `json:"rating" binding:"required,min=1,max=5"`
	Content  string `json:"content" binding:"omitnil"`
	ImageURL string `json:"imageUrl" binding:"omitnil,url"`
}

type UpdateReviewParam struct {
	ReviewID int              `binding:"required"`
	Data     UpdateReviewData `binding:"required"`
}

type UpdateReviewData struct {
	Rating   *int    `json:"rating" binding:"omitnil,min=1,max=5"`
	Content  *string `json:"content" binding:"omitnil"`
	ImageURL *string `json:"imageUrl" binding:"omitnil,url"`
}

type ListReviewsParam struct {
	PaginationParam
	OrderItemIDs     *[]int              `binding:"omitnil"`
	ProductVariantID *int                `binding:"omitnil"`
	UserIDs          *[]int              `binding:"omitnil"`
	Deleted          domain.DeletedParam `binding:"required,oneof=exclude only all"`
}

type GetReviewParam struct {
	ReviewID int `binding:"required"`
}

type DeleteReviewParam struct {
	ReviewID int `binding:"required"`
}
