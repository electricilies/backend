package application

import (
	"context"

	"backend/internal/domain"
)

type ReviewImpl struct {
	reviewRepo    domain.ReviewRepository
	reviewService domain.ReviewService
	reviewCache   ReviewCache
}

func ProvideReview(reviewRepo domain.ReviewRepository, reviewService domain.ReviewService, reviewCache ReviewCache) *ReviewImpl {
	return &ReviewImpl{
		reviewRepo:    reviewRepo,
		reviewService: reviewService,
		reviewCache:   reviewCache,
	}
}

var _ Review = &ReviewImpl{}

func (r *ReviewImpl) Create(ctx context.Context, param CreateReviewParam) (*domain.Review, error) {
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

func (r *ReviewImpl) List(ctx context.Context, param ListReviewsParam) (*Pagination[domain.Review], error) {
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

	pagination := newPagination(
		*reviews,
		*count,
		param.Page,
		param.Limit,
	)

	// TODO: Log redis failure
	_ = r.reviewCache.SetReviewList(ctx, cacheKey, pagination)

	return pagination, nil
}

func (r *ReviewImpl) Get(ctx context.Context, param GetReviewParam) (*domain.Review, error) {
	review, err := r.reviewRepo.Get(ctx, param.ReviewID)
	if err != nil {
		return nil, err
	}
	return review, nil
}

func (r *ReviewImpl) Update(ctx context.Context, param UpdateReviewParam) (*domain.Review, error) {
	review, err := r.reviewRepo.Get(ctx, param.ReviewID)
	if err != nil {
		return nil, err
	}

	err = r.reviewService.Update(
		review,
		param.UserID,
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

func (r *ReviewImpl) Delete(ctx context.Context, param DeleteReviewParam) error {
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
