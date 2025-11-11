package category

import (
	"time"

	"backend/internal/domain/pagination"
)

type Model struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
}

type Pagination struct {
	Metadata   *pagination.Metadata
	Categories *[]Model
}

type QueryParams struct {
	PaginationParams pagination.Params
}
