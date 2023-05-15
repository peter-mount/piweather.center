package measurment

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"testing"
)

func Test_speed_transforms(t *testing.T) {
	tests := []struct {
		from    value.Value
		to      value.Value
		wantErr bool
	}{
		// Base transforms where one side is always MetersPerSecond
		{MetersPerSecond.Value(1), KilometersPerHour.Value(3.6), false},
		{KilometersPerHour.Value(3.6), MetersPerSecond.Value(1), false},
		{Knots.Value(1), MetersPerSecond.Value(0.5144444444), false},

		// Generated transforms which go via MetersPerSecond
		{MilesPerHour.Value(1), KilometersPerHour.Value(1.609344), false},
		{KilometersPerHour.Value(1.609344), MilesPerHour.Value(1), false},
		{Knots.Value(1), MilesPerHour.Value(1.1507794480), false},
		{MilesPerHour.Value(1.1507794480), Knots.Value(1), false},
		{Knots.Value(1), KilometersPerHour.Value(1.852), false},
		{KilometersPerHour.Value(1.852), Knots.Value(1), false},
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
