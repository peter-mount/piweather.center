package weathercalc

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/broker"
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
	Latest         memory.Latest         `kernel:"inject"`
	Calculations   *Calculations         `kernel:"inject"`
	mutex          sync.Mutex
	metrics        map[string][]*Calculation // Map of Calculation's by their dependencies
}

func (calc *Calculator) Start() error {
	calc.metrics = make(map[string][]*Calculation)

	// Load the calculations
	if err := calc.Calculations.Script().Accept(lang.NewBuilder().
		Calculation(calc.addCalculation).
		Build()); err != nil {
		return err
	}

	return nil
}

func (calc *Calculator) addMetric(n string, c *lang.Calculation) {
	calc.mutex.Lock()
	defer calc.mutex.Unlock()

	log.Printf("addMetric %q -> %q", n, c.Target)

	metrics := calc.metrics[n]
	if metrics != nil {
		for _, e := range metrics {
			if e.ID() == c.Target {
				return
			}
		}
	}
	calc.metrics[n] = append(metrics, NewCalculation(c))
}

func (calc *Calculator) addCalculation(_ lang.Visitor, c *lang.Calculation) error {
	//c.time = value.BasicTime(time.Time{},
	//	stn.LatLong().Coord(),
	//	stn.Location.Altitude)

	f := func(_ lang.Visitor, m *lang.Metric) error {
		calc.addMetric(m.Name, c)
		return nil
	}
	if err := c.Accept(lang.NewBuilder().
		Metric(f).
		UseFirst(f).
		Build()); err != nil {
		return err
	}

	// If Default and Use set then set the default value
	//if c.src.Default != nil && c.src.Use != "" {
	//	if v, exists := value.GetUnit(c.src.Use); exists {
	//		calc.Latest.Append(
	//			c.ID(),
	//			record.Record{
	//				Value: v.Value(*c.src.Default),
	//				Time:  time.Now(),
	//			})
	//	}
	//}

	return nil
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

	if c.lastUpdate.After(metric.Time) {
		return false
	}

	//cv.metric = metric
	//cv.ready = true
	//
	//// Return true only if all metrics are ready
	//for _, m := range c.metrics {
	//	if !m.ready && m.metric.Metric != "" {
	//		return false
	//	}
	//}
	return true
}
