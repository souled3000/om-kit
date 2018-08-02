package service

import "context"

// AttendanceService describes the service.
type AttendanceService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicAttendanceService struct{}

func (b *basicAttendanceService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicAttendanceService returns a naive, stateless implementation of AttendanceService.
func NewBasicAttendanceService() AttendanceService {
	return &basicAttendanceService{}
}

// New returns a AttendanceService with all of the expected middleware wired in.
func New(middleware []Middleware) AttendanceService {
	var svc AttendanceService = NewBasicAttendanceService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
