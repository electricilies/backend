package domain

import "github.com/google/uuid"

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
