package calculator

import (
	"github.com/peter-mount/go-build/application"
	"github.com/soniakeys/meeus/v3/planetposition"
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
