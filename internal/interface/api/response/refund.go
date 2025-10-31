package response

import "time"

type Refund struct {
	ID              int        `json:"id"`
	StatusID        int        `json:"status_id"`
	PaymentID       int        `json:"payment_id"`
	ReturnRequestID int        `json:"return_request_id"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
