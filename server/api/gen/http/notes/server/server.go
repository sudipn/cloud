// Code generated by goa v3.1.2, DO NOT EDIT.
//
// notes HTTP server
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"context"
	"io"
	"net/http"
	"regexp"

	notes "github.com/fieldkit/cloud/server/api/gen/notes"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the notes service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Update http.Handler
	Get    http.Handler
	Media  http.Handler
	Upload http.Handler
	CORS   http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the notes service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *notes.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Update", "PATCH", "/stations/{stationId}/notes"},
			{"Get", "GET", "/stations/{stationId}/notes"},
			{"Media", "GET", "/notes/media/{mediaId}"},
			{"Upload", "POST", "/notes/media"},
			{"CORS", "OPTIONS", "/stations/{stationId}/notes"},
			{"CORS", "OPTIONS", "/notes/media/{mediaId}"},
			{"CORS", "OPTIONS", "/notes/media"},
		},
		Update: NewUpdateHandler(e.Update, mux, decoder, encoder, errhandler, formatter),
		Get:    NewGetHandler(e.Get, mux, decoder, encoder, errhandler, formatter),
		Media:  NewMediaHandler(e.Media, mux, decoder, encoder, errhandler, formatter),
		Upload: NewUploadHandler(e.Upload, mux, decoder, encoder, errhandler, formatter),
		CORS:   NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "notes" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Update = m(s.Update)
	s.Get = m(s.Get)
	s.Media = m(s.Media)
	s.Upload = m(s.Upload)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the notes endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountUpdateHandler(mux, h.Update)
	MountGetHandler(mux, h.Get)
	MountMediaHandler(mux, h.Media)
	MountUploadHandler(mux, h.Upload)
	MountCORSHandler(mux, h.CORS)
}

// MountUpdateHandler configures the mux to serve the "notes" service "update"
// endpoint.
func MountUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleNotesOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PATCH", "/stations/{stationId}/notes", f)
}

// NewUpdateHandler creates a HTTP handler which loads the HTTP request and
// calls the "notes" service "update" endpoint.
func NewUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateRequest(mux, decoder)
		encodeResponse = EncodeUpdateResponse(encoder)
		encodeError    = EncodeUpdateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "notes")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetHandler configures the mux to serve the "notes" service "get"
// endpoint.
func MountGetHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleNotesOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/stations/{stationId}/notes", f)
}

// NewGetHandler creates a HTTP handler which loads the HTTP request and calls
// the "notes" service "get" endpoint.
func NewGetHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetRequest(mux, decoder)
		encodeResponse = EncodeGetResponse(encoder)
		encodeError    = EncodeGetError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get")
		ctx = context.WithValue(ctx, goa.ServiceKey, "notes")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountMediaHandler configures the mux to serve the "notes" service "media"
// endpoint.
func MountMediaHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleNotesOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/notes/media/{mediaId}", f)
}

// NewMediaHandler creates a HTTP handler which loads the HTTP request and
// calls the "notes" service "media" endpoint.
func NewMediaHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMediaRequest(mux, decoder)
		encodeResponse = EncodeMediaResponse(encoder)
		encodeError    = EncodeMediaError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "media")
		ctx = context.WithValue(ctx, goa.ServiceKey, "notes")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		o := res.(*notes.MediaResponseData)
		defer o.Body.Close()
		if err := encodeResponse(ctx, w, o.Result); err != nil {
			errhandler(ctx, w, err)
			return
		}
		if _, err := io.Copy(w, o.Body); err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
		}
	})
}

// MountUploadHandler configures the mux to serve the "notes" service "upload"
// endpoint.
func MountUploadHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleNotesOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/notes/media", f)
}

// NewUploadHandler creates a HTTP handler which loads the HTTP request and
// calls the "notes" service "upload" endpoint.
func NewUploadHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUploadRequest(mux, decoder)
		encodeResponse = EncodeUploadResponse(encoder)
		encodeError    = EncodeUploadError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "upload")
		ctx = context.WithValue(ctx, goa.ServiceKey, "notes")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		data := &notes.UploadRequestData{Payload: payload.(*notes.UploadPayload), Body: r.Body}
		res, err := endpoint(ctx, data)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service notes.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleNotesOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/stations/{stationId}/notes", f)
	mux.Handle("OPTIONS", "/notes/media/{mediaId}", f)
	mux.Handle("OPTIONS", "/notes/media", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleNotesOrigin applies the CORS response headers corresponding to the
// origin for the service notes.
func handleNotesOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile("(.+[.])?127.0.0.1:\\d+")
	spec1 := regexp.MustCompile("(.+[.])?fieldkit.org:\\d+")
	spec2 := regexp.MustCompile("(.+[.])?local.fkdev.org:\\d+")
	spec3 := regexp.MustCompile("(.+[.])?localhost:\\d+")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec1) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec2) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec3) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org:8080") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://*.fkdev.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org:8080") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://fkdev.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
