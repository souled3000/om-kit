package service

import "context"

// ReportService describes the service.
type ReportService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicReportService struct{}

func (b *basicReportService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicReportService returns a naive, stateless implementation of ReportService.
func NewBasicReportService() ReportService {
	return &basicReportService{}
}

// New returns a ReportService with all of the expected middleware wired in.
func New(middleware []Middleware) ReportService {
	var svc ReportService = NewBasicReportService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
