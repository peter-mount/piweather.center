package measurement

import "github.com/peter-mount/piweather.center/weather/value"

var (
	// TimeGroup groups together various time values
	TimeGroup *value.Group
	// ModifiedJD number of Julian Days since 1858-11-17T00:00Z
	ModifiedJD *value.Unit
	// JulianDate number of days since -4713-1-1T12:00Z
	JulianDate *value.Unit
	// MarsSolDate is the number of Martian days (Sols) since 1873-12-29T12:00Z
	MarsSolDate *value.Unit
	// RataDie is the number of days since 1-1-1T00:00Z
	RataDie *value.Unit
	// UnixTime in seconds since 1970-1-1T00:00Z
	UnixTime *value.Unit
)

const (
	unixJDEpoch        = 2440587.5
	mjdEpoch           = 2400000.5
	rataDieEpoch       = 1721424.5
	martianEpoch       = 2451549.5
	martianK           = -4.5
	martianDayLen      = 1.027491252
	martianEpochOffset = 44796.0 - 0.00096
)

func init() {
	ModifiedJD = value.NewUnit("ModifiedJD", "Modified Julian Date", "", 5, nil)
	JulianDate = value.NewUnit("JulianDate", "Julian Date", "", 5, nil)
	MarsSolDate = value.NewUnit("MarsSolDate", "Martian Sol Date", "M.S.D.", 5, nil)
	RataDie = value.NewUnit("RataDie", "Rata Die", "R.D.", 5, nil)
	UnixTime = value.NewUnit("UnixTime", "Unix Time", "", 0, nil)

	// Unix is special where we can transform with an Integer
	value.NewBasicBiTransform(value.Integer, UnixTime, 1)

	epocTransform(JulianDate, UnixTime, unixJDEpoch, 86400)
	epocTransform(JulianDate, ModifiedJD, mjdEpoch, 1)
	epocTransform(JulianDate, RataDie, rataDieEpoch, 1)

	// FIXME this should be working but it gives the wrong dates
	// ref: https://marsclock.com/
	//		https://en.wikipedia.org/wiki/Timekeeping_on_Mars#Mars_Sol_Date
	//		https://en.wikipedia.org/wiki/Julian_day#Variants
	value.NewBiTransform(JulianDate, MarsSolDate,
		func(f float64) (float64, error) {
			return ((f - martianEpoch + martianK) / martianDayLen) + martianEpochOffset, nil
		},
		func(f float64) (float64, error) {
			return ((f - martianEpochOffset) * martianDayLen) + martianEpoch - martianK, nil
		})

	TimeGroup = value.NewGroup("Time", JulianDate, UnixTime, MarsSolDate, ModifiedJD, RataDie)
}

func epocTransform(src, dst *value.Unit, epoch, factor float64) {
	// factor=1 then just use the epoch.
	if factor == 1 || factor == 0 {
		value.NewBiTransform(src, dst,
			func(f float64) (float64, error) {
				return f - epoch, nil
			}, func(f float64) (float64, error) {
				return f + epoch, nil
			})
	} else {
		value.NewBiTransform(src, dst,
			func(f float64) (float64, error) {
				return (f - epoch) * factor, nil
			}, func(f float64) (float64, error) {
				return (f / factor) + epoch, nil
			})
	}
}
