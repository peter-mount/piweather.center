package graph

import (
	svg "github.com/ajstarks/svgo/float"
	"github.com/peter-mount/piweather.center/weather/value"
)

func CalculateStep(min, max float64) float64 {
	if max < min {
		min, max = max, min
	}

	vRange := max - min
	e := 1.0
	for i := 1; i < 10; i++ {
		if vRange < e {
			// return previous e value
			return e / 10
		}
		e = e * 10.0
	}
	// range is >= 1e10 so default to 1e10
	return e
}

func DrawYAxisGrid(canvas *svg.SVG, proj *Projection, stepFactor float64, styles ...string) {

	minY, maxY, stepY := proj.YAxisTicks()

	// Apply stepFactor
	if value.NotEqual(stepFactor, 0.0) && value.NotEqual(stepFactor, 1.0) {
		stepY = stepY * stepFactor
	}

	if len(styles) > 0 {
		canvas.Group(styles...)
	}

	p := &Path{}
	for i := minY; i <= maxY; i += stepY {
		_, y := proj.Project(0, i)
		p.Line(proj.X0(), y, proj.X1(), y)
	}
	p.Draw(canvas)

	if len(styles) > 0 {
		canvas.Gend()
	}
}

func GetLabelFormat(step float64) string {
	switch {
	case value.LessThanEqual(step, 0.3):
		return "%.3f"
	case value.LessThanEqual(step, 0.1):
		return "%.2f"
	case value.LessThanEqual(step, 1):
		return "%.1f"
	default:
		return "%.0f"
	}
}

func DrawYAxisLegend(title, subTitle string, canvas *svg.SVG, proj *Projection, styles ...string) {
	minY, maxY, stepY := proj.YAxisTicks()

	x0, yc := proj.X0(), proj.Yc()

	canvas.Group(styles...)
	canvas.Gend()

	x := x0 - LabelSize

	f := GetLabelFormat(stepY)
	p := &Path{}
	for i := minY; i <= maxY; i += stepY {
		_, y := proj.Project(0, i)
		p.MoveTo(x, y).LineTo(x0, y)
		Text(canvas, x, y, -90, "labelY", f, i)
	}
	p.Draw(canvas, StrokeBlack, StrokeWidth1)

	if subTitle != "" {
		x = x - SubTitleSize
		Text(canvas, x, yc, -90, "subTitleY", subTitle)
	}

	if title != "" {
		x = x - TitleSize
		Text(canvas, x, yc, -90, "titleY", title)
	}
}
