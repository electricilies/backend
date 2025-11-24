package http

import (
	"strconv"

	"backend/internal/application"
	"backend/internal/domain"
	"backend/internal/helper/ptr"

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

func queryToUUID(ctx *gin.Context, key string) (*uuid.UUID, bool) {
	idStr := ctx.Query(key)
	if idStr == "" {
		return nil, false
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, false
	}
	return &id, true
}

func pathToUUID(ctx *gin.Context, key string) (*uuid.UUID, bool) {
	idStr := ctx.Param(key)
	if idStr == "" {
		return ptr.To(uuid.Nil), false
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, false
	}
	return &id, true
}

func createPaginationParamsFromQuery(ctx *gin.Context) (*application.PaginationParam, error) {
	page := 1
	limit := 20
	var err error

	pageQuery := ctx.Query("page")
	if pageQuery != "" {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			return nil, domain.ErrInvalid
		}
	}

	limitQuery := ctx.Query("limit")
	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			return nil, domain.ErrInvalid
		}
	}
	return &application.PaginationParam{
		Page:  page,
		Limit: limit,
	}, nil
}
