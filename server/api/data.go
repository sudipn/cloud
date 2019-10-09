package api

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/goadesign/goa"
	"github.com/iancoleman/strcase"

	"github.com/conservify/sqlxcache"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/fieldkit/cloud/server/api/app"
	"github.com/fieldkit/cloud/server/backend"
	"github.com/fieldkit/cloud/server/backend/repositories"
	"github.com/fieldkit/cloud/server/data"
	"github.com/fieldkit/cloud/server/jobs"
	"github.com/fieldkit/cloud/server/messages"

	pb "github.com/fieldkit/data-protocol"
)

const (
	MetaTypeName = "meta"
	DataTypeName = "data"
)

type DataControllerOptions struct {
	Config    *ApiConfiguration
	Session   *session.Session
	Database  *sqlxcache.DB
	Publisher jobs.MessagePublisher
}

type DataController struct {
	options DataControllerOptions
	*goa.Controller
}

func NewDataController(ctx context.Context, service *goa.Service, options DataControllerOptions) *DataController {
	return &DataController{
		options:    options,
		Controller: service.NewController("DataController"),
	}
}

func (c *DataController) Process(ctx *app.ProcessDataContext) error {
	log := Logger(ctx).Sugar()

	ir, err := repositories.NewIngestionRepository(c.options.Database)
	if err != nil {
		return err
	}

	pending, err := ir.QueryPending(ctx)
	if err != nil {
		return err
	}

	recordAdder := backend.NewRecordAdder(c.options.Session, c.options.Database)

	for _, i := range pending {
		log.Infow("pending", "ingestion", i)

		err := recordAdder.WriteRecords(ctx, i)
		if err != nil {
			log.Errorw("error", "error", err)
			err := ir.MarkProcessedHasErrors(ctx, i.ID)
			if err != nil {
				return err
			}
		} else {
			err := ir.MarkProcessedDone(ctx, i.ID)
			if err != nil {
				return err
			}
		}
	}

	log.Infow("done", "elapsed", 0)

	return nil
}

func (c *DataController) ProcessIngestion(ctx *app.ProcessIngestionDataContext) error {
	p, err := NewPermissions(ctx)
	if err != nil {
		return err
	}

	log := Logger(ctx).Sugar()

	ir, err := repositories.NewIngestionRepository(c.options.Database)
	if err != nil {
		return err
	}

	log.Infow("processing", "ingestion_id", ctx.IngestionID)

	i, err := ir.QueryByID(ctx, int64(ctx.IngestionID))
	if err != nil {
		return err
	}
	if i == nil {
		return ctx.NotFound()
	}

	err = p.CanModifyStationByDeviceID(i.DeviceID)
	if err != nil {
		return err
	}

	c.options.Publisher.Publish(ctx, &messages.IngestionReceived{
		Time: i.Time,
		ID:   i.ID,
		URL:  i.URL,
	})

	return ctx.OK([]byte("queued"))
}

func (c *DataController) Delete(ctx *app.DeleteDataContext) error {
	p, err := NewPermissions(ctx)
	if err != nil {
		return err
	}

	log := Logger(ctx).Sugar()

	ir, err := repositories.NewIngestionRepository(c.options.Database)
	if err != nil {
		return err
	}

	log.Infow("deleting", "ingestion_id", ctx.IngestionID)

	i, err := ir.QueryByID(ctx, int64(ctx.IngestionID))
	if err != nil {
		return err
	}
	if i == nil {
		return ctx.NotFound()
	}

	err = p.CanModifyStationByDeviceID(i.DeviceID)
	if err != nil {
		return err
	}

	svc := s3.New(c.options.Session)

	object, err := backend.GetBucketAndKey(i.URL)
	if err != nil {
		return fmt.Errorf("Error parsing URL: %v", err)
	}

	log.Infow("deleting", "url", i.URL)

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(object.Bucket), Key: aws.String(object.Key)})
	if err != nil {
		return fmt.Errorf("Unable to delete object %q from bucket %q, %v", object.Key, object.Bucket, err)
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(object.Bucket),
		Key:    aws.String(object.Key),
	})

	if err := ir.Delete(ctx, int64(ctx.IngestionID)); err != nil {
		return err
	}

	return ctx.OK([]byte("deleted"))
}

type ProvisionSummaryRow struct {
	Generation []byte
	Type       string
	Blocks     data.Int64Range
}

type BlocksSummaryRow struct {
	Size   int64
	Blocks data.Int64Range
}

func (c *DataController) DeviceSummary(ctx *app.DeviceSummaryDataContext) error {
	p, err := NewPermissions(ctx)
	if err != nil {
		return err
	}

	deviceIdBytes, err := data.DecodeBinaryString(ctx.DeviceID)
	if err != nil {
		return err
	}

	err = p.CanViewStationByDeviceID(deviceIdBytes)
	if err != nil {
		return err
	}

	log := Logger(ctx).Sugar()

	_ = log

	log.Infow("summary", "device_id", deviceIdBytes)

	provisions := make([]*data.Provision, 0)
	if err := c.options.Database.SelectContext(ctx, &provisions, `SELECT * FROM fieldkit.provision WHERE (device_id = $1)`, deviceIdBytes); err != nil {
		return err
	}

	var rows = make([]ProvisionSummaryRow, 0)
	if err := c.options.Database.SelectContext(ctx, &rows, `SELECT generation, type, range_merge(blocks) AS blocks FROM fieldkit.ingestion WHERE (device_id = $1) GROUP BY type, generation`, deviceIdBytes); err != nil {
		return err
	}

	var blockSummaries = make(map[string]map[string]BlocksSummaryRow)
	for _, p := range provisions {
		blockSummaries[hex.EncodeToString(p.Generation)] = make(map[string]BlocksSummaryRow)
	}

	for _, row := range rows {
		g := hex.EncodeToString(row.Generation)
		blockSummaries[g][row.Type] = BlocksSummaryRow{
			Blocks: row.Blocks,
		}
	}

	provisionVms := make([]*app.DeviceProvisionSummary, len(provisions))
	for i, p := range provisions {
		g := hex.EncodeToString(p.Generation)
		byType := blockSummaries[g]
		provisionVms[i] = &app.DeviceProvisionSummary{
			Generation: g,
			Created:    p.Created,
			Updated:    p.Updated,
			Meta: &app.DeviceMetaSummary{
				Size:  int(0),
				First: int(byType[MetaTypeName].Blocks[0]),
				Last:  int(byType[MetaTypeName].Blocks[1]),
			},
			Data: &app.DeviceDataSummary{
				Size:  int(0),
				First: int(byType[DataTypeName].Blocks[0]),
				Last:  int(byType[DataTypeName].Blocks[1]),
			},
		}
	}

	return ctx.OK(&app.DeviceDataSummaryResponse{
		Provisions: provisionVms,
	})
}

func (c *DataController) DeviceData(ctx *app.DeviceDataDataContext) error {
	p, err := NewPermissions(ctx)
	if err != nil {
		return err
	}

	deviceIdBytes, err := data.DecodeBinaryString(ctx.DeviceID)
	if err != nil {
		return err
	}

	err = p.CanViewStationByDeviceID(deviceIdBytes)
	if err != nil {
		return err
	}

	log := Logger(ctx).Sugar()

	_ = log

	rr, err := repositories.NewRecordRepository(c.options.Database)
	if err != nil {
		return err
	}

	pageNumber := 0
	if ctx.PageNumber != nil {
		pageNumber = *ctx.PageNumber
	}

	pageSize := 200
	if ctx.PageSize != nil {
		pageSize = *ctx.PageSize
	}

	page, err := rr.QueryDevice(ctx, ctx.DeviceID, pageNumber, pageSize)
	if err != nil {
		return err
	}

	dataVms := make([]*app.DeviceDataRecord, len(page.Data))
	for i, r := range page.Data {
		data, err := r.GetData()
		if err != nil {
			return err
		}

		coordinates := []float64{}
		if r.Location != nil {
			coordinates = r.Location.Coordinates()
		}

		dataVms[i] = &app.DeviceDataRecord{
			ID:       int(r.ID),
			Time:     r.Time,
			Record:   int(r.Number),
			Meta:     int(r.Meta),
			Location: coordinates,
			Data:     data,
		}
	}

	metaVms := make([]*app.DeviceMetaRecord, len(page.Meta))
	for i, r := range page.Meta {
		data, err := r.GetData()
		if err != nil {
			return err
		}

		metaVms[i] = &app.DeviceMetaRecord{
			ID:     int(r.ID),
			Time:   r.Time,
			Record: int(r.Number),
			Data:   data,
		}
	}

	return ctx.OK(&app.DeviceDataRecordsResponse{
		Meta: metaVms,
		Data: dataVms,
	})
}

type JSONDataController struct {
	options DataControllerOptions
	*goa.Controller
}

func NewJSONDataController(ctx context.Context, service *goa.Service, options DataControllerOptions) *JSONDataController {
	return &JSONDataController{
		options:    options,
		Controller: service.NewController("JSONDataController"),
	}
}

func getLocation(l *pb.DeviceLocation) []float64 {
	if l == nil {
		return nil
	}
	if l.Longitude > 180 || l.Longitude < -180 {
		return nil
	}
	return []float64{
		float64(l.Longitude),
		float64(l.Latitude),
		float64(l.Altitude),
	}
}

func isInternalModule(m *pb.ModuleInfo) bool {
	return m.Name == "random" || m.Name == "diagnostics"
}

func (c *JSONDataController) Get(ctx *app.GetJSONDataContext) error {
	log := Logger(ctx).Sugar()

	p, err := NewPermissions(ctx)
	if err != nil {
		return err
	}

	deviceIdBytes, err := data.DecodeBinaryString(ctx.DeviceID)
	if err != nil {
		return err
	}

	err = p.CanViewStationByDeviceID(deviceIdBytes)
	if err != nil {
		return err
	}

	rr, err := repositories.NewRecordRepository(c.options.Database)
	if err != nil {
		return err
	}

	pageNumber := 0
	if ctx.PageNumber != nil {
		pageNumber = *ctx.PageNumber
	}

	pageSize := 200
	if ctx.PageSize != nil {
		pageSize = *ctx.PageSize
	}

	internal := false
	if ctx.Internal != nil {
		internal = *ctx.Internal
	}

	log.Infow("querying", "device_id", ctx.DeviceID, "page_number", pageNumber, "page_size", pageSize, "internal", internal)

	page, err := rr.QueryDevice(ctx, ctx.DeviceID, pageNumber, pageSize)
	if err != nil {
		return err
	}

	log.Infow("querying", "nmeta", len(page.Meta), "ndata", len(page.Data))

	byMeta := make(map[int64][]*data.DataRecord)
	for _, d := range page.Data {
		if byMeta[d.Meta] == nil {
			byMeta[d.Meta] = make([]*data.DataRecord, 0)
		}
		byMeta[d.Meta] = append(byMeta[d.Meta], d)
	}

	sr, err := repositories.NewStationRepository(c.options.Database)
	if err != nil {
		return err
	}

	station, err := sr.QueryStationByDeviceID(ctx, deviceIdBytes)
	if err != nil {
		return err
	}

	log.Infow("querying", "station_id", station.ID, "station_name", station.Name)

	versions := make([]*app.JSONDataVersion, 0)
	for _, m := range page.Meta {
		dataRecords := byMeta[m.ID]

		var metaRecord pb.DataRecord
		err := m.Unmarshal(&metaRecord)
		if err != nil {
			return err
		}

		name := metaRecord.Identity.Name
		if name == "" {
			name = station.Name
		}

		modules := make([]*app.JSONDataMetaModule, 0)
		for _, module := range metaRecord.Modules {
			if internal || !isInternalModule(module) {
				sensors := make([]*app.JSONDataMetaSensor, 0)
				for _, sensor := range module.Sensors {
					sensors = append(sensors, &app.JSONDataMetaSensor{
						Name:  sensor.Name,
						Key:   strcase.ToLowerCamel(sensor.Name),
						Units: sensor.UnitOfMeasure,
					})
				}

				modules = append(modules, &app.JSONDataMetaModule{
					Name:         module.Name,
					ID:           hex.EncodeToString(module.Id),
					Manufacturer: int(module.Header.Manufacturer),
					Kind:         int(module.Header.Kind),
					Version:      int(module.Header.Version),
					Sensors:      sensors,
				})
			}
		}

		rows := make([]*app.JSONDataRow, 0)
		for _, d := range dataRecords {
			var dataRecord pb.DataRecord
			err := d.Unmarshal(&dataRecord)
			if err != nil {
				return err
			}

			data := make(map[string]interface{})
			for mi, sg := range dataRecord.Readings.SensorGroups {
				if sg.Module == 255 {
					continue
				}
				if mi >= len(metaRecord.Modules) {
					log.Warnw("module index range", "module_index", mi, "modules", len(metaRecord.Modules), "meta", metaRecord.Modules)
					continue
				}
				module := metaRecord.Modules[mi]
				if internal || !isInternalModule(module) {
					for si, r := range sg.Readings {
						sensor := module.Sensors[si]
						key := strcase.ToLowerCamel(sensor.Name)
						data[key] = r.Value
					}
				}
			}

			l := dataRecord.Readings.Location
			rows = append(rows, &app.JSONDataRow{
				ID:       int(d.ID),
				Time:     int(dataRecord.Readings.Time),
				Location: getLocation(l),
				D:        data,
			})
		}

		if len(rows) > 0 {
			versions = append(versions, &app.JSONDataVersion{
				Meta: &app.JSONDataMeta{
					ID: int(m.ID),
					Station: &app.JSONDataMetaStation{
						ID:      hex.EncodeToString(metaRecord.Metadata.DeviceId),
						Name:    name,
						Modules: modules,
						Firmware: &app.JSONDataMetaStationFirmware{
							Git:   metaRecord.Metadata.Firmware.Git,
							Build: metaRecord.Metadata.Firmware.Build,
						},
					},
				},
				Data: rows,
			})
		}
	}

	return ctx.OK(&app.JSONDataResponse{
		Versions: versions,
	})
}
