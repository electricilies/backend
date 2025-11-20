package domain

type CreateReviewParam struct {
	OrderItemID int              `binding:"required"`
	UserID      int              `binding:"required"`
	Data        CreateReviewData `binding:"required"`
}

type CreateReviewData struct {
	Rating   int    `json:"rating"   binding:"required,gte=1,lte=5"`
	Content  string `json:"content"  binding:"omitnil"`
	ImageURL string `json:"imageUrl" binding:"omitnil,url"`
}

type UpdateReviewParam struct {
	ReviewID int              `binding:"required"`
	Data     UpdateReviewData `binding:"required"`
}

type UpdateReviewData struct {
	Rating   *int    `json:"rating"   binding:"omitnil,gte=1,lte=5"`
	Content  *string `json:"content"  binding:"omitnil"`
	ImageURL *string `json:"imageUrl" binding:"omitnil,url"`
}

type GetReviewParam struct {
	ReviewID int `binding:"required"`
}

type DeleteReviewParam struct {
	ReviewID int `binding:"required"`
}
