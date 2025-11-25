package http

import "github.com/google/uuid"

type ListCategoryRequestDto struct {
	PaginationRequestDto
	Search *string `json:"search" binding:"omitnil"`
}

type CreateCategoryRequestDto struct {
	Data CreateCategoryData `binding:"required"`
}

type CreateCategoryData struct {
	Name string `json:"name" binding:"required"`
}

type GetCategoryRequestDto struct {
	CategoryID uuid.UUID `binding:"required"`
}

type UpdateCategoryRequestDto struct {
	CategoryID uuid.UUID          `binding:"required"`
	Data       UpdateCategoryData `binding:"required"`
}

type UpdateCategoryData struct {
	Name *string `json:"name" binding:"omitnil"`
}
