package forecast

import (
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
	"time"
)

var (
	zHpaFactor = []float64{
		6, 5, 5, 2, -0.5, -2, -5, -8.5, -12, -10, -6, -4.5, -3, -0.5, 1.5, 3,
	}
)

const (
	// Range of local weather window, 950...1050 hPa for the UK
	// TODO make this configurable
	hpaMin   = 950.0
	hpaMax   = 1050.0
	hpaRange = hpaMax - hpaMin
	// the number of steps within hpaRange
	hpaSteps = 22
	// factor to convert pressure to index
	hpaConstant = hpaRange / hpaSteps
)

// CalculateZambrettiForecast will return the ZambrettiSeverity representing a weather forecast.
//
// t is the Time of the forecast.
//
// h is the Hemisphere the weather station is located within.
//
// pressure0 is the previous pressure at mean sea level, ideally 3 hours before t.
//
// pressure is the  pressure at mean sea level at t.
//
// windDirection is the wind direction at t.
func CalculateZambrettiForecast(t time.Time, h Hemisphere, pressure0, pressure value.Value, windDirection value.Value) (ZambrettiSeverity, error) {
	trend, err := GetTrend(pressure0, pressure)
	if err != nil {
		return 0, err
	}

	wind, err := util.WindCompassDirectionDegrees(windDirection)
	if err != nil {
		return 0, err
	}
	if h == SouthernHemisphere {
		// we need to rotate the compass by 180 degrees so S is 0 and N is 8
		wind = wind.Add(util.WindS)
	}

	pf := zHpaFactor[wind.Int()]
	if isSummer(t, h) {
		switch trend {
		case TrendRising:
			pf += 7
		case TrendFalling:
			pf -= 7
		default:
		}
	}

	p0, err := pressure.As(measurement.PressureHPA)
	if err != nil {
		return 0, err
	}
	p := value.Enforce(p0.Float()+(pf/100.0*hpaRange), hpaMin, hpaMax)

	z := int(value.Enforce(math.Floor((p-hpaMin)/hpaConstant), 0, hpaSteps-1))
	switch trend {
	case TrendRising:
		return riseOptions[z], nil
	case TrendFalling:
		return fallOptions[z], nil
	default:
		return steadyOptions[z], nil
	}
}

func isSummer(t time.Time, h Hemisphere) bool {
	// true if Summer which here is April to September inclusive for the northern hemisphere
	// and October to March for the southern hemisphere
	month := t.Month()
	summer := month >= time.April && month <= time.September
	if h == SouthernHemisphere {
		summer = !summer
	}
	return summer
}
