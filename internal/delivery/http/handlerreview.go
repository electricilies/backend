package http

import (
	"net/http"

	"backend/internal/application"
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReviewHandler interface {
	Get(*gin.Context)
	List(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type GinReviewHandler struct {
	reviewApp           application.Review
	ErrRequiredReviewID string
	ErrInvalidReviewID  string
}

var _ ReviewHandler = &GinReviewHandler{}

func ProvideReviewHandler(reviewApp application.Review) *GinReviewHandler {
	return &GinReviewHandler{
		reviewApp:           reviewApp,
		ErrRequiredReviewID: "review_id is required",
		ErrInvalidReviewID:  "invalid review_id",
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
//	@Success		200			{object}	domain.Review
//	@Failure		404			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/reviews/{review_id} [get]
func (h *GinReviewHandler) Get(ctx *gin.Context) {
	reviewIDString := ctx.Param("review_id")
	if reviewIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredReviewID))
		return
	}
	reviewID, err := uuid.Parse(reviewIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidReviewID))
		return
	}
	review, err := h.reviewApp.Get(ctx, application.GetReviewParam{
		ReviewID: reviewID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, review)
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
//	@Param			page		query		int		false	"Page for pagination"		default(1)
//	@Param			limit		query		int		false	"Limit for pagination"		default(20)
//	@Success		200			{object}	application.Pagination[domain.Review]
//	@Failure		500			{object}	Error
//	@Router			/reviews [get]
func (h *GinReviewHandler) List(ctx *gin.Context) {
	paginateParam, err := createPaginationParamsFromQuery(ctx)
	if err != nil {
		SendError(ctx, err)
		return
	}

	var orderItemIDs *[]uuid.UUID
	if orderItemIDsQuery, ok := queryArrayToUUIDSlice(ctx, "order_item_ids"); ok {
		orderItemIDs = orderItemIDsQuery
	}

	var productVariantID *uuid.UUID
	if productVariantIDQuery, ok := ctx.GetQuery("product_variant_id"); ok {
		parsedID, err := uuid.Parse(productVariantIDQuery)
		if err == nil {
			productVariantID = &parsedID
		}
	}

	var userIDs *[]uuid.UUID
	if userIDsQuery, ok := queryArrayToUUIDSlice(ctx, "user_ids"); ok {
		userIDs = userIDsQuery
	}

	deleted := domain.DeletedExcludeParam
	if deletedQuery, ok := ctx.GetQuery("deleted"); ok {
		deleted = domain.DeletedParam(deletedQuery)
	}

	reviews, err := h.reviewApp.List(ctx, application.ListReviewsParam{
		PaginationParam:  *paginateParam,
		OrderItemIDs:     orderItemIDs,
		ProductVariantID: productVariantID,
		UserIDs:          userIDs,
		Deleted:          deleted,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, reviews)
}

// CreateReview godoc
//
//	@Summary		Create a review
//	@Description	Create a new review
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			review	body		application.CreateReviewData	true	"Review request"
//	@Success		201		{object}	domain.Review
//	@Failure		400		{object}	Error
//	@Failure		409		{object}	Error
//	@Failure		500		{object}	Error
//	@Router			/reviews [post]
func (h *GinReviewHandler) Create(ctx *gin.Context) {
	var data application.CreateReviewData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	// TODO: Get orderItemID and userID from request/context
	orderItemID := uuid.New() // Placeholder
	userID := uuid.New()      // Placeholder

	review, err := h.reviewApp.Create(ctx, application.CreateReviewParam{
		OrderItemID: orderItemID,
		UserID:      userID,
		Data:        data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, review)
}

// UpdateReview godoc
//
//	@Summary		Update a review
//	@Description	Update review by ID
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			review_id	path		int								true	"Review ID"
//	@Param			review		body		application.UpdateReviewData	true	"Update review request"
//	@Success		204			{object}	domain.Review
//	@Failure		400			{object}	Error
//	@Failure		404			{object}	Error
//	@Failure		409			{object}	Error
//	@Failure		500			{object}	Error
//	@Router			/reviews/{review_id} [patch]
func (h *GinReviewHandler) Update(ctx *gin.Context) {
	reviewIDString := ctx.Param("review_id")
	if reviewIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredReviewID))
		return
	}
	reviewID, err := uuid.Parse(reviewIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidReviewID))
		return
	}

	var data application.UpdateReviewData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(err.Error()))
		return
	}

	review, err := h.reviewApp.Update(ctx, application.UpdateReviewParam{
		ReviewID: reviewID,
		Data:     data,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, review)
}

// DeleteReview godoc
//
//	@Summary		Delete a review
//	@Description	Delete review by ID
//	@Tags			Review
//	@Accept			json
//	@Produce		json
//	@Param			review_id	path	int	true	"Review ID"
//	@Success		204
//	@Failure		404	{object}	Error
//	@Failure		500	{object}	Error
//	@Router			/reviews/{review_id} [delete]
func (h *GinReviewHandler) Delete(ctx *gin.Context) {
	reviewIDString := ctx.Param("review_id")
	if reviewIDString == "" {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrRequiredReviewID))
		return
	}
	reviewID, err := uuid.Parse(reviewIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewError(h.ErrInvalidReviewID))
		return
	}

	err = h.reviewApp.Delete(ctx, application.DeleteReviewParam{
		ReviewID: reviewID,
	})
	if err != nil {
		SendError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}
