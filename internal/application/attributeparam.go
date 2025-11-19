package application

type ListAttributesParam struct {
	PaginationParam
	IDs     *[]int       `binding:"omitnil"`
	Search  *string      `binding:"omitnil"`
	Deleted DeletedParam `binding:"required,oneof=exclude only all"`
}

type ListAttributeValuesParam struct {
	PaginationParam
	AttributeValueIDs *[]int  `binding:"omitnil"`
	AttributeIDs      *[]int  `binding:"omitnil"`
	Search            *string `binding:"omitnil"`
}
