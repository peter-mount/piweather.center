package graph

import (
	"github.com/peter-mount/piweather.center/station"
	"net/http"
)

func (s *SVG) registerSvgEndpoints(sensor *station.Sensors, graph *station.Graph, id, name string, generator Generator, factories ...GeneratorFactory) error {
	for _, factory := range factories {
		if err := s.registerSvgEndpoint(sensor, graph, id, name, generator, factory); err != nil {
			return err
		}
	}
	return nil
}

func (s *SVG) registerSvgEndpoint(sensor *station.Sensors, graph *station.Graph, id, name string, generator Generator, factory GeneratorFactory) error {
	path, suffix, task := factory(graph.Path, generator)
	return s.Inbound.RegisterHttpEndpoint(
		"svg "+sensor.Name,
		path,
		id, name+suffix,
		http.MethodGet,
		"svg",
		task.WithValue("id", id).
			Using(graph.WithContext),
	)
}
