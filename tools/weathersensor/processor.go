package weathersensor

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"github.com/peter-mount/piweather.center/tools/weathersensor/payload"
	"github.com/peter-mount/piweather.center/util/strings"
)

func (s *Service) processHttp(id string, d *station.Sensor, body []byte) error {
	return s.process(id, d.Http.Format.GetType(), d.Http.Timestamp, d, body)
}

func (s *Service) process(id string, format station.HttpFormatType, timestamp *station.SourcePath, d *station.Sensor, body []byte) error {
	p, err := payload.FromBytes(id, format, timestamp, body)
	if err == nil && p != nil {
		err = s.processPayload(id, p, d)
	}
	return err
}

func (s *Service) processPayload(id string, p *payload.Payload, d *station.Sensor) error {
	return processorVisitor.Clone().
		Set(&processor{
			id:      id,
			service: s,
			payload: p,
		}).
		Sensor(d)
}

type processor struct {
	service *Service
	id      string
	payload *payload.Payload
	sensor  *station.Sensor
	r       *reading.Reading
}

var (
	processorVisitor = station.NewBuilder[*processor]().
		Sensor(processSensor).
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

	var err error
	switch {
	case d.Http != nil:
		err = v.SourceParameterList(d.Http.SourceParameters)
	case d.Rtl433 != nil:
		err = v.SourceParameterList(d.Rtl433.SourceParameters)
	}

	if err == nil {
		err = s.service.GetPublisher(s.r.ID).Do(s.r)
	}

	if err == nil {
		err = errors.VisitorStop
	}
	return err
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
