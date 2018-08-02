package service

import "context"

// GroupService describes the service.
type GroupService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicGroupService struct{}

func (b *basicGroupService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicGroupService returns a naive, stateless implementation of GroupService.
func NewBasicGroupService() GroupService {
	return &basicGroupService{}
}

// New returns a GroupService with all of the expected middleware wired in.
func New(middleware []Middleware) GroupService {
	var svc GroupService = NewBasicGroupService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
