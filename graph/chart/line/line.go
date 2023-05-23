package line

import (
	"fmt"
	"github.com/peter-mount/piweather.center/graph"
	"github.com/peter-mount/piweather.center/graph/chart"
	"github.com/peter-mount/piweather.center/graph/svg"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

// Line handles a simple Line plot
type Line struct {
	chart.AbstractChart                 // Common implementation
	proj                *svg.Projection // Projection of the plot
}

func New() *Line {
	return &Line{}
}

func (l *Line) Draw(s svg.SVG, styles ...string) {
	if len(styles) > 0 {
		s.Group(l.draw, styles...)
	} else {
		l.draw(s)
	}
}

func (l *Line) draw(s svg.SVG) {
	period := l.Period()

	// Nearest 10min or 1hour?
	xStep := 60.0
	if period.Duration() <= time.Hour {
		xStep = 10.0
	}

	yRange, _ := l.GetYRange()
	unit := yRange.Unit()

	bounds := l.Bounds()

	// Plot area accounting for axes
	src := l.Sources()
	src1 := l.Sources()[0]
	src1unit := src1.DataSource().GetUnit()
	plotArea := bounds.Reduce(10+graph.LabelSize, 10, 10, 10+graph.LabelSize)
	if src1unit.Name() != "" {
		plotArea = plotArea.Reduce(graph.TitleSize, 0, 0, 0)
	}
	if src1unit.Unit() != "" {
		plotArea = plotArea.Reduce(graph.SubTitleSize, 0, 0, 0)
	}
	// TODO decide on how to set x-axis title, for now presume it's always there
	//if src1.SubTitle() != "" {
	plotArea = plotArea.Reduce(0, 0, 0, graph.SubTitleSize)
	/*
		}
		if src1.Title() != "" {
			plotArea = plotArea.Reduce(0, 0, 0, graph.TitleSize)
		}
	*/

	proj := plotArea.Projection().
		SetPeriod(period).
		SetValueRange(yRange).
		ZeroY().
		NearestY(10.0).
		NearestX(xStep).
		Build()

	s.Group(func(_ svg.SVG) {
		graph.DrawYAxisGrid(s, proj, 0.2)
		graph.DrawXAxisGrid(s, proj, 0.25)
	}, "class=\"grid1\"")

	s.Group(func(_ svg.SVG) {
		graph.DrawYAxisGrid(s, proj, 1.0)
		graph.DrawXAxisGrid(s, proj, 1.0)
	}, "class=\"grid0\"")

	s.Group(func(_ svg.SVG) {
		graph.DrawYAxisLegend(s, proj, unit.Name(), unit.Unit())
		graph.DrawXAxisLegend(s, proj,
			"",
			fmt.Sprintf("Time %s", time2.Zone(period.Start())),
			func(f float64) string {
				t := period.Start().Add(time.Minute * time.Duration(f))
				return fmt.Sprintf("%d", t.Hour())
			})
	}, "class=\"txt\"")

	// Single source so show at top left of graph
	if len(src) == 1 {
		s.Text(proj.X0()+5, proj.Y0()+15, 0, src1.Name(), "class=\"graphId\"")
	}

	for _, src := range src {
		p := &svg.Path{}
		src.DataSource().ForEach(func(i int, t time.Time, v value.Value) {
			if period.Contains(t) {
				p.AddProjectX(period.MinutesFromStart(t), v.Float(), proj)
			}
		})
		s.Draw(p, graph.StrokeRed, graph.StrokeWidth1, graph.FillNone)
	}

	s.Draw(plotArea, graph.StrokeBlack, graph.StrokeWidth1, graph.FillNone)
}
