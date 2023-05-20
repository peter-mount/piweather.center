package graph

import (
	"bytes"
	"context"
	"fmt"
	svg "github.com/ajstarks/svgo/float"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/graph"
	"github.com/peter-mount/piweather.center/server/store"
	"math"
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

	readings := s.Store.GetHistory(id)
	if readings == nil {
		r.Status(http.StatusNotFound)
		return nil
	}

	var buf bytes.Buffer
	s.draw(&buf, id, start, end, readings)

	r.Status(http.StatusOK).
		ContentType("image/svg+xml").
		Value(buf.Bytes())

	return nil
}

func (s *SVG) draw(buf *bytes.Buffer, id string, start, end time.Time, readings []*store.Reading) {

	if start.After(end) {
		start, end = end, start
	}

	// Get min/max
	minVal, maxVal := math.MaxFloat64, -math.MaxFloat64
	for _, reading := range readings {
		v := reading.Value.Float()
		minVal, maxVal = math.Min(minVal, v), math.Max(maxVal, v)
	}

	// Nearest 10 min or 1 hour?
	xStep := 60.0
	if end.Sub(start) < time.Hour {
		xStep = 10.0
	}

	proj := graph.NewProjection(svgWidth-960-1, 1, 960, svgHeight-10).
		SetXRange(0, end.Sub(start).Minutes()).
		SetYRange(minVal, maxVal).
		ZeroY().
		NearestY(10.0).
		NearestX(xStep).
		Build()

	canvas := svg.New(buf)
	canvas.Startview(svgWidth, svgHeight, 0, 0, svgWidth, svgHeight)
	defer canvas.End()

	canvas.Def()
	graph.CSS(canvas)
	canvas.DefEnd()

	canvas.Rect(proj.X0(), proj.Y0(), proj.Width()-1, proj.Height()-1, graph.StrokeBlack, graph.StrokeWidth1, graph.FillWhite)

	graph.Text(canvas, proj.X0()+5, proj.Y0()+15, 0, "graphId", id)

	unit := readings[0].Value.Unit()
	graph.DrawYAxisLegend(unit.Name(), unit.Unit(), canvas, proj, graph.StrokeBlack)
	graph.DrawYAxisGrid(canvas, proj, 0.2, graph.StrokeLightGrey, graph.StrokeWidth1)
	graph.DrawYAxisGrid(canvas, proj, 1.0, graph.StrokeGrey, graph.StrokeWidth1)

	p := graph.Path{}
	for _, reading := range readings {
		p.AddProjectX(reading.Time.Sub(start).Minutes(), reading.Value.Float(), proj)
	}
	p.Draw(canvas, graph.StrokeRed, graph.StrokeWidth1, graph.FillNone)

}

func (s *SVG) plotXAxis(canvas *svg.SVG, plotWidth, xScale, x0, y0, y1 float64, start, len, step int) {
	var p []string
	for i := 0; i < len; i += step {
		x := x0 + (xScale * float64(i))
		p = append(p, fmt.Sprintf("M%.2f %.2fl0 %.2f", x, y0, y1-y0))
		x += xScale
	}
	canvas.Path(strings.Join(p, " "))
}
