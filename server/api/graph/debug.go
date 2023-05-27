package graph

import (
	"bytes"
	"context"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/station"
	"net/http"
	"path"
	"sort"
	"strings"
)

// FIXME this is for development purposes only - to be removed when dashboards implemented

// registerSensors adds endpoints for a Sensors object
func (s *SVG) registerSensors(ctx context.Context) error {
	sensors := station.SensorsFromContext(ctx)
	sensorsPath := path.Join("/svg", path.Join(strings.Split(sensors.ID, ".")...))

	// Work out what indices we want available
	lineCount := 0
	_ = station.NewVisitor().
		Graph(func(ctx context.Context) error {
			g := station.GraphFromContext(ctx)
			switch {
			case g.Line != nil:
				lineCount++
			}
			return nil
		}).
		WithContext(context.Background()).
		VisitSensors(sensors)

	if lineCount > 0 {
		err := s.Inbound.RegisterHttpEndpoint(
			"svg "+sensors.Name,
			path.Join(sensorsPath, "day"),
			sensors.ID,
			"Index last 24 hours",
			http.MethodGet,
			"html",
			task.Of(s.serveSensorDay).Using(sensors.WithContext),
		)
		if err != nil {
			return err
		}

		err = s.Inbound.RegisterHttpEndpoint(
			"svg "+sensors.Name,
			path.Join(sensorsPath, "today"),
			sensors.ID,
			"Index since midnight",
			http.MethodGet,
			"html",
			task.Of(s.serveSensorToday).Using(sensors.WithContext),
		)
		if err != nil {
			return err
		}
	}

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

	sensors := station.SensorsFromContext(ctx)

	// Sort keys so we have some sort of order with the results
	keys := append(sensors.ReadingsKeys(), sensors.CalculationsKeys()...)
	sort.SliceStable(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})

	var buf bytes.Buffer

	// Only include Graph with a path defined
	buf.WriteString("<html><body>")
	for _, k := range keys {
		graphs := sensors.GetGraph(k)
		if graphs != nil {
			for _, g := range graphs {
				if g.Path != "" {
					buf.WriteString("<object type=\"image/svg+xml\" data=\"" + g.Path + "/" + img + "\"></object>")
				}
			}
		}
	}
	buf.WriteString("</body></html>")

	r.Status(http.StatusOK).
		HTML().
		Value(buf.Bytes())

	return nil
}
