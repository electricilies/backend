package application

type ListOrderParam struct {
	PaginationParam
	IDs       *[]int    `binding:"omitnil"`
	UserIDs   *[]string `binding:"omitnil"`
	StatusIDs *[]int    `binding:"omitnil"`
}
