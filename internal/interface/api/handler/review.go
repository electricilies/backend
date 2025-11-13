package handler

import (
	"net/http"
	"strconv"

	"backend/internal/application"
	"backend/internal/domain/review"
	"backend/internal/interface/api/request"
	"backend/internal/interface/api/response"

	"github.com/gin-gonic/gin"
)

type Review interface {
	Get(ctx *gin.Context)
	ListReviewsByProducts(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type reviewHandler struct {
	app application.Review
}

func NewReview(app application.Review) Review {
	return &reviewHandler{
		app: app,
	}
}

// GetReview godoc
//
//	@Summary		Get review by ID
//	@Description	Get review details by ID
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			review_id	path		int	true	"Review ID"
//	@Success		200			{object}	response.Review
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/reviews/{review_id} [get]
func (h *reviewHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListReviews godoc
//
//	@Summary		List all reviews
//	@Description	Get all reviews
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			product_ids	query		[]int	false	"Product IDs"	collectionFormat(csv)
//	@Param			offset		query		int		false	"Offset for pagination"
//	@Param			limit		query		int		false	"Limit for pagination"
//	@Success		200			{object}	response.DataPagination{data=[]response.Review}
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/reviews [get]
func (h *reviewHandler) ListReviewsByProducts(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset")) // TODO: check, now it not required
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	productID, _ := strconv.Atoi(ctx.Query("product_id"))
	pagination, err := h.app.ListReviewsByProductID(ctx, productID, &review.QueryParams{PaginationParams: *request.PaginationToDomain(offset, limit)})
	if err != nil {
		response.ErrorFromDomain(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response.DataPaginationFromDomain(pagination.Reviews, pagination.Metadata))
}

// CreateReview godoc
//
//	@Summary		Create a review
//	@Description	Create a new review
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			review	body		request.CreateReview	true	"Review request"
//	@Success		201		{object}	response.Review
//	@Failure		400		{object}	response.BadRequestError
//	@Failure		409		{object}	response.ConflictError
//	@Failure		500		{object}	response.InternalServerError
//	@Router			/reviews [post]
func (h *reviewHandler) Create(ctx *gin.Context) {
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
//	@Param			review		body		request.UpdateReview	true	"Update review request"
//	@Success		204			{string}	string					"no content"
//	@Failure		400			{object}	response.BadRequestError
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		409			{object}	response.ConflictError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/reviews/{review_id} [put]
func (h *reviewHandler) Update(ctx *gin.Context) {
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
//	@Success		204			{string}	string	"no content"
//	@Failure		404			{object}	response.NotFoundError
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/reviews/{review_id} [delete]
func (h *reviewHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
