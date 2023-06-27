// Code generated by goa v3.11.3, DO NOT EDIT.
//
// sop endpoints
//
// Command:
// $ goa gen stocktrader/design

package sop

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "sop" service endpoints.
type Endpoints struct {
	Plan goa.Endpoint
}

// NewEndpoints wraps the methods of the "sop" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Plan: NewPlanEndpoint(s),
	}
}

// Use applies the given middleware to all the "sop" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Plan = m(e.Plan)
}

// NewPlanEndpoint returns an endpoint function that calls the method "plan" of
// service "sop".
func NewPlanEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*PlanPayload)
		return s.Plan(ctx, p)
	}
}