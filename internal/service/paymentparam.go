package service

type CreatePaymentParam struct {
	Data CreatePaymentData
}

type CreatePaymentData struct {
	Provider string `json:"paymentProvider" binding:"required"`
	OrderID  int    `json:"orderId" binding:"required"`
}

type UpdatePaymentParam struct {
	PaymentID int
	Data      UpdatePaymentData
}

type UpdatePaymentData struct {
	Status *string `json:"status" binding:"required"`
}

type ListPaymentParam struct {
	PaginationParam
}

type GetPaymentParam struct {
	PaymentID int
}
