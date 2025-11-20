package domain

import "github.com/go-playground/validator/v10"

// LICENSE to Claude sonet 4.5 hehehe
func ProductVariantStructLevel(sl validator.StructLevel) {
	variant := sl.Current().Interface().(ProductVariant)

	// Skip validation if Product or OptionValues are nil
	if variant.Product == nil || variant.OptionValues == nil {
		return
	}

	// Skip validation if Product.Options is nil
	if variant.Product.Options == nil {
		return
	}

	optionValues := *variant.OptionValues
	productOptions := *variant.Product.Options

	// Validate: number of option values must equal number of product options
	if len(optionValues) != len(productOptions) {
		sl.ReportError(variant.OptionValues, "OptionValues", "OptionValues", "optionvalueslength", "")
		return
	}

	// Validate: each option value must belong to a different option
	optionIDsSeen := make(map[int]bool)

	for _, optionValue := range optionValues {
		// Find which option this option value belongs to
		optionID := findOptionIDForValue(productOptions, optionValue.ID)

		if optionID == 0 {
			sl.ReportError(variant.OptionValues, "OptionValues", "OptionValues", "invalidoptionvalue", "")
			return
		}

		// Check if we've already seen this option ID
		if optionIDsSeen[optionID] {
			sl.ReportError(variant.OptionValues, "OptionValues", "OptionValues", "duplicateoption", "")
			return
		}

		optionIDsSeen[optionID] = true
	}
}

// findOptionIDForValue finds which option contains the given option value ID
func findOptionIDForValue(options []Option, optionValueID int) int {
	for _, option := range options {
		if option.Values == nil {
			continue
		}

		for _, value := range *option.Values {
			if value.ID == optionValueID {
				return option.ID
			}
		}
	}
	return 0
}
