package renderer

import (
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/util/html"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"strconv"
)

func Gauge(v station.Visitor[*State], d *station.Gauge) error {
	err := util.VisitorStop

	switch d.GetType() {
	case "barometer":
		err = barometer(v, d)
	case "compass":
		err = compass(v, d)
	case "gauge":
		err = gauge(v, d)
	case "inclinometer":
		err = inclinometer(v, d)
	}

	return err
}

var (
	barometerAngles = []float64{975, 990, 1010, 1026, 1045}
	barometerLabels = []string{"Stormy", "Rain", "Change", "Fair", "Very Dry"}
)

func barometer(v station.Visitor[*State], d *station.Gauge) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			dash := s.Dashboard()
			stn := dash.Station()
			metricValues := d.Metrics.GetValues(stn)
			axis := d.Axis.GenAxis(225)

			svg := s.Builder().
				Svg().
				ViewBox(0, 0, 250, 250).
				Attr("role", "img")

			if s.IsLive() {
				// Add data attributes for the javascript to use
				svg = svg.Attr(ID, "%s.svg", d.GetID()).
					Attr("data-min", fix(d.Axis.Min)).
					Attr("data-max", fix(d.Axis.Max)).
					Attr("data-delta", fix(axis.Delta))
				for i, v := range metricValues {
					if v.IsValid() {
						svg = svg.Attr("data-d"+strconv.Itoa(i),
							fix(d.Axis.EnsureWithin(v.Float(), axis.Delta, -112.5)))
					}
				}
			}

			svg = svg.G().
				Attr(Transform, Translate(125, 125)).
				Attr(DominantBaseline, Middle).
				Attr(TextAnchor, Middle).
				Circle().
				CX(0).CY(0).R(90).
				Fill(None).Stroke(Black).StrokeWidth("3px").
				StrokeDasharray(35, 212, 353).
				End() // circle

			for _, a := range axis.Ticks {
				svg = svg.G().
					Attr(Transform, Rotate(a.Angle-112.5)).
					Path().Attr(D, "M0,-90 l0,10").Fill(None).Stroke(Black).StrokeWidth("2px").End().
					SvgText().Y(-100).Text(a.Label).End().
					End()
			}

			// condition labels
			svg = svg.G().Attr(FontSize, "80%")
			for i, l := range barometerLabels {
				svg = svg.G().
					Attr(Transform, Rotate(d.Axis.EnsureWithin(barometerAngles[i], axis.Delta, -112.5))).
					SvgText().Y(-70).Text(l).End().
					End() // g
			}
			svg = svg.End() // g wrapping labels

			if len(metricValues) > 0 {
				svg = svg.SvgText()
				if s.IsLive() {
					svg = svg.Attr(ID, "%s.txt0", d.GetID())
				}
				svg = svg.Y(35).Attr(FontSize, "150%").Text(metricValues[0].String()).End()
			}

			if d.Label != "" {
				svg = svg.SvgText().Y(55).Attr(FontSize, "150%").Text(d.Label).End()
			}

			// The display hands
			GaugeHands(s, d, axis, svg, metricValues)

			return nil
		})

	if err != nil {
		return err
	}

	return util.VisitorStop
}

var (
	compassBearings = []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
)

func compass(v station.Visitor[*State], d *station.Gauge) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			dash := s.Dashboard()
			stn := dash.Station()
			metricValues := d.Metrics.GetValues(stn)

			svg := s.Builder().
				Svg().
				ViewBox(0, 0, 250, 250).
				Attr("role", "img")

			if s.IsLive() {
				// Add data attributes for the javascript to use
				svg = svg.Attr(ID, "%s.svg", d.GetID()).
					Attr("data-min", fix(d.Axis.Min)).
					Attr("data-max", fix(d.Axis.Max))
				for i, v := range metricValues {
					if v.IsValid() {
						svg = svg.Attr("data-metric"+strconv.Itoa(i), fix(v.Float()))
					}
				}
			}

			svg = svg.G().
				Attr(Transform, Translate(125, 125)).
				StrokeWidth("3px").
				Circle().R(100).Fill(None).Stroke(Black).End()

			for d := 0.0; d < 360.0; d += 22.5 {
				svg = svg.G().
					Attr(Transform, Rotate(d)).
					Fill(None).Stroke(Black).
					Path().Attr(D, "M-85 0 l-15 0").End().
					End()
			}

			svg = svg.G().
				Attr(DominantBaseline, Middle).
				Attr(TextAnchor, Middle)
			for i, l := range compassBearings {
				svg = svg.G().Attr(Transform, Rotate(float64(i)*45)).
					SvgText().Y(-110).Text(l).End().
					End()
			}
			if d.Label != "" {
				svg = svg.SvgText().Y(-45).Attr(FontSize, "150%").Text(d.Label).End()
			}
			if len(metricValues) > 0 {
				svg = svg.SvgText()
				if s.IsLive() {
					svg = svg.Attr(ID, "%s.txt0", d.GetID())
				}
				svg = svg.Y(45).Attr(FontSize, "150%").Text(fix(metricValues[0].Float())).End()
			}
			svg = svg.End() // g

			if len(metricValues) > 0 {
				svg = svg.G().
					Fill(Red).
					Attr(Transform, Rotate(metricValues[0].Float())).
					Circle().R(15).End().
					Path().Attr(D, "M-15 0 l30 0 l-15 -80 z")
				if s.IsLive() {
					svg = svg.AnimateTransform().
						Attr(ID, "%s.ptr0", d.GetID()).
						Attr("attributeName", "transform").
						Attr("attributeType", "XML").
						Attr("type", "rotate").
						Attr("from", "0").
						Attr("to", "0").
						Attr("dur", "1s").
						Fill("freeze").
						End()
				}
				svg = svg.End(). // path
							End() // g
			}

			svg = svg.Circle().R(10).Fill(Black).End()
			return nil
		})

	if err != nil {
		return err
	}

	return util.VisitorStop
}

func gauge(v station.Visitor[*State], d *station.Gauge) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			dash := s.Dashboard()
			stn := dash.Station()
			metricValues := d.Metrics.GetValues(stn)
			axis := d.Axis.GenAxis(180)

			svg := s.Builder().
				Svg().
				ViewBox(0, 0, 250, 250).
				Attr("role", "img")

			if s.IsLive() {
				// Add data attributes for the javascript to use
				svg = svg.Attr(ID, "%s.svg", d.GetID()).
					Attr("data-min", fix(d.Axis.Min)).
					Attr("data-max", fix(d.Axis.Max)).
					Attr("data-delta", fix(axis.Delta))
				for i, v := range metricValues {
					if v.IsValid() {
						svg = svg.Attr("data-metric"+strconv.Itoa(i), fix(v.Float()))
					}
				}
			}

			svg = svg.G().
				Attr(Transform, Translate(125, 125)).
				Attr(DominantBaseline, Middle).
				Attr(TextAnchor, Middle).
				// Outer dial
				Path().Attr(D, "M-90,0 a1,1 0 0,1 180,0").Fill(None).Stroke(Black).StrokeWidth("3px").End()

			// Ticks and their labels
			for _, a := range axis.Ticks {
				svg = svg.G().
					Attr(Transform, Rotate(a.Angle-90)).
					Path().Attr(D, "M0,-90 l0,10").Fill(None).Stroke(Black).StrokeWidth("2px").End().
					SvgText().Y(-100).Text(a.Label).End().
					End()
			}

			// Display the main metric value
			if len(metricValues) > 0 {
				svg = svg.SvgText().
					Attr(ID, "%s.txt%i", d.Component.GetID(), 0).
					Y(35).
					Text(metricValues[0].String()).
					End()
			}

			if d.Label != "" {
				svg = svg.SvgText().Y(55).Attr(FontSize, "150%").Text(d.Label).End()
			}

			// The display hands
			GaugeHands(s, d, axis, svg, metricValues)

			// Center dial cap
			svg = svg.Circle().R(10).Fill(Black).End()

			return nil
		})

	if err != nil {
		return err
	}

	return util.VisitorStop
}

func inclinometer(v station.Visitor[*State], d *station.Gauge) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			dash := s.Dashboard()
			stn := dash.Station()
			metricValues := d.Metrics.GetValues(stn)
			axis := station.GenAxis(-90, 90, 9, 180)

			svg := s.Builder().
				Svg().
				ViewBox(0, 0, 250, 250).
				Attr("role", "img")

			if s.IsLive() {
				// Add data attributes for the javascript to use
				svg = svg.Attr(ID, "%s.svg", d.GetID()).
					Attr("data-min", fix(d.Axis.Min)).
					Attr("data-max", fix(d.Axis.Max)).
					Attr("data-delta", fix(axis.Delta))
				for i, v := range metricValues {
					if v.IsValid() {
						svg = svg.Attr("data-metric"+strconv.Itoa(i), fix(v.Float()))
					}
				}
			}

			svg = svg.G().
				Attr(Transform, Translate(125, 125)).
				Attr(DominantBaseline, Middle).
				Attr(TextAnchor, Middle).
				Circle().
				CX(0).CY(0).R(90).
				Fill(None).Stroke(Black).StrokeWidth("3px").
				StrokeDasharray(141, 283, 142).
				End(). // circle
				Path().Class("dash-inclinometer-horizon").Attr(D, "M0,90L0,-90M90,0L0,0L90,9M88,19L0,0L85,28").End()

			for _, a := range axis.Ticks {
				svg = svg.G().
					Attr(Transform, Rotate(180-a.Angle)).
					Path().Attr(D, "M0,-90 l0,10").Fill(None).Stroke(Black).StrokeWidth("2px").End().
					SvgText().Y(-100).Text(a.Label).End().
					End()
			}

			svg = svg.G().Attr(Transform, Rotate(-90))
			if len(metricValues) > 0 {
				svg = svg.SvgText()
				if s.IsLive() {
					svg = svg.Attr(ID, "%s.txt0", d.GetID())
				}
				svg = svg.Y(-35).Attr(FontSize, "150%").Text(metricValues[0].String()).End()
			}

			if d.Label != "" {
				svg = svg.SvgText().Y(-55).Attr(FontSize, "150%").Text(d.Label).End()
			}
			svg = svg.End() // g rot -90

			// The display hands
			for i, m := range metricValues {
				ang := 180 - ((station.EnsureWithin(m.Float(), -90, 90) + 90) * axis.Delta)
				svg = svg.G().
					Attr(ID, "%s.ptr%d", d.GetID(), i).
					Attr(Transform, Rotate(ang)).
					Path().
					Class("dash-h%d", i).
					Attr(D, "M0,0l0%.0f", (math.Min(float64(i), 5)*10)-80).
					End(). // path
					End()  // g
			}

			svg = svg.Circle().R(10).Fill(Black).End()

			return nil
		})

	if err != nil {
		return err
	}

	return util.VisitorStop
}

func GaugeHands(s *State, d *station.Gauge, axis station.AxisDef, svg *html.Element, metricValues []value.Value) {
	for i, v := range metricValues {
		ang := d.Axis.EnsureWithin(v.Float(), axis.Delta, -90)
		svg = svg.G().
			Attr(Transform, Rotate(ang)).
			Path().
			Class("dash-h%d", i).
			Attr("d", "M0,10l0,%s", fix((math.Min(5, float64(i))*10)-80))
		if s.IsLive() {
			svg = svg.AnimateTransform().
				Attr(ID, "%s.ptr%d", d.GetID(), i).
				Attr("attributeName", "transform").
				Attr("attributeType", "XML").
				Attr("type", "rotate").
				Attr("from", "0").
				Attr("to", "0").
				Attr("dur", "1s").
				Fill("freeze").
				End()
		}
		svg = svg.End(). // path
					End(). // g
			// Center dial cap
			Circle().R(10).Fill(Black).End()
	}
}
