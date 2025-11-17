package request

type PaymentProvider string

const (
	PaymentProviderCOD     PaymentProvider = "COD"
	PaymentProviderVNPAY   PaymentProvider = "VNPAY"
	PaymentProviderMOMO    PaymentProvider = "MOMO"
	PaymentProviderZALOPAY PaymentProvider = "ZALO"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusFailed    PaymentStatus = "FAILED"
)

type CreatePayment struct {
	Provider PaymentProvider `json:"paymentProvider" binding:"required"`
	OrderID  int             `json:"orderId" binding:"required"`
}

type UpdatePayment struct {
	Status PaymentStatus `json:"status" binding:"required"`
}
