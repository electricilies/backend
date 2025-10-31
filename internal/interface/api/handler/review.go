package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Review interface {
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type reviewHandler struct{}

func NewReview() Review { return &reviewHandler{} }

// GetReview godoc
//
// @Summary      Get review by ID
// @Description  Get review details by ID
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        id path int true "Review ID"
// @Success      200 {object} response.ReviewResponse
// @Failure      404 {object} mapper.NotFoundError
// @Failure      500 {object} mapper.InternalServerError
// @Router       /reviews/{id} [get]
func (h *reviewHandler) Get(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// ListReviews godoc
//
// @Summary      List all reviews
// @Description  Get all reviews
// @Tags         Review
// @Accept       json
// @Produce      json
// @Success      200 {object} response.ReviewListResponse
// @Failure      500 {object} mapper.InternalServerError
// @Router       /reviews [get]
func (h *reviewHandler) List(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// CreateReview godoc
//
// @Summary      Create a review
// @Description  Create a new review
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        review body request.CreateReview true "Review request"
// @Success      201 {object} response.ReviewResponse
// @Failure      400 {object} mapper.BadRequestError
// @Failure      409 {object} mapper.ConflictError
// @Failure      500 {object} mapper.InternalServerError
// @Router       /reviews [post]
func (h *reviewHandler) Create(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// UpdateReview godoc
//
// @Summary      Update a review
// @Description  Update review by ID
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        id path int true "Review ID"
// @Param        review body request.UpdateReview true "Update review request"
// @Success      204 {string} string "no content"
// @Failure      400 {object} mapper.BadRequestError
// @Failure      404 {object} mapper.NotFoundError
// @Failure      409 {object} mapper.ConflictError
// @Failure      500 {object} mapper.InternalServerError
// @Router       /reviews/{id} [put]
func (h *reviewHandler) Update(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

// DeleteReview godoc
//
// @Summary      Delete a review
// @Description  Delete review by ID
// @Tags         Review
// @Accept       json
// @Produce      json
// @Param        id path int true "Review ID"
// @Success      204 {string} string "no content"
// @Failure      404 {object} mapper.NotFoundError
// @Failure      500 {object} mapper.InternalServerError
// @Router       /reviews/{id} [delete]
func (h *reviewHandler) Delete(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
