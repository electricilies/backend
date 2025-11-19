package service

type CreateUserParam struct {
	Data CreateUserData `json:"data" binding:"required"`
}

type CreateUserData struct {
	ID string `json:"id" binding:"required"`
}

type UpdateUserParam struct {
	UserID string         `json:"userId" binding:"required"`
	Data   UpdateUserData `json:"data" binding:"required"`
}

type UpdateUserData struct {
	FirstName   *string `json:"firstName" binding:"required"`
	LastName    *string `json:"lastName" binding:"required"`
	Email       *string `json:"email" binding:"required,email"`
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Address     *string `json:"address,omitempty"`
}

type ListUsersParam struct {
	Limit  int
	Offset int
}

type GetUserParam struct {
	UserID string `json:"userId" binding:"required"`
}

type GetUserCartParam struct {
	CartID string `json:"cartId" binding:"required"`
}
