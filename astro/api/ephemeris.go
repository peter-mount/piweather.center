package api

import (
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/sidereal"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/unit"
	"time"
)

// EphemerisResult of a specific object at a specific time.
//
// Note: this replaces the old ephemeris package
type EphemerisResult interface {
	Name() string
	Time() time.Time
	JD() julian.Day
	SiderialTime() unit.Time

	GetDistance() value.Value
	SetDistance(value.Value) EphemerisResult

	GetEcliptic() *coord.Ecliptic
	SetEcliptic(lat, lon unit.Angle) EphemerisResult

	GetEquatorial() *coord.Equatorial
	SetEquatorial(ra unit.RA, dec unit.Angle) EphemerisResult

	GetGalactic() *coord.Galactic
	SetGalactic(lat, lon unit.Angle) EphemerisResult

	GetHorizontal() *coord.Horizontal
	SetHorizontal(az, alt unit.Angle) EphemerisResult

	// Value returns the appropriate entry as a Value.
	// If the entry is not present, or the calculated value is invalid then the returned value is invalid.
	Value(t EphemerisOption) value.Value

	// ToMetrics returns a slice of metrics from this result, dependent on the requested options.
	ToMetrics(prefix string, opts EphemerisOption) []api.Metric
}

var (
	obliquity = coord.NewObliquity(unit.AngleFromDeg(23.4392911))
)

type ephemerisResult struct {
	name       string            // name of object
	time       time.Time         // time of result
	jd         julian.Day        // date of result
	siderial   unit.Time         // siderial time
	loc        *globe.Coord      // Location of observer
	distance   value.Value       // distance
	ecliptic   *coord.Ecliptic   // ecliptic coordinates
	equatorial *coord.Equatorial // equatorial coordinates
	galactic   *coord.Galactic   // galactic coordinates
	horizontal *coord.Horizontal // horizontal coordinates at observers location
}

func NewEphemerisResult(name string, t value.Time) EphemerisResult {
	tm := t.Time()
	jd := julian.FromTime(tm)
	st := sidereal.FromJD(jd)
	return &ephemerisResult{
		name:       name,
		time:       tm,
		jd:         jd,
		siderial:   st,
		loc:        t.Location(),
		ecliptic:   &coord.Ecliptic{},
		equatorial: &coord.Equatorial{},
		galactic:   &coord.Galactic{},
		horizontal: &coord.Horizontal{},
	}
}

func (r *ephemerisResult) Name() string {
	return r.name
}

func (r *ephemerisResult) Time() time.Time {
	return r.time
}

func (r *ephemerisResult) JD() julian.Day {
	return r.jd
}

func (r *ephemerisResult) SiderialTime() unit.Time {
	return r.siderial
}

func (r *ephemerisResult) GetDistance() value.Value {
	return r.distance
}

func (r *ephemerisResult) SetDistance(v value.Value) EphemerisResult {
	if err := measurement.Length.AssertValue(v); err != nil {
		panic(err)
	}
	r.distance = v
	return r
}

func (r *ephemerisResult) GetEcliptic() *coord.Ecliptic {
	return r.ecliptic
}

func (r *ephemerisResult) SetEcliptic(lat, lon unit.Angle) EphemerisResult {
	r.ecliptic = &coord.Ecliptic{Lat: lat, Lon: lon}
	r.equatorial = r.equatorial.EclToEq(r.ecliptic, obliquity)
	r.horizontal = r.horizontal.EqToHz(r.equatorial, r.loc, r.siderial)
	r.galactic = r.galactic.EqToGal(r.equatorial)
	return r
}

func (r *ephemerisResult) GetEquatorial() *coord.Equatorial {
	return r.equatorial
}

func (r *ephemerisResult) SetEquatorial(ra unit.RA, dec unit.Angle) EphemerisResult {
	r.equatorial = &coord.Equatorial{RA: ra, Dec: dec}
	r.ecliptic = r.ecliptic.EqToEcl(r.equatorial, obliquity)
	r.horizontal = r.horizontal.EqToHz(r.equatorial, r.loc, r.siderial)
	r.galactic = r.galactic.EqToGal(r.equatorial)
	return r
}

func (r *ephemerisResult) GetGalactic() *coord.Galactic {
	return r.galactic
}

func (r *ephemerisResult) SetGalactic(lat, lon unit.Angle) EphemerisResult {
	r.galactic = &coord.Galactic{Lat: lat, Lon: lon}
	r.equatorial = r.equatorial.GalToEq(r.galactic)
	r.ecliptic = r.ecliptic.EqToEcl(r.equatorial, obliquity)
	r.horizontal = r.horizontal.EqToHz(r.equatorial, r.loc, r.siderial)
	return r
}

func (r *ephemerisResult) GetHorizontal() *coord.Horizontal {
	return r.horizontal
}

func (r *ephemerisResult) SetHorizontal(az, alt unit.Angle) EphemerisResult {
	r.horizontal = &coord.Horizontal{Az: az, Alt: alt}
	r.equatorial = r.equatorial.HzToEq(r.horizontal, *r.loc, r.siderial)
	r.ecliptic = r.ecliptic.EqToEcl(r.equatorial, obliquity)
	r.galactic = r.galactic.EqToGal(r.equatorial)
	return r
}

func (r *ephemerisResult) Value(t EphemerisOption) value.Value {
	switch t {
	case HorizonAltitude:
		return measurement.Degree.Value(r.horizontal.Alt.Deg())

	case HorizonAzimuth:
		return measurement.Degree.Value(r.horizontal.Az.Deg())

	case EquatorialRA:
		return measurement.RA.Value(r.equatorial.RA.Hour())

	case EquatorialDec:
		return measurement.Degree.Value(r.equatorial.Dec.Deg())

	case EclipticLatitude:
		return measurement.Degree.Value(r.ecliptic.Lat.Deg())

	case EclipticLongitude:
		return measurement.Degree.Value(r.ecliptic.Lon.Deg())

	case GalacticLatitude:
		return measurement.Degree.Value(r.galactic.Lat.Deg())

	case GalacticLongitude:
		return measurement.Degree.Value(r.galactic.Lon.Deg())

	case Distance:
		return r.distance
	default:
	}

	return value.Value{}
}

func (r *ephemerisResult) eqToHz(eq *coord.Equatorial) *coord.Horizontal {
	if eq == nil {
		return nil
	}

	hz := coord.Horizontal{}
	return hz.EqToHz(eq, r.loc, r.siderial)
}

func (r *ephemerisResult) hzToEq(hz *coord.Horizontal) *coord.Equatorial {
	if hz == nil {
		return nil
	}

	eq := coord.Equatorial{}
	return eq.HzToEq(hz, *r.loc, r.siderial)
}

func (r *ephemerisResult) eclToEq(ecl *coord.Ecliptic) *coord.Equatorial {
	if ecl == nil {
		return nil
	}

	eq := coord.Equatorial{}
	return eq.EclToEq(ecl, obliquity)
}

func (r *ephemerisResult) galToEq(gal *coord.Galactic) *coord.Equatorial {
	if gal == nil {
		return nil
	}

	eq := coord.Equatorial{}
	return eq.GalToEq(gal)
}

func (r *ephemerisResult) ToMetrics(prefix string, opts EphemerisOption) []api.Metric {
	var metrics []api.Metric
	t := r.Time()
	for _, opt := range opts.Split() {
		val := r.Value(opt)
		if val.IsValid() {
			metric := api.Metric{
				Metric:    prefix + "." + opt.MetricSuffix(),
				Time:      t,
				Unit:      val.Unit().ID(),
				Value:     val.Float(),
				Formatted: val.String(),
				Unix:      t.Unix(),
			}
			if metric.IsValid() {
				metrics = append(metrics, metric)
			}
		}
	}
	return metrics
}
