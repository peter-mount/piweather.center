package renderer

import (
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"time"
)

func MultiValue(v station.Visitor[*State], d *station.MultiValue) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			dash := s.Dashboard()
			stn := dash.Station()

			metrics := stn.AcceptMetrics(d)
			var metricValues []value.Value
			var metricTimes []string
			for _, n := range metrics {
				m, _ := stn.GetMetric(n)
				if str, ok := m.ToValue(); ok {
					metricValues = append(metricValues, str)
					metricTimes = append(metricTimes, m.Time.Format(time.RFC3339))
				} else {
					metricValues = append(metricValues, value.Value{})
					metricTimes = append(metricTimes, "")
				}
			}

			e := s.Builder()

			e = e.Span().Class("label")
			for _, m := range metrics {
				e = e.Span().TextNbsp(strings.SplitN(m, ".", 2)[1]).End()
			}
			e = e.End()

			e = e.Span().Class("metric")
			for i, mv := range metricValues {
				t := ""
				if mv.IsValid() {
					t = mv.String()
				}
				e = e.Span()
				if s.IsLive() {
					e = e.Attr("id", "%s.txt%d", d.Component.GetID(), i)
				}
				e = e.TextNbsp(t).End()
			}
			e = e.End() // span

			if d.Time {
				e = e.Span().Class("metric-time")
				for i, t := range metricTimes {
					e = e.Span()
					if s.IsLive() {
						e = e.Attr("id", "%s.txt%dT", d.Component.GetID(), i)
					}
					e = e.TextNbsp(t).End()
				}
				e = e.End() // span
			}

			return nil
		})

	if err != nil {
		return err
	}

	return util.VisitorStop
}
