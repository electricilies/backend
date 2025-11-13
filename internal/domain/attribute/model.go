package attribute

import "backend/internal/domain/param"

type Model struct {
	ID              *string
	Code            *string
	Name            *string
	AttributeValues *[]ValueModel
}

type ValueModel struct {
	ID    *string
	Value *string
}

type QueryParams struct {
	PaginationParams *param.Pagination
	ProductID        *int
	Search           *string
	Deleted          *param.Deleted
}

type PaginationModel struct {
	Attributes *[]Model
	Metadata   *param.PaginationMetadata
}
