package station

import (
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/store/api"
	"sync"
)

type Calculation struct {
	mutex       sync.Mutex
	station     *Station              // Containing Station
	calculation *station.Calculation  // Associated Calculation
	metrics     map[string]api.Metric // Metrics required and current values
}

func newCalculation(station *Station, calculation *station.Calculation) *Calculation {
	return &Calculation{
		station:     station,
		calculation: calculation,
		metrics:     make(map[string]api.Metric),
	}
}

func (c *Calculation) Station() *Station {
	return c.station
}

func (c *Calculation) Calculation() *station.Calculation {
	return c.calculation
}

func (c *Calculation) AddMetric(metric string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, exists := c.metrics[metric]; !exists {
		c.metrics[metric] = api.Metric{Metric: metric}
	}
}
