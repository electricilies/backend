package http

import (
	"backend/internal/domain"
	"github.com/google/uuid"
)

type ListAttributesRequestDto struct {
	PaginationRequestDto
	AttributeIDs []uuid.UUID
	ProductIDs   []uuid.UUID
	Deleted      domain.DeletedParam
	Search       string
}

type ListAttributeValuesRequestDto struct {
	PaginationRequestDto
	AttributeID       uuid.UUID
	AttributeValueIDs []uuid.UUID
	Deleted           domain.DeletedParam
	Search            string
}

type GetAttributeRequestDto struct {
	AttributeID uuid.UUID
}

type CreateAttributeRequestDto struct {
	Data CreateAttributeData
}

type CreateAttributeData struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateAttributeRequestDto struct {
	AttributeID uuid.UUID
	Data        UpdateAttributeData
}

type UpdateAttributeData struct {
	Name string `json:"name,omitempty"`
}

type DeleteAttributeRequestDto struct {
	AttributeID uuid.UUID
}

type CreateAttributeValueRequestDto struct {
	AttributeID uuid.UUID
	Data        CreateAttributeValueData
}

type CreateAttributeValueData struct {
	Value string `json:"value" binding:"required"`
}

type UpdateAttributeValueRequestDto struct {
	AttributeID      uuid.UUID
	AttributeValueID uuid.UUID
	Data             UpdateAttributeValueData
}

type UpdateAttributeValueData struct {
	Value string `json:"value" binding:"required"`
}

type DeleteAttributeValueRequestDto struct {
	AttributeID      uuid.UUID
	AttributeValueID uuid.UUID
}
