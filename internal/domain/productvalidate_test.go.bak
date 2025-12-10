// vim: tabstop=4:
package domain_test

import (
	"testing"

	"backend/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type ProductTestSuite struct {
	suite.Suite
}

func (suite *ProductTestSuite) TestValidateProductVariantStructure_MultiVariants() {
	options := []domain.Option{
		{
			ID:   uuid.New(),
			Name: "Size",
			Values: []domain.OptionValue{
				{ID: uuid.New(), Value: "Small"},
				{ID: uuid.New(), Value: "Medium"},
			},
		},
		{
			ID:   uuid.New(),
			Name: "Color",
			Values: []domain.OptionValue{
				{ID: uuid.New(), Value: "Red"},
				{ID: uuid.New(), Value: "Blue"},
			},
		},
	}

	testcases := []struct {
		name     string
		variants []domain.ProductVariant
		expectOk bool
	}{
		{
			name: "Valid variants with correct option values",
			variants: []domain.ProductVariant{
				{
					ID:  uuid.New(),
					SKU: "SKU1",
					OptionValues: []domain.OptionValue{
						options[0].Values[0], // Small
						options[1].Values[0], // Red
					},
				},
				{
					ID:  uuid.New(),
					SKU: "SKU2",
					OptionValues: []domain.OptionValue{
						options[0].Values[1], // Medium
						options[1].Values[1], // Blue
					},
				},
			},
			expectOk: true,
		},
		{
			name: "Variant with missing option value",
			variants: []domain.ProductVariant{
				{
					ID:  uuid.New(),
					SKU: "SKU1",
					OptionValues: []domain.OptionValue{
						options[0].Values[0], // Small
					},
				},
			},
			expectOk: false,
		},
	}
	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			product := &domain.Product{
				Options:  options,
				Variants: tc.variants,
			}
			ok := domain.ValidateProductVariantStructure(product)
			suite.Equal(tc.expectOk, ok)
		})
	}
}

func (suite *ProductTestSuite) TestValidateProductVariantStructure_SingleVariant() {
	testcases := []struct {
		name     string
		product  domain.Product
		expectOk bool
	}{
		{
			name: "No options with single variant and no option values",
			product: domain.Product{
				Options: []domain.Option{},
				Variants: []domain.ProductVariant{
					{
						ID:           uuid.New(),
						OptionValues: []domain.OptionValue{},
					},
				},
			},
			expectOk: true,
		},
		{
			name: "No option but multiple variants",
			product: domain.Product{
				Options: []domain.Option{},
				Variants: []domain.ProductVariant{
					{
						ID:           uuid.New(),
						OptionValues: []domain.OptionValue{},
					},
					{
						ID:           uuid.New(),
						OptionValues: []domain.OptionValue{},
					},
				},
			},
			expectOk: false,
		},
		{
			name: "No options with single variant but has option values",
			product: domain.Product{
				Options: []domain.Option{},
				Variants: []domain.ProductVariant{
					{
						ID:  uuid.New(),
						SKU: "SKU1",
						OptionValues: []domain.OptionValue{
							{ID: uuid.New(), Value: "SomeValue"},
						},
					},
				},
			},
			expectOk: false,
		},
	}
	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			ok := domain.ValidateProductVariantStructure(&tc.product)
			suite.Equal(tc.expectOk, ok)
		})
	}
}

func TestProductTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ProductTestSuite))
}
