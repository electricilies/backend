package domain

import (
	"context"
)

type AttributeService interface {
	Get(ctx context.Context, param GetAttributeParam) (*Attribute, error)
	Create(ctx context.Context, param CreateAttributeParam) (*Attribute, error)
	Update(ctx context.Context, param UpdateAttributeParam) (*Attribute, error)
	Delete(ctx context.Context, param DeleteAttributeParam) error
	CreateValue(ctx context.Context, param CreateAttributeValueParam) (*AttributeValue, error)
	UpdateValue(ctx context.Context, param UpdateAttributeValueParam) (*AttributeValue, error)
	DeleteValue(ctx context.Context, param DeleteAttributeValueParam) error
}
