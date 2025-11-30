package domain

func ValidateOrderTotalAmount(order *Order) bool {
	var sum int64
	for _, item := range order.Items {
		sum += item.Price
	}
	return order.TotalAmount == sum
}
