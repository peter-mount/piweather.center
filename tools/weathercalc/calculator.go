package weathercalc

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/service"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/broker"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/store/memory"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"sync"
	"time"
)

// Calculator is a service which accepts metrics and then performs any defined calculations.
// However, it only does the calculation once it gets all the values the calculation requires.
type Calculator struct {
	DatabaseBroker broker.DatabaseBroker `kernel:"inject"`
	Latest         memory.Latest         `kernel:"inject"`
	Config         service.Config        `kernel:"inject"`
	mutex          sync.Mutex
	calculations   map[string]*Calculation   // Map of Calculation's by their ID's
	metrics        map[string][]*Calculation // Map of Calculation's by their dependencies
}

func (calc *Calculator) Start() error {
	calc.calculations = make(map[string]*Calculation)
	calc.metrics = make(map[string][]*Calculation)

	// Load the calculations
	if err := calc.Config.Accept(station.NewVisitor().
		CalculatedValue(calc.addCalculation).
		WithContext(context.Background())); err != nil {
		return err
	}

	return nil
}

func (calc *Calculator) addCalculation(ctx context.Context) error {
	cDef := station.CalculatedValueFromContext(ctx)
	if cDef == nil {
		return nil
	}

	c := calc.getCalculation(cDef.ID)
	if c != nil {
		return fmt.Errorf("calculation for %q already defined", cDef.ID)
	}

	c = NewCalculation(cDef)

	stn := station.StationFromContext(ctx)
	c.time = value.BasicTime(time.Time{},
		stn.LatLong().Coord(),
		stn.Location.Altitude)

	n := strings.ToLower(cDef.ID)
	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	calc.calculations[n] = c

	for _, s := range cDef.Source {
		switch s {
		case "current":
		default:
			calc.metrics[s] = append(calc.metrics[s], c)
		}
	}

	// If Default and Use set then set the default value
	if c.src.Default != nil && c.src.Use != "" {
		if v, exists := value.GetUnit(c.src.Use); exists {
			calc.Latest.Append(
				c.ID(),
				record.Record{
					Value: v.Value(*c.src.Default),
					Time:  time.Now(),
				})
		}
	}

	return nil
}

func (calc *Calculator) getCalculation(n string) *Calculation {
	n = strings.ToLower(n)
	calc.mutex.Lock()
	defer calc.mutex.Unlock()
	return calc.calculations[n]
}

func (calc *Calculator) getCalculationByMetric(n string) ([]*Calculation, bool) {
	n = strings.ToLower(n)
	calc.mutex.Lock()
	defer calc.mutex.Unlock()
	c, exists := calc.metrics[n]
	return c, exists
}

func (calc *Calculator) Seed() {
	for _, m := range calc.Latest.Metrics() {
		if r, ok := calc.Latest.Latest(m); ok {
			calc.accept(api.Metric{
				Metric: m,
				Time:   r.Time,
				Unit:   r.Value.Unit().ID(),
				Value:  r.Value.Float(),
			}, false)
		}
	}
}

func (calc *Calculator) Accept(metric api.Metric) {
	calc.accept(metric, true)
}

func (calc *Calculator) accept(metric api.Metric, post bool) {
	if m, exists := calc.getCalculationByMetric(metric.Metric); exists {
		for _, c := range m {
			if c.Accept(metric) {
				calc.calculate(c, post)
			}
		}
	}
}

func (calc *Calculator) calculate(c *Calculation, post bool) {
	result, t, err := calc.calculateResult(c)
	if err != nil {
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

		if post {
			if err = calc.DatabaseBroker.PublishMetric(metric); err != nil {
				log.Printf("post %q failed %v", c.ID(), metric)
			}
		}

		calc.accept(metric, post)
	}
}

func (calc *Calculator) calculateResult(c *Calculation) (value.Value, time.Time, error) {
	if _, exists := calc.Latest.Latest(c.ID()); !exists && c.src.UseFirst {
		for _, s := range c.src.Source {
			switch s {
			case "current":
			default:
				r, ok := calc.Latest.Latest(s)
				if ok && r.IsValid() {
					return r.Value, r.Time, nil
				}
				return value.Value{}, time.Time{}, nil
			}
		}

		return value.Value{}, time.Time{}, nil
	}

	var t time.Time
	var args []value.Value
	for _, s := range c.src.Source {
		switch s {
		case "current":
			s = c.ID()
		}

		r, ok := calc.Latest.Latest(s)
		if !ok {
			return value.Value{}, time.Time{}, nil
		}
		args = append(args, r.Value)

		if t.IsZero() || t.Before(r.Time) {
			t = r.Time
		}
	}

	r, err := c.src.Calculate(c.time.Clone().SetTime(t), args...)
	return r, t, err
}

type Calculation struct {
	mutex      sync.Mutex
	src        *station.CalculatedValue     // Link to definition
	metrics    map[string]*CalculationValue // Most recent values received
	lastUpdate time.Time                    // Time calculation last run
	lastValue  value.Value                  // Last value
	time       value.Time                   // Time with location
}

func NewCalculation(src *station.CalculatedValue) *Calculation {
	calc := &Calculation{
		src:     src,
		metrics: make(map[string]*CalculationValue),
	}

	for _, s := range src.Source {
		switch s {
		case "current":
		default:
			calc.metrics[s] = &CalculationValue{}
		}
	}

	return calc
}

type CalculationValue struct {
	metric api.Metric // Last metric received
	ready  bool       // true if we have received this value since the last Calculation
}

func (c *Calculation) ID() string {
	return c.src.ID
}

func (c *Calculation) Accept(metric api.Metric) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	cv := c.metrics[metric.Metric]
	// Do nothing if metric is newer
	if cv == nil || cv.metric.Time.After(metric.Time) {
		return false
	}

	cv.metric = metric
	cv.ready = true

	// Return true only if all metrics are ready
	for _, m := range c.metrics {
		if !m.ready && m.metric.Metric != "" {
			return false
		}
	}
	return true
}
