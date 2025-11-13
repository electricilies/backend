package request

import "backend/internal/domain/param"

func PaginationToDomain(limit int, offset int) *param.Pagination {
	return &param.Pagination{
		Limit:  &limit,
		Offset: &offset,
	}
}
