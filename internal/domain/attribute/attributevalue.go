package attribute

type AttributeValue struct {
	id    string
	value string
}

func NewAttributeValue(value string) *AttributeValue {
	return &AttributeValue{
		value: value,
	}
}

func (av *AttributeValue) GetID() string {
	return av.id
}

func (av *AttributeValue) GetValue() string {
	return av.value
}

func (av *AttributeValue) SetValue(value string) {
	av.value = value
}
