package service

import "context"

// KindergartenService describes the service.
type KindergartenService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicKindergartenService struct{}

func (b *basicKindergartenService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicKindergartenService returns a naive, stateless implementation of KindergartenService.
func NewBasicKindergartenService() KindergartenService {
	return &basicKindergartenService{}
}

// New returns a KindergartenService with all of the expected middleware wired in.
func New(middleware []Middleware) KindergartenService {
	var svc KindergartenService = NewBasicKindergartenService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
