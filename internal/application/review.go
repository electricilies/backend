package application

import (
	"context"

	"backend/internal/domain"
)

type Review interface {
	List(context.Context, ListReviewsParam) (*Pagination[domain.Review], error)
}
