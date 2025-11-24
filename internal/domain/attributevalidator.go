package domain

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// RegisterAttributeValidators registers custom validators for Attribute validation
func RegisterAttributeValidators(v *validator.Validate) error {
	if err := v.RegisterValidation("unique_attribute_values", validateUniqueAttributeValues); err != nil {
		return err
	}
	return nil
}

// validateUniqueAttributeValues validates that all values within an attribute are unique
func validateUniqueAttributeValues(fl validator.FieldLevel) bool {
	values, ok := fl.Field().Interface().([]AttributeValue)
	if !ok {
		return true // skip validation if not the right type
	}

	seen := make(map[string]struct{})
	seenIDs := make(map[uuid.UUID]struct{})
	
	for _, val := range values {
		// Check for duplicate values (case-insensitive)
		valueLower := val.Value
		if _, exists := seen[valueLower]; exists {
			return false
		}
		seen[valueLower] = struct{}{}

		// Check for duplicate IDs
		if _, exists := seenIDs[val.ID]; exists {
			return false
		}
		seenIDs[val.ID] = struct{}{}
	}

	return true
}

// ValidateUniqueAttributeIDs validates that all attribute IDs in a slice are unique
func ValidateUniqueAttributeIDs(attributes []Attribute) error {
	seen := make(map[uuid.UUID]struct{})
	for _, attr := range attributes {
		if _, exists := seen[attr.ID]; exists {
			return fmt.Errorf("duplicate attribute ID found: %s", attr.ID)
		}
		seen[attr.ID] = struct{}{}
	}
	return nil
}

// ValidateUniqueAttributeCodes validates that all attribute codes in a slice are unique
func ValidateUniqueAttributeCodes(attributes []Attribute) error {
	seen := make(map[string]struct{})
	for _, attr := range attributes {
		if _, exists := seen[attr.Code]; exists {
			return fmt.Errorf("duplicate attribute code found: %s", attr.Code)
		}
		seen[attr.Code] = struct{}{}
	}
	return nil
}

// ValidateAttributeValueUniqueness validates that values within an attribute are unique
func ValidateAttributeValueUniqueness(attr *Attribute) error {
	seen := make(map[string]struct{})
	seenIDs := make(map[uuid.UUID]struct{})

	for _, val := range attr.Values {
		// Check for duplicate values
		if _, exists := seen[val.Value]; exists {
			return fmt.Errorf("duplicate attribute value found in attribute %s: %s", attr.Code, val.Value)
		}
		seen[val.Value] = struct{}{}

		// Check for duplicate IDs
		if _, exists := seenIDs[val.ID]; exists {
			return fmt.Errorf("duplicate attribute value ID found in attribute %s: %s", attr.Code, val.ID)
		}
		seenIDs[val.ID] = struct{}{}
	}

	return nil
}
