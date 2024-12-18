package renderer

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/util/html"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"time"
)

func init() {
	registerJs("multivalue", `Object.keys(idx).forEach(i=>{setText(id,i,idx[i].formatted)})`)
}

func MultiValue(v station.Visitor[*State], d *station.MultiValue) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			dash := s.Dashboard()
			stn := dash.Station()

			comp := s.Dashboard().GetComponent(d.GetID())
			if comp == nil {
				return nil
			}

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

			s.Builder().
				Table().
				THead().TR().
				TH().Text("Metric").End().
				TH().Text("Latest Value").End().
				If(d.Time, func(e *html.Element) *html.Element {
					return e.TH().Text("Time updated").End()
				}).
				End().End(). // tr thead
				TBody().
				Exec(func(e *html.Element) *html.Element {
					for i, m := range metrics {
						cei := comp.GetMetrics(m)
						// cei can be empty for a metric who's not fully configured so we ignore them here
						if len(cei) > 0 {
							idx := cei[0].Index

							mv := metricValues[i]
							t := ""
							if mv.IsValid() {
								t = mv.String()
							}
							e = e.TR().
								// Metric column
								TD().Text(strings.SplitN(m, ".", 2)[1]).End(). // td
								// Value column
								TD().
								If(s.IsLive(), func(e *html.Element) *html.Element {
									return e.Attr("id", "%s.txt%d", d.Component.GetID(), idx)
								}).
								Text(t).
								End(). // td
								// Time column
								If(d.Time, func(e *html.Element) *html.Element {
									return e.TD().
										If(s.IsLive(), func(e *html.Element) *html.Element {
											return e.Attr("id", "%s.txt%dT", d.Component.GetID(), idx)
										}).
										Text(metricTimes[i]).
										End() // td
								}).
								End() // tr
						}
					}
					return e
				})

			return nil
		})

	if err != nil {
		return err
	}

	return errors.VisitorStop
}
