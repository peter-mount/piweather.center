package station

import (
	"context"
)

type Output struct {
	InfluxDB *InfluxDB `json:"influxdb,omitempty" xml:"influxdb,omitempty" yaml:"influxdb,omitempty"`
}

func (s *Output) Accept(v Visitor) error {
	return v.VisitOutput(s)
}

type InfluxDB struct {
	Name        string `yaml:"name" xml:"name,attr" json:"name"`                      // InfluxDB name
	Measurement string `yaml:"measurement" xml:"measurement,attr" json:"measurement"` // measurement
}

const (
	outputKey = "piweather.center/output"
)

func OutputFromContext(ctx context.Context) *Output {
	return ctx.Value(outputKey).(*Output)
}

func (g *Output) WithContext(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, outputKey, g), nil
}
