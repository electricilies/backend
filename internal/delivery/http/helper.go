package http

import (
	"strconv"

	"backend/internal/application"
	"backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func queryArrayToUUIDSlice(ctx *gin.Context, key string) (*[]uuid.UUID, bool) {
	queryArr := ctx.QueryArray(key)
	if len(queryArr) == 0 {
		return nil, false
	}
	ids := make([]uuid.UUID, 0, len(queryArr))
	for _, idStr := range queryArr {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, false
		}
		ids = append(ids, id)
	}
	return &ids, true
}

func createPaginationParamsFromQuery(ctx *gin.Context) (*application.PaginationParam, error) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		return nil, domain.ErrInvalid
	}

	limitQuery := ctx.Query("limit")
	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		return nil, domain.ErrInvalid
	}
	return &application.PaginationParam{
		Page:  &page,
		Limit: &limit,
	}, nil
}
