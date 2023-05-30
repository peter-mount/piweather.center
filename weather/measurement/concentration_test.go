package measurement

import (
	"fmt"
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
				t.Errorf("from %s got = %.10f, want %s (%.10f)", tt.from.String(), got.Float(), tt.to.String(), tt.to.Float())
			}
		})
	}
}
