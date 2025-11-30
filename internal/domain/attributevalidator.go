package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func RegisterAttributeValidators(v *validator.Validate) error {
	if err := v.RegisterValidation("uniqueAttributeValues", validateUniqueAttributeValues); err != nil {
		return err
	}
	return nil
}

func validateUniqueAttributeValues(fl validator.FieldLevel) bool {
	values, ok := fl.Field().Interface().([]AttributeValue) // 1
	if !ok {                                                // 2
		return true // 3
	}

	seen := make(map[string]struct{})       // 4
	seenIDs := make(map[uuid.UUID]struct{}) // 4

	for _, val := range values { // 5
		v := val.Value                    // 6
		if _, exists := seen[v]; exists { // 6, 7
			return false // 8
		}
		seen[v] = struct{}{} // 9

		if _, exists := seenIDs[val.ID]; exists { // 9, 10
			return false // 11
		}
		seenIDs[val.ID] = struct{}{} // 12
	}

	return true // 13
}
