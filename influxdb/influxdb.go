package influxdb

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/util/config"
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	kernel.RegisterAPI((*Pool)(nil), &pool{})
}

type Pool interface {
	GetDB(string) *Config
	StoreReading(ctx context.Context) error
}

type pool struct {
	ConfigManager config.Manager `kernel:"inject"`
	Worker        task.Queue     `kernel:"worker"`
	Brokers       *map[string]*Config
}

func (p *pool) Start() error {
	m := make(map[string]*Config)
	p.Brokers = &m
	return p.ConfigManager.ReadYamlOptional("influxdb.yaml", p.Brokers)
}

func (p *pool) GetDB(n string) *Config {
	return (*p.Brokers)[n]
}

func (p *pool) StoreReading(ctx context.Context) error {
	s := station.SensorsFromContext(ctx)
	if s != nil && s.Output != nil && s.Output.InfluxDB != nil {
		payloadEntry := payload.GetPayload(ctx)

		db := s.Output.InfluxDB
		broker := (*p.Brokers)[db.Name]
		if broker == nil {
			log.Printf("InfluxDB %q undefined!", db.Name)
		}

		values := value.MapFromContext(ctx)
		for _, key := range values.GetKeys() {
			val := values.Get(key)
			fields := map[string]interface{}{
				"value": val.Float(),
			}

			tags := map[string]string{
				"stationId": s.Station().ID,
				"sensorId":  s.ID,
				"unit":      val.Unit().ID(),
			}

			measurement := key
			if db.Measurement != "" {
				measurement = db.Measurement + "." + key
			}

			point := write.NewPoint(measurement, tags, fields, payloadEntry.Time().UTC())

			// Put onto worker queue, so we don't hold things up
			p.Worker.AddTask(func(ctx context.Context) error {
				err := broker.WriteAPIBlocking().WritePoint(context.Background(), point)
				if err != nil {
					log.Printf("[InfluxDB] WritePoint: %v", err)
				}
				return nil
			})
		}
	}

	return nil
}
