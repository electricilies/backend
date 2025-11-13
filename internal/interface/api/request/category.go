package request

import "backend/internal/domain/category"

type CreateCategory struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategory struct {
	Name string `json:"name" binding:"required"`
}

type CategoryQueryParams struct {
	Limit  int
	Offset int
}

func CategoryQueryParamsToDomain(categoryQueryParams *CategoryQueryParams) *category.QueryParams {
	return &category.QueryParams{
		PaginationParams: PaginationParamsToDomain(
			categoryQueryParams.Limit,
			categoryQueryParams.Offset,
		),
	}
}
