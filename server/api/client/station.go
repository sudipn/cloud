// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": station Resource Client
//
// Command:
// $ main

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// AddStationPath computes a request path to the add action of station.
func AddStationPath() string {

	return fmt.Sprintf("/stations")
}

// Add a station
func (c *Client) AddStation(ctx context.Context, path string, payload *AddStationPayload) (*http.Response, error) {
	req, err := c.NewAddStationRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddStationRequest create the request corresponding to the add action endpoint of the station resource.
func (c *Client) NewAddStationRequest(ctx context.Context, path string, payload *AddStationPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetStationPath computes a request path to the get action of station.
func GetStationPath(station string) string {
	param0 := station

	return fmt.Sprintf("/stations/@/%s", param0)
}

// Get a station
func (c *Client) GetStation(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetStationRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetStationRequest create the request corresponding to the get action endpoint of the station resource.
func (c *Client) NewGetStationRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateStationPath computes a request path to the update action of station.
func UpdateStationPath(stationID int) string {
	param0 := strconv.Itoa(stationID)

	return fmt.Sprintf("/stations/%s", param0)
}

// Update a station
func (c *Client) UpdateStation(ctx context.Context, path string, payload *AddStationPayload) (*http.Response, error) {
	req, err := c.NewUpdateStationRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateStationRequest create the request corresponding to the update action endpoint of the station resource.
func (c *Client) NewUpdateStationRequest(ctx context.Context, path string, payload *AddStationPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
