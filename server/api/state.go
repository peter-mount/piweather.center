package api

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/store"
	"net/http"
	"strings"
)

// Api general APIs
type Api struct {
	Config    station.Config   `kernel:"inject"`
	Endpoints *EndpointManager `kernel:"inject"`
	State     *store.State     `kernel:"inject"`
}

func (s *Api) Start() error {
	for _, stn := range *s.Config.Stations() {
		if err := s.Endpoints.RegisterHttpEndpoint(
			"api",
			"/api/station/"+stn.ID+".json",
			stn.ID,
			"Station State",
			http.MethodGet,
			"json",
			s.stationState(stn.ID),
		); err != nil {
			return err
		}

		for _, sensor := range stn.Sensors {
			if err := s.Endpoints.RegisterHttpEndpoint(
				"api",
				"/api/station/"+strings.Join(strings.Split(sensor.ID, "."), "/")+".json",
				sensor.ID,
				"Sensor State",
				http.MethodGet,
				"json",
				s.sensorState(stn.ID, sensor.ID),
			); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Api) stationState(id string) task.Task {
	return func(ctx context.Context) error {
		r := rest.GetRest(ctx)

		stn := s.State.GetStation(id)
		if stn == nil {
			r.Status(http.StatusNotFound)
		} else {
			r.JSON().
				CacheControl(60).
				AccessControlAllowOrigin("*").
				Value(stn)
		}

		return nil
	}
}

func (s *Api) sensorState(stationId, sensorId string) task.Task {
	return func(ctx context.Context) error {
		r := rest.GetRest(ctx)

		stn := s.State.GetStation(stationId)
		if stn == nil {
			r.Status(http.StatusNotFound)
		} else {
			prefix := sensorId + "."

			// Duplicate station but with just this sensor's measurements
			sensor := stn.Clone()
			for _, m := range stn.Measurements {
				if strings.HasPrefix(m.ID, prefix) {
					sensor.Measurements = append(sensor.Measurements, m)
				}
			}

			r.JSON().
				CacheControl(60).
				AccessControlAllowOrigin("*").
				Value(sensor)
		}

		return nil
	}
}
