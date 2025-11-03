package request

type CreatePayment struct {
	Amount            float64 `json:"amount" binding:"required"`
	PaymentMethodID   int     `json:"paymentMethodId" binding:"required"`
	PaymentProviderID int     `json:"paymentProviderId" binding:"required"`
}
