package weathercalc

import (
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	station2 "github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/util/config"
	"github.com/peter-mount/piweather.center/weather/value"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Calculator is a service which accepts metrics and then performs any defined calculations.
// However, it only does the calculation once it gets all the values the calculation requires.
type Calculator struct {
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	Config         config.Manager        `kernel:"inject"`
	Cron           *cron.CronService     `kernel:"inject"`
	Latest         memory.Latest         `kernel:"inject"`
	DBServer       *string               `kernel:"flag,metric-db,DB url"`
	Stations       *station.Stations     `kernel:"inject"`
	mutex          sync.Mutex
	dashDir        string

	//script         *lang2.Script
	targets      map[string]*Calculation   // Map of Calculations by target
	metrics      map[string][]*Calculation // Map of Calculation's by their dependencies
	calculations []*Calculation            // Calculation's in sequence
}

const (
	dashDir    = "stations"
	fileSuffix = ".calc"
)

func (calc *Calculator) Start() error {

	calc.dashDir = filepath.Join(calc.Config.EtcDir(), dashDir)

	// Load existing dashboards
	stations, err := calc.Stations.LoadDirectory(calc.dashDir, fileSuffix, station.CalculationOption)
	if err != nil {
		return err
	}

	calc.targets = make(map[string]*Calculation)
	calc.metrics = make(map[string][]*Calculation)

	// Load the calculations
	if err := station2.NewBuilder[*calcState]().
		Station(visitStation).
		Calculation(addCalculation).
		Metric(addMetric).
		Build().
		Set(&calcState{calc: calc}).
		Stations(stations); err != nil {
		return err
	}

	// Get latest metrics from DB
	if err := calc.loadLatestMetrics(); err != nil {
		return err
	}

	// Now run through all calculations for the first time
	for _, c := range calc.calculations {
		calc.calculate(c)
	}

	return nil
}

// loadLatestMetrics retrieves the current metrics from the DB server
func (calc *Calculator) loadLatestMetrics() error {
	if *calc.DBServer != "" {
		c := &client.Client{Url: *calc.DBServer}

		r, err := c.LatestMetrics()
		if err != nil {
			return err
		}

		// r can be nil if there are no results returned from the DB
		if r != nil {
			for _, m := range r.Metrics {
				u, ok := value.GetUnit(m.Unit)
				if ok {
					calc.Latest.Append(m.Metric, record.Record{
						Time:  m.Time,
						Value: u.Value(m.Value),
					})
				}
			}
		}

	}
	return nil
}

func (calc *Calculator) addMetric(n string, c *station2.Calculation, s *station2.Station) {
	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	metrics := calc.metrics[n]
	if metrics != nil {
		for _, e := range metrics {
			if e.ID() == c.Target {
				return
			}
		}
	}

	nc := NewCalculation(c, s)
	calc.targets[c.Target] = nc
	calc.metrics[n] = append(metrics, nc)
	calc.calculations = append(calc.calculations, nc)
}

func (calc *Calculator) getCalculationByMetric(n string) ([]*Calculation, bool) {
	n = strings.ToLower(n)
	calc.mutex.Lock()
	defer calc.mutex.Unlock()
	c, exists := calc.metrics[n]
	return c, exists
}

func (calc *Calculator) addCalculationByTarget(c *Calculation) {
	calc.mutex.Lock()
	defer calc.mutex.Unlock()
	calc.targets[c.ID()] = c
	calc.calculations = append(calc.calculations, c)
}

func (calc *Calculator) getCalculationByTarget(n string) *Calculation {
	calc.mutex.Lock()
	defer calc.mutex.Unlock()
	return calc.targets[n]
}

func (calc *Calculator) Accept(metric api.Metric) {
	// Only process the metric if it's not one we are calculating.
	// This is done as we would receive a duplicate from Rabbit after we have
	// made a calculation which. As we already pass the calculation result back
	// into the calculator locally, we don't need this duplication.
	target := calc.getCalculationByTarget(metric.Metric)
	if target == nil {
		calc.accept(metric)
	}
}

func (calc *Calculator) accept(metric api.Metric) {
	if m, exists := calc.getCalculationByMetric(metric.Metric); exists {
		for _, c := range m {
			if c.Accept(metric) {
				calc.calculate(c)
			}
		}
	}
}

func (calc *Calculator) calculateTarget(n string) {
	cd := calc.getCalculationByTarget(n)
	if cd != nil {
		calc.calculate(cd)
	}
}

func (calc *Calculator) calculate(c *Calculation) {
	result, t, err := calc.calculateResult(c)
	if err != nil {
		log.Println(c.Src().Pos, err)
		return
	}

	if result.IsValid() {
		c.lastValue = result
		c.lastUpdate = t

		calc.Latest.Append(c.ID(), record.Record{Value: result, Time: t})

		metric := api.Metric{
			Metric: c.ID(),
			Time:   t,
			Unit:   result.Unit().ID(),
			Value:  result.Float(),
		}

		if err = calc.DatabaseBroker.PublishMetric(metric); err != nil {
			log.Printf("post %q failed %v", c.ID(), metric)
		}

		// Pass the calculated result back into the calculator so any dependencies
		// may then be calculated immediately
		if c.Src().Expression != nil {
			calc.accept(metric)
		}
	}
}

type Calculation struct {
	mutex      sync.Mutex
	src        *station2.Calculation // Link to definition
	station    *station2.Station     // Link to station
	lastUpdate time.Time             // Time calculation last run
	lastValue  value.Value           // Last value
	time       value.Time            // Time with location
}

func NewCalculation(src *station2.Calculation, station *station2.Station) *Calculation {
	return &Calculation{src: src, station: station}
}

type CalculationValue struct {
	metric api.Metric // Last metric received
	ready  bool       // true if we have received this value since the last Calculation
}

func (c *Calculation) ID() string {
	return c.src.Target
}

func (c *Calculation) Src() *station2.Calculation {
	return c.src
}

func (c *Calculation) Accept(metric api.Metric) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Note: !After and not Before as they are NOT the same thing!
	return !c.lastUpdate.After(metric.Time)
}

// Station this Calculation is part of
func (c *Calculation) Station() *station2.Station {
	return c.station
}

// LastValue from previous calculation
func (c *Calculation) LastValue() value.Value {
	return c.lastValue
}

// LastUpdate time
func (c *Calculation) LastUpdate() time.Time {
	return c.lastUpdate
}

// Time with location
func (c *Calculation) Time() value.Time {
	return c.time
}
