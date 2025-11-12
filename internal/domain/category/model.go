package category

import (
	"time"

	"backend/internal/domain/pagination"
)

type Model struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

type PaginationModel struct {
	Metadata   *pagination.Metadata
	Categories *[]Model
}

type QueryParams struct {
	PaginationParams pagination.Params
}
