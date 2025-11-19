package application

import (
	"context"

	"backend/internal/domain"
)

type Product interface {
	List(context.Context, ListProductParam) (*Pagination[domain.Product], error)
	GetDeleteImageURL(context.Context, int) (*DeleteImageURL, error)
	GetUploadImageURL(context.Context) (*UploadImageURL, error)
}
