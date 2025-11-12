package attribute

import "backend/internal/domain/param"

type Deleted string

const (
	Exclude Deleted = "exclude"
	Only    Deleted = "only"
	All     Deleted = "all"
)

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
	PaginationParams *param.Params
	ProductID        int
	Search           string
	Deleted          *Deleted
}

type PaginationModel struct {
	Attributes *[]Model
	Metadata   *param.Metadata
}
