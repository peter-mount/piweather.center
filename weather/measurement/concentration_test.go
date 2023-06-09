package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"testing"
)

func Test_concentration(t *testing.T) {
	tests := []struct {
		from    value.Value
		to      value.Value
		wantErr bool
	}{
		{PartsPerMillion.Value(1), MicrogramsPerCubicMeter.Value(1000), false},
		{MicrogramsPerCubicMeter.Value(1000), PartsPerMillion.Value(1), false},
		{PartsPerMillion.Value(500), MicrogramsPerCubicMeter.Value(500000), false},
		{MicrogramsPerCubicMeter.Value(500000), PartsPerMillion.Value(500), false},
	}
	for _, tt := range tests {
		testConversion(t, tt.from, tt.to, tt.wantErr)
		testConversion(t, tt.to, tt.from, tt.wantErr)
	}
}
