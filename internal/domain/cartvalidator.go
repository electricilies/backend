package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// RegisterCartValidators registers custom validators for Cart validation
func RegisterCartValidators(v *validator.Validate) error {
	if err := v.RegisterValidation("unique_cart_items", validateUniqueCartItems); err != nil {
		return err
	}
	return nil
}

// validateUniqueCartItems validates that all items within a cart have unique (ProductID, ProductVariantID) pairs
func validateUniqueCartItems(fl validator.FieldLevel) bool {
	items, ok := fl.Field().Interface().([]CartItem)
	if !ok {
		return true // skip validation if not the right type
	}

	type itemKey struct {
		productID        uuid.UUID
		productVariantID uuid.UUID
	}

	seen := make(map[itemKey]struct{})

	for _, item := range items {
		key := itemKey{
			productID:        item.ProductID,
			productVariantID: item.ProductVariantID,
		}

		if _, exists := seen[key]; exists {
			return false
		}
		seen[key] = struct{}{}
	}

	return true
}
