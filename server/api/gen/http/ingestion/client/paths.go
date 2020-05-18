// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the ingestion service.
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

// ProcessPendingIngestionPath returns the URL path to the ingestion service process pending HTTP endpoint.
func ProcessPendingIngestionPath() string {
	return "/data/process"
}

// ProcessStationIngestionPath returns the URL path to the ingestion service process station HTTP endpoint.
func ProcessStationIngestionPath() string {
	return "/data/stations/:stationId/process"
}

// ProcessIngestionIngestionPath returns the URL path to the ingestion service process ingestion HTTP endpoint.
func ProcessIngestionIngestionPath() string {
	return "/data/ingestions/:ingestionId/process"
}

// DeleteIngestionPath returns the URL path to the ingestion service delete HTTP endpoint.
func DeleteIngestionPath() string {
	return "/data/ingestions/:ingestionId"
}
