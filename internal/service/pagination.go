package service

type PaginationParam struct {
	Page  int `json:"page" example:"1"`
	Limit int `json:"limit" example:"10"`
}
