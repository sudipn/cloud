// Code generated by goa v3.1.2, DO NOT EDIT.
//
// notes HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"

	notes "github.com/fieldkit/cloud/server/api/gen/notes"
	notesviews "github.com/fieldkit/cloud/server/api/gen/notes/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUpdateResponse returns an encoder for responses returned by the notes
// update endpoint.
func EncodeUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*notesviews.FieldNotes)
		enc := encoder(ctx, w)
		body := NewUpdateResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateRequest returns a decoder for requests sent to the notes update
// endpoint.
func DecodeUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			stationID int32
			auth      string

			params = mux.Vars(r)
		)
		{
			stationIDRaw := params["stationId"]
			v, err2 := strconv.ParseInt(stationIDRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("stationID", stationIDRaw, "integer"))
			}
			stationID = int32(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdatePayload(&body, stationID, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeUpdateError returns an encoder for errors returned by the update notes
// endpoint.
func EncodeUpdateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(notes.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			res := v.(notes.Forbidden)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(notes.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(notes.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetResponse returns an encoder for responses returned by the notes get
// endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*notesviews.FieldNotes)
		enc := encoder(ctx, w)
		body := NewGetResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRequest returns a decoder for requests sent to the notes get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			stationID int32
			auth      string
			err       error

			params = mux.Vars(r)
		)
		{
			stationIDRaw := params["stationId"]
			v, err2 := strconv.ParseInt(stationIDRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("stationID", stationIDRaw, "integer"))
			}
			stationID = int32(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetPayload(stationID, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeGetError returns an encoder for errors returned by the get notes
// endpoint.
func EncodeGetError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(notes.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			res := v.(notes.Forbidden)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(notes.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(notes.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeMediaResponse returns an encoder for responses returned by the notes
// media endpoint.
func EncodeMediaResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*notes.MediaResult)
		val := res.Length
		lengths := strconv.FormatInt(val, 10)
		w.Header().Set("Content-Length", lengths)
		w.Header().Set("Content-Type", res.ContentType)
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeMediaRequest returns a decoder for requests sent to the notes media
// endpoint.
func DecodeMediaRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			mediaID int32
			auth    string
			err     error

			params = mux.Vars(r)
		)
		{
			mediaIDRaw := params["mediaId"]
			v, err2 := strconv.ParseInt(mediaIDRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("mediaID", mediaIDRaw, "integer"))
			}
			mediaID = int32(v)
		}
		auth = r.URL.Query().Get("token")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "query string"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewMediaPayload(mediaID, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeMediaError returns an encoder for errors returned by the media notes
// endpoint.
func EncodeMediaError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(notes.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMediaBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			res := v.(notes.Forbidden)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMediaForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(notes.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMediaNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(notes.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMediaUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUploadResponse returns an encoder for responses returned by the notes
// upload endpoint.
func EncodeUploadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*notesviews.NoteMedia)
		enc := encoder(ctx, w)
		body := NewUploadResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUploadRequest returns a decoder for requests sent to the notes upload
// endpoint.
func DecodeUploadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			auth          string
			contentType   string
			contentLength int64
			err           error
		)
		auth = r.URL.Query().Get("token")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "query string"))
		}
		contentType = r.Header.Get("Content-Type")
		if contentType == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Content-Type", "header"))
		}
		{
			contentLengthRaw := r.Header.Get("Content-Length")
			if contentLengthRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Content-Length", "header"))
			}
			v, err2 := strconv.ParseInt(contentLengthRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("contentLength", contentLengthRaw, "integer"))
			}
			contentLength = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewUploadPayload(auth, contentType, contentLength)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeUploadError returns an encoder for errors returned by the upload notes
// endpoint.
func EncodeUploadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(notes.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			res := v.(notes.Forbidden)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(notes.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(notes.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// unmarshalFieldNoteUpdateRequestBodyToNotesFieldNoteUpdate builds a value of
// type *notes.FieldNoteUpdate from a value of type *FieldNoteUpdateRequestBody.
func unmarshalFieldNoteUpdateRequestBodyToNotesFieldNoteUpdate(v *FieldNoteUpdateRequestBody) *notes.FieldNoteUpdate {
	res := &notes.FieldNoteUpdate{}
	res.Notes = make([]*notes.ExistingFieldNote, len(v.Notes))
	for i, val := range v.Notes {
		res.Notes[i] = unmarshalExistingFieldNoteRequestBodyToNotesExistingFieldNote(val)
	}
	res.Creating = make([]*notes.NewFieldNote, len(v.Creating))
	for i, val := range v.Creating {
		res.Creating[i] = unmarshalNewFieldNoteRequestBodyToNotesNewFieldNote(val)
	}

	return res
}

// unmarshalExistingFieldNoteRequestBodyToNotesExistingFieldNote builds a value
// of type *notes.ExistingFieldNote from a value of type
// *ExistingFieldNoteRequestBody.
func unmarshalExistingFieldNoteRequestBodyToNotesExistingFieldNote(v *ExistingFieldNoteRequestBody) *notes.ExistingFieldNote {
	res := &notes.ExistingFieldNote{
		ID:      *v.ID,
		Key:     v.Key,
		Body:    v.Body,
		MediaID: v.MediaID,
	}

	return res
}

// unmarshalNewFieldNoteRequestBodyToNotesNewFieldNote builds a value of type
// *notes.NewFieldNote from a value of type *NewFieldNoteRequestBody.
func unmarshalNewFieldNoteRequestBodyToNotesNewFieldNote(v *NewFieldNoteRequestBody) *notes.NewFieldNote {
	res := &notes.NewFieldNote{
		Key:     v.Key,
		Body:    v.Body,
		MediaID: v.MediaID,
	}

	return res
}

// marshalNotesviewsFieldNoteViewToFieldNoteResponseBody builds a value of type
// *FieldNoteResponseBody from a value of type *notesviews.FieldNoteView.
func marshalNotesviewsFieldNoteViewToFieldNoteResponseBody(v *notesviews.FieldNoteView) *FieldNoteResponseBody {
	res := &FieldNoteResponseBody{
		ID:        *v.ID,
		CreatedAt: *v.CreatedAt,
		Key:       v.Key,
		Body:      v.Body,
		MediaID:   v.MediaID,
	}

	return res
}
