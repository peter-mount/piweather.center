package chart

import (
	"errors"
	"github.com/peter-mount/piweather.center/graph/svg"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
)

type ChartFactory func() Chart

func Factory(c ChartFactory) ChartFactory { return c }

type Chart interface {
	// Type returns the type of chart, e.g. "line", "gauge" etc.
	// This is used in forming url paths
	Type() string
	// Add a Source to the Chart
	Add(Source) Chart
	// Period of the Chart
	Period() time.Period
	// SetPeriod sets the overall Period of the chart.
	// By default, it consists of the union of each datasource's Period
	SetPeriod(time.Period) Chart
	// Bounds of the chart
	Bounds() svg.Rect
	// SetBounds of the chart
	SetBounds(svg.Rect) Chart
	// Draw the Chart
	Draw(svg.SVG, ...string)
	SetDefinition(*station.Graph) Chart
	Definition() *station.Graph
}

// AbstractChart is an implementation of Chart which other implementations
// can use
type AbstractChart struct {
	sources    []Source
	period     time.Period
	bounds     svg.Rect
	definition *station.Graph
}

func (c *AbstractChart) Type() string { return "???" }

func (c *AbstractChart) Sources() []Source { return c.sources }

func (c *AbstractChart) Add(s Source) Chart {
	c.sources = append(c.sources, s)
	c.period = c.period.Include(s.DataSource().Period())
	return c
}

func (c *AbstractChart) GetYRange() (*value.Range, error) {
	var r *value.Range
	for i, s := range c.sources {
		if i == 0 {
			r = s.DataSource().GetYRange().Clone()
		} else {
			if err := r.Include(s.DataSource().GetYRange()); err != nil {
				return nil, err
			}
		}
	}
	return r, nil
}

func (c *AbstractChart) Period() time.Period { return c.period }

func (c *AbstractChart) SetPeriod(s time.Period) Chart {
	c.period = s
	return c
}

func (c *AbstractChart) Bounds() svg.Rect { return c.bounds }

func (c *AbstractChart) SetBounds(r svg.Rect) Chart {
	c.bounds = r
	return c
}

func (c *AbstractChart) SetDefinition(definition *station.Graph) Chart {
	c.definition = definition
	return c
}

func (c *AbstractChart) Definition() *station.Graph { return c.definition }

// Draw from svg.Drawable.
// This panics nothing as it's down to Chart implementors to implement.
//
// Note: Don't do this: l:=NewLine().Add(xxxx); s.Draw(l) as this will call this function.
//
// Instead use: l:=NewLine(); l.Add(xxxx);s.Draw(l) as that will call the Draw in Line.
func (c *AbstractChart) Draw(_ svg.SVG, _ ...string) {
	panic(abstractErr)
}

var (
	abstractErr = errors.New("AbstractChart.Draw() invoked. Use actual type when Drawing to get correct function")
)
