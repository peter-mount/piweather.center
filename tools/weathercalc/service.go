package weathercalc

import (
	"context"
	"encoding/json"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/mq/amqp"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/tools/weathercalc/lang"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/rabbitmq/amqp091-go"
	"time"
)

type Service struct {
	Latest         memory.Latest         `kernel:"inject"`
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	Calculations   *Calculations         `kernel:"inject"`
	Calculator     *Calculator           `kernel:"inject"`
	Cron           *cron.CronService     `kernel:"inject"`
	Daemon         *kernel.Daemon        `kernel:"inject"`
	//Config         service.Config        `kernel:"inject"`
	QueueName *string `kernel:"flag,metric-queue,DB queue name,database.calc"`
	DBServer  *string `kernel:"flag,metric-db,DB url"`
	mqQueue   *amqp.Queue
}

func (s *Service) Start() error {

	if *s.DBServer != "" {
		s.initFromDB()
		// Reload from the DB at 00:01
		// This allows for 1 minute for some data to arrive before
		// we refresh the metrics
		if _, err := s.Cron.AddFunc("1 0 * * *", s.initFromDB); err != nil {
			return err
		}
	}

	// Get latest metrics from DB
	if err := s.loadLatestMetrics(); err != nil {
		return err
	}

	// Seed the Calculator with the latest metrics
	s.Calculator.Seed()

	s.mqQueue = &amqp.Queue{
		Name:       *s.QueueName,
		Durable:    true,
		AutoDelete: false,
	}

	err := s.DatabaseBroker.ConsumeKeys(s.mqQueue, "calc", s.processMetricUpdate, "metric.#")

	if err == nil {
		log.Println(version.Version)
	}

	// Mark the application as a daemon
	s.Daemon.SetDaemon()

	return nil
}

func (s *Service) initFromDB() {
	if *s.DBServer != "" {
		err := s.Calculations.Script().Accept(lang.NewBuilder().
			Calculation(s.addCalculation).
			Build())

		if err != nil {
			panic(err)
		}
	}
}

// processMetricUpdate accepts a metric from RabbitMQ, updates it in Latest
// then forwards it to any calculations
func (s *Service) processMetricUpdate(delivery amqp091.Delivery) error {
	var metric api.Metric
	err := json.Unmarshal(delivery.Body, &metric)
	if err == nil {
		s.storeLatest(metric)
	}
	return err
}

// loadLatestMetrics retrieves the current metrics from the DB server
func (s *Service) loadLatestMetrics() error {
	if *s.DBServer != "" {
		c := &client.Client{Url: *s.DBServer}

		r, err := c.LatestMetrics()
		if err != nil {
			return err
		}

		for _, m := range r.Metrics {
			s.storeLatest(m)
		}

	}
	return nil
}

func (s *Service) storeLatest(metric api.Metric) {
	u, ok := value.GetUnit(metric.Unit)
	if ok {
		updated := s.Latest.Append(metric.Metric, record.Record{
			Time:  metric.Time,
			Value: u.Value(metric.Value),
		})

		if updated {
			s.Calculator.Accept(metric)
		}
	}
}

func (s *Service) addCalculation(_ lang.Visitor, c *lang.Calculation) error {
	return nil
}

func (s *Service) addCalculationX(ctx context.Context) error {
	cDef := station.CalculatedValueFromContext(ctx)
	if cDef != nil && cDef.Query != "" {
		c := &client.Client{Url: *s.DBServer}
		res, err := c.Query(cDef.Query)
		if err != nil {
			return err
		}

		// Use the first valid value as the metric
		for _, t := range res.Table {
			for _, r := range t.Rows {
				for _, c := range *r {
					if c.Value.IsValid() {
						s.storeLatest(api.Metric{
							Metric: cDef.ID,
							Time:   time.Now(),
							Unit:   c.Value.Unit().ID(),
							Value:  c.Value.Float(),
						})
						return nil
					}
				}
			}
		}

	}

	return nil
}
