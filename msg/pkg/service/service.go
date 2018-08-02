package service

import "context"

// MsgService describes the service.
type MsgService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicMsgService struct{}

func (b *basicMsgService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicMsgService returns a naive, stateless implementation of MsgService.
func NewBasicMsgService() MsgService {
	return &basicMsgService{}
}

// New returns a MsgService with all of the expected middleware wired in.
func New(middleware []Middleware) MsgService {
	var svc MsgService = NewBasicMsgService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
