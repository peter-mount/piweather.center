package graph

import (
	"context"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/server/api"
	"github.com/peter-mount/piweather.center/server/store"
	"github.com/peter-mount/piweather.center/station"
	"path"
	"strings"
)

func init() {
	kernel.Register(&SVG{})
}

const (
	svgWidth  = 1024
	svgHeight = 132
)

// SVG provides the /api/svg endpoint which displays svg graphs for a metric
type SVG struct {
	Inbound *api.EndpointManager `kernel:"inject"`
	Store   *store.Store         `kernel:"inject"`
	Config  station.Config       `kernel:"inject"`
}

func (s *SVG) Start() error {
	return s.Config.Accept(station.NewVisitor().
		Sensors(s.registerSensors).
		Graph(s.registerGraph).
		WithContext(context.Background()))
}

// registerGraph adds endpoints for a Graph object
func (s *SVG) registerGraph(ctx context.Context) error {
	sensors := station.SensorsFromContext(ctx)
	g := station.GraphFromContext(ctx)

	id := g.Sensor().GetID()
	g.Path = path.Join("/svg", path.Join(strings.Split(id, ".")...))

	switch {
	case g.Line != nil:
		return s.registerSvgEndpoints(sensors, g, id, "Line graph", s.serveLine, ServeDay, ServeToday)

	default:
		// No Chart defined so remove path, so that we don't use it elsewhere
		g.Path = ""
	}
	return nil
}
