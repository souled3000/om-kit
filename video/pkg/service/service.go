package service

import "context"

// VideoService describes the service.
type VideoService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicVideoService struct{}

func (b *basicVideoService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicVideoService returns a naive, stateless implementation of VideoService.
func NewBasicVideoService() VideoService {
	return &basicVideoService{}
}

// New returns a VideoService with all of the expected middleware wired in.
func New(middleware []Middleware) VideoService {
	var svc VideoService = NewBasicVideoService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
