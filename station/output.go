package station

import (
	"context"
)

type Output struct {
	InfluxDB *InfluxDB `yaml:"influxdb,omitempty"`
}

func (s *Output) Accept(v Visitor) error {
	return v.VisitOutput(s)
}

type InfluxDB struct {
	Name        string `yaml:"name"`        // InfluxDB name
	Measurement string `yaml:"measurement"` // measurement
}

const (
	outputKey = "piweather.center/output"
)

func OutputFromContext(ctx context.Context) *Output {
	v := ctx.Value(outputKey)
	if o, ok := v.(*Output); ok {
		return o
	}
	return nil
}

func (g *Output) WithContext(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, outputKey, g), nil
}
