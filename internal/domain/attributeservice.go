package domain

type AttributeService interface {
	Create(
		code string,
		name string,
	) (*Attribute, error)

	Update(
		attribute *Attribute,
		name *string,
	) error

	CreateValue(
		value string,
	) (*AttributeValue, error)
}
