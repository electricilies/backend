package domain

type User struct {
	ID string `json:"id" binding:"required" validate:"required,uuid"`
}
