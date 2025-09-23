package user

import "context"

type Service interface {
	//This cotains  business logic like use case in Clean Arch
	DoSomething(ctx context.Context) (*User, error)
}

type service struct {
	userRepo Repository
}

func NewService(userRepo Repository) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s service) DoSomething(ctx context.Context) (*User, error) {
	return s.userRepo.Get(ctx, "")
}
