package measurment

import (
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	HeatIndex = value.NewBoundedUnit("HeatIndex", "%", value.Dp0, 0, 100)
}

var (
	HeatIndex value.Unit
)

func IsHeatIndex(v value.Value) bool {
	return v.Unit() == HeatIndex
}

const (
	hic1 = -42.379
	hic2 = 2.04901523
	hic3 = 10.14333127
	hic4 = -0.22475541
	hic5 = -6.83783e-3
	hic6 = -5.481717e-2
	hic7 = 1.22874e-3
	hic8 = 8.5282e-4
	hic9 = -1.99e-6
)

func GetHeatIndex(temp value.Value, relHumidity value.Value) (value.Value, error) {
	return TemperatureRelativeHumidityCalculation(temp, relHumidity, Fahrenheit,
		func(temp, relHumidity value.Value) (value.Value, error) {
			T, RH := temp.Float(), relHumidity.Float()

			// try simplified formula first (used for HI < 80)
			HI := 0.5 * (T + 61.0 + (T-68.0)*1.2 + RH*0.094)

			if HI >= 80.0 {
				// use Rothfusz regression
				T2, RH2 := T*T, RH*RH
				HI = util.Fsum(
					hic1,
					hic2*T,
					hic3*RH,
					hic4*T*RH,
					hic5*T2,
					hic6*RH2,
					hic7*T2*RH,
					hic8*T*RH2,
					hic9*T2*RH2,
				)
			}
			return Fahrenheit.Value(HI), nil
		})
}

/*
func (hi HeatIndex) String() string {
	return util.ToString(float64(hi))
}

func (hi HeatIndex) Status() string {
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

func (hi HeatIndex) IsValid() bool {
	return true
}

func GetHeatIndex(temp Temperature, relHumidity Humidity) (HeatIndex, error) {
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
