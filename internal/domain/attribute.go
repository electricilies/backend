package domain

import (
	"backend/internal/domain/pagination"
	"backend/internal/domain/param"
)

type Attribute struct {
	ID              string
	Code            string
	Name            string
	AttributeValues []AttributeValue
}

type AttributeValue struct {
	ID    string
	Value string
}

type AttributePagination struct {
	Attributes []Attribute
	Metadata   pagination.Metadata
}

type CreateAttribute struct {
	Code string
	Name string
}

type ListAttributes struct {
	param.Pagination
	ProductID int
	Search    string
	Deleted   param.Deleted
}

type CreateAttributeValues struct {
	AttributeID string
	Values      []string
}

type GetAttribute struct {
	ID string
}

type UpdateAttribute struct {
	ID   string
	Name string
}

type UpdateAttributeValues struct {
	AttributeID string
	Values      []AttributeValue
}
