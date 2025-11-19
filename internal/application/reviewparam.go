package application

type ListReviewsParam struct {
	PaginationParam
	OrderItemIDs     *[]int       `binding:"omitnil"`
	ProductVariantID *int         `binding:"omitnil"`
	UserIDs          *[]int       `binding:"omitnil"`
	Deleted          DeletedParam `binding:"required,oneof=exclude only all"`
}
