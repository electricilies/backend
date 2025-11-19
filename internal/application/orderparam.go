package application

type ListOrderParam struct {
	PaginationParam
	OrderIDs  *[]int    `binding:"omitnil"`
	UserIDs   *[]string `binding:"omitnil"`
	StatusIDs *[]int    `binding:"omitnil"`
}
