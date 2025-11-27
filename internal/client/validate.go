package client

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
)

func NewValidate() *validator.Validate {
	validate := validator.New(
		validator.WithRequiredStructEnabled(),
	)
	if err := domain.RegisterAttributeValidators(validate); err != nil {
		panic(err)
	}
	if err := domain.RegisterCartValidators(validate); err != nil {
		panic(err)
	}
	if err := domain.RegisterProductValidators(validate); err != nil {
		panic(err)
	}
	if err := domain.RegisterOrderValidators(validate); err != nil {
		panic(err)
	}
	return validate
}
