package domain

import "github.com/go-playground/validator/v10"

func RegisterOrderValidators(v *validator.Validate) error {
	if err := v.RegisterValidation("orderTotalAmount", orderTotalAmountValidator); err != nil {
		return err
	}
	return nil
}

func orderTotalAmountValidator(fl validator.FieldLevel) bool {
	order, ok := fl.Parent().Interface().(Order)
	if !ok {
		return true
	}
	var sum int64
	for _, item := range order.Items {
		sum += item.Price * int64(item.Quantity)
	}
	return order.TotalAmount == sum
}
