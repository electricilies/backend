package domain

type CategoryService interface {
	Create(
		name string,
	) (*Category, error)
}
