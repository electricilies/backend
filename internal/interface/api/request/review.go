package request

import _ "github.com/go-playground/validator/v10"

type CreateReview struct {
	ProductID int    `json:"product_id" binding:"required"`
	UserID    string `json:"user_id" binding:"required"`
	Rate      int    `json:"rate" binding:"required" validator:"min=1,max=5"`
	Content   string `json:"content,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

type UpdateReview struct {
	Rate     int    `json:"rate,omitempty"`
	Content  string `json:"content,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}
