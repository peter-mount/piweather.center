package sensors

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/sensors"
	"github.com/peter-mount/piweather.center/config/util/time"
	"strings"
)

func NewParser() util.Parser[sensors.Sensors] {
	return util.NewParser[sensors.Sensors](nil, nil, sensorsInit)
}

func sensorsInit(q *sensors.Sensors, err error) (*sensors.Sensors, error) {
	if err == nil {
		s := &state{
			ids: make(map[string]*sensors.Sensor),
		}

		err = NewBuilder[*state]().
			Sensors(s.sensors).
			Sensor(s.sensor).
			I2C(s.i2c).
			Serial(s.serial).
			Poll(s.poll).
			Build().
			Sensors(q)
	}
	return q, err
}

type state struct {
	ids map[string]*sensors.Sensor
}

func (s *state) sensors(_ sensors.SensorVisitor[*state], q *sensors.Sensors) error {
	if len(q.Sensors) == 0 {
		return participle.Errorf(q.Pos, "No sensors defined")
	}
	return nil
}

func (s *state) sensor(_ sensors.SensorVisitor[*state], q *sensors.Sensor) error {
	// IDs are lower case & must be unique within this instance
	q.ID = strings.TrimSpace(strings.ToLower(q.ID))

	if e, exists := s.ids[strings.ToLower(q.ID)]; exists {
		return participle.Errorf(q.Pos, "sensor %q already defined at %s", q.ID, e.Pos.String())
	}

	s.ids[q.ID] = q

	return nil
}

func (s *state) i2c(_ sensors.SensorVisitor[*state], q *sensors.I2C) error {
	if q.Bus < 1 || q.Device < 0 {
		return participle.Errorf(q.Pos, "i2c address %d:%02x is invalid", q.Bus, q.Device)
	}
	if q.Device <= 7 || (q.Device&0x78) != 0 {
		return participle.Errorf(q.Pos, "i2c address %d:%02x is reserved", q.Bus, q.Device)
	}
	return nil
}

func (s *state) serial(_ sensors.SensorVisitor[*state], q *sensors.Serial) error {
	q.Port = strings.TrimSpace(q.Port)
	if q.Port == "" {
		return participle.Errorf(q.Pos, "serial port is invalid")
	}

	if q.BaudRate < 1 {
		return participle.Errorf(q.Pos, "serial baudrate %d is invalid", q.BaudRate)
	}

	return nil
}

func (s *state) poll(_ sensors.SensorVisitor[*state], q *time.CronTab) error {
	return q.Init()
}
