package service

import "context"

// MqttService describes the service.
type MqttService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicMqttService struct{}

func (b *basicMqttService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicMqttService returns a naive, stateless implementation of MqttService.
func NewBasicMqttService() MqttService {
	return &basicMqttService{}
}

// New returns a MqttService with all of the expected middleware wired in.
func New(middleware []Middleware) MqttService {
	var svc MqttService = NewBasicMqttService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
