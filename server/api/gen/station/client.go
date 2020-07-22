// Code generated by goa v3.1.2, DO NOT EDIT.
//
// station client
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package station

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "station" service client.
type Client struct {
	AddEndpoint         goa.Endpoint
	GetEndpoint         goa.Endpoint
	UpdateEndpoint      goa.Endpoint
	ListMineEndpoint    goa.Endpoint
	ListProjectEndpoint goa.Endpoint
	PhotoEndpoint       goa.Endpoint
	ListAllEndpoint     goa.Endpoint
}

// NewClient initializes a "station" service client given the endpoints.
func NewClient(add, get, update, listMine, listProject, photo, listAll goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:         add,
		GetEndpoint:         get,
		UpdateEndpoint:      update,
		ListMineEndpoint:    listMine,
		ListProjectEndpoint: listProject,
		PhotoEndpoint:       photo,
		ListAllEndpoint:     listAll,
	}
}

// Add calls the "add" endpoint of the "station" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res *StationFull, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StationFull), nil
}

// Get calls the "get" endpoint of the "station" service.
func (c *Client) Get(ctx context.Context, p *GetPayload) (res *StationFull, err error) {
	var ires interface{}
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StationFull), nil
}

// Update calls the "update" endpoint of the "station" service.
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (res *StationFull, err error) {
	var ires interface{}
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StationFull), nil
}

// ListMine calls the "list mine" endpoint of the "station" service.
func (c *Client) ListMine(ctx context.Context, p *ListMinePayload) (res *StationsFull, err error) {
	var ires interface{}
	ires, err = c.ListMineEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StationsFull), nil
}

// ListProject calls the "list project" endpoint of the "station" service.
func (c *Client) ListProject(ctx context.Context, p *ListProjectPayload) (res *StationsFull, err error) {
	var ires interface{}
	ires, err = c.ListProjectEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StationsFull), nil
}

// Photo calls the "photo" endpoint of the "station" service.
func (c *Client) Photo(ctx context.Context, p *PhotoPayload) (res *PhotoResult, resp io.ReadCloser, err error) {
	var ires interface{}
	ires, err = c.PhotoEndpoint(ctx, p)
	if err != nil {
		return
	}
	o := ires.(*PhotoResponseData)
	return o.Result, o.Body, nil
}

// ListAll calls the "list all" endpoint of the "station" service.
func (c *Client) ListAll(ctx context.Context, p *ListAllPayload) (res *PageOfStations, err error) {
	var ires interface{}
	ires, err = c.ListAllEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PageOfStations), nil
}
