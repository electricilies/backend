package client

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
)

func NewValidate() *validator.Validate {
	validate := validator.New(
		validator.WithRequiredStructEnabled(),
	)
	if err := domain.RegisterProductValidates(validate); err != nil {
		panic(err)
	}
	if err := domain.RegisterOrderValidates(validate); err != nil {
		panic(err)
	}
	return validate
}
