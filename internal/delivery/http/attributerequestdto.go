package http

import (
	"backend/internal/domain"
	"github.com/google/uuid"
)

type ListAttributesRequestDto struct {
	PaginationRequestDto
	AttributeIDs *[]uuid.UUID        `binding:"omitnil"`
	ProductIDs   *[]uuid.UUID        `binding:"omitnil"`
	Deleted      domain.DeletedParam `binding:"required,oneof=exclude only all"`
	Search       *string             `binding:"omitnil"`
}

type ListAttributeValuesRequestDto struct {
	PaginationRequestDto
	AttributeID       uuid.UUID
	AttributeValueIDs *[]uuid.UUID        `binding:"omitnil"`
	Deleted           domain.DeletedParam `binding:"required,oneof=exclude only all"`
	Search            *string             `binding:"omitnil"`
}

type GetAttributeRequestDto struct {
	AttributeID uuid.UUID `json:"attributeId" binding:"required"`
}

type CreateAttributeRequestDto struct {
	Data CreateAttributeData `binding:"required"`
}

type CreateAttributeData struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateAttributeRequestDto struct {
	AttributeID uuid.UUID           `binding:"required"`
	Data        UpdateAttributeData `binding:"required"`
}

type UpdateAttributeData struct {
	Name *string `json:"name" binding:"omitnil"`
}

type DeleteAttributeRequestDto struct {
	AttributeID uuid.UUID `json:"attributeId" binding:"required"`
}

type CreateAttributeValueRequestDto struct {
	AttributeID uuid.UUID                `binding:"required"`
	Data        CreateAttributeValueData `binding:"required"`
}

type CreateAttributeValueData struct {
	Value string `json:"value" binding:"required"`
}

type UpdateAttributeValueRequestDto struct {
	AttributeID      uuid.UUID                `binding:"required"`
	AttributeValueID uuid.UUID                `binding:"required"`
	Data             UpdateAttributeValueData `binding:"required,dive"`
}

type UpdateAttributeValueData struct {
	Value *string `json:"value" binding:"required"`
}

type DeleteAttributeValueRequestDto struct {
	AttributeID      uuid.UUID `json:"attributeId"      binding:"required"`
	AttributeValueID uuid.UUID `json:"attributeValueId" binding:"required"`
}
