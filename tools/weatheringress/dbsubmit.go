package weatheringress

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/tools/weatheringress/model"
	"github.com/peter-mount/piweather.center/tools/weatheringress/payload"
	"github.com/peter-mount/piweather.center/util"
)

func (s *Ingress) processReading(ctx context.Context) error {
	r := model.ReadingFromContext(ctx)
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

			if v.IsValid() {

				metric := api.Metric{Time: p.Time().UTC()}

				metric.Metric = r.ID
				metric.Value = v.Float()
				metric.Unit = v.Unit().ID()

				_ = s.DatabaseBroker.PublishMetric(metric)
			} else {
				// Has happened, not sure if down to altered ID's etc
				log.Printf("Invalid Metric %q", r.ID)
			}
		}
	}
	return nil
}
