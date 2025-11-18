package service

type GetAttributeParam struct {
	AttributeID int `json:"attributeId" binding:"required"`
}

type ListAttributesParam struct {
	PaginationParam
	AttributeIDs *[]int
	Search       *string
	Deleted      *string
}

type CreateAttributeParam struct {
	Data CreateAttributeData
}

type CreateAttributeData struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateAttributeParam struct {
	AttributeID int
	Data        UpdateAttributeData
}

type UpdateAttributeData struct {
	Name *string `json:"name"`
}

type DeleteAttributeParam struct {
	AttributeID int `json:"attributeId" binding:"required"`
}

type CreateAttributeValueParam struct {
	AttributeID int
	Data        CreateAttributeValueData
}

type CreateAttributeValueData struct {
	Value string `json:"value" binding:"required"`
}

type UpdateAttributeValueParam struct {
	AttributeValueIDs int
	Data              UpdateAttributeValueData
}

type UpdateAttributeValueData struct {
	Value *string `json:"value" binding:"required"`
}
