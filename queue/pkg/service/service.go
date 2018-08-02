package service

import "context"

// QueueService describes the service.
type QueueService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicQueueService struct{}

func (b *basicQueueService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicQueueService returns a naive, stateless implementation of QueueService.
func NewBasicQueueService() QueueService {
	return &basicQueueService{}
}

// New returns a QueueService with all of the expected middleware wired in.
func New(middleware []Middleware) QueueService {
	var svc QueueService = NewBasicQueueService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
