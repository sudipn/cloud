// Code generated by goa v3.1.2, DO NOT EDIT.
//
// project HTTP client types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	project "github.com/fieldkit/cloud/server/api/gen/project"
	projectviews "github.com/fieldkit/cloud/server/api/gen/project/views"
	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "project" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	Body string `form:"body" json:"body" xml:"body"`
}

// InvitesResponseBody is the type of the "project" service "invites" endpoint
// HTTP response body.
type InvitesResponseBody struct {
	Pending []*PendingInviteResponseBody `form:"pending,omitempty" json:"pending,omitempty" xml:"pending,omitempty"`
}

// LookupInviteResponseBody is the type of the "project" service "lookup
// invite" endpoint HTTP response body.
type LookupInviteResponseBody struct {
	Pending []*PendingInviteResponseBody `form:"pending,omitempty" json:"pending,omitempty" xml:"pending,omitempty"`
}

// UpdateNotFoundResponseBody is the type of the "project" service "update"
// endpoint HTTP response body for the "not-found" error.
type UpdateNotFoundResponseBody string

// UpdateUnauthorizedResponseBody is the type of the "project" service "update"
// endpoint HTTP response body for the "unauthorized" error.
type UpdateUnauthorizedResponseBody string

// InvitesNotFoundResponseBody is the type of the "project" service "invites"
// endpoint HTTP response body for the "not-found" error.
type InvitesNotFoundResponseBody string

// InvitesUnauthorizedResponseBody is the type of the "project" service
// "invites" endpoint HTTP response body for the "unauthorized" error.
type InvitesUnauthorizedResponseBody string

// LookupInviteNotFoundResponseBody is the type of the "project" service
// "lookup invite" endpoint HTTP response body for the "not-found" error.
type LookupInviteNotFoundResponseBody string

// LookupInviteUnauthorizedResponseBody is the type of the "project" service
// "lookup invite" endpoint HTTP response body for the "unauthorized" error.
type LookupInviteUnauthorizedResponseBody string

// AcceptInviteNotFoundResponseBody is the type of the "project" service
// "accept invite" endpoint HTTP response body for the "not-found" error.
type AcceptInviteNotFoundResponseBody string

// AcceptInviteUnauthorizedResponseBody is the type of the "project" service
// "accept invite" endpoint HTTP response body for the "unauthorized" error.
type AcceptInviteUnauthorizedResponseBody string

// RejectInviteNotFoundResponseBody is the type of the "project" service
// "reject invite" endpoint HTTP response body for the "not-found" error.
type RejectInviteNotFoundResponseBody string

// RejectInviteUnauthorizedResponseBody is the type of the "project" service
// "reject invite" endpoint HTTP response body for the "unauthorized" error.
type RejectInviteUnauthorizedResponseBody string

// PendingInviteResponseBody is used to define fields on response body types.
type PendingInviteResponseBody struct {
	ID      *int64                      `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Project *ProjectSummaryResponseBody `form:"project,omitempty" json:"project,omitempty" xml:"project,omitempty"`
	Time    *int64                      `form:"time,omitempty" json:"time,omitempty" xml:"time,omitempty"`
}

// ProjectSummaryResponseBody is used to define fields on response body types.
type ProjectSummaryResponseBody struct {
	ID   *int64  `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "project" service.
func NewUpdateRequestBody(p *project.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Body: p.Body,
	}
	return body
}

// NewUpdateNotFound builds a project service update endpoint not-found error.
func NewUpdateNotFound(body UpdateNotFoundResponseBody) project.NotFound {
	v := project.NotFound(body)
	return v
}

// NewUpdateUnauthorized builds a project service update endpoint unauthorized
// error.
func NewUpdateUnauthorized(body UpdateUnauthorizedResponseBody) project.Unauthorized {
	v := project.Unauthorized(body)
	return v
}

// NewInvitesPendingInvitesOK builds a "project" service "invites" endpoint
// result from a HTTP "OK" response.
func NewInvitesPendingInvitesOK(body *InvitesResponseBody) *projectviews.PendingInvitesView {
	v := &projectviews.PendingInvitesView{}
	v.Pending = make([]*projectviews.PendingInviteView, len(body.Pending))
	for i, val := range body.Pending {
		v.Pending[i] = unmarshalPendingInviteResponseBodyToProjectviewsPendingInviteView(val)
	}

	return v
}

// NewInvitesNotFound builds a project service invites endpoint not-found error.
func NewInvitesNotFound(body InvitesNotFoundResponseBody) project.NotFound {
	v := project.NotFound(body)
	return v
}

// NewInvitesUnauthorized builds a project service invites endpoint
// unauthorized error.
func NewInvitesUnauthorized(body InvitesUnauthorizedResponseBody) project.Unauthorized {
	v := project.Unauthorized(body)
	return v
}

// NewLookupInvitePendingInvitesOK builds a "project" service "lookup invite"
// endpoint result from a HTTP "OK" response.
func NewLookupInvitePendingInvitesOK(body *LookupInviteResponseBody) *projectviews.PendingInvitesView {
	v := &projectviews.PendingInvitesView{}
	v.Pending = make([]*projectviews.PendingInviteView, len(body.Pending))
	for i, val := range body.Pending {
		v.Pending[i] = unmarshalPendingInviteResponseBodyToProjectviewsPendingInviteView(val)
	}

	return v
}

// NewLookupInviteNotFound builds a project service lookup invite endpoint
// not-found error.
func NewLookupInviteNotFound(body LookupInviteNotFoundResponseBody) project.NotFound {
	v := project.NotFound(body)
	return v
}

// NewLookupInviteUnauthorized builds a project service lookup invite endpoint
// unauthorized error.
func NewLookupInviteUnauthorized(body LookupInviteUnauthorizedResponseBody) project.Unauthorized {
	v := project.Unauthorized(body)
	return v
}

// NewAcceptInviteNotFound builds a project service accept invite endpoint
// not-found error.
func NewAcceptInviteNotFound(body AcceptInviteNotFoundResponseBody) project.NotFound {
	v := project.NotFound(body)
	return v
}

// NewAcceptInviteUnauthorized builds a project service accept invite endpoint
// unauthorized error.
func NewAcceptInviteUnauthorized(body AcceptInviteUnauthorizedResponseBody) project.Unauthorized {
	v := project.Unauthorized(body)
	return v
}

// NewRejectInviteNotFound builds a project service reject invite endpoint
// not-found error.
func NewRejectInviteNotFound(body RejectInviteNotFoundResponseBody) project.NotFound {
	v := project.NotFound(body)
	return v
}

// NewRejectInviteUnauthorized builds a project service reject invite endpoint
// unauthorized error.
func NewRejectInviteUnauthorized(body RejectInviteUnauthorizedResponseBody) project.Unauthorized {
	v := project.Unauthorized(body)
	return v
}

// ValidatePendingInviteResponseBody runs the validations defined on
// PendingInviteResponseBody
func ValidatePendingInviteResponseBody(body *PendingInviteResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Project == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("project", "body"))
	}
	if body.Time == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("time", "body"))
	}
	if body.Project != nil {
		if err2 := ValidateProjectSummaryResponseBody(body.Project); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProjectSummaryResponseBody runs the validations defined on
// ProjectSummaryResponseBody
func ValidateProjectSummaryResponseBody(body *ProjectSummaryResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}
