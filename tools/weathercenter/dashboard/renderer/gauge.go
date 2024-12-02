package renderer

import (
	"fmt"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/util/html"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"strconv"
	"strings"
)

const (
	_barometer    = "barometer"
	_compass      = "compass"
	_gauge        = "gauge"
	_inclinometer = "inclinometer"
	_rainGauge    = "raingauge"
)

func init() {
	registerJs(_barometer, `Object.keys(idx).forEach(i=>{`+
		`let m=idx[i],d=document.getElementById(id+".svg"),`+
		`min=d.dataset.min,`+
		`max=d.dataset.max,`+
		`delta=d.dataset.delta,`+
		`ofs=d.dataset["d"+i],`+
		`v=ensureWithin(m.value,min,max);setText(id,i,m.formatted);`+
		`setRotate(id,i,((v-min)*delta)-112.5-ofs)`+
		`})`)
	registerJs(_compass, `Object.keys(idx).forEach(i=>{`+
		`let m=idx[i],d=document.getElementById(id+".svg"),`+
		`v=m.value,`+
		`a=v-d.dataset["d"+i];`+
		`setRotate(id,i,a);`+
		`setText(id,i,""+Math.floor(v)+'°')`+
		`})`)
	registerJs(_gauge, `Object.keys(idx).forEach(i=>{`+
		`let m=idx[i],d=document.getElementById(id+".svg"),`+
		`min=d.dataset.min,`+
		`max=d.dataset.max,`+
		`delta=d.dataset.delta,`+
		`ofs=d.dataset["d"+i],`+
		`v=ensureWithin(m.value,min,max);`+
		`setText(id,i,m.formatted);`+
		`setRotate(id,i,((v-min)*delta)-90-ofs)`+
		`})`)
	registerJs(_inclinometer, `Object.keys(idx).forEach(i=>{`+
		`let m=idx[i],e=document.getElementById(id+".ptr"+i),`+
		`v=90-ensureWithin(m.value,-90,90);`+
		`setText(id,i,m.formatted);`+
		`e.setAttribute("transform","rotate("+v+")")`+
		`})`)
	registerJs(_rainGauge, `Object.keys(idx).forEach(i=>{`+
		`let m=idx[i],d=document.getElementById(id+".svg"),`+
		`min=d.dataset.min,`+
		`max=d.dataset.max,`+
		`scale=d.dataset.scale,`+
		`height=d.dataset.height,`+
		`v=m.value,`+
		`y=scale*(v-min);`+
		// Update means we exceed the axis so reload to get a new axis
		`if(v>max){location.reload();return}`+
		`let e=document.getElementById(id+".rect");`+
		`e.setAttribute("y",height-y);`+
		`e.setAttribute("height",y);`+
		`setText(id,i,m.formatted)`+
		`})`)
}

func Gauge(v station.Visitor[*State], d *station.Gauge) error {
	err := util.VisitorStop

	switch d.GetType() {
	case _barometer:
		err = barometer(v, d)
	case _compass:
		err = compass(v, d)
	case _gauge:
		err = gauge(v, d)
	case _inclinometer:
		err = inclinometer(v, d)
	case _rainGauge:
		err = rainGauge(v, d)
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
			metricValues, err := d.ConvertAll(d.Metrics.GetValues(stn))
			if err != nil {
				return err
			}
			axis := d.Axis.GenAxis(225)

			s.Builder().
				Svg().
				ViewBox(0, 0, 250, 250).
				Attr("role", "img").
				If(s.IsLive(), func(svg *html.Element) *html.Element {
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
					return svg
				}).
				G().
				Attr(Transform, Translate(125, 125)).
				Attr(DominantBaseline, Middle).
				Attr(TextAnchor, Middle).
				ExecEnd(func(svg *html.Element) {
					svg = svg.Circle().
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
					svg = svg.G().Attr(FontSize, "80%").
						ExecEnd(func(svg *html.Element) {
							for i, l := range barometerLabels {
								svg = svg.G().
									Attr(Transform, Rotate(d.Axis.EnsureWithin(barometerAngles[i], axis.Delta, -112.5))).
									SvgText().Y(-70).Text(l).End().
									End() // g
							}
						})

					if len(metricValues) > 0 {
						svg = svg.SvgText().
							If(s.IsLive(), func(svg *html.Element) *html.Element {
								return svg.Attr(ID, "%s.txt0", d.GetID())
							}).
							Y(35).Attr(FontSize, "150%").Text(metricValues[0].String()).End()
					}

					if d.Label != "" {
						svg = svg.SvgText().Y(55).Attr(FontSize, "150%").Text(d.Label).End()
					}

					// The display hands
					GaugeHands(s, d, axis, svg, metricValues)

					svg = svg.Circle().R(10).Fill(Black).End()
				})

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
			metricValues, err := d.ConvertAll(d.Metrics.GetValues(stn))
			if err != nil {
				return err
			}

			s.Builder().
				Svg().
				ViewBox(0, 0, 250, 250).
				Attr("role", "img").
				If(s.IsLive(), func(svg *html.Element) *html.Element {
					// Add data attributes for the javascript to use
					svg = svg.Attr(ID, "%s.svg", d.GetID()).
						Attr("data-min", fix(d.Axis.Min)).
						Attr("data-max", fix(d.Axis.Max))
					for i, v := range metricValues {
						if v.IsValid() {
							svg = svg.Attr("data-metric"+strconv.Itoa(i), fix(v.Float()))
						}
					}
					return svg
				}).
				G().
				Attr(Transform, Translate(125, 125)).
				StrokeWidth("3px").
				ExecEnd(func(svg *html.Element) {
					svg = svg.Circle().R(100).Fill(None).Stroke(Black).End()

					for d := 0.0; d < 360.0; d += 22.5 {
						svg = svg.G().
							Attr(Transform, Rotate(d)).
							Fill(None).Stroke(Black).
							Path().Attr(D, "M-85 0 l-15 0").End().
							End()
					}

					svg = svg.G().
						Attr(DominantBaseline, Middle).
						Attr(TextAnchor, Middle).
						ExecEnd(func(svg *html.Element) {
							for i, l := range compassBearings {
								svg = svg.G().Attr(Transform, Rotate(float64(i)*45)).
									SvgText().Y(-110).Text(l).End().
									End()
							}
							if d.Label != "" {
								svg = svg.SvgText().Y(-45).Attr(FontSize, "150%").Text(d.Label).End()
							}
							if len(metricValues) > 0 {
								svg = svg.SvgText().
									If(s.IsLive(), func(svg *html.Element) *html.Element {
										return svg.Attr(ID, "%s.txt0", d.GetID())
									}).
									Y(45).
									Attr(FontSize, "150%").
									Text(fix(metricValues[0].Float())).
									End()
							}
						})

					if len(metricValues) > 0 {
						svg = svg.G().
							Fill(Red).
							Attr(Transform, Rotate(metricValues[0].Float())).
							Circle().R(15).End().
							Path().Attr(D, "M-15 0 l30 0 l-15 -80 z").
							If(s.IsLive(), func(svg *html.Element) *html.Element {
								return svg.AnimateTransform().
									Attr(ID, "%s.ptr0", d.GetID()).
									Attr("attributeName", "transform").
									Attr("attributeType", "XML").
									Attr("type", "rotate").
									Attr("from", "0").
									Attr("to", "0").
									Attr("dur", "1s").
									Fill("freeze").
									End()
							}).
							End(). // path
							End()  // g
					}

					svg = svg.Circle().R(10).Fill(Black).End()
				})
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
			metricValues, err := d.ConvertAll(d.Metrics.GetValues(stn))
			if err != nil {
				return err
			}
			axis := d.Axis.GenAxis(180)

			s.Builder().
				Svg().
				ViewBox(0, 0, 250, 250).
				Attr("role", "img").
				If(s.IsLive(), func(svg *html.Element) *html.Element {
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
					return svg
				}).
				G().
				Attr(Transform, Translate(125, 125)).
				Attr(DominantBaseline, Middle).
				Attr(TextAnchor, Middle).
				ExecEnd(func(svg *html.Element) {
					// Outer dial
					svg = svg.Path().Attr(D, "M-90,0 a1,1 0 0,1 180,0").Fill(None).Stroke(Black).StrokeWidth("3px").End()

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
				})

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
			metricValues, err := d.ConvertAll(d.Metrics.GetValues(stn))
			if err != nil {
				return err
			}
			axis := station.GenAxis(-90, 90, 9, 180)

			s.Builder().
				Svg().
				ViewBox(0, 0, 250, 250).
				Attr("role", "img").
				If(s.IsLive(), func(svg *html.Element) *html.Element {
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
					return svg
				}).
				G().
				Attr(Transform, Translate(125, 125)).
				Attr(DominantBaseline, Middle).
				Attr(TextAnchor, Middle).
				ExecEnd(func(svg *html.Element) {
					svg = svg.Circle().
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

					svg = svg.G().Attr(Transform, Rotate(-90)).
						ExecEnd(func(svg *html.Element) {
							if len(metricValues) > 0 {
								svg = svg.SvgText().
									If(s.IsLive(), func(svg *html.Element) *html.Element {
										return svg.Attr(ID, "%s.txt0", d.GetID())
									}).
									Y(-35).
									Attr(FontSize, "150%").
									Text(metricValues[0].String()).
									End()
							}

							if d.Label != "" {
								svg = svg.SvgText().Y(-55).Attr(FontSize, "150%").Text(d.Label).End()
							}
						})

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
				})

			return nil
		})

	if err != nil {
		return err
	}

	return util.VisitorStop
}

const (
	rainHeight = 200.0
)

func rainGauge(v station.Visitor[*State], d *station.Gauge) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			dash := s.Dashboard()
			stn := dash.Station()
			metricValues, err := d.ConvertAll(d.Metrics.GetValues(stn))
			if err != nil {
				return err
			}
			metric := metricValues[0]
			val := metric.Float()

			axis := station.AutoScale(0, math.Max(2, val), rainHeight-10)

			valOff := axis.Scale * (val - axis.Min)
			// Axes tick marks
			var points []string
			for _, p := range axis.Points {
				points = append(points, fmt.Sprintf("M80,%d l10,0", int(rainHeight-p)))
			}

			s.Builder().
				Svg().
				ViewBox(0, 0, 125, 250).
				Attr("role", "img").
				If(s.IsLive(), func(e *html.Element) *html.Element {
					return e.Attr(ID, "%s.svg", d.GetID()).
						Attr("data-min", fix(axis.Min)).
						Attr("data-max", fix(axis.Max)).
						Attr("data-scale", fix(axis.Scale)).
						Attr("data-height", fix(rainHeight))
				}).
				// The actual blue water level in the gauge
				Rect().
				If(s.IsLive(), func(e *html.Element) *html.Element {
					return e.Attr(ID, "%s.rect", d.GetID())
				}).
				X(20).
				Y(int(rainHeight-valOff)).
				Width(60).
				Height(int(valOff)).
				Fill(LightBlue).
				End().
				// Axes tick marks
				G().Fill(None).Stroke(Black).
				Path().Attr(D, strings.Join(points, " ")).StrokeWidth("1px").End().
				Path().Attr(D, "M10,10 l10,0 l0,%d l60,0 l0,%d l10,0", int(rainHeight-10), int(10-rainHeight)).
				StrokeWidth("3px").End().
				End().
				// Labels
				G().Attr(DominantBaseline, Middle).
				ExecEnd(func(svg *html.Element) {
					// tick mark labels
					svg = svg.G().
						Attr(Transform, Translate(95, 0)).
						ExecEnd(func(svg *html.Element) {
							for i, y := range axis.Points {
								svg = svg.SvgText().Y(int(rainHeight - y)).Text(axis.Labels[i]).End()
							}
						}).
						// Labels
						G().
						Attr(Transform, Translate(50, rainHeight+15)).
						Attr(TextAnchor, Middle).
						ExecEnd(func(svg *html.Element) {
							if d.Label != "" {
								svg = svg.SvgText().Text(d.Label).End()
							}
							svg = svg.SvgText().
								If(s.IsLive(), func(e *html.Element) *html.Element {
									return e.Attr(ID, "%s.txt", d.GetID())
								}).
								Y(15).Text(metric.String()).
								End()
						})
				})

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
			Attr("d", "M0,10l0,%s", fix((math.Min(5, float64(i))*10)-80)).
			If(s.IsLive(), func(svg *html.Element) *html.Element {
				return svg.AnimateTransform().
					Attr(ID, "%s.ptr%d", d.GetID(), i).
					Attr("attributeName", "transform").
					Attr("attributeType", "XML").
					Attr("type", "rotate").
					Attr("from", "0").
					Attr("to", "0").
					Attr("dur", "1s").
					Fill("freeze").
					End()
			}).
			End(). // path
			End()  // g
	}
}
