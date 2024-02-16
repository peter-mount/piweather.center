package weathercalc

import (
	"flag"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/tools/weathercalc/lang"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"sync"
	"time"
)

// Calculator is a service which accepts metrics and then performs any defined calculations.
// However, it only does the calculation once it gets all the values the calculation requires.
type Calculator struct {
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	Cron           *cron.CronService     `kernel:"inject"`
	Latest         memory.Latest         `kernel:"inject"`
	DBServer       *string               `kernel:"flag,metric-db,DB url"`
	mutex          sync.Mutex
	script         *lang.Script
	targets        map[string]*Calculation   // Map of Calculations by target
	metrics        map[string][]*Calculation // Map of Calculation's by their dependencies
	calculations   []*Calculation            // Calculation's in sequence
}

func (calc *Calculator) Start() error {
	p := lang.NewParser()
	script, err := p.ParseFiles(flag.Args()...)
	if err != nil {
		return err
	}

	calc.script = script
	calc.targets = make(map[string]*Calculation)
	calc.metrics = make(map[string][]*Calculation)

	// Load the calculations
	if err := lang.NewBuilder[*Calculator]().
		Calculation(calc.addCalculation).
		Build().
		Script(calc.script); err != nil {
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

func (calc *Calculator) Script() *lang.Script {
	return calc.script
}

// loadLatestMetrics retrieves the current metrics from the DB server
func (calc *Calculator) loadLatestMetrics() error {
	if *calc.DBServer != "" {
		c := &client.Client{Url: *calc.DBServer}

		r, err := c.LatestMetrics()
		if err != nil {
			return err
		}

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
	return nil
}

func (calc *Calculator) addMetric(n string, c *lang.Calculation) {
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

	nc := NewCalculation(c)
	calc.targets[c.Target] = nc
	calc.metrics[n] = append(metrics, nc)
	calc.calculations = append(calc.calculations, nc)
}

func (calc *Calculator) addCalculation(_ lang.Visitor[*Calculator], c *lang.Calculation) error {
	if err := lang.NewBuilder[*Calculator]().
		Metric(func(_ lang.Visitor[*Calculator], m *lang.Metric) error {
			calc.addMetric(m.Name, c)
			return nil
		}).
		Build().
		SetData(calc).
		Calculation(c); err != nil {
		return err
	}

	// RESET EVERY cron definition
	if c.ResetEvery != nil {
		if _, err := calc.Cron.AddFunc(c.ResetEvery.Definition, func() {
			calc.Latest.Remove(c.Target)
			if c.Load != nil {
				_ = calc.loadFromDB(c)
			}
			calc.calculateTarget(c.Target)
		}); err != nil {
			return err
		}
	}

	// Every definition
	if c.Every != nil {
		if _, err := calc.Cron.AddFunc(c.Every.Definition, func() {
			if c.Load != nil {
				_ = calc.loadFromDB(c)
			}
			calc.calculateTarget(c.Target)
		}); err != nil {
			return err
		}
	}

	// If the target still has no Calculation registered then create it.
	// This will happen when a calculation is defined that doesn't
	// reference any metrics. e.g. SolarAltitude which uses just location and time
	if calc.getCalculationByTarget(c.Target) == nil {
		calc.addCalculationByTarget(NewCalculation(c))
	}

	if c.Load != nil {
		if err := calc.loadFromDB(c); err != nil {
			return err
		}
	}
	return nil
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
	src        *lang.Calculation // Link to definition
	lastUpdate time.Time         // Time calculation last run
	lastValue  value.Value       // Last value
	time       value.Time        // Time with location
}

func NewCalculation(src *lang.Calculation) *Calculation {
	return &Calculation{src: src}
}

type CalculationValue struct {
	metric api.Metric // Last metric received
	ready  bool       // true if we have received this value since the last Calculation
}

func (c *Calculation) ID() string {
	return c.src.Target
}

func (c *Calculation) Src() *lang.Calculation {
	return c.src
}

func (c *Calculation) Accept(metric api.Metric) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Note: !After and not Before as they are NOT the same thing!
	return !c.lastUpdate.After(metric.Time)
}
