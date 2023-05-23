package graph

import (
	"bytes"
	"context"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/graph"
	"github.com/peter-mount/piweather.center/graph/chart"
	"github.com/peter-mount/piweather.center/graph/chart/line"
	"github.com/peter-mount/piweather.center/graph/svg"
	"github.com/peter-mount/piweather.center/server/store"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"net/http"
	"strings"
	"time"
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
	Rest  *rest.Server `kernel:"inject"`
	Store *store.Store `kernel:"inject"`
}

func (s *SVG) Start() error {
	s.Rest.Do("/api/svg/{stationId}/{sensorId}/{readingId}/hour.svg", s.serveHour).Methods(http.MethodGet)
	s.Rest.Do("/api/svg/{stationId}/{sensorId}/{readingId}/day.svg", s.serveDay).Methods(http.MethodGet)
	s.Rest.Do("/api/svg/{stationId}/{sensorId}/{readingId}/today.svg", s.serveToday).Methods(http.MethodGet)

	s.Rest.Do("/api/svg/{stationId}/{sensorId}/day", s.serveSensorDay).Methods(http.MethodGet)
	s.Rest.Do("/api/svg/{stationId}/{sensorId}/today", s.serveSensorToday).Methods(http.MethodGet)
	return nil
}

func (s *SVG) serveSensorDay(ctx context.Context) error {
	return s.serveSensor(ctx, "day.svg")
}

func (s *SVG) serveSensorToday(ctx context.Context) error {
	return s.serveSensor(ctx, "today.svg")
}

func (s *SVG) serveSensor(ctx context.Context, img string) error {
	r := rest.GetRest(ctx)
	prefix := r.Var("stationId") + "." + r.Var("sensorId") + "."

	var buf bytes.Buffer

	buf.WriteString("<html><body>")
	for _, k := range s.Store.GetKeys() {
		if strings.HasPrefix(k, prefix) {
			buf.WriteString("<img src=\"/api/svg/" + strings.ReplaceAll(k, ".", "/") + "/" + img + "\"/>")
		}
	}
	buf.WriteString("</body></html>")

	r.Status(http.StatusOK).
		HTML().
		Value(buf.Bytes())

	return nil
}

func (s *SVG) serveHour(ctx context.Context) error {
	now := time.Now()
	start := now.Truncate(time.Hour)
	end := start.Add(time.Hour)
	return s.serve(start, end, ctx)
}

func (s *SVG) serveDay(ctx context.Context) error {
	// End is end of current hour
	end := time.Now().Truncate(time.Hour).Add(time.Hour)
	start := end.Add(-24 * time.Hour)
	return s.serve(start, end, ctx)
}

func (s *SVG) serveToday(ctx context.Context) error {
	now := time.Now()
	// Start at beginning of the current local day
	//
	// Note: truncate to hour then subtract hours to get the start.
	// It might look weird when you could truncate to day, but that truncate
	// seems to set it to 0h UTC, so if we are in BST (UTC+1) then the day
	// starts at 0100 and not 0000 midnight.
	//
	// TODO check this works for other timezones
	start := now.Truncate(time.Hour)
	start = start.Add(time.Hour * time.Duration(-start.Hour()))

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

	l := line.New()
	l.Add(chart.NewUnitSource(id, readings)).
		SetPeriod(time2.PeriodOf(start, end)).
		SetBounds(svg.NewRect(0, 0, svgWidth, svgHeight))

	svg.New(&buf, svgWidth, svgHeight, func(s svg.SVG) {
		graph.CSS(s)
		s.Draw(l)
	})

	r.Status(http.StatusOK).
		ContentType("image/svg+xml").
		Value(buf.Bytes())

	return nil
}
