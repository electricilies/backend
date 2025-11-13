package request

import (
	"backend/internal/domain/attribute"
)

type CreateAttribute struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateAttribute struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAttributeValue struct {
	TargetID int    `json:"targetId" binding:"required"`
	NewValue string `json:"newValue" binding:"required"`
}

type CreateAttributeValue struct {
	Value string `json:"value" binding:"required"`
}

type AttributeQueryParams struct {
	Limit     int
	Offset    int
	ProductID int
	Search    string
	Deleted   string
}

func AttributeQueryParamsToDomain(attributeQueryParams *AttributeQueryParams) *attribute.QueryParams {
	return &attribute.QueryParams{
		PaginationParams: PaginationParamsToDomain(attributeQueryParams.Limit, attributeQueryParams.Offset),
		ProductID:        &attributeQueryParams.ProductID,
		Search:           &attributeQueryParams.Search,
		Deleted:          DeletedParamToDomain(attributeQueryParams.Deleted),
	}
}
