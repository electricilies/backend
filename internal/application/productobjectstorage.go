package application

import (
	"context"

	"github.com/google/uuid"
)

type ProductObjectStorage interface {
	GetUploadURL(ctx context.Context) (*UploadImageURL, error)
	GetDeleteURL(ctx context.Context, imageID uuid.UUID) (*DeleteImageURL, error)
	PersistImageFromTemp(ctx context.Context, key string, imageID uuid.UUID) error
}
