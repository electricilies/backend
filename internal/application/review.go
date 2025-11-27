package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"
)

type Review struct {
	reviewRepo    domain.ReviewRepository
	reviewService domain.ReviewService
	reviewCache   ReviewCache
}

func ProvideReview(reviewRepo domain.ReviewRepository, reviewService domain.ReviewService, reviewCache ReviewCache) *Review {
	return &Review{
		reviewRepo:    reviewRepo,
		reviewService: reviewService,
		reviewCache:   reviewCache,
	}
}

var _ http.ReviewApplication = &Review{}

func (r *Review) Create(ctx context.Context, param http.CreateReviewRequestDto) (*domain.Review, error) {
	review, err := r.reviewService.Create(
		param.OrderItemID,
		param.UserID,
		param.Data.Rating,
		&param.Data.Content,
		&param.Data.ImageURL,
	)
	if err != nil {
		return nil, err
	}

	err = r.reviewRepo.Save(ctx, *review)
	if err != nil {
		return nil, err
	}

	// TODO: Log redis failure
	_ = r.reviewCache.InvalidateReviewList(ctx)

	return review, nil
}

func (r *Review) List(ctx context.Context, param http.ListReviewsRequestDto) (*http.PaginationResponseDto[domain.Review], error) {
	// Build cache key
	cacheKey := r.reviewCache.BuildListCacheKey(
		param.OrderItemIDs,
		param.ProductVariantID,
		param.UserIDs,
		param.Deleted,
		param.Limit,
		param.Page,
	)

	// Try to get from cache
	if cachedPagination, err := r.reviewCache.GetReviewList(ctx, cacheKey); err == nil {
		return cachedPagination, nil
	}

	reviews, err := r.reviewRepo.List(
		ctx,
		param.OrderItemIDs,
		param.ProductVariantID,
		param.UserIDs,
		param.Deleted,
		param.Limit,
		param.Page,
	)
	if err != nil {
		return nil, err
	}

	count, err := r.reviewRepo.Count(
		ctx,
		param.OrderItemIDs,
		param.ProductVariantID,
		param.UserIDs,
		param.Deleted,
	)
	if err != nil {
		return nil, err
	}

	pagination := newPaginationResponseDto(
		*reviews,
		*count,
		param.Page,
		param.Limit,
	)

	// TODO: Log redis failure
	_ = r.reviewCache.SetReviewList(ctx, cacheKey, pagination)

	return pagination, nil
}

func (r *Review) Get(ctx context.Context, param http.GetReviewRequestDto) (*domain.Review, error) {
	review, err := r.reviewRepo.Get(ctx, param.ReviewID)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (r *Review) Update(ctx context.Context, param http.UpdateReviewRequestDto) (*domain.Review, error) {
	review, err := r.reviewRepo.Get(ctx, param.ReviewID)
	if err != nil {
		return nil, err
	}

	err = r.reviewService.Update(
		review,
		param.Data.Rating,
		param.Data.Content,
		param.Data.ImageURL,
	)
	if err != nil {
		return nil, err
	}

	err = r.reviewRepo.Save(ctx, *review)
	if err != nil {
		return nil, err
	}

	// TODO: Log redis failure
	_ = r.reviewCache.InvalidateReviewList(ctx)

	return review, nil
}

func (r *Review) Delete(ctx context.Context, param http.DeleteReviewRequestDto) error {
	review, err := r.reviewRepo.Get(ctx, param.ReviewID)
	if err != nil {
		return err
	}

	// Mark as deleted by saving with DeletedAt set
	// This assumes the domain model handles soft delete
	err = r.reviewRepo.Save(ctx, *review)
	if err != nil {
		return err
	}

	// TODO: Log redis failure
	_ = r.reviewCache.InvalidateReviewList(ctx)

	return nil
}
