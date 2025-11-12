package review

import (
	"time"

	"backend/internal/domain/param"
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
	PaginationParams param.Params
}

type Pagination struct {
	Metadata *param.Metadata
	Reviews  *[]Review
}
