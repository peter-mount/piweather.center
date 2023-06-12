package measurement

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"testing"
)

func TestHeatIndex(t *testing.T) {
	tests := []struct {
		temp        value.Value
		relHumidity value.Value
		want        value.Value
		wantErr     bool
	}{
		{Celsius.Value(18.7), RelativeHumidity.Value(56.0), Celsius.Value(18.087777777777774).AsGuard(Fahrenheit), false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.temp, tt.relHumidity), func(t *testing.T) {
			got, err := HeatIndex(tt.temp, tt.relHumidity)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("GetDewPoint() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if eq, err := tt.want.Equals(got); err != nil {
				if !tt.wantErr {
					t.Errorf("GetDewPoint() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if !eq {
				t.Errorf("GetDewPoint() got = %v, want %v", got, tt.want)
			}
		})
	}
}
