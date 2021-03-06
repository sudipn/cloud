// Code generated by goa v3.1.2, DO NOT EDIT.
//
// information client
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package information

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "information" service client.
type Client struct {
	DeviceLayoutEndpoint       goa.Endpoint
	FirmwareStatisticsEndpoint goa.Endpoint
}

// NewClient initializes a "information" service client given the endpoints.
func NewClient(deviceLayout, firmwareStatistics goa.Endpoint) *Client {
	return &Client{
		DeviceLayoutEndpoint:       deviceLayout,
		FirmwareStatisticsEndpoint: firmwareStatistics,
	}
}

// DeviceLayout calls the "device layout" endpoint of the "information" service.
func (c *Client) DeviceLayout(ctx context.Context, p *DeviceLayoutPayload) (res *DeviceLayoutResponse, err error) {
	var ires interface{}
	ires, err = c.DeviceLayoutEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DeviceLayoutResponse), nil
}

// FirmwareStatistics calls the "firmware statistics" endpoint of the
// "information" service.
func (c *Client) FirmwareStatistics(ctx context.Context, p *FirmwareStatisticsPayload) (res *FirmwareStatisticsResult, err error) {
	var ires interface{}
	ires, err = c.FirmwareStatisticsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*FirmwareStatisticsResult), nil
}
