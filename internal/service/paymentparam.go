package service

import "backend/internal/domain"

type CreatePaymentParam struct {
	Data CreatePaymentData
}

type CreatePaymentData struct {
	Provider domain.PaymentProvider `json:"paymentProvider" binding:"required,oneof=COD VNPAY MOMO ZALOPAY"`
	OrderID  int                    `json:"orderId" binding:"required"`
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
