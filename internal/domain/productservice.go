package domain

type ProductService interface {
	Validate(
		product Product,
	) error

	CreateOptionsWithOptionValues(
		optionsWithOptionValues map[string][]string,
	) (*[]Option, error)
}
