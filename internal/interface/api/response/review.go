package response

import "time"

type ReviewResponse struct {
	ID        int        `json:"id"`
	Rate      int        `json:"rate"`
	Content   string     `json:"content,omitempty"`
	ImageURL  string     `json:"image_url,omitempty"`
	UserID    string     `json:"user_id"`
	ProductID int        `json:"product_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ReviewListResponse struct {
	Reviews []ReviewResponse `json:"reviews"`
}
