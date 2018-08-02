package service

import "context"

// FoodService describes the service.
type FoodService interface {
	// Add your methods here
	Foo(ctx context.Context, s string) (rs string, err error)
}

type basicFoodService struct{}

func (b *basicFoodService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	return rs, err
}

// NewBasicFoodService returns a naive, stateless implementation of FoodService.
func NewBasicFoodService() FoodService {
	return &basicFoodService{}
}

// New returns a FoodService with all of the expected middleware wired in.
func New(middleware []Middleware) FoodService {
	var svc FoodService = NewBasicFoodService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
