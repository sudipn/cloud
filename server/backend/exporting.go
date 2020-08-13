package backend

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/conservify/sqlxcache"

	"github.com/fieldkit/cloud/server/backend/repositories"
	"github.com/fieldkit/cloud/server/common/logging"
	"github.com/fieldkit/cloud/server/data"
	"github.com/fieldkit/cloud/server/files"
	"github.com/fieldkit/cloud/server/messages"
)

type ExportCsvHandler struct {
	db      *sqlxcache.DB
	files   files.FileArchive
	metrics *logging.Metrics
}

func NewExportCsvHandler(db *sqlxcache.DB, files files.FileArchive, metrics *logging.Metrics) *ExportCsvHandler {
	return &ExportCsvHandler{
		db:      db,
		files:   files,
		metrics: metrics,
	}
}

func (h *ExportCsvHandler) Handle(ctx context.Context, m *messages.ExportCsv) error {
	log := Logger(ctx).Sugar().Named("exporting").With("data_export_id", m.ID).With("user_id", m.UserID)

	log.Infow("processing")

	r, err := repositories.NewExportRepository(h.db)
	if err != nil {
		return err
	}

	de, err := r.QueryByID(ctx, m.ID)
	if err != nil {
		return err
	}

	rawParams := &RawQueryParams{}
	if err := json.Unmarshal(de.Args, rawParams); err != nil {
		return err
	}

	qp, err := rawParams.BuildQueryParams()
	if err != nil {
		return fmt.Errorf("invalid query params: %v", err)
	}

	aggregateName := "1m"
	qp.Aggregate = aggregateName

	log.Infow("query_parameters", "start", qp.Start, "end", qp.End, "sensors", qp.Sensors, "stations", qp.Stations,
		"resolution", qp.Resolution, "aggregate", qp.Aggregate, "complete", qp.Complete)

	aqp, err := NewAggregateQueryParams(qp, aggregateName, nil)
	if err != nil {
		return err
	}

	dq := NewDataQuerier(h.db)

	qm, err := dq.QueryMeta(ctx, qp)
	if err != nil {
		return err
	}

	queried, err := dq.QueryAggregate(ctx, aqp)
	if err != nil {
		return err
	}

	defer queried.Close()

	// One end we write to and the reading ends ends up in S3.
	pipeReader, pipeWriter := io.Pipe()

	// Spawn a go routine that will read the other end of the pip and
	// archive the file to S3. We also wait on this goroutine to
	// finish, and then propagate any error upwards.
	var wg sync.WaitGroup
	var archiveError error

	wg.Add(1)

	go func() {
		log.Infow("archiver:ready")

		metadata := make(map[string]string)
		contentType := "text/csv"
		af, err := h.files.Archive(ctx, contentType, metadata, pipeReader)
		if err != nil {
			log.Errorw("archiver:error", "error", err)
			archiveError = err
		} else {
			log.Infow("archiver:done", "key", af.Key)
		}

		now := time.Now()
		size := int32(af.BytesRead)

		de.DownloadURL = &af.URL
		de.CompletedAt = &now
		de.Progress = 100
		de.Size = &size

		log.Infow("done", "bytes", size)

		if _, err := r.UpdateDataExport(ctx, de); err != nil {
			archiveError = err
		}

		wg.Done()
	}()

	// Walk the rows and dump their CSV representation.
	writer := csv.NewWriter(pipeWriter)
	records := 0

	writer.Write(makeHeader())

	for queried.Next() {
		row := &DataRow{}
		if err = queried.StructScan(row); err != nil {
			return err
		}
		writer.Write(makeRow(qm, row))
		records += 1
	}

	writer.Flush()

	pipeWriter.Close()

	log.Infow("waiting")

	wg.Wait()

	log.Infow("done", "records", records)

	if archiveError != nil {
		return archiveError
	}

	return nil
}

const (
	NumberOfColumns = 7
)

func makeHeader() []string {
	cols := make([]string, NumberOfColumns)
	cols[0] = "Time (JS)"
	cols[1] = "Station"
	cols[2] = "Sensor"
	cols[3] = "Latitude"
	cols[4] = "Longitude"
	cols[5] = "Reading"
	cols[6] = "Time Group"
	return cols
}

func makeRow(qm *QueryMeta, row *DataRow) []string {
	cols := make([]string, NumberOfColumns)
	cols[0] = fmt.Sprintf("%v", row.Time.Time().Unix()*1000)
	if row.StationID != nil {
		cols[1] = fmt.Sprintf("%v", qm.Stations[*row.StationID].Name)
	}
	if row.SensorID != nil {
		cols[2] = fmt.Sprintf("%v", qm.Sensors[*row.SensorID].Key)
	}
	if row.Location != nil {
		cols[3] = fmt.Sprintf("%v", row.Location.Latitude())
		cols[4] = fmt.Sprintf("%v", row.Location.Longitude())
	}
	if row.Value != nil {
		cols[5] = fmt.Sprintf("%v", *row.Value)
	}
	if row.TimeGroup != nil {
		cols[6] = fmt.Sprintf("%v", *row.TimeGroup)
	}
	return cols
}

type DataRow struct {
	Time      data.NumericWireTime `db:"time" json:"time"`
	ID        *int64               `db:"id" json:"-"`
	StationID *int32               `db:"station_id" json:"stationId,omitempty"`
	SensorID  *int64               `db:"sensor_id" json:"sensorId,omitempty"`
	Location  *data.Location       `db:"location" json:"location,omitempty"`
	Value     *float64             `db:"value" json:"value,omitempty"`
	TimeGroup *int32               `db:"time_group" json:"tg,omitempty"`
}
