package calculator

import (
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/planetposition"
	"github.com/soniakeys/unit"
	"os"
	"path"
	"path/filepath"
	"sync"
)

func init() {
	kernel.RegisterAPI((*Calculator)(nil), &calculator{})
}

// Calculator provides a series of value.Calculator's for various
// astronomical entities
type Calculator interface {
	// Planet returns the planetposition.V87Planet element set for each of
	// the 8 major Planets, loading from disk as necessary
	Planet(i int) (*planetposition.V87Planet, error)

	CalculateMoon(t value.Time) (api.EphemerisResult, error)

	CalculatePlanet(planetId station.EphemerisTargetType, t value.Time) (api.EphemerisResult, error)

	// CalculateSun returns an EphemerisResult for a specified time.
	// Note: the RA/Dec epoch be for the date, not J2000
	CalculateSun(t value.Time) (api.EphemerisResult, error)

	// SolarAltitudeCalculator is a calculator that calculates
	// the sun's altitude at a station at a specific time
	// use CalculateSun instead
	// Deprecated
	SolarAltitudeCalculator() value.Calculator

	// SolarAzimuthCalculator is a calculator that calculates
	// the sun's azimuth at a station at a specific time
	//	use CalculateSun instead
	// Deprecated
	SolarAzimuthCalculator() value.Calculator

	// SolarHZ calculate the Azimuth and Altitude of the sun
	//	use CalculateSun instead
	// Deprecated
	SolarHZ(t value.Time) (unit.Angle, unit.Angle, error)
}

type calculator struct {
	mutex           sync.Mutex
	rootDir         string
	planetPositions []*planetposition.V87Planet
}

func (c *calculator) Start() error {
	// There are 8 major Planets ;-)
	c.planetPositions = make([]*planetposition.V87Planet, 8)

	// Path to script directory for data lookup
	c.rootDir = path.Join(filepath.Dir(os.Args[0]), "../script")

	// Register calculators
	value.NewCalculator("SolarAltitude", c.SolarAltitudeCalculator())
	value.NewCalculator("SolarAzimuth", c.SolarAzimuthCalculator())

	return nil
}
