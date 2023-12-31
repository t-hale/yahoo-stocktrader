// Code generated by goa v3.11.3, DO NOT EDIT.
//
// sop client HTTP transport
//
// Command:
// $ goa gen stocktrader/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the sop service endpoint HTTP clients.
type Client struct {
	// Plan Doer is the HTTP client used to make requests to the plan endpoint.
	PlanDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the sop service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		PlanDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Plan returns an endpoint that makes HTTP requests to the sop service plan
// server.
func (c *Client) Plan() goa.Endpoint {
	var (
		decodeResponse = DecodePlanResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildPlanRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PlanDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("sop", "plan", err)
		}
		return decodeResponse(resp)
	}
}
