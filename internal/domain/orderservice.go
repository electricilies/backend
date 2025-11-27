package domain

type OrderService interface {
	Validate(
		order Order,
	) error
}
