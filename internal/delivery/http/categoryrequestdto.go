package http

import "github.com/google/uuid"

type ListCategoryRequestDto struct {
	PaginationRequestDto
	Search string
}

type CreateCategoryRequestDto struct {
	Data CreateCategoryData
}

type CreateCategoryData struct {
	Name string `json:"name" binding:"required"`
}

type GetCategoryRequestDto struct {
	CategoryID uuid.UUID
}

type UpdateCategoryRequestDto struct {
	CategoryID uuid.UUID
	Data       UpdateCategoryData
}

type UpdateCategoryData struct {
	Name string `json:"name"`
}
