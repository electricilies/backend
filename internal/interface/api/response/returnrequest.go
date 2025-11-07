package response

import "time"

type ReturnRequestStatus string

type ReturnRequest struct {
	ID        int                 `json:"id" binding:"required"`
	Reason    string              `json:"reason" binding:"required"`
	Status    ReturnRequestStatus `json:"status" binding:"required"`
	User      User                `json:"user" binding:"required"`
	OrderItem OrderItem           `json:"orderItem" binding:"required"`
	CreatedAt time.Time           `json:"createdAt" binding:"required"`
	UpdatedAt time.Time           `json:"updatedAt" binding:"required"`
}
