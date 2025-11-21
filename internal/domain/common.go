package domain

type DeletedParam string

const (
	DeletedExcludeParam DeletedParam = "exclude"
	DeletedOnlyParam    DeletedParam = "only"
	DeletedAllParam     DeletedParam = "all"
)

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleStaff    UserRole = "staff"
	RoleCustomer UserRole = "customer"
)
