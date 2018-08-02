package service

import "context"

// HealthyService describes the service.
type HealthyService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicHealthyService struct{}

func (b *basicHealthyService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicHealthyService returns a naive, stateless implementation of HealthyService.
func NewBasicHealthyService() HealthyService {
	return &basicHealthyService{}
}

// New returns a HealthyService with all of the expected middleware wired in.
func New(middleware []Middleware) HealthyService {
	var svc HealthyService = NewBasicHealthyService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
