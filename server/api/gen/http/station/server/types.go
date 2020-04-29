// Code generated by goa v3.1.2, DO NOT EDIT.
//
// station HTTP server types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	station "github.com/fieldkit/cloud/server/api/gen/station"
	stationviews "github.com/fieldkit/cloud/server/api/gen/station/views"
)

// StationResponseBody is the type of the "station" service "station" endpoint
// HTTP response body.
type StationResponseBody struct {
	ID       int32                        `form:"id" json:"id" xml:"id"`
	Name     string                       `form:"name" json:"name" xml:"name"`
	Owner    *StationOwnerResponseBody    `form:"owner" json:"owner" xml:"owner"`
	DeviceID string                       `form:"device_id" json:"device_id" xml:"device_id"`
	Uploads  []*StationUploadResponseBody `form:"uploads" json:"uploads" xml:"uploads"`
	Images   []*ImageRefResponseBody      `form:"images" json:"images" xml:"images"`
	Photos   *StationPhotosResponseBody   `form:"photos" json:"photos" xml:"photos"`
	ReadOnly bool                         `form:"read_only" json:"read_only" xml:"read_only"`
}

// StationNotFoundResponseBody is the type of the "station" service "station"
// endpoint HTTP response body for the "not-found" error.
type StationNotFoundResponseBody string

// StationUnauthorizedResponseBody is the type of the "station" service
// "station" endpoint HTTP response body for the "unauthorized" error.
type StationUnauthorizedResponseBody string

// StationOwnerResponseBody is used to define fields on response body types.
type StationOwnerResponseBody struct {
	ID   int32  `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// StationUploadResponseBody is used to define fields on response body types.
type StationUploadResponseBody struct {
	ID       int64   `form:"id" json:"id" xml:"id"`
	Time     int64   `form:"time" json:"time" xml:"time"`
	UploadID string  `form:"upload_id" json:"upload_id" xml:"upload_id"`
	Size     int64   `form:"size" json:"size" xml:"size"`
	URL      string  `form:"url" json:"url" xml:"url"`
	Type     string  `form:"type" json:"type" xml:"type"`
	Blocks   []int64 `form:"blocks" json:"blocks" xml:"blocks"`
}

// ImageRefResponseBody is used to define fields on response body types.
type ImageRefResponseBody struct {
	URL string `form:"url" json:"url" xml:"url"`
}

// StationPhotosResponseBody is used to define fields on response body types.
type StationPhotosResponseBody struct {
	Small string `form:"small" json:"small" xml:"small"`
}

// NewStationResponseBody builds the HTTP response body from the result of the
// "station" endpoint of the "station" service.
func NewStationResponseBody(res *stationviews.StationFullView) *StationResponseBody {
	body := &StationResponseBody{
		ID:       *res.ID,
		Name:     *res.Name,
		DeviceID: *res.DeviceID,
		ReadOnly: *res.ReadOnly,
	}
	if res.Owner != nil {
		body.Owner = marshalStationviewsStationOwnerViewToStationOwnerResponseBody(res.Owner)
	}
	if res.Uploads != nil {
		body.Uploads = make([]*StationUploadResponseBody, len(res.Uploads))
		for i, val := range res.Uploads {
			body.Uploads[i] = marshalStationviewsStationUploadViewToStationUploadResponseBody(val)
		}
	}
	if res.Images != nil {
		body.Images = make([]*ImageRefResponseBody, len(res.Images))
		for i, val := range res.Images {
			body.Images[i] = marshalStationviewsImageRefViewToImageRefResponseBody(val)
		}
	}
	if res.Photos != nil {
		body.Photos = marshalStationviewsStationPhotosViewToStationPhotosResponseBody(res.Photos)
	}
	return body
}

// NewStationNotFoundResponseBody builds the HTTP response body from the result
// of the "station" endpoint of the "station" service.
func NewStationNotFoundResponseBody(res station.NotFound) StationNotFoundResponseBody {
	body := StationNotFoundResponseBody(res)
	return body
}

// NewStationUnauthorizedResponseBody builds the HTTP response body from the
// result of the "station" endpoint of the "station" service.
func NewStationUnauthorizedResponseBody(res station.Unauthorized) StationUnauthorizedResponseBody {
	body := StationUnauthorizedResponseBody(res)
	return body
}

// NewStationPayload builds a station service station endpoint payload.
func NewStationPayload(id int32, auth string) *station.StationPayload {
	v := &station.StationPayload{}
	v.ID = id
	v.Auth = auth

	return v
}