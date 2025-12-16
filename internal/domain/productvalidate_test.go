// vim: tabstop=4 shiftwidth=4:
package domain_test

import (
	"testing"

	"backend/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateProductVariantStructure_NoOptions_SingleVariant_NoOptionValues(t *testing.T) {
	product := &domain.Product{
		Options: []domain.Option{},
		Variants: []domain.ProductVariant{
			{
				ID:           uuid.New(),
				OptionValues: []domain.OptionValue{},
			},
		},
	}
	ok := domain.ValidateProductVariantStructure(product)
	assert.True(t, ok)
}

func TestValidateProductVariantStructure_NoOptions_MultipleVariants(t *testing.T) {
	product := &domain.Product{
		Options: []domain.Option{},
		Variants: []domain.ProductVariant{
			{ID: uuid.New(), OptionValues: []domain.OptionValue{}},
			{ID: uuid.New(), OptionValues: []domain.OptionValue{}},
		},
	}
	ok := domain.ValidateProductVariantStructure(product)
	assert.False(t, ok)
}

func TestValidateProductVariantStructure_NoOptions_SingleVariant_WithOptionValues(t *testing.T) {
	product := &domain.Product{
		Options: []domain.Option{},
		Variants: []domain.ProductVariant{
			{
				ID: uuid.New(),
				OptionValues: []domain.OptionValue{
					{ID: uuid.New(), Value: "value1"},
				},
			},
		},
	}
	ok := domain.ValidateProductVariantStructure(product)
	assert.False(t, ok)
}

func TestValidateProductVariantStructure_WithOptions_MatchingOptionValues(t *testing.T) {
	product := &domain.Product{
		Options: []domain.Option{
			{ID: uuid.New(), Name: "Color"},
			{ID: uuid.New(), Name: "Size"},
		},
		Variants: []domain.ProductVariant{
			{
				ID: uuid.New(),
				OptionValues: []domain.OptionValue{
					{ID: uuid.New(), Value: "Red"},
					{ID: uuid.New(), Value: "Large"},
				},
			},
			{
				ID: uuid.New(),
				OptionValues: []domain.OptionValue{
					{ID: uuid.New(), Value: "Blue"},
					{ID: uuid.New(), Value: "Small"},
				},
			},
		},
	}
	ok := domain.ValidateProductVariantStructure(product)
	assert.True(t, ok)
}

func TestValidateProductVariantStructure_WithOptions_MismatchedOptionValues(t *testing.T) {
	product := &domain.Product{
		Options: []domain.Option{
			{ID: uuid.New(), Name: "Color"},
			{ID: uuid.New(), Name: "Size"},
		},
		Variants: []domain.ProductVariant{
			{
				ID: uuid.New(),
				OptionValues: []domain.OptionValue{
					{ID: uuid.New(), Value: "Red"},
				},
			},
		},
	}
	ok := domain.ValidateProductVariantStructure(product)
	assert.False(t, ok)
}

func TestValidateProductVariantStructure_WithOptions_NoOptionValues(t *testing.T) {
	product := &domain.Product{
		Options: []domain.Option{
			{ID: uuid.New(), Name: "Color"},
		},
		Variants: []domain.ProductVariant{
			{
				ID:           uuid.New(),
				OptionValues: []domain.OptionValue{},
			},
		},
	}
	ok := domain.ValidateProductVariantStructure(product)
	assert.False(t, ok)
}
