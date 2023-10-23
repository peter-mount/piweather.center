package service

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/station"
	configManager "github.com/peter-mount/piweather.center/util/config"
)

func init() {
	kernel.RegisterAPI((*Config)(nil), &config{})
}

// Config provides access to the Stations config
type Config interface {
	Accept(v station.Visitor) error
	Stations() *station.Stations
}

type config struct {
	ConfigManager configManager.Manager `kernel:"inject"`
	Config        *station.Stations
}

func (c *config) Stations() *station.Stations {
	return c.Config
}

func (c *config) Accept(v station.Visitor) error {
	return v.VisitStations(c.Config)
}

func (c *config) Start() error {
	{
		m := station.Stations(make(map[string]*station.Station))
		c.Config = &m
		if err := c.ConfigManager.ReadYaml("station.yaml", &c.Config); err != nil {
			return err
		}
	}

	if c.Config == nil || len(*c.Config) == 0 {
		return fmt.Errorf("no configuration provided")
	}

	// Once loaded ensure the structure is intact and ids are set up
	return c.Config.Init()
}
