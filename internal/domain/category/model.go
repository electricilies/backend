package category

import (
	"time"

	"backend/internal/domain/param"
)

type Model struct {
	ID        int
	Name      string
	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}

type PaginationModel struct {
	Metadata   *param.PaginationMetadata
	Categories *[]Model
}

type QueryParams struct {
	PaginationParams param.Pagination
}
