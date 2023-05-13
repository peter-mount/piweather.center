package measurment

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

func GetDewPoint(temp value.Value, relHumidity value.Value) (value.Value, error) {
	return TemperatureRelativeHumidityCalculation(temp, relHumidity, Celsius,
		func(temp, relHumidity value.Value) (value.Value, error) {
			t0, rh := temp.Float(), relHumidity.Float()

			b, c := 17.368, 238.88
			if t0 <= 0 {
				b, c = 17.966, 247.15
			}
			pa := math.Log(rh / 100.0 * math.Exp(b*t0/(c+t0)))
			dp := c * pa / (b - pa)
			return Celsius.Value(dp), nil
		})
}

/*
func (hi DewPoint) String() string {
	return util.ToString(float64(hi))
}

func (hi DewPoint) Status() string {
	switch {
	case hi < 27.0:
		return value.Safe
	case hi < 32.0:
		return value.Caution
	case hi < 41.0:
		return value.Danger
	case hi >= 54.0:
		return value.ExtremeDanger
	default:
		return value.Safe
	}
}

func (hi DewPoint) IsValid() bool {
	return true
}

func GetDewPoint(temp Temperature, relHumidity Humidity) (HeatIndex, error) {
	if !relHumidity.IsValid() {
		return 0, fmt.Errorf("invalid humidity %f must be within 1..100", relHumidity)
	}
	if !temp.IsValid() {
		return 0, fmt.Errorf("invalid temperature %f", temp)
	}
	b, c := 17.368, 238.88
	if temp <= 0 {
		b, c = 17.966, 247.15
	}
	pa := float64(temp) / 100.0 * math.Exp(b*float64(temp)/(c+float64(temp)))
	dp := c * math.Log(pa) / (b - math.Log(pa))
	return HeatIndex(dp), nil
}
*/
