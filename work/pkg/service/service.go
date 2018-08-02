package service

import "context"

// WorkService describes the service.
type WorkService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicWorkService struct{}

func (b *basicWorkService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicWorkService returns a naive, stateless implementation of WorkService.
func NewBasicWorkService() WorkService {
	return &basicWorkService{}
}

// New returns a WorkService with all of the expected middleware wired in.
func New(middleware []Middleware) WorkService {
	var svc WorkService = NewBasicWorkService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
