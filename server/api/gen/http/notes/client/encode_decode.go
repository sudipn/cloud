// Code generated by goa v3.1.2, DO NOT EDIT.
//
// notes HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	notes "github.com/fieldkit/cloud/server/api/gen/notes"
	notesviews "github.com/fieldkit/cloud/server/api/gen/notes/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "notes" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		stationID int32
	)
	{
		p, ok := v.(*notes.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("notes", "update", "*notes.UpdatePayload", v)
		}
		stationID = p.StationID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateNotesPath(stationID)}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("notes", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the notes update
// server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*notes.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("notes", "update", "*notes.UpdatePayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("notes", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the notes
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUpdateResponse may return the following errors:
//	- "bad-request" (type notes.BadRequest): http.StatusBadRequest
//	- "forbidden" (type notes.Forbidden): http.StatusForbidden
//	- "not-found" (type notes.NotFound): http.StatusNotFound
//	- "unauthorized" (type notes.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "update", err)
			}
			p := NewUpdateFieldNotesOK(&body)
			view := "default"
			vres := &notesviews.FieldNotes{Projected: p, View: view}
			if err = notesviews.ValidateFieldNotes(vres); err != nil {
				return nil, goahttp.ErrValidationError("notes", "update", err)
			}
			res := notes.NewFieldNotes(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body UpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "update", err)
			}
			return nil, NewUpdateBadRequest(body)
		case http.StatusForbidden:
			var (
				body UpdateForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "update", err)
			}
			return nil, NewUpdateForbidden(body)
		case http.StatusNotFound:
			var (
				body UpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "update", err)
			}
			return nil, NewUpdateNotFound(body)
		case http.StatusUnauthorized:
			var (
				body UpdateUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "update", err)
			}
			return nil, NewUpdateUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("notes", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "notes" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		stationID int32
	)
	{
		p, ok := v.(*notes.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("notes", "get", "*notes.GetPayload", v)
		}
		stationID = p.StationID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetNotesPath(stationID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("notes", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the notes get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*notes.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("notes", "get", "*notes.GetPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeGetResponse returns a decoder for responses returned by the notes get
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeGetResponse may return the following errors:
//	- "bad-request" (type notes.BadRequest): http.StatusBadRequest
//	- "forbidden" (type notes.Forbidden): http.StatusForbidden
//	- "not-found" (type notes.NotFound): http.StatusNotFound
//	- "unauthorized" (type notes.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "get", err)
			}
			p := NewGetFieldNotesOK(&body)
			view := "default"
			vres := &notesviews.FieldNotes{Projected: p, View: view}
			if err = notesviews.ValidateFieldNotes(vres); err != nil {
				return nil, goahttp.ErrValidationError("notes", "get", err)
			}
			res := notes.NewFieldNotes(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body GetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "get", err)
			}
			return nil, NewGetBadRequest(body)
		case http.StatusForbidden:
			var (
				body GetForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "get", err)
			}
			return nil, NewGetForbidden(body)
		case http.StatusNotFound:
			var (
				body GetNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "get", err)
			}
			return nil, NewGetNotFound(body)
		case http.StatusUnauthorized:
			var (
				body GetUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "get", err)
			}
			return nil, NewGetUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("notes", "get", resp.StatusCode, string(body))
		}
	}
}

// BuildMediaRequest instantiates a HTTP request object with method and path
// set to call the "notes" service "media" endpoint
func (c *Client) BuildMediaRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		mediaID int32
	)
	{
		p, ok := v.(*notes.MediaPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("notes", "media", "*notes.MediaPayload", v)
		}
		mediaID = p.MediaID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MediaNotesPath(mediaID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("notes", "media", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMediaRequest returns an encoder for requests sent to the notes media
// server.
func EncodeMediaRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*notes.MediaPayload)
		if !ok {
			return goahttp.ErrInvalidType("notes", "media", "*notes.MediaPayload", v)
		}
		values := req.URL.Query()
		values.Add("token", p.Auth)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeMediaResponse returns a decoder for responses returned by the notes
// media endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeMediaResponse may return the following errors:
//	- "bad-request" (type notes.BadRequest): http.StatusBadRequest
//	- "forbidden" (type notes.Forbidden): http.StatusForbidden
//	- "not-found" (type notes.NotFound): http.StatusNotFound
//	- "unauthorized" (type notes.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeMediaResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				length      int64
				contentType string
				err         error
			)
			{
				lengthRaw := resp.Header.Get("Content-Length")
				if lengthRaw == "" {
					return nil, goahttp.ErrValidationError("notes", "media", goa.MissingFieldError("Content-Length", "header"))
				}
				v, err2 := strconv.ParseInt(lengthRaw, 10, 64)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("length", lengthRaw, "integer"))
				}
				length = v
			}
			contentTypeRaw := resp.Header.Get("Content-Type")
			if contentTypeRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Content-Type", "header"))
			}
			contentType = contentTypeRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("notes", "media", err)
			}
			res := NewMediaResultOK(length, contentType)
			return res, nil
		case http.StatusBadRequest:
			var (
				body MediaBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "media", err)
			}
			return nil, NewMediaBadRequest(body)
		case http.StatusForbidden:
			var (
				body MediaForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "media", err)
			}
			return nil, NewMediaForbidden(body)
		case http.StatusNotFound:
			var (
				body MediaNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "media", err)
			}
			return nil, NewMediaNotFound(body)
		case http.StatusUnauthorized:
			var (
				body MediaUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "media", err)
			}
			return nil, NewMediaUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("notes", "media", resp.StatusCode, string(body))
		}
	}
}

// BuildUploadRequest instantiates a HTTP request object with method and path
// set to call the "notes" service "upload" endpoint
func (c *Client) BuildUploadRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		body io.Reader
	)
	rd, ok := v.(*notes.UploadRequestData)
	if !ok {
		return nil, goahttp.ErrInvalidType("notes", "upload", "notes.UploadRequestData", v)
	}
	body = rd.Body
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UploadNotesPath()}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("notes", "upload", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUploadRequest returns an encoder for requests sent to the notes upload
// server.
func EncodeUploadRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		data, ok := v.(*notes.UploadRequestData)
		if !ok {
			return goahttp.ErrInvalidType("notes", "upload", "*notes.UploadRequestData", v)
		}
		p := data.Payload
		{
			head := p.ContentType
			req.Header.Set("Content-Type", head)
		}
		{
			head := p.ContentLength
			headStr := strconv.FormatInt(head, 10)
			req.Header.Set("Content-Length", headStr)
		}
		values := req.URL.Query()
		values.Add("token", p.Auth)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeUploadResponse returns a decoder for responses returned by the notes
// upload endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUploadResponse may return the following errors:
//	- "bad-request" (type notes.BadRequest): http.StatusBadRequest
//	- "forbidden" (type notes.Forbidden): http.StatusForbidden
//	- "not-found" (type notes.NotFound): http.StatusNotFound
//	- "unauthorized" (type notes.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeUploadResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UploadResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "upload", err)
			}
			p := NewUploadNoteMediaOK(&body)
			view := "default"
			vres := &notesviews.NoteMedia{Projected: p, View: view}
			if err = notesviews.ValidateNoteMedia(vres); err != nil {
				return nil, goahttp.ErrValidationError("notes", "upload", err)
			}
			res := notes.NewNoteMedia(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body UploadBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "upload", err)
			}
			return nil, NewUploadBadRequest(body)
		case http.StatusForbidden:
			var (
				body UploadForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "upload", err)
			}
			return nil, NewUploadForbidden(body)
		case http.StatusNotFound:
			var (
				body UploadNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "upload", err)
			}
			return nil, NewUploadNotFound(body)
		case http.StatusUnauthorized:
			var (
				body UploadUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("notes", "upload", err)
			}
			return nil, NewUploadUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("notes", "upload", resp.StatusCode, string(body))
		}
	}
}

// // BuildUploadStreamPayload creates a streaming endpoint request payload from
// the method payload and the path to the file to be streamed
func BuildUploadStreamPayload(payload interface{}, fpath string) (*notes.UploadRequestData, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	return &notes.UploadRequestData{
		Payload: payload.(*notes.UploadPayload),
		Body:    f,
	}, nil
}

// marshalNotesFieldNoteUpdateToFieldNoteUpdateRequestBody builds a value of
// type *FieldNoteUpdateRequestBody from a value of type *notes.FieldNoteUpdate.
func marshalNotesFieldNoteUpdateToFieldNoteUpdateRequestBody(v *notes.FieldNoteUpdate) *FieldNoteUpdateRequestBody {
	res := &FieldNoteUpdateRequestBody{}
	if v.Notes != nil {
		res.Notes = make([]*ExistingFieldNoteRequestBody, len(v.Notes))
		for i, val := range v.Notes {
			res.Notes[i] = marshalNotesExistingFieldNoteToExistingFieldNoteRequestBody(val)
		}
	}
	if v.Creating != nil {
		res.Creating = make([]*NewFieldNoteRequestBody, len(v.Creating))
		for i, val := range v.Creating {
			res.Creating[i] = marshalNotesNewFieldNoteToNewFieldNoteRequestBody(val)
		}
	}

	return res
}

// marshalNotesExistingFieldNoteToExistingFieldNoteRequestBody builds a value
// of type *ExistingFieldNoteRequestBody from a value of type
// *notes.ExistingFieldNote.
func marshalNotesExistingFieldNoteToExistingFieldNoteRequestBody(v *notes.ExistingFieldNote) *ExistingFieldNoteRequestBody {
	res := &ExistingFieldNoteRequestBody{
		ID:      v.ID,
		Key:     v.Key,
		Body:    v.Body,
		MediaID: v.MediaID,
	}

	return res
}

// marshalNotesNewFieldNoteToNewFieldNoteRequestBody builds a value of type
// *NewFieldNoteRequestBody from a value of type *notes.NewFieldNote.
func marshalNotesNewFieldNoteToNewFieldNoteRequestBody(v *notes.NewFieldNote) *NewFieldNoteRequestBody {
	res := &NewFieldNoteRequestBody{
		Key:     v.Key,
		Body:    v.Body,
		MediaID: v.MediaID,
	}

	return res
}

// marshalFieldNoteUpdateRequestBodyToNotesFieldNoteUpdate builds a value of
// type *notes.FieldNoteUpdate from a value of type *FieldNoteUpdateRequestBody.
func marshalFieldNoteUpdateRequestBodyToNotesFieldNoteUpdate(v *FieldNoteUpdateRequestBody) *notes.FieldNoteUpdate {
	res := &notes.FieldNoteUpdate{}
	if v.Notes != nil {
		res.Notes = make([]*notes.ExistingFieldNote, len(v.Notes))
		for i, val := range v.Notes {
			res.Notes[i] = marshalExistingFieldNoteRequestBodyToNotesExistingFieldNote(val)
		}
	}
	if v.Creating != nil {
		res.Creating = make([]*notes.NewFieldNote, len(v.Creating))
		for i, val := range v.Creating {
			res.Creating[i] = marshalNewFieldNoteRequestBodyToNotesNewFieldNote(val)
		}
	}

	return res
}

// marshalExistingFieldNoteRequestBodyToNotesExistingFieldNote builds a value
// of type *notes.ExistingFieldNote from a value of type
// *ExistingFieldNoteRequestBody.
func marshalExistingFieldNoteRequestBodyToNotesExistingFieldNote(v *ExistingFieldNoteRequestBody) *notes.ExistingFieldNote {
	res := &notes.ExistingFieldNote{
		ID:      v.ID,
		Key:     v.Key,
		Body:    v.Body,
		MediaID: v.MediaID,
	}

	return res
}

// marshalNewFieldNoteRequestBodyToNotesNewFieldNote builds a value of type
// *notes.NewFieldNote from a value of type *NewFieldNoteRequestBody.
func marshalNewFieldNoteRequestBodyToNotesNewFieldNote(v *NewFieldNoteRequestBody) *notes.NewFieldNote {
	res := &notes.NewFieldNote{
		Key:     v.Key,
		Body:    v.Body,
		MediaID: v.MediaID,
	}

	return res
}

// unmarshalFieldNoteResponseBodyToNotesviewsFieldNoteView builds a value of
// type *notesviews.FieldNoteView from a value of type *FieldNoteResponseBody.
func unmarshalFieldNoteResponseBodyToNotesviewsFieldNoteView(v *FieldNoteResponseBody) *notesviews.FieldNoteView {
	res := &notesviews.FieldNoteView{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		Key:       v.Key,
		Body:      v.Body,
		MediaID:   v.MediaID,
	}

	return res
}
