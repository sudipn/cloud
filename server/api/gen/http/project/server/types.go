// Code generated by goa v3.1.2, DO NOT EDIT.
//
// project HTTP server types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	project "github.com/fieldkit/cloud/server/api/gen/project"
	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "project" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	Body *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
}

// UpdateUnauthorizedResponseBody is the type of the "project" service "update"
// endpoint HTTP response body for the "unauthorized" error.
type UpdateUnauthorizedResponseBody string

// NewUpdateUnauthorizedResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "project" service.
func NewUpdateUnauthorizedResponseBody(res project.Unauthorized) UpdateUnauthorizedResponseBody {
	body := UpdateUnauthorizedResponseBody(res)
	return body
}

// NewUpdatePayload builds a project service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id int64, auth *string) *project.UpdatePayload {
	v := &project.UpdatePayload{
		Body: *body.Body,
	}
	v.ID = id
	v.Auth = auth

	return v
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Body == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("body", "body"))
	}
	return
}