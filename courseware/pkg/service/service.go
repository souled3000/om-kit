package service

import "context"

// CoursewareService describes the service.
type CoursewareService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicCoursewareService struct{}

func (b *basicCoursewareService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicCoursewareService returns a naive, stateless implementation of CoursewareService.
func NewBasicCoursewareService() CoursewareService {
	return &basicCoursewareService{}
}

// New returns a CoursewareService with all of the expected middleware wired in.
func New(middleware []Middleware) CoursewareService {
	var svc CoursewareService = NewBasicCoursewareService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
