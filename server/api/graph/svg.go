package graph

import (
	"bytes"
	"context"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/graph"
	"github.com/peter-mount/piweather.center/graph/chart"
	svg2 "github.com/peter-mount/piweather.center/graph/svg"
	"github.com/peter-mount/piweather.center/server/store"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"net/http"
	"strings"
	"time"
)

func init() {
	kernel.Register(&SVG{})
}

const (
	svgWidth  = 1024
	svgHeight = 200
)

// SVG provides the /api/svg endpoint which displays svg graphs for a metric
type SVG struct {
	Rest  *rest.Server `kernel:"inject"`
	Store *store.Store `kernel:"inject"`
}

func (s *SVG) Start() error {
	s.Rest.Do("/api/svg/{stationId}/{sensorId}/{readingId}/hour.svg", s.serveHour).Methods(http.MethodGet)
	s.Rest.Do("/api/svg/{stationId}/{sensorId}/{readingId}/day.svg", s.serveDay).Methods(http.MethodGet)
	s.Rest.Do("/api/svg/{stationId}/{sensorId}/{readingId}/today.svg", s.serveToday).Methods(http.MethodGet)
	return nil
}

func (s *SVG) serveHour(ctx context.Context) error {
	now := time.Now()
	start := now.Truncate(time.Hour)
	end := start.Add(time.Hour)
	return s.serve(start, end, ctx)
}

func (s *SVG) serveDay(ctx context.Context) error {
	now := time.Now()
	start := now.Add(-24 * time.Hour)
	return s.serve(start, now, ctx)
}

func (s *SVG) serveToday(ctx context.Context) error {
	now := time.Now()
	start := now.Truncate(time.Hour * 24)
	end := start.Add(time.Hour * 24)
	return s.serve(start, end, ctx)
}

func (s *SVG) serve(start, end time.Time, ctx context.Context) error {
	r := rest.GetRest(ctx)
	stationId := r.Var("stationId")
	sensorId := r.Var("sensorId")
	readingId := r.Var("readingId")

	id := strings.ToLower(strings.Join([]string{stationId, sensorId, readingId}, "."))

	readings := s.Store.GetHistoryBetween(id, start, end)
	if readings == nil {
		r.Status(http.StatusNotFound)
		return nil
	}

	var buf bytes.Buffer
	s.draw(&buf, id, readings)

	r.Status(http.StatusOK).
		ContentType("image/svg+xml").
		Value(buf.Bytes())

	return nil
}

func (s *SVG) draw(buf *bytes.Buffer, id string, readings chart.DataSource) {

	// Get min/max
	start, end := readings.GetXRange()
	minVal, maxVal := readings.GetYRange()

	// Nearest 10 min or 1 hour?
	xStep := 60.0
	if end.Sub(start) < time.Hour {
		xStep = 10.0
	}

	proj := svg2.NewProjection(svgWidth-960-1, 1, 960, svgHeight-10).
		SetXRange(0, end.Sub(start).Minutes()).
		SetYRange(minVal.Float(), maxVal.Float()).
		ZeroY().
		NearestY(10.0).
		NearestX(xStep).
		Build()

	svg2.New(buf, svgWidth, svgHeight, func(s svg2.SVG) {
		//s.Defs()
		graph.CSS(s)

		s.Rect(proj.X0(), proj.Y0(), proj.Width()-1, proj.Height()-1, graph.StrokeBlack, graph.StrokeWidth1, graph.FillWhite)

		s.Text(proj.X0()+5, proj.Y0()+15, 0, id, "class=\"graphId\"")

		unit := readings.GetUnit()

		graph.DrawYAxisLegend(s, proj, unit.Name(), unit.Unit(), graph.StrokeBlack)
		graph.DrawYAxisGrid(s, proj, 0.2, graph.StrokeLightGrey, graph.StrokeWidth1)
		graph.DrawYAxisGrid(s, proj, 1.0, graph.StrokeGrey, graph.StrokeWidth1)

		p := &svg2.Path{}
		readings.ForEach(func(i int, t time.Time, value value.Value) {
			if util.TimeBetween(t, start, end) {
				p.AddProjectX(t.Sub(start).Minutes(), value.Float(), proj)
			}
		})
		s.Draw(p, graph.StrokeRed, graph.StrokeWidth1, graph.FillNone)

	})
}
