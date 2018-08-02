package service

import "context"

// MaterialService describes the service.
type MaterialService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicMaterialService struct{}

func (b *basicMaterialService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicMaterialService returns a naive, stateless implementation of MaterialService.
func NewBasicMaterialService() MaterialService {
	return &basicMaterialService{}
}

// New returns a MaterialService with all of the expected middleware wired in.
func New(middleware []Middleware) MaterialService {
	var svc MaterialService = NewBasicMaterialService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
