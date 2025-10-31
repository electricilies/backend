package response

import "time"

type Return struct {
	ID          int        `json:"id"`
	Reason      string     `json:"reason"`
	StatusID    int        `json:"status_id"`
	UserID      string     `json:"user_id"`
	OrderItemID int        `json:"order_item_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
