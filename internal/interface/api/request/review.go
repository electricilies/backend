package request

type CreateReview struct {
	ProductID int    `json:"productId" binding:"required"`
	UserID    string `json:"userId" binding:"required"`
	Rate      int    `json:"rate" binding:"required,min=1,max=5"`
	Content   string `json:"content,omitempty"`
	ImageURL  string `json:"imageUrl,omitempty"`
}

type UpdateReview struct {
	Rate     int    `json:"rate,omitempty"`
	Content  string `json:"content,omitempty"`
	ImageURL string `json:"imageUrl,omitempty"`
}
