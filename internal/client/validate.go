package client

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
)

func NewValidate() *validator.Validate {
	validate := validator.New(
		validator.WithRequiredStructEnabled(),
	)
	validate.RegisterStructValidation(
		domain.ProductStructLevel,
		domain.Product{},
	)
	return validate
}
