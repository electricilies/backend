package request

type PaymentProvider string

type CreatePayment struct {
	Amount          int64           `json:"amount" binding:"required"`
	PaymentProvider PaymentProvider `json:"paymentProvider" binding:"required"`
	OrderID         int             `json:"orderId" binding:"required"`
}
