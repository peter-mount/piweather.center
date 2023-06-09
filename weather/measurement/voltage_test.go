package measurement

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"testing"
)

func Test_voltage_transforms(t *testing.T) {
	tests := []struct {
		from    value.Value
		to      value.Value
		wantErr bool
	}{
		{Volt.Value(1), MilliVolt.Value(1000.0), false},
		{Volt.Value(1), MicroVolt.Value(1000000.0), false},
		// dBV reference
		{Volt.Value(1), DecibelVolt.Value(0), false},
		// V -> dBV
		{MilliVolt.Value(100), DecibelVolt.Value(-20), false},
		{MilliVolt.Value(1), DecibelVolt.Value(-60), false},
		{Volt.Value(10), DecibelVolt.Value(20), false},
		{Volt.Value(100), DecibelVolt.Value(40), false},
	}

	for _, tt := range tests {
		testConversion(t, tt.from, tt.to, tt.wantErr)
		testConversion(t, tt.to, tt.from, tt.wantErr)
	}
}

// Used by multiple tests, test we can convert between two values
func testConversion(t *testing.T, from, to value.Value, wantErr bool) {
	t.Run(fmt.Sprintf("%s %s %s", from.Unit().Name(), to.Unit().Name(), from), func(t *testing.T) {

		got, err := from.As(to.Unit())
		if err != nil {
			if !wantErr {
				t.Errorf("error = %v, wantErr %v", err, wantErr)
			}
			return
		}

		if eq, err := to.Equals(got); err != nil {
			if !wantErr {
				t.Errorf("Value Equals error = %v", err)
			}
			return
		} else if !eq {
			t.Errorf("from %s got = %f, want %s", from.String(), got.Float(), to.String())
		}
	})
}
