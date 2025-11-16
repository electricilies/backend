package domain

import "context"

type AttributeRepository interface {
	CreateAttribute(context.Context, CreateAttribute) (*Attribute, error)
	CreateAttributeValues(context.Context, CreateAttributeValues) (*[]AttributeValue, error)
	ListAttributes(context.Context, ListAttributes) (*[]Attribute, error)
	GetAttribute(context.Context, GetAttribute) (*Attribute, error)
	UpdateAttribute(context.Context, UpdateAttribute) (*Attribute, error)
	UpdateAttributeValues(context.Context, UpdateAttributeValues) (*[]AttributeValue, error)
	Delete(context.Context, int) error
}
