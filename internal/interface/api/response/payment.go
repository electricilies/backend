package response

import "time"

type (
	PaymentMethod   string
	PaymentStatus   string
	PaymentProvider string
)

type Payment struct {
    ID                int             `json:"id" binding:"required"`
    Amount            float64         `json:"amount" binding:"required"`
    PaymentMethod     PaymentMethod   `json:"paymentMethodId" binding:"required"`
    PaymentStatus     PaymentStatus   `json:"paymentStatusId" binding:"required"`
    PaymentProviderID PaymentProvider `json:"paymentProviderId" binding:"omitempty"`
    UpdatedAt         time.Time       `json:"updatedAt" binding:"required"`
}
