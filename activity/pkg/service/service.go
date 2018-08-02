package service

import "context"

// ActivityService describes the service.
type ActivityService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicActivityService struct{}

func (b *basicActivityService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	rs = s + " word!"
	return rs, err
}

// NewBasicActivityService returns a naive, stateless implementation of ActivityService.
func NewBasicActivityService() ActivityService {
	return &basicActivityService{}
}

// New returns a ActivityService with all of the expected middleware wired in.
func New(middleware []Middleware) ActivityService {
	var svc ActivityService = NewBasicActivityService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
