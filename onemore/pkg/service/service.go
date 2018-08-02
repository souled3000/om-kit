package service

import "context"

// OnemoreService describes the service.
type OnemoreService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicOnemoreService struct{}

func (b *basicOnemoreService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicOnemoreService returns a naive, stateless implementation of OnemoreService.
func NewBasicOnemoreService() OnemoreService {
	return &basicOnemoreService{}
}

// New returns a OnemoreService with all of the expected middleware wired in.
func New(middleware []Middleware) OnemoreService {
	var svc OnemoreService = NewBasicOnemoreService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
