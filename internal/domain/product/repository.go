package product

import (
	"context"
)

type Repository interface {
	// Get(ctx context.Context, id string) (*User, error)
	// List(ctx context.Context) ([]*User, error)
	// Create(ctx context.Context, user *User) (*User, error)
	// Update(ctx context.Context, user *User) error
	// Delete(ctx context.Context, id string) error
	GetUploadImageURL(ctx context.Context) (*UploadImageURLModel, error)
	GetDeleteImageURL(ctx context.Context, id int) (string, error)
	MoveImage(ctx context.Context, key string) error
	ListProducts(ctx context.Context, queryParams QueryParams) ([]Model, error)
}
