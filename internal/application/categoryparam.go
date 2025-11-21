package application

import "github.com/google/uuid"

type ListCategoryParam struct {
	PaginationParam
	Search *string `json:"search" binding:"omitnil"`
}

type CreateCategoryParam struct {
	Data CreateCategoryData `binding:"required"`
}

type CreateCategoryData struct {
	Name string `json:"name" binding:"required"`
}

type GetCategoryParam struct {
	CategoryID uuid.UUID `binding:"required"`
}

type UpdateCategoryParam struct {
	CategoryID uuid.UUID          `binding:"required"`
	Data       UpdateCategoryData `binding:"required"`
}

type UpdateCategoryData struct {
	Name *string `json:"name" binding:"omitnil"`
}
