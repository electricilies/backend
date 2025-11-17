package service

import (
	"context"

	"backend/internal/domain"
)

type GetAttributeParam struct {
	AttributeID int `json:"attributeId" binding:"required"`
}

type ListAttributesParam struct {
	Limit       int    `json:"limit" binding:"required"`
	Offset      int    `json:"offset" binding:"required"`
	AttributeID int    `json:"attributeId"`
	Search      string `json:"search"`
	Deleted     string `json:"deleted"`
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
	GetAttribute(context.Context, GetAttributeParam) (*domain.Attribute, error)
	ListAttributes(context.Context, ListAttributesParam) (*domain.DataPagination, error)
	CreateAttribute(context.Context, CreateAttributeParam) (*domain.Attribute, error)
	UpdateAttribute(context.Context, UpdateAttributeParam) (*domain.Attribute, error)
	DeleteAttribute(context.Context, DeleteAttributeParam) error
	UpdateAttributeValues(context.Context, []UpdateAttributeValueParam) error
}

type AttributeImpl struct{}

func ProvideAttribute() *AttributeImpl {
	return &AttributeImpl{}
}

var _ Attribute = &AttributeImpl{}

func (s *AttributeImpl) GetAttribute(ctx context.Context, param GetAttributeParam) (*domain.Attribute, error) {
	return nil, nil
}

func (s *AttributeImpl) ListAttributes(ctx context.Context, param ListAttributesParam) (*domain.DataPagination, error) {
	return nil, nil
}

func (s *AttributeImpl) CreateAttribute(ctx context.Context, param CreateAttributeParam) (*domain.Attribute, error) {
	return nil, nil
}

func (s *AttributeImpl) UpdateAttribute(ctx context.Context, param UpdateAttributeParam) (*domain.Attribute, error) {
	return nil, nil
}

func (s *AttributeImpl) DeleteAttribute(ctx context.Context, param DeleteAttributeParam) error {
	return nil
}

func (s *AttributeImpl) UpdateAttributeValues(ctx context.Context, param []UpdateAttributeValueParam) error {
	return nil
}
