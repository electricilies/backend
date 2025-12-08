package application

import (
	"context"

	"backend/internal/delivery/http"

	"github.com/google/uuid"
)

type ProductObjectStorage interface {
	GetUploadImageURL(ctx context.Context) (*http.UploadImageURLResponseDto, error)
	GetDeleteImageURL(ctx context.Context, imageID uuid.UUID) (*http.DeleteImageURLResponseDto, error)
	PersistImageFromTemp(ctx context.Context, key string) error
}
