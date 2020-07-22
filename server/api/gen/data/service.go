// Code generated by goa v3.1.2, DO NOT EDIT.
//
// data service
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package data

import (
	"context"

	dataviews "github.com/fieldkit/cloud/server/api/gen/data/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the data service interface.
type Service interface {
	// DeviceSummary implements device summary.
	DeviceSummary(context.Context, *DeviceSummaryPayload) (res *DeviceDataSummaryResponse, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "data"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"device summary"}

// DeviceSummaryPayload is the payload type of the data service device summary
// method.
type DeviceSummaryPayload struct {
	Auth     *string
	DeviceID string
}

// DeviceDataSummaryResponse is the result type of the data service device
// summary method.
type DeviceDataSummaryResponse struct {
	Provisions DeviceProvisionSummaryCollection
}

type DeviceProvisionSummaryCollection []*DeviceProvisionSummary

type DeviceProvisionSummary struct {
	Generation string
	Created    int64
	Updated    int64
	Meta       *DeviceMetaSummary
	Data       *DeviceDataSummary
}

type DeviceMetaSummary struct {
	Size  int64
	First int64
	Last  int64
}

type DeviceDataSummary struct {
	Size  int64
	First int64
	Last  int64
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeForbidden builds a goa.ServiceError from an error.
func MakeForbidden(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "forbidden",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not-found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad-request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewDeviceDataSummaryResponse initializes result type
// DeviceDataSummaryResponse from viewed result type DeviceDataSummaryResponse.
func NewDeviceDataSummaryResponse(vres *dataviews.DeviceDataSummaryResponse) *DeviceDataSummaryResponse {
	return newDeviceDataSummaryResponse(vres.Projected)
}

// NewViewedDeviceDataSummaryResponse initializes viewed result type
// DeviceDataSummaryResponse from result type DeviceDataSummaryResponse using
// the given view.
func NewViewedDeviceDataSummaryResponse(res *DeviceDataSummaryResponse, view string) *dataviews.DeviceDataSummaryResponse {
	p := newDeviceDataSummaryResponseView(res)
	return &dataviews.DeviceDataSummaryResponse{Projected: p, View: "default"}
}

// newDeviceDataSummaryResponse converts projected type
// DeviceDataSummaryResponse to service type DeviceDataSummaryResponse.
func newDeviceDataSummaryResponse(vres *dataviews.DeviceDataSummaryResponseView) *DeviceDataSummaryResponse {
	res := &DeviceDataSummaryResponse{}
	if vres.Provisions != nil {
		res.Provisions = newDeviceProvisionSummaryCollection(vres.Provisions)
	}
	return res
}

// newDeviceDataSummaryResponseView projects result type
// DeviceDataSummaryResponse to projected type DeviceDataSummaryResponseView
// using the "default" view.
func newDeviceDataSummaryResponseView(res *DeviceDataSummaryResponse) *dataviews.DeviceDataSummaryResponseView {
	vres := &dataviews.DeviceDataSummaryResponseView{}
	if res.Provisions != nil {
		vres.Provisions = newDeviceProvisionSummaryCollectionView(res.Provisions)
	}
	return vres
}

// newDeviceProvisionSummaryCollection converts projected type
// DeviceProvisionSummaryCollection to service type
// DeviceProvisionSummaryCollection.
func newDeviceProvisionSummaryCollection(vres dataviews.DeviceProvisionSummaryCollectionView) DeviceProvisionSummaryCollection {
	res := make(DeviceProvisionSummaryCollection, len(vres))
	for i, n := range vres {
		res[i] = newDeviceProvisionSummary(n)
	}
	return res
}

// newDeviceProvisionSummaryCollectionView projects result type
// DeviceProvisionSummaryCollection to projected type
// DeviceProvisionSummaryCollectionView using the "default" view.
func newDeviceProvisionSummaryCollectionView(res DeviceProvisionSummaryCollection) dataviews.DeviceProvisionSummaryCollectionView {
	vres := make(dataviews.DeviceProvisionSummaryCollectionView, len(res))
	for i, n := range res {
		vres[i] = newDeviceProvisionSummaryView(n)
	}
	return vres
}

// newDeviceProvisionSummary converts projected type DeviceProvisionSummary to
// service type DeviceProvisionSummary.
func newDeviceProvisionSummary(vres *dataviews.DeviceProvisionSummaryView) *DeviceProvisionSummary {
	res := &DeviceProvisionSummary{}
	if vres.Generation != nil {
		res.Generation = *vres.Generation
	}
	if vres.Created != nil {
		res.Created = *vres.Created
	}
	if vres.Updated != nil {
		res.Updated = *vres.Updated
	}
	if vres.Meta != nil {
		res.Meta = newDeviceMetaSummary(vres.Meta)
	}
	if vres.Data != nil {
		res.Data = newDeviceDataSummary(vres.Data)
	}
	return res
}

// newDeviceProvisionSummaryView projects result type DeviceProvisionSummary to
// projected type DeviceProvisionSummaryView using the "default" view.
func newDeviceProvisionSummaryView(res *DeviceProvisionSummary) *dataviews.DeviceProvisionSummaryView {
	vres := &dataviews.DeviceProvisionSummaryView{
		Generation: &res.Generation,
		Created:    &res.Created,
		Updated:    &res.Updated,
	}
	if res.Meta != nil {
		vres.Meta = newDeviceMetaSummaryView(res.Meta)
	}
	if res.Data != nil {
		vres.Data = newDeviceDataSummaryView(res.Data)
	}
	return vres
}

// newDeviceMetaSummary converts projected type DeviceMetaSummary to service
// type DeviceMetaSummary.
func newDeviceMetaSummary(vres *dataviews.DeviceMetaSummaryView) *DeviceMetaSummary {
	res := &DeviceMetaSummary{}
	if vres.Size != nil {
		res.Size = *vres.Size
	}
	if vres.First != nil {
		res.First = *vres.First
	}
	if vres.Last != nil {
		res.Last = *vres.Last
	}
	return res
}

// newDeviceMetaSummaryView projects result type DeviceMetaSummary to projected
// type DeviceMetaSummaryView using the "default" view.
func newDeviceMetaSummaryView(res *DeviceMetaSummary) *dataviews.DeviceMetaSummaryView {
	vres := &dataviews.DeviceMetaSummaryView{
		Size:  &res.Size,
		First: &res.First,
		Last:  &res.Last,
	}
	return vres
}

// newDeviceDataSummary converts projected type DeviceDataSummary to service
// type DeviceDataSummary.
func newDeviceDataSummary(vres *dataviews.DeviceDataSummaryView) *DeviceDataSummary {
	res := &DeviceDataSummary{}
	if vres.Size != nil {
		res.Size = *vres.Size
	}
	if vres.First != nil {
		res.First = *vres.First
	}
	if vres.Last != nil {
		res.Last = *vres.Last
	}
	return res
}

// newDeviceDataSummaryView projects result type DeviceDataSummary to projected
// type DeviceDataSummaryView using the "default" view.
func newDeviceDataSummaryView(res *DeviceDataSummary) *dataviews.DeviceDataSummaryView {
	vres := &dataviews.DeviceDataSummaryView{
		Size:  &res.Size,
		First: &res.First,
		Last:  &res.Last,
	}
	return vres
}
