package ecowitt

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/station"
	"path"
)

type Server struct {
	Rest      *rest.Server `kernel:"inject"`
	endPoints []*Endpoint
}

type Endpoint struct {
	Path    string
	Sensors *station.Sensors // Sensor collection for the endpoint
	Task    task.Task
}

func (s *Server) RegisterEndpoint(sp *station.Sensors, t task.Task) error {
	p := path.Clean(sp.Source.EcoWitt.Path)
	if p == "." || p == "/" {
		return fmt.Errorf("ecowitt path invalid")
	}

	e := &Endpoint{
		Path:    path.Join("/api/ecowitt", p),
		Sensors: sp,
		Task:    t,
	}

	for _, ep := range s.endPoints {
		if e.Path == ep.Path {
			return fmt.Errorf("ecowitt path %q already in use", e.Path)
		}
	}

	s.endPoints = append(s.endPoints, e)

	log.Printf("ecowitt:registering %q", e.Path)
	s.Rest.Do(e.Path, e.Task).Methods("POST")

	return nil
}
