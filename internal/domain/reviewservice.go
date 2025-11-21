package domain

type ReviewService interface {
	Create(
		rating int,
		content string,
		orderItem OrderItem,
		ImageURL string,
	) (*Review, error)
}
