package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

// AttributeResponseDto represents the response structure for an attribute
type AttributeResponseDto struct {
	ID        uuid.UUID                   `json:"id"        binding:"required"`
	Code      string                      `json:"code"      binding:"required"`
	Name      string                      `json:"name"      binding:"required"`
	Values    []AttributeValueResponseDto `json:"values"    binding:"required"`
	DeletedAt *time.Time                  `json:"deletedAt" binding:"required"`
}

// AttributeValueResponseDto represents the response structure for an attribute value
type AttributeValueResponseDto struct {
	ID        uuid.UUID  `json:"id"        binding:"required"`
	Value     string     `json:"value"     binding:"required"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// ToAttributeResponseDto maps a domain.Attribute to AttributeResponseDto
func ToAttributeResponseDto(attr *domain.Attribute) *AttributeResponseDto {
	if attr == nil {
		return nil
	}

	values := make([]AttributeValueResponseDto, 0, len(attr.Values))
	for _, v := range attr.Values {
		values = append(values, AttributeValueResponseDto{
			ID:        v.ID,
			Value:     v.Value,
			DeletedAt: v.DeletedAt,
		})
	}

	return &AttributeResponseDto{
		ID:        attr.ID,
		Code:      attr.Code,
		Name:      attr.Name,
		Values:    values,
		DeletedAt: attr.DeletedAt,
	}
}

// ToAttributeValueResponseDto maps a domain.AttributeValue to AttributeValueResponseDto
func ToAttributeValueResponseDto(val *domain.AttributeValue) *AttributeValueResponseDto {
	if val == nil {
		return nil
	}

	return &AttributeValueResponseDto{
		ID:        val.ID,
		Value:     val.Value,
		DeletedAt: val.DeletedAt,
	}
}

// ToAttributeResponseDtoList maps a slice of domain.Attribute to a slice of AttributeResponseDto
func ToAttributeResponseDtoList(attrs []domain.Attribute) []AttributeResponseDto {
	result := make([]AttributeResponseDto, 0, len(attrs))
	for _, attr := range attrs {
		dto := ToAttributeResponseDto(&attr)
		if dto != nil {
			result = append(result, *dto)
		}
	}
	return result
}

// ToAttributeValueResponseDtoList maps a slice of domain.AttributeValue to a slice of AttributeValueResponseDto
func ToAttributeValueResponseDtoList(vals []domain.AttributeValue) []AttributeValueResponseDto {
	result := make([]AttributeValueResponseDto, 0, len(vals))
	for _, val := range vals {
		dto := ToAttributeValueResponseDto(&val)
		if dto != nil {
			result = append(result, *dto)
		}
	}
	return result
}
