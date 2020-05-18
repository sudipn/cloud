// Code generated by goa v3.1.2, DO NOT EDIT.
//
// ingestion client HTTP transport
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the ingestion service endpoint HTTP clients.
type Client struct {
	// ProcessPending Doer is the HTTP client used to make requests to the process
	// pending endpoint.
	ProcessPendingDoer goahttp.Doer

	// ProcessStation Doer is the HTTP client used to make requests to the process
	// station endpoint.
	ProcessStationDoer goahttp.Doer

	// ProcessIngestion Doer is the HTTP client used to make requests to the
	// process ingestion endpoint.
	ProcessIngestionDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the ingestion service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ProcessPendingDoer:   doer,
		ProcessStationDoer:   doer,
		ProcessIngestionDoer: doer,
		DeleteDoer:           doer,
		CORSDoer:             doer,
		RestoreResponseBody:  restoreBody,
		scheme:               scheme,
		host:                 host,
		decoder:              dec,
		encoder:              enc,
	}
}

// ProcessPending returns an endpoint that makes HTTP requests to the ingestion
// service process pending server.
func (c *Client) ProcessPending() goa.Endpoint {
	var (
		encodeRequest  = EncodeProcessPendingRequest(c.encoder)
		decodeResponse = DecodeProcessPendingResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildProcessPendingRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ProcessPendingDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("ingestion", "process pending", err)
		}
		return decodeResponse(resp)
	}
}

// ProcessStation returns an endpoint that makes HTTP requests to the ingestion
// service process station server.
func (c *Client) ProcessStation() goa.Endpoint {
	var (
		encodeRequest  = EncodeProcessStationRequest(c.encoder)
		decodeResponse = DecodeProcessStationResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildProcessStationRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ProcessStationDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("ingestion", "process station", err)
		}
		return decodeResponse(resp)
	}
}

// ProcessIngestion returns an endpoint that makes HTTP requests to the
// ingestion service process ingestion server.
func (c *Client) ProcessIngestion() goa.Endpoint {
	var (
		encodeRequest  = EncodeProcessIngestionRequest(c.encoder)
		decodeResponse = DecodeProcessIngestionResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildProcessIngestionRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ProcessIngestionDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("ingestion", "process ingestion", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the ingestion service
// delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteRequest(c.encoder)
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("ingestion", "delete", err)
		}
		return decodeResponse(resp)
	}
}
