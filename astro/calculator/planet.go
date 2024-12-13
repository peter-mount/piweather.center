package calculator

import (
	"fmt"
	"github.com/peter-mount/go-build/application"
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/apparent"
	"github.com/soniakeys/meeus/v3/base"
	"github.com/soniakeys/meeus/v3/nutation"
	"github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/meeus/v3/semidiameter"
	"github.com/soniakeys/unit"
	"math"
)

// Planet returns the V87Planet by ID.
func (c *calculator) Planet(i int) (*planetposition.V87Planet, error) {
	if planet := c.getPlanet(i); planet != nil {
		return planet, nil
	}

	planet, err := planetposition.LoadPlanetPath(i, application.FileName(application.STATIC, "vsop87b"))
	if err != nil {
		return nil, err
	}
	c.setPlanet(i, planet)
	return planet, nil
}

func (c *calculator) getPlanet(i int) *planetposition.V87Planet {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.planetPositions[i]
}

func (c *calculator) setPlanet(i int, planet *planetposition.V87Planet) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.planetPositions[i] = planet
}

func (c *calculator) CalculatePlanet(planetId station.EphemerisTargetType, t value.Time) (api.EphemerisResult, error) {
	if !planetId.IsPlanet() {
		return nil, fmt.Errorf("planet id %d is not a planet", int(planetId))
	}
	p, err := c.Planet(int(planetId))
	if err != nil {
		return nil, err
	}

	earth, err := c.Planet(planetposition.Earth)
	if err != nil {
		return nil, err
	}

	jde := julian.FromTime(t.Time()).Float()

	// From this point we are the same as elliptic.Position, but we need more info than just ra,dec so
	// we expand it here
	L0, B0, R0 := earth.Position(jde)
	L, B, R := p.Position(jde)
	sB0, cB0 := B0.Sincos()
	sL0, cL0 := L0.Sincos()
	sB, cB := B.Sincos()
	sL, cL := L.Sincos()
	x := R*cB*cL - R0*cB0*cL0
	y := R*cB*sL - R0*cB0*sL0
	z := R*sB - R0*sB0
	{
		Δ := math.Sqrt(x*x + y*y + z*z) // (33.4) p. 224
		τ := base.LightTime(Δ)
		// repeating with jde-τ
		L, B, R = p.Position(jde - τ)
		sB, cB = B.Sincos()
		sL, cL = L.Sincos()
		x = R*cB*cL - R0*cB0*cL0
		y = R*cB*sL - R0*cB0*sL0
		z = R*sB - R0*sB0
	}
	λ := unit.Angle(math.Atan2(y, x))                // (33.1) p. 223
	β := unit.Angle(math.Atan2(z, math.Hypot(x, y))) // (33.2) p. 223
	Δλ, Δβ := apparent.EclipticAberration(λ, β, jde)
	λ, β = planetposition.ToFK5(λ+Δλ, β+Δβ, jde)
	Δψ, Δε := nutation.Nutation(jde)
	λ += Δψ

	// Note original code did EclToEq but we can do that in the result
	// so do the obliquity slightly differently
	ε := nutation.MeanObliquity(jde) + Δε

	// Our additions are here:

	// Distance from earth (32.4 P210 in my first edition, 33.4 p224 in second edition)
	Δearth := math.Sqrt((x * x) + (y * y) + (z * z))

	// Light Time from earth in days
	τEarth := base.LightTime(Δearth)

	// Finally return the results
	return api.NewEphemerisResult(planetId.String(), t).
			SetObliquity(ε).
			SetEcliptic(λ, β).
			SetDistance(measurement.AU.Value(Δearth)).
			// Apparent semidiameter from earth
			SetSemiDiameter(measurement.AngleRoundDown(measurement.Degree.Value(SemiDiameter(planetId, Δearth).Deg()))).
			// Light Time is in Days but DurationRoundDown makes it more useful
			SetLightTime(measurement.DurationRoundDown(measurement.DurationDay.Value(τEarth))).
			// R is the planets heliocentric distance
			SetDistanceSun(measurement.AU.Value(R)),
		nil
}

// SemiDiameter returns semidiameter at specified distance.
//
// Δ must be observer-body distance in AU.
func SemiDiameter(planetId station.EphemerisTargetType, Δ float64) unit.Angle {
	s0 := getSemiDiameter(planetId)
	return semidiameter.Semidiameter(s0, Δ)
}

func getSemiDiameter(planetId station.EphemerisTargetType) unit.Angle {
	switch planetId {
	case station.EphemerisTargetMercury:
		return semidiameter.Mercury
	case station.EphemerisTargetVenus:
		return semidiameter.VenusCloud
	case station.EphemerisTargetMars:
		return semidiameter.Mars
	case station.EphemerisTargetJupiter:
		return semidiameter.JupiterEquatorial
	case station.EphemerisTargetSaturn:
		return semidiameter.SaturnEquatorial
	case station.EphemerisTargetUranus:
		return semidiameter.Uranus
	case station.EphemerisTargetNeptune:
		return semidiameter.Neptune
	case station.EphemerisTargetMoon:
		return semidiameter.Moon
	case station.EphemerisTargetSun:
		return semidiameter.Sun
	default:
		// Return something but make it tiny so it's a point object
		return 0.000001
	}
}
