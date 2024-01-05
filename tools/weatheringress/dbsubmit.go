package weatheringress

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/payload"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
)

func (s *Ingress) databaseReading(ctx context.Context) error {
	payloadEntry := payload.GetPayload(ctx)

	metric := api.Metric{Time: payloadEntry.Time().UTC()}

	values := value.MapFromContext(ctx)
	for _, key := range values.GetKeys() {
		val := values.Get(key)

		if val.IsValid() {

			metric.Metric = key
			metric.Value = val.Float()
			metric.Unit = val.Unit().ID()

			err := s.DatabaseBroker.PublishMetric(metric)
			if err != nil {
				return err
			}
		} else {
			// Has happened, not sure if down to altered ID's etc
			log.Printf("Invalid Metric %q", key)
		}
	}

	return nil
}

func (s *Ingress) processReading(ctx context.Context) error {
	r := station.ReadingFromContext(ctx)
	values := value.MapFromContext(ctx)
	if r.Unit() != nil {
		p := payload.GetPayload(ctx)

		str, ok := p.Get(r.Source)
		if !ok {
			// FIXME warn/fail if not found?
			return nil
		}

		if f, ok := util.ToFloat64(str); ok {
			// Convert to Type unit then transform to Use unit
			v, err := r.Value(f)
			if err != nil {
				// Ignore, should only happen if the result is
				// invalid as we checked the transform previously
				return nil
			}

			values.Put(r.ID, v)
		}
	}
	return nil
}

func (s *Ingress) calculate(ctx context.Context) error {
	// Get value.Time from Station and Payload
	sensors := station.SensorsFromContext(ctx)
	p := payload.GetPayload(ctx)
	t := sensors.Station().LatLong().Time(p.Time())

	calc := station.CalculatedValueFromContext(ctx)

	values := value.MapFromContext(ctx)
	args := values.GetAll(calc.Source...)

	result, err := calc.Calculate(t, args...)
	if err != nil {
		return err
	}

	values.Put(calc.ID, result)

	return nil
}
