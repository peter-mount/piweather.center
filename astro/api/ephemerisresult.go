package api

import (
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/unit"
)

// EphemerisResult of a specific object at a specific time.
//
// Note: this replaces the old ephemeris package
type EphemerisResult interface {
	EphemerisCommon

	GetDistance() value.Value
	SetDistance(value.Value) EphemerisResult

	GetLightTime() value.Value
	SetLightTime(value.Value) EphemerisResult

	GetDistanceSun() value.Value
	SetDistanceSun(value.Value) EphemerisResult

	GetSemiDiameter() value.Value
	SetSemiDiameter(value.Value) EphemerisResult

	SetObliquity(unit.Angle) EphemerisResult

	GetEcliptic() *coord.Ecliptic
	SetEcliptic(lon, lat unit.Angle) EphemerisResult

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

type ephemerisResult struct {
	ephemerisCommon
	distance     value.Value       // distance from Earth
	lightTime    value.Value       // light time for distance
	distanceSun  value.Value       // distance from Sun
	semiDiameter value.Value       // SemiDiameter of object
	ecliptic     *coord.Ecliptic   // ecliptic coordinates
	equatorial   *coord.Equatorial // equatorial coordinates
	galactic     *coord.Galactic   // galactic coordinates
	horizontal   *coord.Horizontal // horizontal coordinates at observers location
}

func NewEphemerisResult(name string, t value.Time) EphemerisResult {
	r := newEphemerisResult(name, nil)
	r.ephemerisCommon.init(name, t.Time(), t.Location(), coord.NewObliquity(defaultObliquity))
	return r
}

func newEphemerisResult(name string, common *ephemerisCommon) *ephemerisResult {
	r := &ephemerisResult{
		ecliptic:   &coord.Ecliptic{},
		equatorial: &coord.Equatorial{},
		galactic:   &coord.Galactic{},
		horizontal: &coord.Horizontal{},
	}
	if common != nil {
		r.ephemerisCommon = *common
	}
	r.ephemerisCommon.name = name
	return r
}

func (r *ephemerisResult) SetObliquity(ε unit.Angle) EphemerisResult {
	r.ephemerisCommon.setObliquity(ε)
	return r
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

func (r *ephemerisResult) GetLightTime() value.Value {
	return r.lightTime
}

func (r *ephemerisResult) SetLightTime(v value.Value) EphemerisResult {
	r.lightTime = v
	return r
}

func (r *ephemerisResult) GetDistanceSun() value.Value {
	return r.distanceSun
}

func (r *ephemerisResult) SetDistanceSun(v value.Value) EphemerisResult {
	if err := measurement.Length.AssertValue(v); err != nil {
		panic(err)
	}
	r.distanceSun = v
	return r
}

func (r *ephemerisResult) GetSemiDiameter() value.Value {
	return r.semiDiameter
}

func (r *ephemerisResult) SetSemiDiameter(v value.Value) EphemerisResult {
	r.semiDiameter = v
	return r
}

func (r *ephemerisResult) GetEcliptic() *coord.Ecliptic {
	return r.ecliptic
}

func (r *ephemerisResult) SetEcliptic(lon, lat unit.Angle) EphemerisResult {
	r.ecliptic = &coord.Ecliptic{Lat: lat, Lon: lon}
	r.equatorial = r.eclToEq(r.ecliptic)
	r.horizontal = r.eqToHz(r.equatorial)
	r.galactic = r.eqToGal(r.equatorial)
	return r
}

func (r *ephemerisResult) GetEquatorial() *coord.Equatorial {
	return r.equatorial
}

func (r *ephemerisResult) SetEquatorial(ra unit.RA, dec unit.Angle) EphemerisResult {
	r.equatorial = &coord.Equatorial{RA: ra, Dec: dec}
	r.ecliptic = r.eqToEcl(r.equatorial)
	r.horizontal = r.eqToHz(r.equatorial)
	r.galactic = r.eqToGal(r.equatorial)
	return r
}

func (r *ephemerisResult) GetGalactic() *coord.Galactic {
	return r.galactic
}

func (r *ephemerisResult) SetGalactic(lat, lon unit.Angle) EphemerisResult {
	r.galactic = &coord.Galactic{Lat: lat, Lon: lon}
	r.equatorial = r.galToEq(r.galactic)
	r.ecliptic = r.eqToEcl(r.equatorial)
	r.horizontal = r.eqToHz(r.equatorial)
	return r
}

func (r *ephemerisResult) GetHorizontal() *coord.Horizontal {
	return r.horizontal
}

func (r *ephemerisResult) SetHorizontal(az, alt unit.Angle) EphemerisResult {
	r.horizontal = &coord.Horizontal{Az: az, Alt: alt}
	r.equatorial = r.hzToEq(r.horizontal)
	r.ecliptic = r.eqToEcl(r.equatorial)
	r.galactic = r.eqToGal(r.equatorial)
	return r
}

func (r *ephemerisResult) Value(t EphemerisOption) value.Value {
	switch t {
	case HorizonAltitude:
		return measurement.Degree.Value(r.horizontal.Alt.Deg())

	case HorizonAzimuth:
		return measurement.Degree.Value(r.horizontal.Az.Deg())

	case HorizonBearing:
		// Add 180° to azimuth to convert to geographic azimuth from due north
		f := r.horizontal.Az.Deg() + 180.0
		for f < 0.0 {
			f = f + 360
		}
		for f >= 360 {
			f = f - 360
		}
		return measurement.Degree.Value(f)

	case EquatorialRA:
		return measurement.RA.Value(r.equatorial.RA.Hour())

	case EquatorialDec:
		return measurement.Declination.Value(r.equatorial.Dec.Deg())

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

	case DistanceSun:
		return r.distanceSun

	case LightTime:
		return r.lightTime

	case SemiDiameter:
		return r.semiDiameter

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
	return eq.EclToEq(ecl, r.GetObliquity())
}

func (r *ephemerisResult) eqToEcl(eq *coord.Equatorial) *coord.Ecliptic {
	if eq == nil {
		return nil
	}

	ecl := coord.Ecliptic{}
	return ecl.EqToEcl(eq, r.GetObliquity())
}

func (r *ephemerisResult) galToEq(gal *coord.Galactic) *coord.Equatorial {
	if gal == nil {
		return nil
	}

	eq := coord.Equatorial{}
	return eq.GalToEq(gal)
}

func (r *ephemerisResult) eqToGal(eq *coord.Equatorial) *coord.Galactic {
	if eq == nil {
		return nil
	}

	gal := coord.Galactic{}
	return gal.EqToGal(eq)
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
