package graph

import (
	"fmt"
	"github.com/peter-mount/piweather.center/graph/svg"
	"github.com/peter-mount/piweather.center/weather/value"
)

func DrawYAxisGrid(s svg.SVG, proj *svg.Projection, stepFactor float64, styles ...string) {
	if len(styles) > 0 {
		s.Group(func(s svg.SVG) {
			drawYAxisGrid(s, proj, stepFactor)
		}, styles...)
	} else {
		drawYAxisGrid(s, proj, stepFactor)
	}
}

func drawYAxisGrid(s svg.SVG, proj *svg.Projection, stepFactor float64) {
	minY, maxY, stepY := proj.YAxisTicks()

	// Apply stepFactor
	if value.NotEqual(stepFactor, 0.0) && value.NotEqual(stepFactor, 1.0) {
		stepY = stepY * stepFactor
	}

	p := &svg.Path{}
	for i := minY; i <= maxY; i += stepY {
		_, y := proj.Project(0, i)
		p.Line(proj.X0(), y, proj.X1(), y)
	}
	s.Draw(p)
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

func DrawYAxisLegend(s svg.SVG, proj *svg.Projection, title, subTitle string, styles ...string) {
	if len(styles) > 0 {
		s.Group(func(s svg.SVG) {
			drawYAxisLegend(s, proj, title, subTitle)
		}, styles...)
	} else {
		drawYAxisLegend(s, proj, title, subTitle)
	}
}

func drawYAxisLegend(s svg.SVG, proj *svg.Projection, title, subTitle string) {
	minY, maxY, stepY := proj.YAxisTicks()

	x0, yc := proj.X0(), proj.Yc()

	x := x0 - LabelSize

	f := GetLabelFormat(stepY)
	p := &svg.Path{}
	for i := minY; i <= maxY; i += stepY {
		_, y := proj.Project(0, i)
		p.Line(x, y, x0, y)
		s.Text(x, y, -90, fmt.Sprintf(f, i), "class=\"labelY\"")
	}
	s.Draw(p, StrokeBlack, StrokeWidth1)

	if subTitle != "" {
		x = x - SubTitleSize
		s.Text(x, yc, -90, svg.CData(subTitle), "class=\"subTitleY\"")
	}

	if title != "" {
		x = x - TitleSize
		s.Text(x, yc, -90, svg.CData(title), "class= \"titleY\"")
	}
}
