package service

import "context"

// OmsService describes the service.
type OmsService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicOmsService struct{}

func (b *basicOmsService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicOmsService returns a naive, stateless implementation of OmsService.
func NewBasicOmsService() OmsService {
	return &basicOmsService{}
}

// New returns a OmsService with all of the expected middleware wired in.
func New(middleware []Middleware) OmsService {
	var svc OmsService = NewBasicOmsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
