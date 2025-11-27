package domain

type ReviewService interface {
	Validate(review Review) error
}
