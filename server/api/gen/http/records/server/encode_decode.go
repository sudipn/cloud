// Code generated by goa v3.1.2, DO NOT EDIT.
//
// records HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	records "github.com/fieldkit/cloud/server/api/gen/records"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeDataResponse returns an encoder for responses returned by the records
// data endpoint.
func EncodeDataResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*records.DataResult)
		enc := encoder(ctx, w)
		body := res.Object
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDataRequest returns a decoder for requests sent to the records data
// endpoint.
func DecodeDataRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			recordID int64
			auth     *string
			err      error

			params = mux.Vars(r)
		)
		{
			recordIDRaw := params["recordId"]
			v, err2 := strconv.ParseInt(recordIDRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("recordID", recordIDRaw, "integer"))
			}
			recordID = v
		}
		authRaw := r.Header.Get("Authorization")
		if authRaw != "" {
			auth = &authRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewDataPayload(recordID, auth)
		if payload.Auth != nil {
			if strings.Contains(*payload.Auth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Auth, " ", 2)[1]
				payload.Auth = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDataError returns an encoder for errors returned by the data records
// endpoint.
func EncodeDataError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(records.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDataBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			res := v.(records.Forbidden)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDataForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(records.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDataNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(records.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDataUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeMetaResponse returns an encoder for responses returned by the records
// meta endpoint.
func EncodeMetaResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*records.MetaResult)
		enc := encoder(ctx, w)
		body := res.Object
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeMetaRequest returns a decoder for requests sent to the records meta
// endpoint.
func DecodeMetaRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			recordID int64
			auth     *string
			err      error

			params = mux.Vars(r)
		)
		{
			recordIDRaw := params["recordId"]
			v, err2 := strconv.ParseInt(recordIDRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("recordID", recordIDRaw, "integer"))
			}
			recordID = v
		}
		authRaw := r.Header.Get("Authorization")
		if authRaw != "" {
			auth = &authRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewMetaPayload(recordID, auth)
		if payload.Auth != nil {
			if strings.Contains(*payload.Auth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Auth, " ", 2)[1]
				payload.Auth = &cred
			}
		}

		return payload, nil
	}
}

// EncodeMetaError returns an encoder for errors returned by the meta records
// endpoint.
func EncodeMetaError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(records.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMetaBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			res := v.(records.Forbidden)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMetaForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(records.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMetaNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(records.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMetaUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeResolvedResponse returns an encoder for responses returned by the
// records resolved endpoint.
func EncodeResolvedResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*records.ResolvedResult)
		enc := encoder(ctx, w)
		body := res.Object
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeResolvedRequest returns a decoder for requests sent to the records
// resolved endpoint.
func DecodeResolvedRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			recordID int64
			auth     *string
			err      error

			params = mux.Vars(r)
		)
		{
			recordIDRaw := params["recordId"]
			v, err2 := strconv.ParseInt(recordIDRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("recordID", recordIDRaw, "integer"))
			}
			recordID = v
		}
		authRaw := r.Header.Get("Authorization")
		if authRaw != "" {
			auth = &authRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewResolvedPayload(recordID, auth)
		if payload.Auth != nil {
			if strings.Contains(*payload.Auth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Auth, " ", 2)[1]
				payload.Auth = &cred
			}
		}

		return payload, nil
	}
}

// EncodeResolvedError returns an encoder for errors returned by the resolved
// records endpoint.
func EncodeResolvedError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(records.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewResolvedBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			res := v.(records.Forbidden)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewResolvedForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(records.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewResolvedNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(records.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewResolvedUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
