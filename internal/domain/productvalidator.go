package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// RegisterProductValidators registers custom validators for Product validation
func RegisterProductValidators(v *validator.Validate) error {
	if err := v.RegisterValidation("productVariantStructure", validateProductVariantStructure); err != nil {
		return err
	}
	return nil
}

// validateProductVariantStructure validates product-level rules including variants and options
func validateProductVariantStructure(fl validator.FieldLevel) bool {
	product, ok := fl.Parent().Interface().(Product)
	if !ok {
		return true // skip validation if not the right type
	}

	// Rule: Product must have at least 1 variant
	if len(product.Variants) < 1 {
		return false
	}

	// Case 1: No options → must have exactly 1 variant with no option values
	if len(product.Options) == 0 {
		if len(product.Variants) != 1 {
			return false
		}
		if len(product.Variants[0].OptionValues) != 0 {
			return false
		}
		return true
	}

	// Case 2: Has options → validate each variant
	for _, variant := range product.Variants {
		// Each variant must have option values matching the number of product options
		if len(variant.OptionValues) != len(product.Options) {
			return false
		}

		// Validate: each option value must belong to a different option
		optionIDsSeen := make(map[uuid.UUID]bool)

		for _, optionValue := range variant.OptionValues {
			// Find which option this option value belongs to
			optionID := findOptionIDForValue(product.Options, optionValue.ID)

			if optionID == uuid.Nil {
				return false
			}

			// Check if we've already seen this option ID
			if optionIDsSeen[optionID] {
				return false
			}

			optionIDsSeen[optionID] = true
		}
	}

	return true
}

// findOptionIDForValue finds which option contains the given option value ID
func findOptionIDForValue(options []Option, optionValueID uuid.UUID) uuid.UUID {
	for _, option := range options {
		if option.Values == nil {
			continue
		}

		for _, value := range option.Values {
			if value.ID == optionValueID {
				return option.ID
			}
		}
	}
	return uuid.Nil
}
