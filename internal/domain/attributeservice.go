package domain

type AttributeService interface {
	Validate(
		attribute Attribute,
	) error
}
