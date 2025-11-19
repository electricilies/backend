package serviceimpl

type Product struct{}

func ProvideProduct() *Product {
	return &Product{}
}

// var _ domain.Product = &Product{}
