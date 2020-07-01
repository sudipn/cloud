// Code generated by goa v3.1.2, DO NOT EDIT.
//
// notes service
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package notes

import (
	"context"
	"io"

	notesviews "github.com/fieldkit/cloud/server/api/gen/notes/views"
	"goa.design/goa/v3/security"
)

// Service is the notes service interface.
type Service interface {
	// Update implements update.
	Update(context.Context, *UpdatePayload) (res *FieldNotes, err error)
	// Get implements get.
	Get(context.Context, *GetPayload) (res *FieldNotes, err error)
	// Media implements media.
	Media(context.Context, *MediaPayload) (res *MediaResult, body io.ReadCloser, err error)
	// Upload implements upload.
	Upload(context.Context, *UploadPayload, io.ReadCloser) (res *NoteMedia, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "notes"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"update", "get", "media", "upload"}

// UpdatePayload is the payload type of the notes service update method.
type UpdatePayload struct {
	Auth      string
	StationID int32
	Notes     *FieldNoteUpdate
}

// FieldNotes is the result type of the notes service update method.
type FieldNotes struct {
	Notes []*FieldNote
}

// GetPayload is the payload type of the notes service get method.
type GetPayload struct {
	Auth      string
	StationID int32
}

// MediaPayload is the payload type of the notes service media method.
type MediaPayload struct {
	Auth    string
	MediaID int32
}

// MediaResult is the result type of the notes service media method.
type MediaResult struct {
	Length      int64
	ContentType string
}

// UploadPayload is the payload type of the notes service upload method.
type UploadPayload struct {
	Auth          string
	ContentLength int64
	ContentType   string
}

// NoteMedia is the result type of the notes service upload method.
type NoteMedia struct {
	ID  int64
	URL string
}

type FieldNoteUpdate struct {
	Notes    []*ExistingFieldNote
	Creating []*NewFieldNote
}

type ExistingFieldNote struct {
	ID      int64
	Key     *string
	Body    *string
	MediaID *int64
}

type NewFieldNote struct {
	Key     *string
	Body    *string
	MediaID *int64
}

type FieldNote struct {
	ID        int64
	CreatedAt int64
	Key       *string
	Body      *string
	MediaID   *int64
}

// unauthorized
type Unauthorized string

// forbidden
type Forbidden string

// not-found
type NotFound string

// bad-request
type BadRequest string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "unauthorized"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e Forbidden) Error() string {
	return "forbidden"
}

// ErrorName returns "forbidden".
func (e Forbidden) ErrorName() string {
	return "forbidden"
}

// Error returns an error description.
func (e NotFound) Error() string {
	return "not-found"
}

// ErrorName returns "not-found".
func (e NotFound) ErrorName() string {
	return "not-found"
}

// Error returns an error description.
func (e BadRequest) Error() string {
	return "bad-request"
}

// ErrorName returns "bad-request".
func (e BadRequest) ErrorName() string {
	return "bad-request"
}

// NewFieldNotes initializes result type FieldNotes from viewed result type
// FieldNotes.
func NewFieldNotes(vres *notesviews.FieldNotes) *FieldNotes {
	return newFieldNotes(vres.Projected)
}

// NewViewedFieldNotes initializes viewed result type FieldNotes from result
// type FieldNotes using the given view.
func NewViewedFieldNotes(res *FieldNotes, view string) *notesviews.FieldNotes {
	p := newFieldNotesView(res)
	return &notesviews.FieldNotes{Projected: p, View: "default"}
}

// NewNoteMedia initializes result type NoteMedia from viewed result type
// NoteMedia.
func NewNoteMedia(vres *notesviews.NoteMedia) *NoteMedia {
	return newNoteMedia(vres.Projected)
}

// NewViewedNoteMedia initializes viewed result type NoteMedia from result type
// NoteMedia using the given view.
func NewViewedNoteMedia(res *NoteMedia, view string) *notesviews.NoteMedia {
	p := newNoteMediaView(res)
	return &notesviews.NoteMedia{Projected: p, View: "default"}
}

// newFieldNotes converts projected type FieldNotes to service type FieldNotes.
func newFieldNotes(vres *notesviews.FieldNotesView) *FieldNotes {
	res := &FieldNotes{}
	if vres.Notes != nil {
		res.Notes = make([]*FieldNote, len(vres.Notes))
		for i, val := range vres.Notes {
			res.Notes[i] = transformNotesviewsFieldNoteViewToFieldNote(val)
		}
	}
	return res
}

// newFieldNotesView projects result type FieldNotes to projected type
// FieldNotesView using the "default" view.
func newFieldNotesView(res *FieldNotes) *notesviews.FieldNotesView {
	vres := &notesviews.FieldNotesView{}
	if res.Notes != nil {
		vres.Notes = make([]*notesviews.FieldNoteView, len(res.Notes))
		for i, val := range res.Notes {
			vres.Notes[i] = transformFieldNoteToNotesviewsFieldNoteView(val)
		}
	}
	return vres
}

// newNoteMedia converts projected type NoteMedia to service type NoteMedia.
func newNoteMedia(vres *notesviews.NoteMediaView) *NoteMedia {
	res := &NoteMedia{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.URL != nil {
		res.URL = *vres.URL
	}
	return res
}

// newNoteMediaView projects result type NoteMedia to projected type
// NoteMediaView using the "default" view.
func newNoteMediaView(res *NoteMedia) *notesviews.NoteMediaView {
	vres := &notesviews.NoteMediaView{
		ID:  &res.ID,
		URL: &res.URL,
	}
	return vres
}

// transformNotesviewsFieldNoteViewToFieldNote builds a value of type
// *FieldNote from a value of type *notesviews.FieldNoteView.
func transformNotesviewsFieldNoteViewToFieldNote(v *notesviews.FieldNoteView) *FieldNote {
	if v == nil {
		return nil
	}
	res := &FieldNote{
		ID:        *v.ID,
		CreatedAt: *v.CreatedAt,
		Key:       v.Key,
		Body:      v.Body,
		MediaID:   v.MediaID,
	}

	return res
}

// transformFieldNoteToNotesviewsFieldNoteView builds a value of type
// *notesviews.FieldNoteView from a value of type *FieldNote.
func transformFieldNoteToNotesviewsFieldNoteView(v *FieldNote) *notesviews.FieldNoteView {
	res := &notesviews.FieldNoteView{
		ID:        &v.ID,
		CreatedAt: &v.CreatedAt,
		Key:       v.Key,
		Body:      v.Body,
		MediaID:   v.MediaID,
	}

	return res
}
