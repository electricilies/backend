package response

import "time"

type ReturnStatus string

type Return struct {
	ID        int          `json:"id" binding:"required"`
	Reason    string       `json:"reason" binding:"required"`
	Status    ReturnStatus `json:"status" binding:"required"`
	User      User         `json:"user" binding:"required"`
	OrderItem OrderItem    `json:"order_item" binding:"required"`
	CreatedAt time.Time    `json:"created_at" binding:"required"`
	UpdatedAt time.Time    `json:"updated_at" binding:"required"`
}
