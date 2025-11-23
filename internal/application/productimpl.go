package application

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/constant"
	"backend/internal/domain"
	"github.com/redis/go-redis/v9"
)

type ProductImpl struct {
	productRepo    domain.ProductRepository
	productService domain.ProductService
	redisClient    *redis.Client
}

func ProvideProduct(productRepo domain.ProductRepository, productService domain.ProductService, redisClient *redis.Client) *ProductImpl {
	return &ProductImpl{
		productRepo:    productRepo,
		productService: productService,
		redisClient:    redisClient,
	}
}

func (p *ProductImpl) List(ctx context.Context, param ListProductParam) (*Pagination[domain.Product], error) {
	cacheKey := constant.ProductListKey(
		param.ProductIDs,
		param.Search,
		param.MinPrice,
		param.MaxPrice,
		param.Rating,
		param.CategoryIDs,
		param.Deleted,
		param.SortRating,
		param.SortPrice,
		param.Limit,
		param.Page,
	)

	// Try to get from cache
	if p.redisClient != nil {
		cachedData, err := p.redisClient.Get(ctx, cacheKey).Result()
		if err == nil && cachedData != "" {
			var pagination Pagination[domain.Product]
			if err := json.Unmarshal([]byte(cachedData), &pagination); err == nil {
				return &pagination, nil
			}
		}
	}

	products, err := p.productRepo.List(
		ctx,
		param.ProductIDs,
		param.Search,
		param.MinPrice,
		param.MaxPrice,
		param.Rating,
		param.CategoryIDs,
		param.Deleted,
		param.SortRating,
		param.SortPrice,
		param.Limit,
		param.Page*param.Limit,
	)
	if err != nil {
		return nil, err
	}

	count, err := p.productRepo.Count(
		ctx,
		param.ProductIDs,
		param.MinPrice,
		param.MaxPrice,
		param.Rating,
		param.CategoryIDs,
		param.Deleted,
	)
	if err != nil {
		return nil, err
	}

	pagination := newPagination(
		*products,
		*count,
		param.Page,
		param.Limit,
	)

	// Cache the result
	if p.redisClient != nil {
		if data, err := json.Marshal(pagination); err == nil {
			p.redisClient.Set(ctx, cacheKey, data, time.Duration(constant.CacheTTLProduct)*time.Second)
		}
	}

	return pagination, nil
}

func (p *ProductImpl) Get(ctx context.Context, param GetProductParam) (*domain.Product, error) {
	// Build cache key
	cacheKey := constant.ProductGetKey(param.ProductID)

	// Try to get from cache
	if p.redisClient != nil {
		cachedData, err := p.redisClient.Get(ctx, cacheKey).Result()
		if err == nil && cachedData != "" {
			var product domain.Product
			if err := json.Unmarshal([]byte(cachedData), &product); err == nil {
				return &product, nil
			}
		}
	}

	product, err := p.productRepo.Get(ctx, param.ProductID)
	if err != nil {
		return nil, err
	}

	// Cache the result
	if p.redisClient != nil {
		if data, err := json.Marshal(product); err == nil {
			p.redisClient.Set(ctx, cacheKey, data, time.Duration(constant.CacheTTLProduct)*time.Second)
		}
	}

	return product, nil
}

func (p *ProductImpl) Create(ctx context.Context, param CreateProductParam) (*domain.Product, error) {
	// TODO: Implement full create logic with options, variants, images, etc.
	// This is a placeholder implementation
	return nil, domain.ErrNotImplemented
}

func (p *ProductImpl) Update(ctx context.Context, param UpdateProductParam) (*Product, error) {
	// TODO: Implement update logic
	// Invalidate cache after update
	if p.redisClient != nil {
		p.redisClient.Del(ctx, constant.ProductGetKey(param.ProductID))
		iter := p.redisClient.Scan(ctx, 0, constant.ProductListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
	}
	return nil, domain.ErrNotImplemented
}

func (p *ProductImpl) Delete(ctx context.Context, param DeleteProductParam) error {
	// TODO: Implement delete logic (soft delete)
	// Invalidate cache after delete
	if p.redisClient != nil {
		p.redisClient.Del(ctx, constant.ProductGetKey(param.ProductID))
		iter := p.redisClient.Scan(ctx, 0, constant.ProductListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
	}
	return domain.ErrNotImplemented
}

func (p *ProductImpl) AddVariants(ctx context.Context, param AddProductVariantsParam) (*domain.ProductVariant, error) {
	// TODO: Implement add variants logic
	// Invalidate cache after adding variants
	if p.redisClient != nil {
		p.redisClient.Del(ctx, constant.ProductGetKey(param.ProductID))
		iter := p.redisClient.Scan(ctx, 0, constant.ProductListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
	}
	return nil, domain.ErrNotImplemented
}

func (p *ProductImpl) UpdateVariant(ctx context.Context, param UpdateProductVariantParam) (*domain.ProductVariant, error) {
	// TODO: Implement update variant logic
	// Invalidate cache after update
	if p.redisClient != nil {
		// Note: We'd need product ID to invalidate specific get cache
		// For now, just invalidate all list caches
		iter := p.redisClient.Scan(ctx, 0, constant.ProductListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
	}
	return nil, domain.ErrNotImplemented
}

func (p *ProductImpl) AddImages(ctx context.Context, param AddProductImagesParam) (*[]domain.ProductImage, error) {
	// TODO: Implement add images logic
	// Invalidate cache after adding images
	if p.redisClient != nil {
		// Invalidate product caches
		iter := p.redisClient.Scan(ctx, 0, constant.ProductGetPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
		iter = p.redisClient.Scan(ctx, 0, constant.ProductListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
	}
	return nil, domain.ErrNotImplemented
}

func (p *ProductImpl) DeleteImages(ctx context.Context, param DeleteProductImagesParam) error {
	// TODO: Implement delete images logic
	// Invalidate cache after deleting images
	if p.redisClient != nil {
		iter := p.redisClient.Scan(ctx, 0, constant.ProductGetPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
		iter = p.redisClient.Scan(ctx, 0, constant.ProductListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
	}
	return domain.ErrNotImplemented
}

func (p *ProductImpl) UpdateOptions(ctx context.Context, param UpdateProductOptionsParam) error {
	// TODO: Implement update options logic
	// Invalidate cache after update
	if p.redisClient != nil {
		iter := p.redisClient.Scan(ctx, 0, constant.ProductGetPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
		iter = p.redisClient.Scan(ctx, 0, constant.ProductListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
	}
	return domain.ErrNotImplemented
}

func (p *ProductImpl) UpdateOptionValues(ctx context.Context, param UpdateProductOptionValuesParam) error {
	// TODO: Implement update option values logic
	// Invalidate cache after update
	if p.redisClient != nil {
		iter := p.redisClient.Scan(ctx, 0, constant.ProductGetPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
		iter = p.redisClient.Scan(ctx, 0, constant.ProductListPrefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			p.redisClient.Del(ctx, iter.Val())
		}
	}
	return domain.ErrNotImplemented
}

func (p *ProductImpl) GetUploadImageURL(ctx context.Context) (*UploadImageURL, error) {
	// TODO: Implement get upload image URL logic
	return nil, domain.ErrNotImplemented
}

func (p *ProductImpl) GetDeleteImageURL(ctx context.Context, imageID int) (*DeleteImageURL, error) {
	// TODO: Implement get delete image URL logic
	return nil, domain.ErrNotImplemented
}
