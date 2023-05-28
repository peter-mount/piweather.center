package measurment

import (
	"github.com/peter-mount/piweather.center/astro/util"
	"github.com/peter-mount/piweather.center/weather/value"
)

func init() {
	value.NewCalculator("heatIndex", TemperatureRelativeHumidityCalculator(GetHeatIndex))
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

// GetHeatIndex returns the HeatIndex based on Temperature and RelativeHumidity.
func GetHeatIndex(temp value.Value, relHumidity value.Value) (value.Value, error) {
	return TemperatureRelativeHumidityCalculation(temp, relHumidity, Fahrenheit, getHeatIndex)
}

// temp must be Fahrenheit
func getHeatIndex(temp value.Value, relHumidity value.Value) (value.Value, error) {
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
}
