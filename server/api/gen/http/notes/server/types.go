// Code generated by goa v3.1.2, DO NOT EDIT.
//
// notes HTTP server types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	notes "github.com/fieldkit/cloud/server/api/gen/notes"
	notesviews "github.com/fieldkit/cloud/server/api/gen/notes/views"
	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "notes" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	Notes *FieldNoteUpdateRequestBody `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
}

// UpdateResponseBody is the type of the "notes" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	Notes []*FieldNoteResponseBody `form:"notes" json:"notes" xml:"notes"`
}

// GetResponseBody is the type of the "notes" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	Notes []*FieldNoteResponseBody `form:"notes" json:"notes" xml:"notes"`
}

// UploadResponseBody is the type of the "notes" service "upload" endpoint HTTP
// response body.
type UploadResponseBody struct {
	ID  int64  `form:"id" json:"id" xml:"id"`
	URL string `form:"url" json:"url" xml:"url"`
}

// UpdateBadRequestResponseBody is the type of the "notes" service "update"
// endpoint HTTP response body for the "bad-request" error.
type UpdateBadRequestResponseBody string

// UpdateForbiddenResponseBody is the type of the "notes" service "update"
// endpoint HTTP response body for the "forbidden" error.
type UpdateForbiddenResponseBody string

// UpdateNotFoundResponseBody is the type of the "notes" service "update"
// endpoint HTTP response body for the "not-found" error.
type UpdateNotFoundResponseBody string

// UpdateUnauthorizedResponseBody is the type of the "notes" service "update"
// endpoint HTTP response body for the "unauthorized" error.
type UpdateUnauthorizedResponseBody string

// GetBadRequestResponseBody is the type of the "notes" service "get" endpoint
// HTTP response body for the "bad-request" error.
type GetBadRequestResponseBody string

// GetForbiddenResponseBody is the type of the "notes" service "get" endpoint
// HTTP response body for the "forbidden" error.
type GetForbiddenResponseBody string

// GetNotFoundResponseBody is the type of the "notes" service "get" endpoint
// HTTP response body for the "not-found" error.
type GetNotFoundResponseBody string

// GetUnauthorizedResponseBody is the type of the "notes" service "get"
// endpoint HTTP response body for the "unauthorized" error.
type GetUnauthorizedResponseBody string

// MediaBadRequestResponseBody is the type of the "notes" service "media"
// endpoint HTTP response body for the "bad-request" error.
type MediaBadRequestResponseBody string

// MediaForbiddenResponseBody is the type of the "notes" service "media"
// endpoint HTTP response body for the "forbidden" error.
type MediaForbiddenResponseBody string

// MediaNotFoundResponseBody is the type of the "notes" service "media"
// endpoint HTTP response body for the "not-found" error.
type MediaNotFoundResponseBody string

// MediaUnauthorizedResponseBody is the type of the "notes" service "media"
// endpoint HTTP response body for the "unauthorized" error.
type MediaUnauthorizedResponseBody string

// UploadBadRequestResponseBody is the type of the "notes" service "upload"
// endpoint HTTP response body for the "bad-request" error.
type UploadBadRequestResponseBody string

// UploadForbiddenResponseBody is the type of the "notes" service "upload"
// endpoint HTTP response body for the "forbidden" error.
type UploadForbiddenResponseBody string

// UploadNotFoundResponseBody is the type of the "notes" service "upload"
// endpoint HTTP response body for the "not-found" error.
type UploadNotFoundResponseBody string

// UploadUnauthorizedResponseBody is the type of the "notes" service "upload"
// endpoint HTTP response body for the "unauthorized" error.
type UploadUnauthorizedResponseBody string

// FieldNoteResponseBody is used to define fields on response body types.
type FieldNoteResponseBody struct {
	ID        int64   `form:"id" json:"id" xml:"id"`
	CreatedAt int64   `form:"createdAt" json:"createdAt" xml:"createdAt"`
	Key       *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Body      *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	MediaID   *int64  `form:"mediaId,omitempty" json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

// FieldNoteUpdateRequestBody is used to define fields on request body types.
type FieldNoteUpdateRequestBody struct {
	Notes    []*ExistingFieldNoteRequestBody `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
	Creating []*NewFieldNoteRequestBody      `form:"creating,omitempty" json:"creating,omitempty" xml:"creating,omitempty"`
}

// ExistingFieldNoteRequestBody is used to define fields on request body types.
type ExistingFieldNoteRequestBody struct {
	ID      *int64  `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Key     *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Body    *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	MediaID *int64  `form:"mediaId,omitempty" json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

// NewFieldNoteRequestBody is used to define fields on request body types.
type NewFieldNoteRequestBody struct {
	Key     *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Body    *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	MediaID *int64  `form:"mediaId,omitempty" json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "notes" service.
func NewUpdateResponseBody(res *notesviews.FieldNotesView) *UpdateResponseBody {
	body := &UpdateResponseBody{}
	if res.Notes != nil {
		body.Notes = make([]*FieldNoteResponseBody, len(res.Notes))
		for i, val := range res.Notes {
			body.Notes[i] = marshalNotesviewsFieldNoteViewToFieldNoteResponseBody(val)
		}
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "notes" service.
func NewGetResponseBody(res *notesviews.FieldNotesView) *GetResponseBody {
	body := &GetResponseBody{}
	if res.Notes != nil {
		body.Notes = make([]*FieldNoteResponseBody, len(res.Notes))
		for i, val := range res.Notes {
			body.Notes[i] = marshalNotesviewsFieldNoteViewToFieldNoteResponseBody(val)
		}
	}
	return body
}

// NewUploadResponseBody builds the HTTP response body from the result of the
// "upload" endpoint of the "notes" service.
func NewUploadResponseBody(res *notesviews.NoteMediaView) *UploadResponseBody {
	body := &UploadResponseBody{
		ID:  *res.ID,
		URL: *res.URL,
	}
	return body
}

// NewUpdateBadRequestResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "notes" service.
func NewUpdateBadRequestResponseBody(res notes.BadRequest) UpdateBadRequestResponseBody {
	body := UpdateBadRequestResponseBody(res)
	return body
}

// NewUpdateForbiddenResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "notes" service.
func NewUpdateForbiddenResponseBody(res notes.Forbidden) UpdateForbiddenResponseBody {
	body := UpdateForbiddenResponseBody(res)
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "notes" service.
func NewUpdateNotFoundResponseBody(res notes.NotFound) UpdateNotFoundResponseBody {
	body := UpdateNotFoundResponseBody(res)
	return body
}

// NewUpdateUnauthorizedResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "notes" service.
func NewUpdateUnauthorizedResponseBody(res notes.Unauthorized) UpdateUnauthorizedResponseBody {
	body := UpdateUnauthorizedResponseBody(res)
	return body
}

// NewGetBadRequestResponseBody builds the HTTP response body from the result
// of the "get" endpoint of the "notes" service.
func NewGetBadRequestResponseBody(res notes.BadRequest) GetBadRequestResponseBody {
	body := GetBadRequestResponseBody(res)
	return body
}

// NewGetForbiddenResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "notes" service.
func NewGetForbiddenResponseBody(res notes.Forbidden) GetForbiddenResponseBody {
	body := GetForbiddenResponseBody(res)
	return body
}

// NewGetNotFoundResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "notes" service.
func NewGetNotFoundResponseBody(res notes.NotFound) GetNotFoundResponseBody {
	body := GetNotFoundResponseBody(res)
	return body
}

// NewGetUnauthorizedResponseBody builds the HTTP response body from the result
// of the "get" endpoint of the "notes" service.
func NewGetUnauthorizedResponseBody(res notes.Unauthorized) GetUnauthorizedResponseBody {
	body := GetUnauthorizedResponseBody(res)
	return body
}

// NewMediaBadRequestResponseBody builds the HTTP response body from the result
// of the "media" endpoint of the "notes" service.
func NewMediaBadRequestResponseBody(res notes.BadRequest) MediaBadRequestResponseBody {
	body := MediaBadRequestResponseBody(res)
	return body
}

// NewMediaForbiddenResponseBody builds the HTTP response body from the result
// of the "media" endpoint of the "notes" service.
func NewMediaForbiddenResponseBody(res notes.Forbidden) MediaForbiddenResponseBody {
	body := MediaForbiddenResponseBody(res)
	return body
}

// NewMediaNotFoundResponseBody builds the HTTP response body from the result
// of the "media" endpoint of the "notes" service.
func NewMediaNotFoundResponseBody(res notes.NotFound) MediaNotFoundResponseBody {
	body := MediaNotFoundResponseBody(res)
	return body
}

// NewMediaUnauthorizedResponseBody builds the HTTP response body from the
// result of the "media" endpoint of the "notes" service.
func NewMediaUnauthorizedResponseBody(res notes.Unauthorized) MediaUnauthorizedResponseBody {
	body := MediaUnauthorizedResponseBody(res)
	return body
}

// NewUploadBadRequestResponseBody builds the HTTP response body from the
// result of the "upload" endpoint of the "notes" service.
func NewUploadBadRequestResponseBody(res notes.BadRequest) UploadBadRequestResponseBody {
	body := UploadBadRequestResponseBody(res)
	return body
}

// NewUploadForbiddenResponseBody builds the HTTP response body from the result
// of the "upload" endpoint of the "notes" service.
func NewUploadForbiddenResponseBody(res notes.Forbidden) UploadForbiddenResponseBody {
	body := UploadForbiddenResponseBody(res)
	return body
}

// NewUploadNotFoundResponseBody builds the HTTP response body from the result
// of the "upload" endpoint of the "notes" service.
func NewUploadNotFoundResponseBody(res notes.NotFound) UploadNotFoundResponseBody {
	body := UploadNotFoundResponseBody(res)
	return body
}

// NewUploadUnauthorizedResponseBody builds the HTTP response body from the
// result of the "upload" endpoint of the "notes" service.
func NewUploadUnauthorizedResponseBody(res notes.Unauthorized) UploadUnauthorizedResponseBody {
	body := UploadUnauthorizedResponseBody(res)
	return body
}

// NewUpdatePayload builds a notes service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, stationID int32, auth string) *notes.UpdatePayload {
	v := &notes.UpdatePayload{}
	v.Notes = unmarshalFieldNoteUpdateRequestBodyToNotesFieldNoteUpdate(body.Notes)
	v.StationID = stationID
	v.Auth = auth

	return v
}

// NewGetPayload builds a notes service get endpoint payload.
func NewGetPayload(stationID int32, auth string) *notes.GetPayload {
	v := &notes.GetPayload{}
	v.StationID = stationID
	v.Auth = auth

	return v
}

// NewMediaPayload builds a notes service media endpoint payload.
func NewMediaPayload(mediaID int32, auth string) *notes.MediaPayload {
	v := &notes.MediaPayload{}
	v.MediaID = mediaID
	v.Auth = auth

	return v
}

// NewUploadPayload builds a notes service upload endpoint payload.
func NewUploadPayload(auth string, contentType string, contentLength int64) *notes.UploadPayload {
	v := &notes.UploadPayload{}
	v.Auth = auth
	v.ContentType = contentType
	v.ContentLength = contentLength

	return v
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Notes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("notes", "body"))
	}
	if body.Notes != nil {
		if err2 := ValidateFieldNoteUpdateRequestBody(body.Notes); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateFieldNoteUpdateRequestBody runs the validations defined on
// FieldNoteUpdateRequestBody
func ValidateFieldNoteUpdateRequestBody(body *FieldNoteUpdateRequestBody) (err error) {
	if body.Notes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("notes", "body"))
	}
	if body.Creating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("creating", "body"))
	}
	for _, e := range body.Notes {
		if e != nil {
			if err2 := ValidateExistingFieldNoteRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateExistingFieldNoteRequestBody runs the validations defined on
// ExistingFieldNoteRequestBody
func ValidateExistingFieldNoteRequestBody(body *ExistingFieldNoteRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}
