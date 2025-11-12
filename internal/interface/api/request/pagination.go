package request

import "backend/internal/domain/param"

func PaginationToDomain(limit int, offset int) *param.Params {
	return &param.Params{
		Limit:  limit,
		Offset: offset,
	}
}
