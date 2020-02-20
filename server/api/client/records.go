// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": records Resource Client
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

// DataRecordsPath computes a request path to the data action of records.
func DataRecordsPath(recordID int) string {
	param0 := strconv.Itoa(recordID)

	return fmt.Sprintf("/records/data/%s", param0)
}

// DataRecords makes a request to the data action endpoint of the records resource
func (c *Client) DataRecords(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDataRecordsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDataRecordsRequest create the request corresponding to the data action endpoint of the records resource.
func (c *Client) NewDataRecordsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
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

// MetaRecordsPath computes a request path to the meta action of records.
func MetaRecordsPath(recordID int) string {
	param0 := strconv.Itoa(recordID)

	return fmt.Sprintf("/records/meta/%s", param0)
}

// MetaRecords makes a request to the meta action endpoint of the records resource
func (c *Client) MetaRecords(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewMetaRecordsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewMetaRecordsRequest create the request corresponding to the meta action endpoint of the records resource.
func (c *Client) NewMetaRecordsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
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
