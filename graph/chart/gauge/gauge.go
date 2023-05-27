package gauge

import (
	"github.com/peter-mount/piweather.center/graph"
	"github.com/peter-mount/piweather.center/graph/chart"
	"github.com/peter-mount/piweather.center/graph/svg"
)

type Gauge struct {
	chart.AbstractChart // Common implementation
}

func New() chart.Chart {
	return &Gauge{}
}

func (c *Gauge) Type() string { return "gauge" }

func (g *Gauge) Draw(s svg.SVG, styles ...string) {
	if len(styles) > 0 {
		s.Group(g.draw, styles...)
	} else {
		g.draw(s)
	}
}

func (g *Gauge) draw(s svg.SVG) {
	//definition := g.Definition()

	bounds := g.Bounds()

	// Reduce left & bottom to allow for grid labels
	plotArea := bounds.Reduce(10, 10, 10, 10)

	s.Clip(
		func(s svg.SVG) {
			s.Draw(plotArea)
		},
		func(s svg.SVG) {
			s.Circle(plotArea.CX(), plotArea.CY(), plotArea.Radius(),
				graph.StrokeRed, graph.StrokeWidth1, graph.FillNone)

		})
	s.Draw(plotArea, graph.StrokeBlack, graph.StrokeWidth1, graph.FillNone)
}
