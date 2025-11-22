package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ParseIntQuery(query string) (int, error) {
	val, err := strconv.Atoi(query)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func ParseUUIDArrayQuery(queryArr []string, key string) (*[]uuid.UUID, bool) {
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

func GetOptionalQuery(ctx *gin.Context, key string) *string {
	if val, ok := ctx.GetQuery(key); ok {
		return &val
	}
	return nil
}

func GetDeletedParam(ctx *gin.Context, key string, defaultVal string) string {
	if val, ok := ctx.GetQuery(key); ok {
		return val
	}
	return defaultVal
}
