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
		p.Line(x+3, y, x0, y)
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

func DrawXAxisGrid(s svg.SVG, proj *svg.Projection, stepFactor float64, styles ...string) {
	if len(styles) > 0 {
		s.Group(func(s svg.SVG) {
			drawXAxisGrid(s, proj, stepFactor)
		}, styles...)
	} else {
		drawXAxisGrid(s, proj, stepFactor)
	}
}

func drawXAxisGrid(s svg.SVG, proj *svg.Projection, stepFactor float64) {
	minX, maxX, stepX := proj.XAxisTicks()

	// Apply stepFactor
	if value.NotEqual(stepFactor, 0.0) && value.NotEqual(stepFactor, 1.0) {
		stepX = stepX * stepFactor
	}

	p := &svg.Path{}
	for i := minX; i <= maxX; i += stepX {
		x, _ := proj.Project(i, 0)
		p.Line(x, proj.Y0(), x, proj.Y1())
	}
	s.Draw(p)
}

func DrawXAxisLegend(s svg.SVG, proj *svg.Projection, title, subTitle string, f func(float64) string, styles ...string) {
	if len(styles) > 0 {
		s.Group(func(s svg.SVG) {
			drawXAxisLegend(s, proj, title, subTitle, f)
		}, styles...)
	} else {
		drawXAxisLegend(s, proj, title, subTitle, f)
	}
}

func drawXAxisLegend(s svg.SVG, proj *svg.Projection, title, subTitle string, f func(float64) string) {
	minX, maxX, stepX := proj.XAxisTicks()

	y1, xc := proj.Y1(), proj.Xc()

	y := y1 + LabelSize

	p := &svg.Path{}
	for i := minX; i <= maxX; i += stepX {
		x, _ := proj.Project(i, 0)
		p.Line(x, y-3, x, y1)
		s.Text(x, y, 0, f(i), "class=\"labelX\"")
	}
	s.Draw(p, StrokeBlack, StrokeWidth1)

	if subTitle != "" {
		y = y + SubTitleSize
		s.Text(xc, y, 0, svg.CData(subTitle), "class=\"subTitleX\"")
	}

	if title != "" {
		y = y + TitleSize
		s.Text(xc, y, 0, svg.CData(title), "class= \"titleX\"")
	}
}
