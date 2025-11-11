package attribute

import "backend/internal/domain/pagination"

type Model struct {
	ID              string
	Code            string
	Name            string
	AttributeValues *[]ValueModel
}

type ValueModel struct {
	ID    string
	Value string
}

type QueryParams struct {
	PaginationParams *pagination.Params
	ProductID        *int
}

type PaginationModel struct {
	Attributes *[]Model
	Metadata   *pagination.Metadata
}
