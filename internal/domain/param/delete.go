package param

type Deleted string

const (
	DeletedExclude Deleted = "exclude"
	DeletedOnly    Deleted = "only"
	DeletedAll     Deleted = "all"
)
