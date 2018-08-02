package service

import "context"

// EducationService describes the service.
type EducationService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicEducationService struct{}

func (b *basicEducationService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicEducationService returns a naive, stateless implementation of EducationService.
func NewBasicEducationService() EducationService {
	return &basicEducationService{}
}

// New returns a EducationService with all of the expected middleware wired in.
func New(middleware []Middleware) EducationService {
	var svc EducationService = NewBasicEducationService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
