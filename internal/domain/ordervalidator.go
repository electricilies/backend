package domain

import "github.com/go-playground/validator/v10"

func RegisterOrderValidators(v *validator.Validate) error {
	if err := v.RegisterValidation("orderTotalAmount", validateOrderTotalAmount); err != nil {
		return err
	}
	return nil
}

func validateOrderTotalAmount(fl validator.FieldLevel) bool {
	order, ok := fl.Parent().Interface().(Order)
	if !ok {
		return false
	}
	var sum int64
	for _, item := range order.Items {
		sum += item.Price
	}
	return order.TotalAmount == sum
}
