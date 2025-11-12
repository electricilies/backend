package request

import "backend/internal/domain/params"

func PaginationToDomain(limit int, offset int) *params.Params {
	return &params.Params{
		Limit:  limit,
		Offset: offset,
	}
}
