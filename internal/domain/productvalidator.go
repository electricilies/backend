package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// ProductStructLevel validates product-level rules including variants and options
func ProductStructLevel(sl validator.StructLevel) {
	product := sl.Current().Interface().(Product)

	// Rule: Product must have at least 1 variant
	if len(product.Variants) < 1 {
		sl.ReportError(product.Variants, "Variants", "Variants", "minvariants", "")
		return
	}

	// Case 1: No options → must have exactly 1 variant with no option values
	if len(product.Options) == 0 {
		if len(product.Variants) != 1 {
			sl.ReportError(product.Variants, "Variants", "Variants", "nooptionssinglevariant", "")
			return
		}
		if len(product.Variants[0].OptionValues) != 0 {
			sl.ReportError(product.Variants[0].OptionValues, "OptionValues", "OptionValues", "nooptionsnooptionvalues", "")
			return
		}
		return
	}

	// Case 2: Has options → validate each variant
	for i, variant := range product.Variants {
		// Each variant must have option values matching the number of product options
		if len(variant.OptionValues) != len(product.Options) {
			sl.ReportError(variant.OptionValues, "OptionValues", "OptionValues", "optionvalueslength", "")
			return
		}

		// Validate: each option value must belong to a different option
		optionIDsSeen := make(map[uuid.UUID]bool)

		for _, optionValue := range variant.OptionValues {
			// Find which option this option value belongs to
			optionID := findOptionIDForValue(product.Options, optionValue.ID)

			if optionID == uuid.Nil {
				sl.ReportError(product.Variants[i].OptionValues, "OptionValues", "OptionValues", "invalidoptionvalue", "")
				return
			}

			// Check if we've already seen this option ID
			if optionIDsSeen[optionID] {
				sl.ReportError(product.Variants[i].OptionValues, "OptionValues", "OptionValues", "duplicateoption", "")
				return
			}

			optionIDsSeen[optionID] = true
		}
	}
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
