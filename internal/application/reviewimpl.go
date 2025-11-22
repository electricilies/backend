package application

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/constant"
	"backend/internal/domain"

	"github.com/redis/go-redis/v9"
)

type ReviewImpl struct {
	reviewRepo    domain.ReviewRepository
	reviewService domain.ReviewService
	redisClient   *redis.Client
}

func ProvideReview(reviewRepo domain.ReviewRepository, reviewService domain.ReviewService, redisClient *redis.Client) *ReviewImpl {
	return &ReviewImpl{
		reviewRepo:    reviewRepo,
		reviewService: reviewService,
		redisClient:   redisClient,
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

	savedReview, err := r.reviewRepo.Save(ctx, *review)
	if err != nil {
		return nil, err
	}

	// Invalidate list cache
	if r.redisClient != nil {
		iter := r.redisClient.Scan(ctx, 0, constant.ReviewListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			r.redisClient.Del(ctx, iter.Val())
		}
	}

	return savedReview, nil
}

func (r *ReviewImpl) List(ctx context.Context, param ListReviewsParam) (*Pagination[domain.Review], error) {
	// Build cache key
	cacheKey := constant.ReviewListKey(param.OrderItemIDs, param.ProductVariantID, param.UserIDs, string(param.Deleted), *param.Limit, *param.Page)

	// Try to get from cache
	if r.redisClient != nil {
		cachedData, err := r.redisClient.Get(ctx, cacheKey).Result()
		if err == nil && cachedData != "" {
			var pagination Pagination[domain.Review]
			if err := json.Unmarshal([]byte(cachedData), &pagination); err == nil {
				return &pagination, nil
			}
		}
	}

	reviews, err := r.reviewRepo.List(
		ctx,
		param.OrderItemIDs,
		param.ProductVariantID,
		param.UserIDs,
		param.Deleted,
		*param.Limit,
		*param.Page,
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

	pagination := newPagination(*reviews, *count, *param.Page, *param.Limit)

	// Cache the result
	if r.redisClient != nil {
		if data, err := json.Marshal(pagination); err == nil {
			r.redisClient.Set(ctx, cacheKey, data, time.Duration(constant.CacheTTLReview)*time.Second)
		}
	}

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
		param.Data.Rating,
		param.Data.Content,
		param.Data.ImageURL,
	)
	if err != nil {
		return nil, err
	}

	savedReview, err := r.reviewRepo.Save(ctx, *review)
	if err != nil {
		return nil, err
	}

	// Invalidate list cache
	if r.redisClient != nil {
		iter := r.redisClient.Scan(ctx, 0, constant.ReviewListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			r.redisClient.Del(ctx, iter.Val())
		}
	}

	return savedReview, nil
}

func (r *ReviewImpl) Delete(ctx context.Context, param DeleteReviewParam) error {
	review, err := r.reviewRepo.Get(ctx, param.ReviewID)
	if err != nil {
		return err
	}

	// Mark as deleted by saving with DeletedAt set
	// This assumes the domain model handles soft delete
	_, err = r.reviewRepo.Save(ctx, *review)
	if err != nil {
		return err
	}

	// Invalidate list cache
	if r.redisClient != nil {
		iter := r.redisClient.Scan(ctx, 0, constant.ReviewListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			r.redisClient.Del(ctx, iter.Val())
		}
	}

	return nil
}
