package forecast

import "github.com/peter-mount/piweather.center/weather/value"

func init() {
	Zambretti = value.NewUnit("zambretti", "Zambretti Forecast", "Zambretti Severity", 0,
		func(f float64) string {
			return ZambrettiSeverity(f).String()
		})
	ForecastGroup = value.NewGroup("Forecast", Zambretti)
}

var (
	Zambretti     *value.Unit
	ForecastGroup *value.Group
)

var (
	forecast = []string{
		"Settled fine",
		"Fine weather",
		"Becoming fine",
		"Fine, becoming less settled",
		"Fine, possible showers",
		"Fairly fine, improving",
		"Fairly fine, possible showers early",
		"Fairly fine, showery later",
		"Showery early, improving",
		"Changeable, mending",
		"Fairly fine, showers likely",
		"Rather unsettled clearing later",
		"Unsettled, probably improving",
		"Showery, bright intervals",
		"Showery, becoming less settled",
		"Changeable, some rain",
		"Unsettled, short fine intervals",
		"Unsettled, rain later",
		"Unsettled, some rain",
		"Mostly very unsettled",
		"Occasional rain, worsening",
		"Rain at times, very unsettled",
		"Rain at frequent intervals",
		"Rain, very unsettled",
		"Stormy, may improve",
		"Stormy, much rain",
	}

	// equivalents of Zambretti 'dial window' letters A - Z
	riseOptions   = []ZambrettiSeverity{25, 25, 25, 24, 24, 19, 16, 12, 11, 9, 8, 6, 5, 2, 1, 1, 0, 0, 0, 0, 0, 0}
	steadyOptions = []ZambrettiSeverity{25, 25, 25, 25, 25, 25, 23, 23, 22, 18, 15, 13, 10, 4, 1, 1, 0, 0, 0, 0, 0, 0}
	fallOptions   = []ZambrettiSeverity{25, 25, 25, 25, 25, 25, 25, 25, 23, 23, 21, 20, 17, 14, 7, 3, 1, 1, 1, 0, 0, 0}
)

type ZambrettiSeverity int

// String returns the text for a ZambrettiSeverity
func (z ZambrettiSeverity) String() string {
	return forecast[z.Index()]
}

// Index returns the ZambrettiSeverity as a value between 0..25
func (z ZambrettiSeverity) Index() int {
	if z < 0 {
		return 0
	}
	if z > 25 {
		return 25
	}
	return int(z)
}
