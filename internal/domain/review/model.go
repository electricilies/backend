package review

import (
	"time"

	"backend/internal/domain/params"
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
	PaginationParams params.Params
}

type Pagination struct {
	Metadata *params.Metadata
	Reviews  *[]Review
}
