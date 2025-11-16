package application

type ProductService interface {
	ListAttributesByProductIDs(productIDs []int)
}
