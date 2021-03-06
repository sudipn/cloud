// Code generated by goa v3.1.2, DO NOT EDIT.
//
// export client
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package export

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "export" service client.
type Client struct {
	ListMineEndpoint  goa.Endpoint
	StatusEndpoint    goa.Endpoint
	DownloadEndpoint  goa.Endpoint
	CsvEndpoint       goa.Endpoint
	JSONLinesEndpoint goa.Endpoint
}

// NewClient initializes a "export" service client given the endpoints.
func NewClient(listMine, status, download, csv, jSONLines goa.Endpoint) *Client {
	return &Client{
		ListMineEndpoint:  listMine,
		StatusEndpoint:    status,
		DownloadEndpoint:  download,
		CsvEndpoint:       csv,
		JSONLinesEndpoint: jSONLines,
	}
}

// ListMine calls the "list mine" endpoint of the "export" service.
func (c *Client) ListMine(ctx context.Context, p *ListMinePayload) (res *UserExports, err error) {
	var ires interface{}
	ires, err = c.ListMineEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UserExports), nil
}

// Status calls the "status" endpoint of the "export" service.
func (c *Client) Status(ctx context.Context, p *StatusPayload) (res *ExportStatus, err error) {
	var ires interface{}
	ires, err = c.StatusEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ExportStatus), nil
}

// Download calls the "download" endpoint of the "export" service.
func (c *Client) Download(ctx context.Context, p *DownloadPayload) (res *DownloadResult, resp io.ReadCloser, err error) {
	var ires interface{}
	ires, err = c.DownloadEndpoint(ctx, p)
	if err != nil {
		return
	}
	o := ires.(*DownloadResponseData)
	return o.Result, o.Body, nil
}

// Csv calls the "csv" endpoint of the "export" service.
func (c *Client) Csv(ctx context.Context, p *CsvPayload) (res *CsvResult, err error) {
	var ires interface{}
	ires, err = c.CsvEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CsvResult), nil
}

// JSONLines calls the "json lines" endpoint of the "export" service.
func (c *Client) JSONLines(ctx context.Context, p *JSONLinesPayload) (res *JSONLinesResult, err error) {
	var ires interface{}
	ires, err = c.JSONLinesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*JSONLinesResult), nil
}
