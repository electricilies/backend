package constant

type UserRole string

const (
	RoleCustomer UserRole = "customer"
	RoleAdmin    UserRole = "admin"
	RoleStaff    UserRole = "staff"
)
