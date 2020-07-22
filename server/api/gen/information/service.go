// Code generated by goa v3.1.2, DO NOT EDIT.
//
// information service
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package information

import (
	"context"

	informationviews "github.com/fieldkit/cloud/server/api/gen/information/views"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the information service interface.
type Service interface {
	// DeviceLayout implements device layout.
	DeviceLayout(context.Context, *DeviceLayoutPayload) (res *DeviceLayoutResponse, err error)
	// FirmwareStatistics implements firmware statistics.
	FirmwareStatistics(context.Context, *FirmwareStatisticsPayload) (res *FirmwareStatisticsResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "information"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"device layout", "firmware statistics"}

// DeviceLayoutPayload is the payload type of the information service device
// layout method.
type DeviceLayoutPayload struct {
	Auth     string
	DeviceID string
}

// DeviceLayoutResponse is the result type of the information service device
// layout method.
type DeviceLayoutResponse struct {
	Configurations []*StationConfiguration
	Sensors        map[string][]*StationSensor
}

// FirmwareStatisticsPayload is the payload type of the information service
// firmware statistics method.
type FirmwareStatisticsPayload struct {
	Auth string
}

// FirmwareStatisticsResult is the result type of the information service
// firmware statistics method.
type FirmwareStatisticsResult struct {
	Object interface{}
}

type StationConfiguration struct {
	ID           int64
	Time         int64
	ProvisionID  int64
	MetaRecordID *int64
	SourceID     *int32
	Modules      []*StationModule
}

type StationModule struct {
	ID           int64
	HardwareID   *string
	MetaRecordID *int64
	Name         string
	Position     int32
	Flags        int32
	Internal     bool
	FullKey      string
	Sensors      []*StationSensor
}

type StationSensor struct {
	Name          string
	UnitOfMeasure string
	Reading       *SensorReading
	Key           string
	FullKey       string
	Ranges        []*SensorRange
}

type SensorReading struct {
	Last float32
	Time int64
}

type SensorRange struct {
	Minimum float32
	Maximum float32
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

// NewDeviceLayoutResponse initializes result type DeviceLayoutResponse from
// viewed result type DeviceLayoutResponse.
func NewDeviceLayoutResponse(vres *informationviews.DeviceLayoutResponse) *DeviceLayoutResponse {
	return newDeviceLayoutResponse(vres.Projected)
}

// NewViewedDeviceLayoutResponse initializes viewed result type
// DeviceLayoutResponse from result type DeviceLayoutResponse using the given
// view.
func NewViewedDeviceLayoutResponse(res *DeviceLayoutResponse, view string) *informationviews.DeviceLayoutResponse {
	p := newDeviceLayoutResponseView(res)
	return &informationviews.DeviceLayoutResponse{Projected: p, View: "default"}
}

// newDeviceLayoutResponse converts projected type DeviceLayoutResponse to
// service type DeviceLayoutResponse.
func newDeviceLayoutResponse(vres *informationviews.DeviceLayoutResponseView) *DeviceLayoutResponse {
	res := &DeviceLayoutResponse{}
	if vres.Configurations != nil {
		res.Configurations = make([]*StationConfiguration, len(vres.Configurations))
		for i, val := range vres.Configurations {
			res.Configurations[i] = transformInformationviewsStationConfigurationViewToStationConfiguration(val)
		}
	}
	if vres.Sensors != nil {
		res.Sensors = make(map[string][]*StationSensor, len(vres.Sensors))
		for key, val := range vres.Sensors {
			tk := key
			tv := make([]*StationSensor, len(val))
			for i, val := range val {
				tv[i] = transformInformationviewsStationSensorViewToStationSensor(val)
			}
			res.Sensors[tk] = tv
		}
	}
	return res
}

// newDeviceLayoutResponseView projects result type DeviceLayoutResponse to
// projected type DeviceLayoutResponseView using the "default" view.
func newDeviceLayoutResponseView(res *DeviceLayoutResponse) *informationviews.DeviceLayoutResponseView {
	vres := &informationviews.DeviceLayoutResponseView{}
	if res.Configurations != nil {
		vres.Configurations = make([]*informationviews.StationConfigurationView, len(res.Configurations))
		for i, val := range res.Configurations {
			vres.Configurations[i] = transformStationConfigurationToInformationviewsStationConfigurationView(val)
		}
	}
	if res.Sensors != nil {
		vres.Sensors = make(map[string][]*informationviews.StationSensorView, len(res.Sensors))
		for key, val := range res.Sensors {
			tk := key
			tv := make([]*informationviews.StationSensorView, len(val))
			for i, val := range val {
				tv[i] = transformStationSensorToInformationviewsStationSensorView(val)
			}
			vres.Sensors[tk] = tv
		}
	}
	return vres
}

// transformInformationviewsStationConfigurationViewToStationConfiguration
// builds a value of type *StationConfiguration from a value of type
// *informationviews.StationConfigurationView.
func transformInformationviewsStationConfigurationViewToStationConfiguration(v *informationviews.StationConfigurationView) *StationConfiguration {
	if v == nil {
		return nil
	}
	res := &StationConfiguration{
		ID:           *v.ID,
		Time:         *v.Time,
		ProvisionID:  *v.ProvisionID,
		MetaRecordID: v.MetaRecordID,
		SourceID:     v.SourceID,
	}
	if v.Modules != nil {
		res.Modules = make([]*StationModule, len(v.Modules))
		for i, val := range v.Modules {
			res.Modules[i] = transformInformationviewsStationModuleViewToStationModule(val)
		}
	}

	return res
}

// transformInformationviewsStationModuleViewToStationModule builds a value of
// type *StationModule from a value of type *informationviews.StationModuleView.
func transformInformationviewsStationModuleViewToStationModule(v *informationviews.StationModuleView) *StationModule {
	res := &StationModule{
		ID:           *v.ID,
		HardwareID:   v.HardwareID,
		MetaRecordID: v.MetaRecordID,
		Name:         *v.Name,
		Position:     *v.Position,
		Flags:        *v.Flags,
		Internal:     *v.Internal,
		FullKey:      *v.FullKey,
	}
	if v.Sensors != nil {
		res.Sensors = make([]*StationSensor, len(v.Sensors))
		for i, val := range v.Sensors {
			res.Sensors[i] = transformInformationviewsStationSensorViewToStationSensor(val)
		}
	}

	return res
}

// transformInformationviewsStationSensorViewToStationSensor builds a value of
// type *StationSensor from a value of type *informationviews.StationSensorView.
func transformInformationviewsStationSensorViewToStationSensor(v *informationviews.StationSensorView) *StationSensor {
	res := &StationSensor{
		Name:          *v.Name,
		UnitOfMeasure: *v.UnitOfMeasure,
		Key:           *v.Key,
		FullKey:       *v.FullKey,
	}
	if v.Reading != nil {
		res.Reading = transformInformationviewsSensorReadingViewToSensorReading(v.Reading)
	}
	if v.Ranges != nil {
		res.Ranges = make([]*SensorRange, len(v.Ranges))
		for i, val := range v.Ranges {
			res.Ranges[i] = transformInformationviewsSensorRangeViewToSensorRange(val)
		}
	}

	return res
}

// transformInformationviewsSensorReadingViewToSensorReading builds a value of
// type *SensorReading from a value of type *informationviews.SensorReadingView.
func transformInformationviewsSensorReadingViewToSensorReading(v *informationviews.SensorReadingView) *SensorReading {
	if v == nil {
		return nil
	}
	res := &SensorReading{
		Last: *v.Last,
		Time: *v.Time,
	}

	return res
}

// transformInformationviewsSensorRangeViewToSensorRange builds a value of type
// *SensorRange from a value of type *informationviews.SensorRangeView.
func transformInformationviewsSensorRangeViewToSensorRange(v *informationviews.SensorRangeView) *SensorRange {
	res := &SensorRange{
		Minimum: *v.Minimum,
		Maximum: *v.Maximum,
	}

	return res
}

// transformStationConfigurationToInformationviewsStationConfigurationView
// builds a value of type *informationviews.StationConfigurationView from a
// value of type *StationConfiguration.
func transformStationConfigurationToInformationviewsStationConfigurationView(v *StationConfiguration) *informationviews.StationConfigurationView {
	res := &informationviews.StationConfigurationView{
		ID:           &v.ID,
		Time:         &v.Time,
		ProvisionID:  &v.ProvisionID,
		MetaRecordID: v.MetaRecordID,
		SourceID:     v.SourceID,
	}
	if v.Modules != nil {
		res.Modules = make([]*informationviews.StationModuleView, len(v.Modules))
		for i, val := range v.Modules {
			res.Modules[i] = transformStationModuleToInformationviewsStationModuleView(val)
		}
	}

	return res
}

// transformStationModuleToInformationviewsStationModuleView builds a value of
// type *informationviews.StationModuleView from a value of type *StationModule.
func transformStationModuleToInformationviewsStationModuleView(v *StationModule) *informationviews.StationModuleView {
	res := &informationviews.StationModuleView{
		ID:           &v.ID,
		HardwareID:   v.HardwareID,
		MetaRecordID: v.MetaRecordID,
		Name:         &v.Name,
		Position:     &v.Position,
		Flags:        &v.Flags,
		Internal:     &v.Internal,
		FullKey:      &v.FullKey,
	}
	if v.Sensors != nil {
		res.Sensors = make([]*informationviews.StationSensorView, len(v.Sensors))
		for i, val := range v.Sensors {
			res.Sensors[i] = transformStationSensorToInformationviewsStationSensorView(val)
		}
	}

	return res
}

// transformStationSensorToInformationviewsStationSensorView builds a value of
// type *informationviews.StationSensorView from a value of type *StationSensor.
func transformStationSensorToInformationviewsStationSensorView(v *StationSensor) *informationviews.StationSensorView {
	res := &informationviews.StationSensorView{
		Name:          &v.Name,
		UnitOfMeasure: &v.UnitOfMeasure,
		Key:           &v.Key,
		FullKey:       &v.FullKey,
	}
	if v.Reading != nil {
		res.Reading = transformSensorReadingToInformationviewsSensorReadingView(v.Reading)
	}
	if v.Ranges != nil {
		res.Ranges = make([]*informationviews.SensorRangeView, len(v.Ranges))
		for i, val := range v.Ranges {
			res.Ranges[i] = transformSensorRangeToInformationviewsSensorRangeView(val)
		}
	}

	return res
}

// transformSensorReadingToInformationviewsSensorReadingView builds a value of
// type *informationviews.SensorReadingView from a value of type *SensorReading.
func transformSensorReadingToInformationviewsSensorReadingView(v *SensorReading) *informationviews.SensorReadingView {
	if v == nil {
		return nil
	}
	res := &informationviews.SensorReadingView{
		Last: &v.Last,
		Time: &v.Time,
	}

	return res
}

// transformSensorRangeToInformationviewsSensorRangeView builds a value of type
// *informationviews.SensorRangeView from a value of type *SensorRange.
func transformSensorRangeToInformationviewsSensorRangeView(v *SensorRange) *informationviews.SensorRangeView {
	res := &informationviews.SensorRangeView{
		Minimum: &v.Minimum,
		Maximum: &v.Maximum,
	}

	return res
}
