package service

import "context"

// SeedlingService describes the service.
type SeedlingService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicSeedlingService struct{}

func (b *basicSeedlingService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicSeedlingService returns a naive, stateless implementation of SeedlingService.
func NewBasicSeedlingService() SeedlingService {
	return &basicSeedlingService{}
}

// New returns a SeedlingService with all of the expected middleware wired in.
func New(middleware []Middleware) SeedlingService {
	var svc SeedlingService = NewBasicSeedlingService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
