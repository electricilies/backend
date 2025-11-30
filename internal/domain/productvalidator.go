package domain

import (
	"github.com/go-playground/validator/v10"
)

func RegisterProductValidators(v *validator.Validate) error {
	if err := v.RegisterValidation("productVariantStructure", productVariantStructureValidator); err != nil {
		return err
	}
	return nil
}

func productVariantStructureValidator(fl validator.FieldLevel) bool {
	product, ok := fl.Parent().Interface().(Product)
	if !ok {
		return true
	}
	return ValidateProductVariantStructure(&product)
}
