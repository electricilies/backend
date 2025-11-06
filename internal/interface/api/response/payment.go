package response

import "time"

type (
	PaymentStatus   string
	PaymentProvider string
)

type Payment struct {
	ID                int             `json:"id" binding:"required"`
	Amount            int64           `json:"amount" binding:"required"`
	PaymentStatus     PaymentStatus   `json:"paymentStatusId" binding:"required"`
	PaymentProviderID PaymentProvider `json:"paymentProviderId" binding:"omitempty"`
	UpdatedAt         time.Time       `json:"updatedAt" binding:"required"`
}
