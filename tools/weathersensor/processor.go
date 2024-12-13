package weathersensor

import (
	"github.com/peter-mount/piweather.center/config/station"
	util2 "github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"github.com/peter-mount/piweather.center/tools/weathersensor/payload"
	"github.com/peter-mount/piweather.center/util/strings"
)

func (s *Service) process(id string, d *station.Sensor, body []byte) error {

	p, err := payload.FromBytes(id, d.Http.Format.GetType(), d.Http.Timestamp, body)
	if err == nil && p != nil {
		err = processorVisitor.Clone().
			Set(&processor{
				id:      id,
				service: s,
				payload: p,
			}).
			Sensor(d)
	}

	return err
}

type processor struct {
	service *Service
	id      string
	payload *payload.Payload
	sensor  *station.Sensor
	http    *station.Http
	r       *reading.Reading
}

var (
	processorVisitor = station.NewBuilder[*processor]().
		Sensor(processSensor).
		Http(processHttp).
		SourceParameter(processSourceParameter).
		Build()
)

func processSensor(v station.Visitor[*processor], d *station.Sensor) error {
	s := v.Get()
	s.sensor = d
	s.r = &reading.Reading{
		ID:   s.id,
		Time: s.payload.Time(),
	}
	defer func() {
		s.sensor = nil
		s.r = nil
	}()

	err := v.SourceParameterList(d.Http.SourceParameters)
	if err == nil {
		err = s.service.GetPublisher(s.r.ID).Do(s.r)
	}

	if err == nil {
		err = util2.VisitorStop
	}
	return err
}

func processHttp(v station.Visitor[*processor], d *station.Http) error {
	v.Get().http = d
	return nil
}

func processSourceParameter(v station.Visitor[*processor], d *station.SourceParameter) error {
	var err error

	s := v.Get()

	src, exists := s.payload.Get(d.Source)
	if exists {
		if f, ok := strings.ToFloat64(src); ok {
			val := d.Unit.Value(f)
			val, err = d.Metric.Convert(val)
			if err == nil && val.IsValid() {
				s.r.Set(d.Metric.Name, val)
			}
		}
	}

	return err
}
