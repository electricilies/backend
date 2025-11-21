package domain

import "github.com/google/uuid"

type GetAttributeParam struct {
	AttributeID uuid.UUID `json:"attributeId" binding:"required"`
}

type CreateAttributeParam struct {
	Data CreateAttributeData `binding:"required"`
}

type CreateAttributeData struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateAttributeParam struct {
	AttributeID uuid.UUID           `binding:"required"`
	Data        UpdateAttributeData `binding:"required"`
}

type UpdateAttributeData struct {
	Name *string `json:"name" binding:"omitnil"`
}

type DeleteAttributeParam struct {
	AttributeID uuid.UUID `json:"attributeId" binding:"required"`
}

type CreateAttributeValueParam struct {
	AttributeID uuid.UUID                `binding:"required"`
	Data        CreateAttributeValueData `binding:"required"`
}

type CreateAttributeValueData struct {
	Value string `json:"value" binding:"required"`
}

type UpdateAttributeValueParam struct {
	AttributeValueID uuid.UUID                  `binding:"required"`
	Data             []UpdateAttributeValueData `binding:"required,dive"`
}

type UpdateAttributeValueData struct {
	Value *string `json:"value" binding:"required"`
}

type DeleteAttributeValueParam struct {
	AttributeValueID uuid.UUID `json:"attributeValueId" binding:"required"`
}
