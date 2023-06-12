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
	// WindChill is valid for temps <=50F (10C) and wind speed >=3mph
	case value.LessThanEqual(temp.Float(), 50.0) && value.GreaterThanEqual(windSpeed.Float(), 1.34112):
		return WindChill(temp, windSpeed)

	// HeatIndex is valid for temps >=80F (26.6C)
	case value.GreaterThanEqual(temp.Float(), 80.0):
		return HeatIndex(temp, relHumidity)

	default:
		return temp, nil
	}
}
