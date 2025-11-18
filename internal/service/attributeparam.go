package service

type GetAttributeParam struct {
	AttributeID int `json:"attributeId" binding:"required"`
}

type ListAttributesParam struct {
	PaginationParam
	AttributeIDs []int
	Search       *string
	Deleted      *string
}

type CreateAttributeData struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CreateAttributeParam struct {
	Data CreateAttributeData
}

type UpdateAttributeData struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAttributeParam struct {
	AttributeID int
	Data        UpdateAttributeData `json:"data" binding:"required"`
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
