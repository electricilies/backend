package request

import "backend/internal/domain/pagination"

func PaginationToDomain(limit int, offset int) *pagination.Params {
	return &pagination.Params{
		Limit:  limit,
		Offset: offset,
	}
}
