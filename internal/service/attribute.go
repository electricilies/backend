package service

import (
	"context"

	"backend/internal/domain"
)

type GetAttributeParam struct {
	AttributeID int `json:"attributeId" binding:"required"`
}

type ListAttributesParam struct {
	PaginationParam
	AttributeIDs *[]int  `json:"attributeId"`
	Search       *string `json:"search"`
	Deleted      string  `json:"deleted"`
}

type CreateAttributeParam struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateAttributeParam struct {
	AttributeID int    `json:"attributeId" binding:"required"`
	Name        string `json:"name" binding:"required"`
}

type DeleteAttributeParam struct {
	AttributeID int `json:"attributeId" binding:"required"`
}

type CreateAttributeValueParam struct {
	AttributeID int    `json:"attributeId" binding:"required"`
	Value       string `json:"value" binding:"required"`
}

type UpdateAttributeValueParam struct {
	AttributeValueIds int    `json:"attributeValueIds" binding:"required"`
	Values            string `json:"values" binding:"required"`
}

type Attribute interface {
	Get(context.Context, GetAttributeParam) (*domain.Attribute, error)
	List(context.Context, ListAttributesParam) (*domain.Pagination[domain.Attribute], error)
	Create(context.Context, CreateAttributeParam) (*domain.Attribute, error)
	Update(context.Context, UpdateAttributeParam) (*domain.Attribute, error)
	Delete(context.Context, DeleteAttributeParam) error
	UpdateValues(context.Context, []UpdateAttributeValueParam) error
}

type AttributeImpl struct{}

func ProvideAttribute() *AttributeImpl {
	return &AttributeImpl{}
}

var _ Attribute = &AttributeImpl{}

func (s *AttributeImpl) Get(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error) {
	return nil, nil
}

func (s *AttributeImpl) List(ctx context.Context, param ListAttributesParam) (*domain.Pagination[domain.Attribute], error) {
	return nil, nil
}

func (s *AttributeImpl) Create(ctx context.Context, param CreateAttributeParam) (*domain.Attribute, error) {
	return nil, nil
}

func (s *AttributeImpl) Update(ctx context.Context, param UpdateAttributeParam) (*domain.Attribute, error) {
	return nil, nil
}

func (s *AttributeImpl) Delete(ctx context.Context, param DeleteAttributeParam) error {
	return nil
}

func (s *AttributeImpl) UpdateValues(ctx context.Context, param []UpdateAttributeValueParam) error {
	return nil
}
