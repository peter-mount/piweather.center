package expression

import (
	station2 "github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/weather/value"
	"sync"
	"time"
)

type Calculation struct {
	mutex      sync.Mutex
	src        *station2.Calculation // Link to definition
	station    *station2.Station     // Link to station
	lastUpdate time.Time             // Time calculation last run
	lastValue  value.Value           // Last value
	time       value.Time            // Time with location
}

func (c *Calculation) String() string {
	return "Calc(" + c.src.Target + ")"
}

func NewCalculation(src *station2.Calculation, station *station2.Station) *Calculation {
	return &Calculation{src: src, station: station}
}

func (c *Calculation) SetLatest(v value.Value, t time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.lastUpdate = t
	c.lastValue = v
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
