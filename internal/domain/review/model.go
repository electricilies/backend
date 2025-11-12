package review

import (
	"time"

	"backend/internal/domain/pagination"
	"backend/internal/domain/user"
)

type Review struct {
	ID        int
	Rating    int
	Content   string
	ImageURL  string
	User      *user.Model
	CreatedAt time.Time
	UpdatedAt time.Time
}

type QueryParams struct {
	PaginationParams pagination.Params
}

type Pagination struct {
	Metadata *pagination.Metadata
	Reviews  *[]Review
}
