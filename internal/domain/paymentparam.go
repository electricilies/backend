package domain

type CreatePaymentParam struct {
	Data CreatePaymentData `binding:"required"`
}

type CreatePaymentData struct {
	Provider PaymentProvider `json:"paymentProvider" binding:"required,oneof=COD VNPAY MOMO ZALOPAY"`
	OrderID  int             `json:"orderId"         binding:"required"`
}

type UpdatePaymentParam struct {
	PaymentID int               `binding:"required"`
	Data      UpdatePaymentData `binding:"required"`
}

type UpdatePaymentData struct {
	Status *string `json:"status" binding:"omitnil,oneof=Pending Completed Failed"`
}

type GetPaymentParam struct {
	PaymentID int `binding:"required"`
}
