package service

import "context"

// MaterialPackService describes the service.
type MaterialPackService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicMaterialPackService struct{}

func (b *basicMaterialPackService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicMaterialPackService returns a naive, stateless implementation of MaterialPackService.
func NewBasicMaterialPackService() MaterialPackService {
	return &basicMaterialPackService{}
}

// New returns a MaterialPackService with all of the expected middleware wired in.
func New(middleware []Middleware) MaterialPackService {
	var svc MaterialPackService = NewBasicMaterialPackService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
