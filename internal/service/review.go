package service

import (
	"backend/internal/domain"
)

type Review struct{}

func ProvideReview() *Review {
	return &Review{}
}

var _ domain.ReviewService = &Review{}
