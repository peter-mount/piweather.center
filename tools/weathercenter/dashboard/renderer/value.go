package renderer

import (
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	"strings"
)

func Value(v station.Visitor[*State], d *station.Value) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			dash := s.Dashboard()
			stn := dash.Station()

			e := s.Builder()

			if d.Label != "" {
				e = e.Span().Class("label")
				for _, l := range strings.Split(d.Label, "_") {
					e = e.Span().TextNbsp(l).End()
				}
				e = e.End()
			}

			e = e.Span().Class("metric")

			for i, m := range d.Metrics.Metrics {
				var t string
				if metric, exists := stn.GetMetric(m.Name); exists {
					if v, ok := metric.ToValue(); ok {
						t = v.String()
					}
				}

				e = e.Span()
				if s.IsLive() {
					e = e.Attr("id", "%s.txt%d", d.Component.GetID(), i)
				}
				e = e.TextNbsp(t).End()
			}

			e = e.End() // span

			return nil
		})

	if err != nil {
		return err
	}

	return util.VisitorStop
}