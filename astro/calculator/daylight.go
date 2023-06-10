package calculator

import (
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/sun"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/unit"
	"math"
	"time"
)

const (
	// SolarStandardAltitude is the geometric altitude of the center of the
	// solar body at the time of apparent rising or setting.
	SolarStandardAltitude = -50.0 / 60.0
	// CivilTwilight is defined as when the geometric center of the Sun is between 6° below the horizon and the horizon itself.
	// This is not the same as the legal definition in some locations, but the generic global definition.
	CivilTwilight = -6.0

	// NauticalTwilight is defined as when the geometric center of the Sun is between 12° and 6° below the horizon.
	NauticalTwilight = -12.0

	// AstronomicalTwilight is defined as when the geometric center of the Sun is between 18° and 12° below the horizon.
	AstronomicalTwilight = -18.0
)

// SolarEphemeris represents details about a specific day at a location
// on the Earth's surface.
//
// Note: At certain times of the year and at higher latitudes some of these
// entries will not be valid because they cannot occur.
//
// For example, in the UK during the summer months AstronomicalTwilight never ends due to the Sun
// never going below -18° altitude. In that instance AstronomicalDawn and AstronomicalDusk are undefined.
//
// The higher the latitude (north or south) then the other entries become undefined, to the point
// near the poles where none of them are defined - e.g. the sun never rises during local winter
// or never sets during local summer.
//
// DayLength is the amount of time the Sun is above the horizon.
// This can be anywhere within 0 and 24 hours long depending on latitude.
type SolarEphemeris struct {
	// JD Julian Day of local Midnight at the start of this day
	JD julian.Day
	// Location on earth where this ephemeris applies
	Location *globe.Coord
	// AstronomicalDawn when the geometric center of the Sun enters morning AstronomicalTwilight.
	AstronomicalDawn AltAz
	// NauticalDawn when the geometric center of the Sun enters morning NauticalTwilight.
	NauticalDawn AltAz
	// CivilDawn when the geometric center of the Sun enters morning CivilTwilight.
	CivilDawn AltAz
	// SunRise when the limb of the sun rises above the horizon.
	SunRise AltAz
	// SunRise when the limb of the sun sets below the horizon.
	SunSet AltAz
	// CivilDusk when the geometric center of the Sun enters evening CivilTwilight.
	CivilDusk AltAz
	// NauticalDusk when the geometric center of the Sun enters evening NauticalTwilight.
	NauticalDusk AltAz
	// AstronomicalDusk when the geometric center of the Sun enters evening AstronomicalTwilight.
	AstronomicalDusk AltAz
	// UpperTransit when the sun is at its highest in the sky
	UpperTransit AltAz
	// LowerTransit when the sun is at its lowest in the sky - usually below the horizon except
	// at extreme latitudes at certain times of the year.
	LowerTransit AltAz
	// DayLength is the amount of time the sun is above the horizon.
	// At extreme latitudes this can be either 0 or 24 hours.
	DayLength time.Duration
}

func (e SolarEphemeris) Latitude() unit.Angle {
	return e.Location.Lat
}

func (e SolarEphemeris) Longitude() unit.Angle {
	return e.Location.Lon
}

func (c *calculator) SolarEphemeris(t0 value.Time) (SolarEphemeris, error) {
	// Copy time and set to start of local day
	// Note we start 1 minute before midnight so that the first value has a
	// previous value. Hence, the duration is also 24h 1m long
	t := t0.Clone()
	midnight := time2.LocalMidnight(t.Time())
	t.SetTime(midnight.Add(-time.Minute))

	r := SolarEphemeris{
		JD:       julian.FromTime(midnight),
		Location: t.Location(),
	}

	earth, err := c.Planet(planetposition.Earth)
	if err != nil {
		return SolarEphemeris{}, err
	}

	var previous AltAz
	started := false
	// For each minute of the day...
	err = t.ForEach(time.Minute, (24*time.Hour)+time.Minute, func(t value.Time) error {
		A, h := sun.ApparentHzVSOP87(julian.FromTime(t.Time()), r.Latitude(), r.Longitude(), earth)

		hD := h.Deg()
		curr := AltAz{
			Time:     t.Time(),
			Altitude: measurement.Degree.Value(hD),
			// A is measured westward from south so convert to a bearing
			Azimuth: measurement.Degree.Value(A.Deg() + 180),
		}

		//fmt.Println(t.Time(), h0, hDeg, value.GreaterThan(hDeg, h0), value.LessThan(hDeg, h0))

		if started {
			switch {
			// Sun is rising in the sky
			case value.GreaterThan(curr.Altitude.Float(), previous.Altitude.Float()):
				r.AstronomicalDawn = SolarEphemerisTime(r.AstronomicalDawn, curr, hD, AstronomicalTwilight)
				r.NauticalDawn = SolarEphemerisTime(r.NauticalDawn, curr, hD, NauticalTwilight)
				r.CivilDawn = SolarEphemerisTime(r.CivilDawn, curr, hD, CivilTwilight)
				r.SunRise = SolarEphemerisTime(r.SunRise, curr, hD, SolarStandardAltitude)

			// Sun is setting in the sky
			case value.LessThan(curr.Altitude.Float(), previous.Altitude.Float()):
				r.SunSet = SolarEphemerisTime(r.SunSet, curr, SolarStandardAltitude, hD)
				r.CivilDusk = SolarEphemerisTime(r.CivilDusk, curr, CivilTwilight, hD)
				r.NauticalDusk = SolarEphemerisTime(r.NauticalDusk, curr, NauticalTwilight, hD)
				r.AstronomicalDusk = SolarEphemerisTime(r.AstronomicalDusk, curr, AstronomicalTwilight, hD)

			// Sun is not moving...
			default:
				// do nothing
			}

			// Handle UpperTransit as the point with the highest altitude
			if value.GreaterThan(hD, r.UpperTransit.Altitude.Float()) {
				r.UpperTransit = curr
			}

			// Handle LowerTransit - always defined byu
			if value.LessThanEqual(hD, r.LowerTransit.Altitude.Float()) {
				r.LowerTransit = curr
			}
		} else {
			started = true

			// For the first calculation set LowerTransit as an initial point. As we use this for the minimum altitude we need an upper bounds
			r.LowerTransit = curr
		}

		previous = curr
		return nil
	})

	// Calculate the day length based on SunRise and SunSet
	switch {
	// Sun rises & sets in the same day
	case r.SunRise.IsValid() && r.SunSet.IsValid():
		r.DayLength = r.SunSet.Time.Sub(r.SunRise.Time)

	// Sun rises but doesn't set in this day - extreme latitudes only
	// So day length is from sun rise to the following local midnight
	case r.SunRise.IsValid():
		r.DayLength = time2.LocalMidnight(t0.Time()).Add(24 * time.Hour).Sub(r.SunRise.Time)

	// Sun sets but didn't rise in this day - extreme latitudes only
	// Day length is from local midnight to when it set
	case r.SunSet.IsValid():
		r.DayLength = r.SunSet.Time.Sub(midnight)

	// The sun is always above the horizon - extreme latitudes only
	case r.UpperTransit.IsValid() && r.UpperTransit.IsVisible():
		r.DayLength = 24 * time.Hour

	// Sun did not rise at all - extreme latitudes only
	default:
		r.DayLength = 0
	}

	return r, err
}

// SolarEphemerisTime will, if a>b and |a-b| is less than 1°,
// return t1 if t0 is zero (e.g. unset).
// This allows us to record the moment an object
func SolarEphemerisTime(t0, t1 AltAz, a, b float64) AltAz {
	if !t0.IsValid() && value.GreaterThan(a, b) && math.Abs(a-b) <= 1 {
		return t1
	}
	return t0
}
