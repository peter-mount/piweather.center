package graph

import (
	"bytes"
	"context"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/graph"
	"github.com/peter-mount/piweather.center/graph/chart"
	"github.com/peter-mount/piweather.center/graph/svg"
	"github.com/peter-mount/piweather.center/station"
	"net/http"
	"path"
	"strings"
	"time"
)

func (s *SVG) registerSvgChartEndpoint(graph *station.Graph, name string, chartFactory chart.ChartFactory, width, height float64, factories ...GeneratorFactory) error {
	return s.registerSvgEndpoint(graph, name, s.createChartGenerator(chartFactory, svg.NewRect(0, 0, width, height)), factories...)
}

func (s *SVG) registerSvgEndpoint(graph *station.Graph, name string, generator Generator, factories ...GeneratorFactory) error {
	id := graph.Sensor().GetID()
	graph.Path = path.Join("/svg", path.Join(strings.Split(id, ".")...))

	sensorName := "svg " + graph.Sensor().Sensors().Name

	for _, factory := range factories {
		urlPath, suffix, task := factory(graph.Path, generator)
		if err := s.Inbound.RegisterHttpEndpoint(
			sensorName,
			urlPath,
			id, name+suffix,
			http.MethodGet,
			"svg",
			task.WithValue("id", id).
				Using(graph.WithContext),
		); err != nil {
			return err
		}
	}
	return nil
}

func (s *SVG) createChartGenerator(factory chart.ChartFactory, bounds svg.Rect) Generator {
	return func(start, end time.Time, ctx context.Context) error {
		r := rest.GetRest(ctx)

		l := factory()

		if ok, err := s.initChart(start, end, ctx, l); err != nil {
			return err
		} else if !ok {
			r.Status(http.StatusNotFound)
			return nil
		}

		l.SetBounds(bounds)

		var buf bytes.Buffer

		svg.New(&buf, bounds.Width(), bounds.Height(), func(s svg.SVG) {
			graph.CSS(s)
			s.Draw(l)
		})

		r.Status(http.StatusOK).
			ContentType("image/svg+xml").
			Value(buf.Bytes())

		return nil
	}
}
