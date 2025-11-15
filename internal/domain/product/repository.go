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
	GetUploadImageURL(context.Context) (*UploadImageURLModel, error)
	GetDeleteImageURL(context.Context, int) (string, error)
	MoveImage(context.Context, string) error
	List(context.Context, *QueryParams) (*PaginationModel, error)
	Create(context.Context, *Model) (*Model, error)
	Update(context.Context, *Model, int) (*Model, error)
	Delete(context.Context, int) error
	AddOption(context.Context, *OptionModel, int) (*OptionModel, error)
	UpdateOption(context.Context, *OptionModel, int) (*OptionModel, error)
	AddVariant(context.Context, *VariantModel, int) (*VariantModel, error)
	UpdateVariant(context.Context, *VariantModel, int) (*VariantModel, error)
	AddImages(context.Context, []*ImageModel, int) error
}
