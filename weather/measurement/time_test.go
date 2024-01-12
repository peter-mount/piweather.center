package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"testing"
)

func Test_Time(t *testing.T) {

	tests := []value.Value{
		UnixTime.Value(1705055520),
		ModifiedJD.Value(60321.43889),
		JulianDate.Value(2460321.93889),
		RataDie.Value(738897.43889),
		//MarsSolDate.Value(53333.72804),
	}

	for _, y := range tests {
		for _, x := range tests {
			if y != x {
				t.Run(
					y.Unit().ID()+" "+x.Unit().ID(),
					func(t *testing.T) {
						got, err := y.As(x.Unit())
						if err != nil {
							t.Error(err)
						} else {
							// Valid as got and x are the same unit here
							if got.PlainString() != x.PlainString() {
								t.Errorf("%q -> %q got %q", y.PlainString(), x.PlainString(), got.PlainString())
							}
						}
					})

			}
		}
	}
}
