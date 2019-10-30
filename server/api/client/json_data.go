// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": jsonData Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetJSONDataPath computes a request path to the get action of jsonData.
func GetJSONDataPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/data/devices/%s/data/json", param0)
}

// Retrieve data
func (c *Client) GetJSONData(ctx context.Context, path string, end *int, internal *bool, pageNumber *int, pageSize *int, start *int) (*http.Response, error) {
	req, err := c.NewGetJSONDataRequest(ctx, path, end, internal, pageNumber, pageSize, start)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetJSONDataRequest create the request corresponding to the get action endpoint of the jsonData resource.
func (c *Client) NewGetJSONDataRequest(ctx context.Context, path string, end *int, internal *bool, pageNumber *int, pageSize *int, start *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if end != nil {
		tmp234 := strconv.Itoa(*end)
		values.Set("end", tmp234)
	}
	if internal != nil {
		tmp235 := strconv.FormatBool(*internal)
		values.Set("internal", tmp235)
	}
	if pageNumber != nil {
		tmp236 := strconv.Itoa(*pageNumber)
		values.Set("pageNumber", tmp236)
	}
	if pageSize != nil {
		tmp237 := strconv.Itoa(*pageSize)
		values.Set("pageSize", tmp237)
	}
	if start != nil {
		tmp238 := strconv.Itoa(*start)
		values.Set("start", tmp238)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetLinesJSONDataPath computes a request path to the get lines action of jsonData.
func GetLinesJSONDataPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/data/devices/%s/data/lines", param0)
}

// Retrieve data
func (c *Client) GetLinesJSONData(ctx context.Context, path string, end *int, internal *bool, pageNumber *int, pageSize *int, start *int) (*http.Response, error) {
	req, err := c.NewGetLinesJSONDataRequest(ctx, path, end, internal, pageNumber, pageSize, start)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetLinesJSONDataRequest create the request corresponding to the get lines action endpoint of the jsonData resource.
func (c *Client) NewGetLinesJSONDataRequest(ctx context.Context, path string, end *int, internal *bool, pageNumber *int, pageSize *int, start *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if end != nil {
		tmp239 := strconv.Itoa(*end)
		values.Set("end", tmp239)
	}
	if internal != nil {
		tmp240 := strconv.FormatBool(*internal)
		values.Set("internal", tmp240)
	}
	if pageNumber != nil {
		tmp241 := strconv.Itoa(*pageNumber)
		values.Set("pageNumber", tmp241)
	}
	if pageSize != nil {
		tmp242 := strconv.Itoa(*pageSize)
		values.Set("pageSize", tmp242)
	}
	if start != nil {
		tmp243 := strconv.Itoa(*start)
		values.Set("start", tmp243)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
