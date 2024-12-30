package weathersensor

import (
	"context"
	errors2 "errors"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-script/errors"
	station2 "github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/sensors/device"
	"github.com/peter-mount/piweather.center/sensors/publisher"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/util/config"
	"github.com/peter-mount/piweather.center/util/table"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Service struct {
	Daemon         *kernel.Daemon        `kernel:"inject"`
	Cron           *cron.CronService     `kernel:"inject"`
	Config         config.Manager        `kernel:"inject"`
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	Stations       *station.Stations     `kernel:"inject"`
	Rest           *rest.Server          `kernel:"inject"`
	WebPrefix      *string               `kernel:"flag,web-prefix,Prefix for http endpoints,/i"`
	// internal from here
	dashDir       string
	mutex         sync.Mutex
	httpSensors   map[string]map[string]map[string]*station2.Sensor // lookup for each http sensor for rest service
	httpPublisher map[string]publisher.Publisher                    // map of publishers
	sensorTable   *table.Table                                      // Used for debugging
	sensorCount   int                                               // Number of sensors defined
}

const (
	dashDir    = "stations"
	fileSuffix = ".sensor"
)

func (s *Service) PostInit() error {
	s.dashDir = filepath.Join(s.Config.EtcDir(), dashDir)
	s.httpSensors = make(map[string]map[string]map[string]*station2.Sensor)
	s.httpPublisher = make(map[string]publisher.Publisher)
	s.sensorTable = table.New("Station", "Sensor", "Type", "Path", "Method", "Options")

	// Load existing dashboards
	stations, err := s.Stations.LoadDirectory(s.dashDir, fileSuffix, station.SensorOption)
	if err != nil {
		return err
	}

	// Configure the sensors
	err = station2.NewBuilder[*state]().
		Http(s.httpSensor).
		I2C(s.i2cSensor).
		Sensor(s.sensor).
		Serial(s.serialSensor).
		Station(s.station).
		Build().
		Set(&state{service: s}).
		Stations(stations)
	if err != nil {
		return err
	}

	// Fail is we have no sensors defined
	if s.sensorCount == 0 {
		return errors2.New("no sensors defined")
	}

	// Disable http if we don't need it
	if len(s.httpSensors) == 0 {
		s.Rest.Disable()
	}

	return nil
}

func (s *Service) webPath(stationId, sensorId string) string {
	return strings.Join([]string{*s.WebPrefix, stationId, sensorId}, "/")
}

func (s *Service) Start() error {

	// Add the web endpoint for each method requested
	if len(s.httpSensors) > 0 {
		webPath := s.webPath("{stationId}", "{sensorId}")
		for k, _ := range s.httpSensors {
			s.Rest.Handle(webPath, s.handleHttp).Methods(k)
		}
	}

	s.Daemon.SetDaemon()

	if log.IsVerbose() {
		log.Println(version.Version)
		log.Printf("Sensors:\n%s", s.sensorTable.SortTable(0, 1, 2).String())
	}
	return nil
}

func (s *Service) GetPublisher(id string) publisher.Publisher {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.httpPublisher[id]
}

func (s *Service) addHttp(method, stationId, sensorId string, d *station2.Sensor) error {
	s.addSensor("http", stationId, sensorId, d.Http.Method, s.webPath(stationId, sensorId), "")
	s.mutex.Lock()
	defer s.mutex.Unlock()
	m1, ok := s.httpSensors[method]
	if !ok {
		m1 = make(map[string]map[string]*station2.Sensor)
		s.httpSensors[method] = m1
	}
	m2, ok := m1[stationId]
	if !ok {
		m2 = make(map[string]*station2.Sensor)
		m1[stationId] = m2
	}

	if e, ok := m2[sensorId]; ok {
		return errors.Errorf(d.Pos, "http %s %q already present at %s", d.Http.Method, sensorId, e.Pos)
	}

	m2[sensorId] = d

	s.httpPublisher[stationId+"."+sensorId] = s.publisher(d)
	return nil
}

func (s *Service) GetHttp(method, stationId, sensorId string) *station2.Sensor {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if m1, ok := s.httpSensors[method]; ok {
		if m2, ok := m1[stationId]; ok {
			return m2[sensorId]
		}
	}
	return nil
}

type state struct {
	service *Service
	station *station2.Station
	sensor  *station2.Sensor
}

func (s *Service) station(v station2.Visitor[*state], d *station2.Station) error {
	v.Get().station = d
	return nil
}

func (s *Service) sensor(v station2.Visitor[*state], d *station2.Sensor) error {
	v.Get().sensor = d
	return nil
}

// PollDevice will configure a task that will poll the given instance based on a cron definition.
// Any errors returned by the device when it's polled will be reported in the log.
func (s *Service) PollDevice(dev device.Device, instance device.Instance, publisher publisher.Publisher, cronDef string) error {
	_, err := s.Cron.AddTask(cronDef, func(_ context.Context) error {
		err := instance.RunDevice(publisher)
		if err != nil {
			log.Printf("device %q error %s",
				dev.Info().ID,
				err.Error())
		}
		return nil
	})
	return err
}

// RunDevice will call the instance in a separate goroutine.
// Any error returned by the device will be logged, and it will retry the device after a short delay.
func (s *Service) RunDevice(dev device.Device, instance device.Instance, publisher publisher.Publisher) {
	go func() {
		for {
			err := instance.RunDevice(publisher)
			if err != nil {
				log.Printf("device %q error %s",
					dev.Info().ID,
					err.Error())
				time.Sleep(time.Second)
			}
		}
	}()
}

func (s *Service) addSensor(bus, stationId, sensorId, mode, path, options string) {
	s.sensorTable.NewRow().Add(stationId).Add(sensorId).Add(bus).Add(path).Add(mode).Add(options)
}

func (s *Service) PublishMetric(m api.Metric) error {
	log.Printf("DB Pub %v", m)
	return nil
}
