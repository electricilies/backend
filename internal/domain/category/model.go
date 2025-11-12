package category

import (
	"time"

	"backend/internal/domain/params"
)

type Model struct {
	ID        int
	Name      string
	CreatedAt time.Time
	DeletedAt *time.Time
	UpdatedAt time.Time
}

type PaginationModel struct {
	Metadata   *params.Metadata
	Categories *[]Model
}

type QueryParams struct {
	PaginationParams params.Params
}
