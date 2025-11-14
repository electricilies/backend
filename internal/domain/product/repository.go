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
	List(ctx context.Context, queryParams *QueryParams) (*PaginationModel, error)
	Create(ctx context.Context, product *Model) (*Model, error)
	Update(ctx context.Context, product *Model, id int) (*Model, error)
	Delete(ctx context.Context, id int) error
	AddOption(ctx context.Context, option  *OptionModel,id int) (*OptionModel, error)
	UpdateOption(ctx context.Context, option  *OptionModel,optionId int) (*OptionModel, error)
	AddVariant(ctx context.Context, variant  *VariantModel,id int) (*VariantModel, error)
	UpdateVariant(ctx context.Context, variant  *VariantModel,variantId int) (*VariantModel, error)
	AddImages(ctx context.Context, images  []*ImageModel,id int) error
}
