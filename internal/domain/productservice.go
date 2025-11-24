package domain

type ProductService interface {
	Validate(
		product Product,
	) error
}
