package user

import (
	"context"

	"github.com/Thiht/transactor"
)

type Service interface {
	// This cotains  business logic like use case in Clean Arch
	DoSomething(ctx context.Context) (*Model, error)
}

type ServiceImpl struct {
	transactor transactor.Transactor
	userRepo   Repository
}

func ProvideService(userRepo Repository, transactor transactor.Transactor) *ServiceImpl {
	return &ServiceImpl{
		userRepo:   userRepo,
		transactor: transactor,
	}
}

func NewService(userRepo Repository, transactor transactor.Transactor) Service {
	return &ServiceImpl{
		userRepo:   userRepo,
		transactor: transactor,
	}
}

func (s ServiceImpl) DoSomething(ctx context.Context) (*Model, error) {
	return &Model{}, nil
}
