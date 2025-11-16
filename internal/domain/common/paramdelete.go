package common

type DeletedParam string

const (
	DeletedParamExclude DeletedParam = "exclude"
	DeletedParamOnly    DeletedParam = "only"
	DeletedParamAll     DeletedParam = "all"
)
