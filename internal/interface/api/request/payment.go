package request

type CreatePayment struct {
	Amount            float64 `json:"amount" binding:"required"`
	PaymentMethodID   int     `json:"payment_method_id" binding:"required"`
	PaymentProviderID int     `json:"payment_provider_id" binding:"required"`
}
