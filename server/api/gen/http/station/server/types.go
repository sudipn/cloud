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
	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "station" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	Name       *string                `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	DeviceID   *string                `form:"device_id,omitempty" json:"device_id,omitempty" xml:"device_id,omitempty"`
	StatusJSON map[string]interface{} `form:"status_json,omitempty" json:"status_json,omitempty" xml:"status_json,omitempty"`
	StatusPb   *string                `form:"status_pb,omitempty" json:"status_pb,omitempty" xml:"status_pb,omitempty"`
}

// UpdateRequestBody is the type of the "station" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	Name       *string                `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	StatusJSON map[string]interface{} `form:"status_json,omitempty" json:"status_json,omitempty" xml:"status_json,omitempty"`
	StatusPb   *string                `form:"status_pb,omitempty" json:"status_pb,omitempty" xml:"status_pb,omitempty"`
}

// AddResponseBody is the type of the "station" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
	ID                 int32                        `form:"id" json:"id" xml:"id"`
	Name               string                       `form:"name" json:"name" xml:"name"`
	Owner              *StationOwnerResponseBody    `form:"owner" json:"owner" xml:"owner"`
	DeviceID           string                       `form:"device_id" json:"device_id" xml:"device_id"`
	Uploads            []*StationUploadResponseBody `form:"uploads" json:"uploads" xml:"uploads"`
	Images             []*ImageRefResponseBody      `form:"images" json:"images" xml:"images"`
	Photos             *StationPhotosResponseBody   `form:"photos" json:"photos" xml:"photos"`
	ReadOnly           bool                         `form:"read_only" json:"read_only" xml:"read_only"`
	StatusJSON         map[string]interface{}       `form:"status_json" json:"status_json" xml:"status_json"`
	Battery            float32                      `form:"battery" json:"battery" xml:"battery"`
	RecordingStartedAt int64                        `form:"recording_started_at" json:"recording_started_at" xml:"recording_started_at"`
	MemoryUsed         int32                        `form:"memory_used" json:"memory_used" xml:"memory_used"`
	MemoryAvailable    int32                        `form:"memory_available" json:"memory_available" xml:"memory_available"`
	FirmwareNumber     int32                        `form:"firmware_number" json:"firmware_number" xml:"firmware_number"`
	FirmwareTime       int32                        `form:"firmware_time" json:"firmware_time" xml:"firmware_time"`
	Modules            []*StationModuleResponseBody `form:"modules" json:"modules" xml:"modules"`
}

// GetResponseBody is the type of the "station" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	ID                 int32                        `form:"id" json:"id" xml:"id"`
	Name               string                       `form:"name" json:"name" xml:"name"`
	Owner              *StationOwnerResponseBody    `form:"owner" json:"owner" xml:"owner"`
	DeviceID           string                       `form:"device_id" json:"device_id" xml:"device_id"`
	Uploads            []*StationUploadResponseBody `form:"uploads" json:"uploads" xml:"uploads"`
	Images             []*ImageRefResponseBody      `form:"images" json:"images" xml:"images"`
	Photos             *StationPhotosResponseBody   `form:"photos" json:"photos" xml:"photos"`
	ReadOnly           bool                         `form:"read_only" json:"read_only" xml:"read_only"`
	StatusJSON         map[string]interface{}       `form:"status_json" json:"status_json" xml:"status_json"`
	Battery            float32                      `form:"battery" json:"battery" xml:"battery"`
	RecordingStartedAt int64                        `form:"recording_started_at" json:"recording_started_at" xml:"recording_started_at"`
	MemoryUsed         int32                        `form:"memory_used" json:"memory_used" xml:"memory_used"`
	MemoryAvailable    int32                        `form:"memory_available" json:"memory_available" xml:"memory_available"`
	FirmwareNumber     int32                        `form:"firmware_number" json:"firmware_number" xml:"firmware_number"`
	FirmwareTime       int32                        `form:"firmware_time" json:"firmware_time" xml:"firmware_time"`
	Modules            []*StationModuleResponseBody `form:"modules" json:"modules" xml:"modules"`
}

// UpdateResponseBody is the type of the "station" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	ID                 int32                        `form:"id" json:"id" xml:"id"`
	Name               string                       `form:"name" json:"name" xml:"name"`
	Owner              *StationOwnerResponseBody    `form:"owner" json:"owner" xml:"owner"`
	DeviceID           string                       `form:"device_id" json:"device_id" xml:"device_id"`
	Uploads            []*StationUploadResponseBody `form:"uploads" json:"uploads" xml:"uploads"`
	Images             []*ImageRefResponseBody      `form:"images" json:"images" xml:"images"`
	Photos             *StationPhotosResponseBody   `form:"photos" json:"photos" xml:"photos"`
	ReadOnly           bool                         `form:"read_only" json:"read_only" xml:"read_only"`
	StatusJSON         map[string]interface{}       `form:"status_json" json:"status_json" xml:"status_json"`
	Battery            float32                      `form:"battery" json:"battery" xml:"battery"`
	RecordingStartedAt int64                        `form:"recording_started_at" json:"recording_started_at" xml:"recording_started_at"`
	MemoryUsed         int32                        `form:"memory_used" json:"memory_used" xml:"memory_used"`
	MemoryAvailable    int32                        `form:"memory_available" json:"memory_available" xml:"memory_available"`
	FirmwareNumber     int32                        `form:"firmware_number" json:"firmware_number" xml:"firmware_number"`
	FirmwareTime       int32                        `form:"firmware_time" json:"firmware_time" xml:"firmware_time"`
	Modules            []*StationModuleResponseBody `form:"modules" json:"modules" xml:"modules"`
}

// ListMineResponseBody is the type of the "station" service "list mine"
// endpoint HTTP response body.
type ListMineResponseBody struct {
	Stations StationFullResponseBodyCollection `form:"stations" json:"stations" xml:"stations"`
}

// ListProjectResponseBody is the type of the "station" service "list project"
// endpoint HTTP response body.
type ListProjectResponseBody struct {
	Stations StationFullResponseBodyCollection `form:"stations" json:"stations" xml:"stations"`
}

// AddBadRequestResponseBody is the type of the "station" service "add"
// endpoint HTTP response body for the "bad-request" error.
type AddBadRequestResponseBody string

// AddForbiddenResponseBody is the type of the "station" service "add" endpoint
// HTTP response body for the "forbidden" error.
type AddForbiddenResponseBody string

// AddNotFoundResponseBody is the type of the "station" service "add" endpoint
// HTTP response body for the "not-found" error.
type AddNotFoundResponseBody string

// AddUnauthorizedResponseBody is the type of the "station" service "add"
// endpoint HTTP response body for the "unauthorized" error.
type AddUnauthorizedResponseBody string

// GetBadRequestResponseBody is the type of the "station" service "get"
// endpoint HTTP response body for the "bad-request" error.
type GetBadRequestResponseBody string

// GetForbiddenResponseBody is the type of the "station" service "get" endpoint
// HTTP response body for the "forbidden" error.
type GetForbiddenResponseBody string

// GetNotFoundResponseBody is the type of the "station" service "get" endpoint
// HTTP response body for the "not-found" error.
type GetNotFoundResponseBody string

// GetUnauthorizedResponseBody is the type of the "station" service "get"
// endpoint HTTP response body for the "unauthorized" error.
type GetUnauthorizedResponseBody string

// UpdateBadRequestResponseBody is the type of the "station" service "update"
// endpoint HTTP response body for the "bad-request" error.
type UpdateBadRequestResponseBody string

// UpdateForbiddenResponseBody is the type of the "station" service "update"
// endpoint HTTP response body for the "forbidden" error.
type UpdateForbiddenResponseBody string

// UpdateNotFoundResponseBody is the type of the "station" service "update"
// endpoint HTTP response body for the "not-found" error.
type UpdateNotFoundResponseBody string

// UpdateUnauthorizedResponseBody is the type of the "station" service "update"
// endpoint HTTP response body for the "unauthorized" error.
type UpdateUnauthorizedResponseBody string

// ListMineBadRequestResponseBody is the type of the "station" service "list
// mine" endpoint HTTP response body for the "bad-request" error.
type ListMineBadRequestResponseBody string

// ListMineForbiddenResponseBody is the type of the "station" service "list
// mine" endpoint HTTP response body for the "forbidden" error.
type ListMineForbiddenResponseBody string

// ListMineNotFoundResponseBody is the type of the "station" service "list
// mine" endpoint HTTP response body for the "not-found" error.
type ListMineNotFoundResponseBody string

// ListMineUnauthorizedResponseBody is the type of the "station" service "list
// mine" endpoint HTTP response body for the "unauthorized" error.
type ListMineUnauthorizedResponseBody string

// ListProjectBadRequestResponseBody is the type of the "station" service "list
// project" endpoint HTTP response body for the "bad-request" error.
type ListProjectBadRequestResponseBody string

// ListProjectForbiddenResponseBody is the type of the "station" service "list
// project" endpoint HTTP response body for the "forbidden" error.
type ListProjectForbiddenResponseBody string

// ListProjectNotFoundResponseBody is the type of the "station" service "list
// project" endpoint HTTP response body for the "not-found" error.
type ListProjectNotFoundResponseBody string

// ListProjectUnauthorizedResponseBody is the type of the "station" service
// "list project" endpoint HTTP response body for the "unauthorized" error.
type ListProjectUnauthorizedResponseBody string

// PhotoBadRequestResponseBody is the type of the "station" service "photo"
// endpoint HTTP response body for the "bad-request" error.
type PhotoBadRequestResponseBody string

// PhotoForbiddenResponseBody is the type of the "station" service "photo"
// endpoint HTTP response body for the "forbidden" error.
type PhotoForbiddenResponseBody string

// PhotoNotFoundResponseBody is the type of the "station" service "photo"
// endpoint HTTP response body for the "not-found" error.
type PhotoNotFoundResponseBody string

// PhotoUnauthorizedResponseBody is the type of the "station" service "photo"
// endpoint HTTP response body for the "unauthorized" error.
type PhotoUnauthorizedResponseBody string

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

// StationModuleResponseBody is used to define fields on response body types.
type StationModuleResponseBody struct {
	ID         int64                        `form:"id" json:"id" xml:"id"`
	HardwareID *string                      `form:"hardware_id,omitempty" json:"hardware_id,omitempty" xml:"hardware_id,omitempty"`
	Name       string                       `form:"name" json:"name" xml:"name"`
	Position   int32                        `form:"position" json:"position" xml:"position"`
	Flags      int32                        `form:"flags" json:"flags" xml:"flags"`
	Internal   bool                         `form:"internal" json:"internal" xml:"internal"`
	Sensors    []*StationSensorResponseBody `form:"sensors" json:"sensors" xml:"sensors"`
}

// StationSensorResponseBody is used to define fields on response body types.
type StationSensorResponseBody struct {
	Name          string                     `form:"name" json:"name" xml:"name"`
	UnitOfMeasure string                     `form:"unit_of_measure" json:"unit_of_measure" xml:"unit_of_measure"`
	Reading       *SensorReadingResponseBody `form:"reading,omitempty" json:"reading,omitempty" xml:"reading,omitempty"`
}

// SensorReadingResponseBody is used to define fields on response body types.
type SensorReadingResponseBody struct {
	Last float32 `form:"last" json:"last" xml:"last"`
	Time int64   `form:"time" json:"time" xml:"time"`
}

// StationFullResponseBodyCollection is used to define fields on response body
// types.
type StationFullResponseBodyCollection []*StationFullResponseBody

// StationFullResponseBody is used to define fields on response body types.
type StationFullResponseBody struct {
	ID                 int32                        `form:"id" json:"id" xml:"id"`
	Name               string                       `form:"name" json:"name" xml:"name"`
	Owner              *StationOwnerResponseBody    `form:"owner" json:"owner" xml:"owner"`
	DeviceID           string                       `form:"device_id" json:"device_id" xml:"device_id"`
	Uploads            []*StationUploadResponseBody `form:"uploads" json:"uploads" xml:"uploads"`
	Images             []*ImageRefResponseBody      `form:"images" json:"images" xml:"images"`
	Photos             *StationPhotosResponseBody   `form:"photos" json:"photos" xml:"photos"`
	ReadOnly           bool                         `form:"read_only" json:"read_only" xml:"read_only"`
	StatusJSON         map[string]interface{}       `form:"status_json" json:"status_json" xml:"status_json"`
	Battery            float32                      `form:"battery" json:"battery" xml:"battery"`
	RecordingStartedAt int64                        `form:"recording_started_at" json:"recording_started_at" xml:"recording_started_at"`
	MemoryUsed         int32                        `form:"memory_used" json:"memory_used" xml:"memory_used"`
	MemoryAvailable    int32                        `form:"memory_available" json:"memory_available" xml:"memory_available"`
	FirmwareNumber     int32                        `form:"firmware_number" json:"firmware_number" xml:"firmware_number"`
	FirmwareTime       int32                        `form:"firmware_time" json:"firmware_time" xml:"firmware_time"`
	Modules            []*StationModuleResponseBody `form:"modules" json:"modules" xml:"modules"`
}

// NewAddResponseBody builds the HTTP response body from the result of the
// "add" endpoint of the "station" service.
func NewAddResponseBody(res *stationviews.StationFullView) *AddResponseBody {
	body := &AddResponseBody{
		ID:                 *res.ID,
		Name:               *res.Name,
		DeviceID:           *res.DeviceID,
		ReadOnly:           *res.ReadOnly,
		Battery:            *res.Battery,
		RecordingStartedAt: *res.RecordingStartedAt,
		MemoryUsed:         *res.MemoryUsed,
		MemoryAvailable:    *res.MemoryAvailable,
		FirmwareNumber:     *res.FirmwareNumber,
		FirmwareTime:       *res.FirmwareTime,
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
	if res.StatusJSON != nil {
		body.StatusJSON = make(map[string]interface{}, len(res.StatusJSON))
		for key, val := range res.StatusJSON {
			tk := key
			tv := val
			body.StatusJSON[tk] = tv
		}
	}
	if res.Modules != nil {
		body.Modules = make([]*StationModuleResponseBody, len(res.Modules))
		for i, val := range res.Modules {
			body.Modules[i] = marshalStationviewsStationModuleViewToStationModuleResponseBody(val)
		}
	}
	return body
}

// NewGetResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "station" service.
func NewGetResponseBody(res *stationviews.StationFullView) *GetResponseBody {
	body := &GetResponseBody{
		ID:                 *res.ID,
		Name:               *res.Name,
		DeviceID:           *res.DeviceID,
		ReadOnly:           *res.ReadOnly,
		Battery:            *res.Battery,
		RecordingStartedAt: *res.RecordingStartedAt,
		MemoryUsed:         *res.MemoryUsed,
		MemoryAvailable:    *res.MemoryAvailable,
		FirmwareNumber:     *res.FirmwareNumber,
		FirmwareTime:       *res.FirmwareTime,
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
	if res.StatusJSON != nil {
		body.StatusJSON = make(map[string]interface{}, len(res.StatusJSON))
		for key, val := range res.StatusJSON {
			tk := key
			tv := val
			body.StatusJSON[tk] = tv
		}
	}
	if res.Modules != nil {
		body.Modules = make([]*StationModuleResponseBody, len(res.Modules))
		for i, val := range res.Modules {
			body.Modules[i] = marshalStationviewsStationModuleViewToStationModuleResponseBody(val)
		}
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "station" service.
func NewUpdateResponseBody(res *stationviews.StationFullView) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:                 *res.ID,
		Name:               *res.Name,
		DeviceID:           *res.DeviceID,
		ReadOnly:           *res.ReadOnly,
		Battery:            *res.Battery,
		RecordingStartedAt: *res.RecordingStartedAt,
		MemoryUsed:         *res.MemoryUsed,
		MemoryAvailable:    *res.MemoryAvailable,
		FirmwareNumber:     *res.FirmwareNumber,
		FirmwareTime:       *res.FirmwareTime,
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
	if res.StatusJSON != nil {
		body.StatusJSON = make(map[string]interface{}, len(res.StatusJSON))
		for key, val := range res.StatusJSON {
			tk := key
			tv := val
			body.StatusJSON[tk] = tv
		}
	}
	if res.Modules != nil {
		body.Modules = make([]*StationModuleResponseBody, len(res.Modules))
		for i, val := range res.Modules {
			body.Modules[i] = marshalStationviewsStationModuleViewToStationModuleResponseBody(val)
		}
	}
	return body
}

// NewListMineResponseBody builds the HTTP response body from the result of the
// "list mine" endpoint of the "station" service.
func NewListMineResponseBody(res *stationviews.StationsFullView) *ListMineResponseBody {
	body := &ListMineResponseBody{}
	if res.Stations != nil {
		body.Stations = make([]*StationFullResponseBody, len(res.Stations))
		for i, val := range res.Stations {
			body.Stations[i] = marshalStationviewsStationFullViewToStationFullResponseBody(val)
		}
	}
	return body
}

// NewListProjectResponseBody builds the HTTP response body from the result of
// the "list project" endpoint of the "station" service.
func NewListProjectResponseBody(res *stationviews.StationsFullView) *ListProjectResponseBody {
	body := &ListProjectResponseBody{}
	if res.Stations != nil {
		body.Stations = make([]*StationFullResponseBody, len(res.Stations))
		for i, val := range res.Stations {
			body.Stations[i] = marshalStationviewsStationFullViewToStationFullResponseBody(val)
		}
	}
	return body
}

// NewAddBadRequestResponseBody builds the HTTP response body from the result
// of the "add" endpoint of the "station" service.
func NewAddBadRequestResponseBody(res station.BadRequest) AddBadRequestResponseBody {
	body := AddBadRequestResponseBody(res)
	return body
}

// NewAddForbiddenResponseBody builds the HTTP response body from the result of
// the "add" endpoint of the "station" service.
func NewAddForbiddenResponseBody(res station.Forbidden) AddForbiddenResponseBody {
	body := AddForbiddenResponseBody(res)
	return body
}

// NewAddNotFoundResponseBody builds the HTTP response body from the result of
// the "add" endpoint of the "station" service.
func NewAddNotFoundResponseBody(res station.NotFound) AddNotFoundResponseBody {
	body := AddNotFoundResponseBody(res)
	return body
}

// NewAddUnauthorizedResponseBody builds the HTTP response body from the result
// of the "add" endpoint of the "station" service.
func NewAddUnauthorizedResponseBody(res station.Unauthorized) AddUnauthorizedResponseBody {
	body := AddUnauthorizedResponseBody(res)
	return body
}

// NewGetBadRequestResponseBody builds the HTTP response body from the result
// of the "get" endpoint of the "station" service.
func NewGetBadRequestResponseBody(res station.BadRequest) GetBadRequestResponseBody {
	body := GetBadRequestResponseBody(res)
	return body
}

// NewGetForbiddenResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "station" service.
func NewGetForbiddenResponseBody(res station.Forbidden) GetForbiddenResponseBody {
	body := GetForbiddenResponseBody(res)
	return body
}

// NewGetNotFoundResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "station" service.
func NewGetNotFoundResponseBody(res station.NotFound) GetNotFoundResponseBody {
	body := GetNotFoundResponseBody(res)
	return body
}

// NewGetUnauthorizedResponseBody builds the HTTP response body from the result
// of the "get" endpoint of the "station" service.
func NewGetUnauthorizedResponseBody(res station.Unauthorized) GetUnauthorizedResponseBody {
	body := GetUnauthorizedResponseBody(res)
	return body
}

// NewUpdateBadRequestResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "station" service.
func NewUpdateBadRequestResponseBody(res station.BadRequest) UpdateBadRequestResponseBody {
	body := UpdateBadRequestResponseBody(res)
	return body
}

// NewUpdateForbiddenResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "station" service.
func NewUpdateForbiddenResponseBody(res station.Forbidden) UpdateForbiddenResponseBody {
	body := UpdateForbiddenResponseBody(res)
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "station" service.
func NewUpdateNotFoundResponseBody(res station.NotFound) UpdateNotFoundResponseBody {
	body := UpdateNotFoundResponseBody(res)
	return body
}

// NewUpdateUnauthorizedResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "station" service.
func NewUpdateUnauthorizedResponseBody(res station.Unauthorized) UpdateUnauthorizedResponseBody {
	body := UpdateUnauthorizedResponseBody(res)
	return body
}

// NewListMineBadRequestResponseBody builds the HTTP response body from the
// result of the "list mine" endpoint of the "station" service.
func NewListMineBadRequestResponseBody(res station.BadRequest) ListMineBadRequestResponseBody {
	body := ListMineBadRequestResponseBody(res)
	return body
}

// NewListMineForbiddenResponseBody builds the HTTP response body from the
// result of the "list mine" endpoint of the "station" service.
func NewListMineForbiddenResponseBody(res station.Forbidden) ListMineForbiddenResponseBody {
	body := ListMineForbiddenResponseBody(res)
	return body
}

// NewListMineNotFoundResponseBody builds the HTTP response body from the
// result of the "list mine" endpoint of the "station" service.
func NewListMineNotFoundResponseBody(res station.NotFound) ListMineNotFoundResponseBody {
	body := ListMineNotFoundResponseBody(res)
	return body
}

// NewListMineUnauthorizedResponseBody builds the HTTP response body from the
// result of the "list mine" endpoint of the "station" service.
func NewListMineUnauthorizedResponseBody(res station.Unauthorized) ListMineUnauthorizedResponseBody {
	body := ListMineUnauthorizedResponseBody(res)
	return body
}

// NewListProjectBadRequestResponseBody builds the HTTP response body from the
// result of the "list project" endpoint of the "station" service.
func NewListProjectBadRequestResponseBody(res station.BadRequest) ListProjectBadRequestResponseBody {
	body := ListProjectBadRequestResponseBody(res)
	return body
}

// NewListProjectForbiddenResponseBody builds the HTTP response body from the
// result of the "list project" endpoint of the "station" service.
func NewListProjectForbiddenResponseBody(res station.Forbidden) ListProjectForbiddenResponseBody {
	body := ListProjectForbiddenResponseBody(res)
	return body
}

// NewListProjectNotFoundResponseBody builds the HTTP response body from the
// result of the "list project" endpoint of the "station" service.
func NewListProjectNotFoundResponseBody(res station.NotFound) ListProjectNotFoundResponseBody {
	body := ListProjectNotFoundResponseBody(res)
	return body
}

// NewListProjectUnauthorizedResponseBody builds the HTTP response body from
// the result of the "list project" endpoint of the "station" service.
func NewListProjectUnauthorizedResponseBody(res station.Unauthorized) ListProjectUnauthorizedResponseBody {
	body := ListProjectUnauthorizedResponseBody(res)
	return body
}

// NewPhotoBadRequestResponseBody builds the HTTP response body from the result
// of the "photo" endpoint of the "station" service.
func NewPhotoBadRequestResponseBody(res station.BadRequest) PhotoBadRequestResponseBody {
	body := PhotoBadRequestResponseBody(res)
	return body
}

// NewPhotoForbiddenResponseBody builds the HTTP response body from the result
// of the "photo" endpoint of the "station" service.
func NewPhotoForbiddenResponseBody(res station.Forbidden) PhotoForbiddenResponseBody {
	body := PhotoForbiddenResponseBody(res)
	return body
}

// NewPhotoNotFoundResponseBody builds the HTTP response body from the result
// of the "photo" endpoint of the "station" service.
func NewPhotoNotFoundResponseBody(res station.NotFound) PhotoNotFoundResponseBody {
	body := PhotoNotFoundResponseBody(res)
	return body
}

// NewPhotoUnauthorizedResponseBody builds the HTTP response body from the
// result of the "photo" endpoint of the "station" service.
func NewPhotoUnauthorizedResponseBody(res station.Unauthorized) PhotoUnauthorizedResponseBody {
	body := PhotoUnauthorizedResponseBody(res)
	return body
}

// NewAddPayload builds a station service add endpoint payload.
func NewAddPayload(body *AddRequestBody, auth string) *station.AddPayload {
	v := &station.AddPayload{
		Name:     *body.Name,
		DeviceID: *body.DeviceID,
		StatusPb: body.StatusPb,
	}
	v.StatusJSON = make(map[string]interface{}, len(body.StatusJSON))
	for key, val := range body.StatusJSON {
		tk := key
		tv := val
		v.StatusJSON[tk] = tv
	}
	v.Auth = auth

	return v
}

// NewGetPayload builds a station service get endpoint payload.
func NewGetPayload(id int32, auth string) *station.GetPayload {
	v := &station.GetPayload{}
	v.ID = id
	v.Auth = auth

	return v
}

// NewUpdatePayload builds a station service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id int32, auth string) *station.UpdatePayload {
	v := &station.UpdatePayload{
		Name:     *body.Name,
		StatusPb: body.StatusPb,
	}
	v.StatusJSON = make(map[string]interface{}, len(body.StatusJSON))
	for key, val := range body.StatusJSON {
		tk := key
		tv := val
		v.StatusJSON[tk] = tv
	}
	v.ID = id
	v.Auth = auth

	return v
}

// NewListMinePayload builds a station service list mine endpoint payload.
func NewListMinePayload(auth string) *station.ListMinePayload {
	v := &station.ListMinePayload{}
	v.Auth = auth

	return v
}

// NewListProjectPayload builds a station service list project endpoint payload.
func NewListProjectPayload(id int32, auth string) *station.ListProjectPayload {
	v := &station.ListProjectPayload{}
	v.ID = id
	v.Auth = auth

	return v
}

// NewPhotoPayload builds a station service photo endpoint payload.
func NewPhotoPayload(id int32, auth string) *station.PhotoPayload {
	v := &station.PhotoPayload{}
	v.ID = id
	v.Auth = auth

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.DeviceID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("device_id", "body"))
	}
	if body.StatusJSON == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status_json", "body"))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.StatusJSON == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status_json", "body"))
	}
	return
}
