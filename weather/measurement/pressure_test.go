package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"testing"
)

func Test_pressure_transforms(t *testing.T) {
	tests := []struct {
		from    value.Value
		to      value.Value
		wantErr bool
	}{
		// Test basic values
		{PressurePA.Value(101325), PressureHPA.Value(1013.25), false},
		{PressurePA.Value(101325), PressureMBar.Value(1013.25), false},
		//
		{PressureHPA.Value(1), PressurePA.Value(100), false},
		{PressureHPA.Value(1), PressureKPA.Value(0.1), false},
		{PressureMBar.Value(1), PressurePA.Value(100), false},
		{PressureMBar.Value(1), PressureBar.Value(0.001), false},
		{PressurePA.Value(100000), PressureBar.Value(1), false},
		// From a real issue, for some reason 29.973inHg came out as 101.5hPa and not 1015hPa
		{PressureInHg.Value(29.973), PressurePA.Value(101500.2267169408), false},
		{PressureInHg.Value(29.973), PressureHPA.Value(1015.002267169408), false},
		{PressureInHg.Value(29.973), PressureKPA.Value(101.5002267169408), false},
	}
	for _, tt := range tests {
		testConversion(t, tt.from, tt.to, tt.wantErr)
		testConversion(t, tt.to, tt.from, tt.wantErr)
	}
}
