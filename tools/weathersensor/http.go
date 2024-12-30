package weathersensor

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
	"io"
	"net/http"
	"strings"
)

func (s *Service) httpSensor(v station.Visitor[*state], d *station.Http) error {
	st := v.Get()

	err := s.addHttp(strings.ToUpper(d.Method), st.station.Name, st.sensor.Target.Name, st.sensor)

	if err == nil {
		s.sensorCount++
	}

	return err
}

func (s *Service) handleHttp(r *rest.Rest) error {
	stationId := r.Var("stationId")
	sensorId := r.Var("sensorId")

	d := s.GetHttp(r.Request().Method, stationId, sensorId)
	if d == nil {
		r.Status(http.StatusNotFound)
		return nil
	}

	var body []byte
	switch d.Http.Method {
	case "get":
		body = []byte(r.Request().URL.RawQuery)
	case "post", "put", "patch":
		body, _ = io.ReadAll(r.Request().Body)
	}
	if len(body) == 0 {
		r.Status(http.StatusBadRequest)
		return nil
	}

	err := s.process(stationId+"."+sensorId, d, body)
	if err != nil {
		r.Status(http.StatusInternalServerError).
			ContentType("text/plain; charset=utf-8").
			Value(errors.Error(d.Pos, err).Error())

		log.Printf("http %s for %s.%s: %v", d.Http.Method, stationId, sensorId, err)
	}

	return nil
}
