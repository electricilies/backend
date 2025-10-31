package user

import (
	"context"

	"github.com/Thiht/transactor"
)

type Service interface {
	// This cotains  business logic like use case in Clean Arch
	DoSomething(ctx context.Context) (*User, error)
}

type service struct {
	transactor transactor.Transactor
	userRepo   Repository
}

func NewService(userRepo Repository, transactor transactor.Transactor) Service {
	return &service{
		userRepo:   userRepo,
		transactor: transactor,
	}
}

func (s service) DoSomething(ctx context.Context) (*User, error) {
	return s.userRepo.Get(ctx, "")
}
