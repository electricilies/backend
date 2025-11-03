package request

type CreateRefund struct {
	PaymentID       int `json:"paymentId" binding:"required"`
	ReturnRequestID int `json:"returnRequestId" binding:"required"`
}

type UpdateRefundStatus struct {
	StatusID int `json:"statusId" binding:"required"`
}
