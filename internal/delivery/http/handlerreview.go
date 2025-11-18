package http

import (
	"net/http"

	_ "backend/internal/domain"
	_ "backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ReviewHandler interface {
	Get(*gin.Context)
	ListReviewsByProducts(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type GinReviewHandler struct{}

var _ ReviewHandler = &GinReviewHandler{}

func ProvideReviewHandler() *GinReviewHandler {
	return &GinReviewHandler{}
}

// GetReview godoc
//
//	@Summary		Get review by ID
//	@Description	Get review details by ID
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			review_id	path		int	true	"Review ID"
//	@Success		200			{object} domain.Review
//	@Failure		404			{object}	service.NotFoundError
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/reviews/{review_id} [get]
func (h *GinReviewHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListReviews godoc
//
//	@Summary		List all reviews
//	@Description	Get all reviews
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			product_ids	query		[]int	false	"Product IDs"				collectionFormat(csv)
//	@Param			deleted		query		string	false	"Include deleted reviews"	Enums(include, only, exclude)
//	@Param			page		query		int		false	"Page for pagination"
//	@Param			limit		query		int		false	"Limit for pagination"	default(10)
//	@Success		200			{object} domain.DataPagination{data=[]domain.Review}
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/reviews [get]
func (h *GinReviewHandler) ListReviewsByProducts(ctx *gin.Context) {
}

// CreateReview godoc
//
//	@Summary		Create a review
//	@Description	Create a new review
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			review	body service.CreateReviewParam	true	"Review request"
//	@Success		201		{object} domain.Review
//	@Failure		400		{object}	service.BadRequestError
//	@Failure		409		{object}	service.ConflictError
//	@Failure		500		{object}	service.InternalServerError
//	@Router			/reviews [post]
func (h *GinReviewHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateReview godoc
//
//	@Summary		Update a review
//	@Description	Update review by ID
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			review_id	path		int						true	"Review ID"
//	@Param			review		body service.UpdateReviewParam	true	"Update review request"
//	@Success		204			{object} domain.Review
//	@Failure		400			{object}	service.BadRequestError
//	@Failure		404			{object}	service.NotFoundError
//	@Failure		409			{object}	service.ConflictError
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/reviews/{review_id} [patch]
func (h *GinReviewHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteReview godoc
//
//	@Summary		Delete a review
//	@Description	Delete review by ID
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			review_id	path		int		true	"Review ID"
//	@Success		204
//	@Failure		404			{object}	service.NotFoundError
//	@Failure		500			{object}	service.InternalServerError
//	@Router			/reviews/{review_id} [delete]
func (h *GinReviewHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
