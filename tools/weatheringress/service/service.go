package service

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/tools/weatheringress/model"
	configManager "github.com/peter-mount/piweather.center/util/config"
)

func init() {
	kernel.RegisterAPI((*Config)(nil), &config{})
}

// Config provides access to the Stations config
type Config interface {
	Accept(v model.Visitor) error
	Stations() *model.Stations
}

type config struct {
	ConfigManager configManager.Manager `kernel:"inject"`
	Config        *model.Stations
}

func (c *config) Stations() *model.Stations {
	return c.Config
}

func (c *config) Accept(v model.Visitor) error {
	return v.VisitStations(c.Config)
}

func (c *config) Start() error {
	{
		m := model.Stations(make(map[string]*model.Station))
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
