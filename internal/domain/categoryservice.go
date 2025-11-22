package domain

type CategoryService interface {
	Create(
		name string,
	) (*Category, error)

	Update(
		category *Category,
		name *string,
	) error
}
