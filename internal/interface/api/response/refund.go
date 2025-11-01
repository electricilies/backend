package response

import "time"

type RefundStatus string

type Refund struct {
	ID            int           `json:"id" binding:"required"`
	Status        RefundStatus  `json:"status_id" binding:"required"`
	Payment       Payment       `json:"Payment" binding:"required"`
	ReturnRequest ReturnRequest `json:"return" binding:"required"`
	CreatedAt     time.Time     `json:"created_at" binding:"required"`
	UpdatedAt     time.Time     `json:"updated_at" binding:"required"`
}
