package request

type RefundStatus string

type CreateRefund struct {
	PaymentID       int `json:"paymentId" binding:"required"`
	ReturnRequestID int `json:"returnRequestId" binding:"required"`
}

type UpdateRefundStatus struct {
	Status RefundStatus `json:"refundStatus" binding:"required"`
}
