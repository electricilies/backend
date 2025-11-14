package review

import (
	"time"

	"backend/internal/domain/param"
	"backend/internal/domain/user"
)

type Model struct {
	ID        *int
	Rating    *int
	Content   *string
	ImageURL  *string
	User      *user.Model
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type QueryParams struct {
	PaginationParams *param.Pagination
	Deleted          *param.Deleted
}

type Pagination struct {
	Metadata *param.PaginationMetadata
	Reviews  *[]Model
}
