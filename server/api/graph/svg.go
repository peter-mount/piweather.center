package graph

import (
	"context"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/graph/chart/gauge"
	"github.com/peter-mount/piweather.center/graph/chart/line"
	"github.com/peter-mount/piweather.center/server/api"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/station/service"
)

func init() {
	kernel.Register(&SVG{})
}

// SVG provides the /api/svg endpoint which displays svg graphs for a metric
type SVG struct {
	Inbound *api.EndpointManager `kernel:"inject"`
	//Store   store.Store          `kernel:"inject"`
	Config service.Config `kernel:"inject"`
}

func (s *SVG) Start() error {
	return s.Config.Accept(station.NewVisitor().
		Sensors(s.registerSensors).
		Graph(s.registerGraph).
		WithContext(context.Background()))
}

// registerGraph adds endpoints for a Graph object
func (s *SVG) registerGraph(ctx context.Context) error {
	g := station.GraphFromContext(ctx)
	switch {
	case g.Gauge != nil:
		return s.registerSvgChartEndpoint(g, "Gauge", gauge.New, gaugeWidth, gaugeHeight, ServeLatest)

	case g.Line != nil:
		return s.registerSvgChartEndpoint(g, "Line graph", line.New, svgWidth, svgHeight, ServeDay, ServeToday)

	default:
		// No Chart defined so remove path, so that we don't use it elsewhere
		g.Path = ""
	}
	return nil
}
