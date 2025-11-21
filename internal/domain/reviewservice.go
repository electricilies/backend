package domain

type ReviewService interface {
	Create(
		orderItem OrderItem,
		rating int,
		content *string,
		ImageURL *string,
	) (*Review, error)
}
