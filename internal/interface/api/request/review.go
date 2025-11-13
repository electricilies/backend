package request

import "backend/internal/domain/review"

type CreateReview struct {
	ProductID int    `json:"productId" binding:"required"`
	UserID    string `json:"userId" binding:"required"`
	Rating    int    `json:"rating" binding:"required,min=1,max=5"`
	Content   string `json:"content,omitempty"`
	ImageURL  string `json:"imageUrl,omitempty"`
}

type UpdateReview struct {
	Rating   int    `json:"rate,omitempty"`
	Content  string `json:"content,omitempty"`
	ImageURL string `json:"imageUrl,omitempty"`
}

type ReviewQueryParams struct {
	Limit   int
	Offset  int
	Deleted string
}

func ReviewQueryParamsToDomain(reviewQueryParams *ReviewQueryParams) *review.QueryParams {
	return &review.QueryParams{
		PaginationParams: PaginationParamsToDomain(
			reviewQueryParams.Limit,
			reviewQueryParams.Offset,
		),
		Deleted: DeletedParamToDomain(reviewQueryParams.Deleted),
	}
}
