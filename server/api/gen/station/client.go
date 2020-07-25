// Code generated by goa v3.1.2, DO NOT EDIT.
//
// station client
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package station

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "station" service client.
type Client struct {
	AddEndpoint           goa.Endpoint
	GetEndpoint           goa.Endpoint
	UpdateEndpoint        goa.Endpoint
	ListMineEndpoint      goa.Endpoint
	ListProjectEndpoint   goa.Endpoint
	DownloadPhotoEndpoint goa.Endpoint
	ListAllEndpoint       goa.Endpoint
	DeleteEndpoint        goa.Endpoint
}

// NewClient initializes a "station" service client given the endpoints.
func NewClient(add, get, update, listMine, listProject, downloadPhoto, listAll, delete_ goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:           add,
		GetEndpoint:           get,
		UpdateEndpoint:        update,
		ListMineEndpoint:      listMine,
		ListProjectEndpoint:   listProject,
		DownloadPhotoEndpoint: downloadPhoto,
		ListAllEndpoint:       listAll,
		DeleteEndpoint:        delete_,
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

// DownloadPhoto calls the "download photo" endpoint of the "station" service.
func (c *Client) DownloadPhoto(ctx context.Context, p *DownloadPhotoPayload) (res *DownloadedPhoto, err error) {
	var ires interface{}
	ires, err = c.DownloadPhotoEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DownloadedPhoto), nil
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

// Delete calls the "delete" endpoint of the "station" service.
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}
