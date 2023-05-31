package measurement

import "github.com/peter-mount/piweather.center/weather/value"

func feelsLike(_ value.Time, args ...value.Value) (value.Value, error) {
	return FeelsLike(args[0], args[1], args[2])
}

// FeelsLike returns feels like temperature based on temperature, relative humidity and windSpeed.
//
// For temperatures below 50F this returns the wind chill. For those above 80F the Heat Index.
// For temperatures between 50 & 80F this returns the temperature.
func FeelsLike(temp, relHumidity, windSpeed value.Value) (value.Value, error) {
	temp, err := temp.As(Fahrenheit)
	if err != nil {
		return value.Value{}, err
	}

	relHumidity, err = relHumidity.As(RelativeHumidity)
	if err != nil {
		return value.Value{}, err
	}

	windSpeed, err = windSpeed.As(MetersPerSecond)
	if err != nil {
		return value.Value{}, err
	}

	switch {
	case temp.Float() <= 50.0 && windSpeed.Float() > 3.0:
		return WindChill(temp, windSpeed)

	case temp.Float() >= 80.0:
		return GetHeatIndex(temp, relHumidity)

	default:
		return temp, nil
	}
}
