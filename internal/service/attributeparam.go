package service

type GetAttributeParam struct {
	AttributeID int `json:"attributeId" binding:"required"`
}

type ListAttributesParam struct {
	PaginationParam
	AttributeIDs *[]int  `binding:"omitnil"`
	Search       *string `binding:"omitnil"`
	Deleted      *string `binding:"omitnil"`
}

type CreateAttributeParam struct {
	Data CreateAttributeData `binding:"required"`
}

type CreateAttributeData struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateAttributeParam struct {
	AttributeID int                 `binding:"required"`
	Data        UpdateAttributeData `binding:"required"`
}

type UpdateAttributeData struct {
	Name *string `json:"name" binding:"omitnil"`
}

type DeleteAttributeParam struct {
	AttributeID int `json:"attributeId" binding:"required"`
}

type CreateAttributeValueParam struct {
	AttributeID int                      `binding:"required"`
	Data        CreateAttributeValueData `binding:"required"`
}

type CreateAttributeValueData struct {
	Value string `json:"value" binding:"required"`
}

type UpdateAttributeValueParam struct {
	AttributeValueID int                      `binding:"required"`
	Data             UpdateAttributeValueData `binding:"required"`
}

type UpdateAttributeValueData struct {
	Value *string `json:"value" binding:"required"`
}
