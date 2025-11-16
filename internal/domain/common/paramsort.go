package common

type SortParam string

const (
	SortParamAscending  SortParam = "asc"
	SortParamDescending SortParam = "desc"
)

type (
	SortParamRating SortParam
	SortParamPrice  SortParam
)
