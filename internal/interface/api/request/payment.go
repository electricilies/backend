package request

type CreatePayment struct {
	Amount            float64 `json:"amount" binding:"required"`
	PaymentProviderID int     `json:"paymentProviderId" binding:"required"`
	OrderID           int     `json:"orderId" binding:"required"`
}
