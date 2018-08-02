// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "om-kit/food/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	FooEndpoint endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.FoodService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{FooEndpoint: MakeFooEndpoint(s)}
	for _, m := range mdw["Foo"] {
		eps.FooEndpoint = m(eps.FooEndpoint)
	}
	return eps
}
