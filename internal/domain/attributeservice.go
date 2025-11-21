package domain

type AttributeService interface {
	Create(
		code string,
		name string,
	) (*Attribute, error)

	CreateValue(
		value string,
	) (*AttributeValue, error)
}
