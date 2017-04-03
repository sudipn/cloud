// Code generated by goagen v1.1.0, command line:
// $ main
//
// API "fieldkit": schema Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// AddSchemaPath computes a request path to the add action of schema.
func AddSchemaPath(projectID int) string {
	param0 := strconv.Itoa(projectID)

	return fmt.Sprintf("/projects/%s/schemas", param0)
}

// Add a schema
func (c *Client) AddSchema(ctx context.Context, path string, payload *AddSchemaPayload) (*http.Response, error) {
	req, err := c.NewAddSchemaRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddSchemaRequest create the request corresponding to the add action endpoint of the schema resource.
func (c *Client) NewAddSchemaRequest(ctx context.Context, path string, payload *AddSchemaPayload) (*http.Request, error) {
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
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListSchemaPath computes a request path to the list action of schema.
func ListSchemaPath(project string) string {
	param0 := project

	return fmt.Sprintf("/projects/@/%s/schemas", param0)
}

// List a project's schemas
func (c *Client) ListSchema(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSchemaRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSchemaRequest create the request corresponding to the list action endpoint of the schema resource.
func (c *Client) NewListSchemaRequest(ctx context.Context, path string) (*http.Request, error) {
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
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListIDSchemaPath computes a request path to the list id action of schema.
func ListIDSchemaPath(projectID int) string {
	param0 := strconv.Itoa(projectID)

	return fmt.Sprintf("/projects/%s/schemas", param0)
}

// List a project's schemas
func (c *Client) ListIDSchema(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListIDSchemaRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListIDSchemaRequest create the request corresponding to the list id action endpoint of the schema resource.
func (c *Client) NewListIDSchemaRequest(ctx context.Context, path string) (*http.Request, error) {
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
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// UpdateSchemaPath computes a request path to the update action of schema.
func UpdateSchemaPath(schemaID int) string {
	param0 := strconv.Itoa(schemaID)

	return fmt.Sprintf("/schemas/%s", param0)
}

// Update a schema
func (c *Client) UpdateSchema(ctx context.Context, path string, payload *UpdateSchemaPayload) (*http.Response, error) {
	req, err := c.NewUpdateSchemaRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateSchemaRequest create the request corresponding to the update action endpoint of the schema resource.
func (c *Client) NewUpdateSchemaRequest(ctx context.Context, path string, payload *UpdateSchemaPayload) (*http.Request, error) {
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
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
