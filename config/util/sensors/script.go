package sensors

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/amqp"
	"github.com/peter-mount/piweather.center/config/util/time"
)

type Sensors struct {
	Pos     lexer.Position
	Sensors []*Sensor `parser:"@@*"`
}

func (s *Sensors) Merge(b *Sensors) (*Sensors, error) {
	if s == nil {
		return b, nil
	}

	if b != nil {
		s.Sensors = append(s.Sensors, b.Sensors...)

		m := make(map[string]*Sensor)
		for _, sensor := range s.Sensors {
			if e, exists := m[sensor.ID]; exists {
				return nil, participle.Errorf(sensor.Pos, "sensor %q already defined at %s", sensor.ID, e.Pos.String())
			}
			m[sensor.ID] = sensor
		}
	}
	return s, nil
}

type Sensor struct {
	Pos       lexer.Position
	ID        string        `parser:"'sensor' @String"`
	Device    string        `parser:"'device' @String"`
	I2C       *I2C          `parser:"( @@"`
	Serial    *Serial       `parser:"| @@ )"`
	Poll      *time.CronTab `parser:"('poll' @@)?"`
	Publisher []*Publisher  `parser:"'publish' @@+"`
}

type Publisher struct {
	Pos         lexer.Position
	Log         bool       `parser:"( @'log'"`
	FilterEmpty bool       `parser:"| @'filterempty'"`
	Amqp        *amqp.Amqp `parser:"| @@ )"`
}
