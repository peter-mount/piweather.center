package ecowitt

import "github.com/peter-mount/piweather.center/util"

func (s *Server) GetEndpoints(r *util.Endpoints) {
	for _, ep := range s.endPoints {
		r.Add(util.EndpointEntry{
			Id:       ep.Sensors.ID,
			Name:     ep.Sensors.Name,
			Endpoint: ep.Path,
			Protocol: "ecowitt",
			Method:   "POST",
		})
	}
}
