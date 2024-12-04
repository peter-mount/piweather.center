package weathersensor

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/tools/weathersensor/payload"
	"io"
	"net/http"
	"strings"
)

func (s *Service) httpSensor(v station.Visitor[*state], d *station.Http) error {
	st := v.Get()

	err := s.addHttp(strings.ToUpper(d.Method), st.station.Name, st.sensor.Target.OriginalName, d)

	if err == nil {
		s.sensorCount++
	}

	return err
}

func (s *Service) handleHttp(r *rest.Rest) error {
	stationId := r.Var("stationId")
	sensorId := r.Var("sensorId")
	log.Printf("http for %q %q", stationId, sensorId)

	d := s.GetHttp(r.Request().Method, stationId, sensorId)
	if d == nil {
		r.Status(http.StatusNotFound)
		return nil
	}

	var body []byte
	switch d.Method {
	case "get":
		body = []byte(r.Request().URL.RawQuery)
	case "post", "put", "patch":
		body, _ = io.ReadAll(r.Request().Body)
	default:
		r.Status(http.StatusBadRequest)
		return nil
	}

	p, err := payload.FromBytes(sensorId, d.Format.GetType(), d.Timestamp, body)
	if err == nil && p != nil {
		//
		log.Printf("Got %s\ntime %v", string(body), p.Time())
	}
	if err != nil {
		log.Println(err)
	}

	return errors.Error(d.Pos, err)
}
