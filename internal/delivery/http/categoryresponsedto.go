package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

// CategoryResponseDto represents the response structure for a category
type CategoryResponseDto struct {
	ID        uuid.UUID  `json:"id"        binding:"required"`
	Name      string     `json:"name"      binding:"required"`
	CreatedAt time.Time  `json:"createdAt" binding:"required"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// ToCategoryResponseDto maps a domain.Category to CategoryResponseDto
func ToCategoryResponseDto(cat *domain.Category) *CategoryResponseDto {
	if cat == nil {
		return nil
	}

	return &CategoryResponseDto{
		ID:        cat.ID,
		Name:      cat.Name,
		CreatedAt: cat.CreatedAt,
		UpdatedAt: cat.UpdatedAt,
		DeletedAt: cat.DeletedAt,
	}
}

// ToCategoryResponseDtoList maps a slice of domain.Category to a slice of CategoryResponseDto
func ToCategoryResponseDtoList(cats []domain.Category) []CategoryResponseDto {
	result := make([]CategoryResponseDto, 0, len(cats))
	for _, cat := range cats {
		dto := ToCategoryResponseDto(&cat)
		if dto != nil {
			result = append(result, *dto)
		}
	}
	return result
}
