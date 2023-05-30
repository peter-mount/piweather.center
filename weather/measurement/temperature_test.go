package measurement

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"testing"
)

func Test_temperature_transforms(t *testing.T) {
	tests := []struct {
		from    value.Value
		to      value.Value
		wantErr bool
	}{
		{Celsius.Value(0), Fahrenheit.Value(32.0), false},
		{Celsius.Value(10), Fahrenheit.Value(50.0), false},
		{Fahrenheit.Value(32), Celsius.Value(0.0), false},
		{Fahrenheit.Value(50), Celsius.Value(10.0), false},
		{Celsius.Value(0), Kelvin.Value(273.15), false},
		{Celsius.Value(10), Kelvin.Value(283.15), false},
		{Kelvin.Value(273.15), Celsius.Value(0.0), false},
		{Kelvin.Value(283.15), Celsius.Value(10.0), false},
		{Kelvin.Value(0), Celsius.Value(-273.15), false},
		// This is invalid as cannot be colder than Absolute Zero so expect an error
		{Kelvin.Value(-1), Celsius.Value(-274.15), true},
		{Celsius.Value(-274.15), Kelvin.Value(-1), true},
		{Kelvin.Value(0), Fahrenheit.Value(-460.67), true},
		{Fahrenheit.Value(-460.67), Kelvin.Value(0), true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s %s", tt.from.Unit().Name(), tt.to.Unit().Name(), tt.from), func(t *testing.T) {

			got, err := tt.from.As(tt.to.Unit())
			if err != nil {
				if !tt.wantErr {
					t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if eq, err := tt.to.Equals(got); err != nil {
				if !tt.wantErr {
					t.Errorf("Value Equals error = %v", err)
				}
				return
			} else if !eq {
				t.Errorf("from %s got = %f, want %s", tt.from.String(), got.Float(), tt.to.String())
			}
		})
	}
}
