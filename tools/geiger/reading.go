package geiger

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"time"
)

// getStats is a Task that sends the relevant commands to the Geiger counter.
func (m *Geiger) getStats(_ context.Context) error {

	now := time.Now().UTC()
	m.publish("cpm", now, float64(m.getCpm()), "CountPerMinute")
	m.publish("cpm", now, m.getTemp(), "Celsius")
	m.publish("cpm", now, m.getVolt(), "Volt")
	x, y, z := m.getGyro()
	m.publish("gyro.x", now, float64(x), "Integer")
	m.publish("gyro.y", now, float64(y), "Integer")
	m.publish("gyro.z", now, float64(z), "Integer")

	return nil
}

func (m *Geiger) publish(suffix string, t time.Time, v float64, unit string) {
	_ = m.DatabaseBroker.PublishMetric(api.Metric{
		Metric: *m.Id + "." + suffix,
		Time:   t,
		Unit:   unit,
		Value:  v,
	})
	if *m.Debug {
		log.Printf("%q %f", suffix, v)
	}
}
