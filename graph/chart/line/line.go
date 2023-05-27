package line

import (
	"fmt"
	"github.com/peter-mount/piweather.center/graph"
	"github.com/peter-mount/piweather.center/graph/chart"
	"github.com/peter-mount/piweather.center/graph/svg"
	"github.com/peter-mount/piweather.center/util"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

// Line handles a simple Line plot
type Line struct {
	chart.AbstractChart                 // Common implementation
	proj                *svg.Projection // Projection of the plot
}

func New() chart.Chart {
	return &Line{}
}

func (c *Line) Type() string { return "line" }

func (l *Line) Draw(s svg.SVG, styles ...string) {
	if len(styles) > 0 {
		s.Group(l.draw, styles...)
	} else {
		l.draw(s)
	}
}

func (l *Line) draw(s svg.SVG) {
	definition := l.Definition()

	period := l.Period()

	yRange, _ := l.GetYRange()
	//unit := yRange.Unit()

	bounds := l.Bounds()

	// Plot area accounting for axes
	src := l.Sources()
	src1 := l.Sources()[0]
	src1unit := src1.DataSource().GetUnit()

	// Reduce left & bottom to allow for grid labels
	plotArea := bounds.Reduce(10+graph.LabelSize, 10, 10, 10+graph.LabelSize)

	// Account for top title
	if definition.Title != "" {
		plotArea = plotArea.Reduce(0, graph.TitleSize, 0, 0)
	}

	// y-axis title
	if definition.Line.YTitle != "" {
		plotArea = plotArea.Reduce(graph.TitleSize, 0, 0, 0)
	}

	// y-axis subtitle
	if util.StringDefault(definition.Line.YSubTitle, src1unit.Unit()) != "" {
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
		NearestY(util.FloatDefault(definition.Line.YStep, 10.0))

	if definition.Line.Min != nil || definition.Line.Max != nil {
		min, max := proj.GetYRange()
		proj.SetYRange(
			util.FloatDefault(definition.Line.Min, min),
			util.FloatDefault(definition.Line.Max, max),
		)
	}
	if definition.Line.ZeroY {
		proj.ZeroY()
	}

	if definition.Line.XStep != nil {
		proj.NearestX(*definition.Line.XStep)
	} else {
		switch {
		// Every 10 minutes if under an hour
		case period.Duration() <= time.Hour:
			proj.NearestX(10.0)

		// Every day if over 1 day
		case period.Duration() > (24 * time.Hour):
			proj.NearestX(1440.0)

		// Default every hour
		default:
			proj.NearestX(60.0)
		}
	}

	proj.Build()

	// Minor grid lines if defined
	if definition.Line.XSubStep != nil || definition.Line.YSubStep != nil {
		s.Group(func(_ svg.SVG) {
			if definition.Line.XSubStep != nil {
				graph.DrawXAxisGrid(s, proj, *definition.Line.XSubStep)
			}
			if definition.Line.YSubStep != nil {
				graph.DrawYAxisGrid(s, proj, *definition.Line.YSubStep)
			}
		}, "class=\"grid1\"")
	}

	s.Group(func(_ svg.SVG) {
		graph.DrawYAxisGrid(s, proj, util.FloatDefault(definition.Line.YStep, 1.0))
		graph.DrawXAxisGrid(s, proj, util.FloatDefault(definition.Line.XStep, 1.0))
	}, "class=\"grid0\"")

	s.Group(func(_ svg.SVG) {
		graph.DrawYAxisLegend(s, proj,
			definition.Line.YTitle,
			util.StringDefault(definition.Line.YSubTitle, src1unit.Unit()))

		graph.DrawXAxisLegend(s, proj,
			period.Start(),
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

	s.Clip(
		func(s svg.SVG) {
			s.Draw(plotArea)
		},
		func(s svg.SVG) {
			for _, src := range src {
				p := &svg.Path{}
				src.DataSource().ForEach(func(i int, t time.Time, v value.Value) {
					if period.Contains(t) {
						p.AddProjectX(period.MinutesFromStart(t), v.Float(), proj)
					}
				})
				s.Draw(p, graph.StrokeRed, graph.StrokeWidth1, graph.FillNone)
			}
		})

	s.Draw(plotArea, graph.StrokeBlack, graph.StrokeWidth1, graph.FillNone)
}
