// Code generated by goa v3.1.2, DO NOT EDIT.
//
// station views
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// StationFull is the viewed result type that is projected based on a view.
type StationFull struct {
	// Type to project
	Projected *StationFullView
	// View to render
	View string
}

// StationsFull is the viewed result type that is projected based on a view.
type StationsFull struct {
	// Type to project
	Projected *StationsFullView
	// View to render
	View string
}

// PageOfStations is the viewed result type that is projected based on a view.
type PageOfStations struct {
	// Type to project
	Projected *PageOfStationsView
	// View to render
	View string
}

// StationFullView is a type that runs validations on a projected type.
type StationFullView struct {
	ID                 *int32
	Name               *string
	Owner              *StationOwnerView
	DeviceID           *string
	Uploads            []*StationUploadView
	Images             []*ImageRefView
	Photos             *StationPhotosView
	ReadOnly           *bool
	Battery            *float32
	RecordingStartedAt *int64
	MemoryUsed         *int32
	MemoryAvailable    *int32
	FirmwareNumber     *int32
	FirmwareTime       *int64
	Configurations     *StationConfigurationsView
	Updated            *int64
	LocationName       *string
	PlaceNameOther     *string
	PlaceNameNative    *string
	Location           *StationLocationView
}

// StationOwnerView is a type that runs validations on a projected type.
type StationOwnerView struct {
	ID   *int32
	Name *string
}

// StationUploadView is a type that runs validations on a projected type.
type StationUploadView struct {
	ID       *int64
	Time     *int64
	UploadID *string
	Size     *int64
	URL      *string
	Type     *string
	Blocks   []int64
}

// ImageRefView is a type that runs validations on a projected type.
type ImageRefView struct {
	URL *string
}

// StationPhotosView is a type that runs validations on a projected type.
type StationPhotosView struct {
	Small *string
}

// StationConfigurationsView is a type that runs validations on a projected
// type.
type StationConfigurationsView struct {
	All []*StationConfigurationView
}

// StationConfigurationView is a type that runs validations on a projected type.
type StationConfigurationView struct {
	ID           *int64
	Time         *int64
	ProvisionID  *int64
	MetaRecordID *int64
	SourceID     *int32
	Modules      []*StationModuleView
}

// StationModuleView is a type that runs validations on a projected type.
type StationModuleView struct {
	ID           *int64
	HardwareID   *string
	MetaRecordID *int64
	Name         *string
	Position     *int32
	Flags        *int32
	Internal     *bool
	FullKey      *string
	Sensors      []*StationSensorView
}

// StationSensorView is a type that runs validations on a projected type.
type StationSensorView struct {
	Name          *string
	UnitOfMeasure *string
	Reading       *SensorReadingView
	Key           *string
	FullKey       *string
	Ranges        []*SensorRangeView
}

// SensorReadingView is a type that runs validations on a projected type.
type SensorReadingView struct {
	Last *float32
	Time *int64
}

// SensorRangeView is a type that runs validations on a projected type.
type SensorRangeView struct {
	Minimum *float32
	Maximum *float32
}

// StationLocationView is a type that runs validations on a projected type.
type StationLocationView struct {
	Latitude  *float64
	Longitude *float64
}

// StationsFullView is a type that runs validations on a projected type.
type StationsFullView struct {
	Stations StationFullCollectionView
}

// StationFullCollectionView is a type that runs validations on a projected
// type.
type StationFullCollectionView []*StationFullView

// PageOfStationsView is a type that runs validations on a projected type.
type PageOfStationsView struct {
	Stations []*EssentialStationView
}

// EssentialStationView is a type that runs validations on a projected type.
type EssentialStationView struct {
	ID                 *int64
	DeviceID           *string
	Name               *string
	Owner              *StationOwnerView
	CreatedAt          *int64
	UpdatedAt          *int64
	RecordingStartedAt *int64
	MemoryUsed         *int32
	MemoryAvailable    *int32
	FirmwareNumber     *int32
	FirmwareTime       *int64
	Location           *StationLocationView
	LastIngestionAt    *int64
}

var (
	// StationFullMap is a map of attribute names in result type StationFull
	// indexed by view name.
	StationFullMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"owner",
			"deviceId",
			"uploads",
			"images",
			"photos",
			"readOnly",
			"battery",
			"recordingStartedAt",
			"memoryUsed",
			"memoryAvailable",
			"firmwareNumber",
			"firmwareTime",
			"configurations",
			"updated",
			"location",
			"locationName",
			"placeNameOther",
			"placeNameNative",
		},
	}
	// StationsFullMap is a map of attribute names in result type StationsFull
	// indexed by view name.
	StationsFullMap = map[string][]string{
		"default": []string{
			"stations",
		},
	}
	// PageOfStationsMap is a map of attribute names in result type PageOfStations
	// indexed by view name.
	PageOfStationsMap = map[string][]string{
		"default": []string{
			"stations",
		},
	}
	// StationFullCollectionMap is a map of attribute names in result type
	// StationFullCollection indexed by view name.
	StationFullCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"owner",
			"deviceId",
			"uploads",
			"images",
			"photos",
			"readOnly",
			"battery",
			"recordingStartedAt",
			"memoryUsed",
			"memoryAvailable",
			"firmwareNumber",
			"firmwareTime",
			"configurations",
			"updated",
			"location",
			"locationName",
			"placeNameOther",
			"placeNameNative",
		},
	}
)

// ValidateStationFull runs the validations defined on the viewed result type
// StationFull.
func ValidateStationFull(result *StationFull) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStationFullView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStationsFull runs the validations defined on the viewed result type
// StationsFull.
func ValidateStationsFull(result *StationsFull) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStationsFullView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidatePageOfStations runs the validations defined on the viewed result
// type PageOfStations.
func ValidatePageOfStations(result *PageOfStations) (err error) {
	switch result.View {
	case "default", "":
		err = ValidatePageOfStationsView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStationFullView runs the validations defined on StationFullView
// using the "default" view.
func ValidateStationFullView(result *StationFullView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Owner == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("owner", "result"))
	}
	if result.DeviceID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("deviceId", "result"))
	}
	if result.Uploads == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("uploads", "result"))
	}
	if result.Images == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("images", "result"))
	}
	if result.Photos == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("photos", "result"))
	}
	if result.ReadOnly == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("readOnly", "result"))
	}
	if result.Configurations == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("configurations", "result"))
	}
	if result.Updated == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated", "result"))
	}
	if result.Owner != nil {
		if err2 := ValidateStationOwnerView(result.Owner); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range result.Uploads {
		if e != nil {
			if err2 := ValidateStationUploadView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range result.Images {
		if e != nil {
			if err2 := ValidateImageRefView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Photos != nil {
		if err2 := ValidateStationPhotosView(result.Photos); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Configurations != nil {
		if err2 := ValidateStationConfigurationsView(result.Configurations); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Location != nil {
		if err2 := ValidateStationLocationView(result.Location); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStationOwnerView runs the validations defined on StationOwnerView.
func ValidateStationOwnerView(result *StationOwnerView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}

// ValidateStationUploadView runs the validations defined on StationUploadView.
func ValidateStationUploadView(result *StationUploadView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Time == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("time", "result"))
	}
	if result.UploadID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("uploadId", "result"))
	}
	if result.Size == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("size", "result"))
	}
	if result.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Blocks == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("blocks", "result"))
	}
	return
}

// ValidateImageRefView runs the validations defined on ImageRefView.
func ValidateImageRefView(result *ImageRefView) (err error) {
	if result.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "result"))
	}
	return
}

// ValidateStationPhotosView runs the validations defined on StationPhotosView.
func ValidateStationPhotosView(result *StationPhotosView) (err error) {
	if result.Small == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("small", "result"))
	}
	return
}

// ValidateStationConfigurationsView runs the validations defined on
// StationConfigurationsView.
func ValidateStationConfigurationsView(result *StationConfigurationsView) (err error) {
	if result.All == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("all", "result"))
	}
	for _, e := range result.All {
		if e != nil {
			if err2 := ValidateStationConfigurationView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStationConfigurationView runs the validations defined on
// StationConfigurationView.
func ValidateStationConfigurationView(result *StationConfigurationView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.ProvisionID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("provisionId", "result"))
	}
	if result.Time == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("time", "result"))
	}
	if result.Modules == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("modules", "result"))
	}
	for _, e := range result.Modules {
		if e != nil {
			if err2 := ValidateStationModuleView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStationModuleView runs the validations defined on StationModuleView.
func ValidateStationModuleView(result *StationModuleView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Position == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("position", "result"))
	}
	if result.Flags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("flags", "result"))
	}
	if result.Internal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("internal", "result"))
	}
	if result.FullKey == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fullKey", "result"))
	}
	if result.Sensors == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sensors", "result"))
	}
	for _, e := range result.Sensors {
		if e != nil {
			if err2 := ValidateStationSensorView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStationSensorView runs the validations defined on StationSensorView.
func ValidateStationSensorView(result *StationSensorView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.UnitOfMeasure == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("unitOfMeasure", "result"))
	}
	if result.Key == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("key", "result"))
	}
	if result.FullKey == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fullKey", "result"))
	}
	if result.Ranges == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ranges", "result"))
	}
	if result.Reading != nil {
		if err2 := ValidateSensorReadingView(result.Reading); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range result.Ranges {
		if e != nil {
			if err2 := ValidateSensorRangeView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateSensorReadingView runs the validations defined on SensorReadingView.
func ValidateSensorReadingView(result *SensorReadingView) (err error) {
	if result.Last == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last", "result"))
	}
	if result.Time == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("time", "result"))
	}
	return
}

// ValidateSensorRangeView runs the validations defined on SensorRangeView.
func ValidateSensorRangeView(result *SensorRangeView) (err error) {
	if result.Minimum == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("minimum", "result"))
	}
	if result.Maximum == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("maximum", "result"))
	}
	return
}

// ValidateStationLocationView runs the validations defined on
// StationLocationView.
func ValidateStationLocationView(result *StationLocationView) (err error) {
	if result.Latitude == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("latitude", "result"))
	}
	if result.Longitude == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("longitude", "result"))
	}
	return
}

// ValidateStationsFullView runs the validations defined on StationsFullView
// using the "default" view.
func ValidateStationsFullView(result *StationsFullView) (err error) {

	if result.Stations != nil {
		if err2 := ValidateStationFullCollectionView(result.Stations); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStationFullCollectionView runs the validations defined on
// StationFullCollectionView using the "default" view.
func ValidateStationFullCollectionView(result StationFullCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStationFullView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePageOfStationsView runs the validations defined on
// PageOfStationsView using the "default" view.
func ValidatePageOfStationsView(result *PageOfStationsView) (err error) {
	if result.Stations == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("stations", "result"))
	}
	for _, e := range result.Stations {
		if e != nil {
			if err2 := ValidateEssentialStationView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateEssentialStationView runs the validations defined on
// EssentialStationView.
func ValidateEssentialStationView(result *EssentialStationView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.DeviceID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("deviceId", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Owner == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("owner", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("createdAt", "result"))
	}
	if result.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updatedAt", "result"))
	}
	if result.Owner != nil {
		if err2 := ValidateStationOwnerView(result.Owner); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Location != nil {
		if err2 := ValidateStationLocationView(result.Location); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
