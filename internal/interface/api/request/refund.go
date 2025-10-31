package request

type CreateRefund struct {
	PaymentID       int `json:"payment_id" binding:"required"`
	ReturnRequestID int `json:"return_request_id" binding:"required"`
}

type UpdateRefundStatus struct {
	StatusID int `json:"status_id" binding:"required"`
}
