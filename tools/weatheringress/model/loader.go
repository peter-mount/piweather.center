package model

import (
	"errors"
	"github.com/peter-mount/go-kernel/v2"
	"github.com/peter-mount/piweather.center/io"
)

func init() {
	kernel.RegisterAPI((*Loader)(nil), &loader{})
}

// Loader provides access to the Stations loader
type Loader interface {
	Accept(v Visitor) error
	Stations() *Stations
}

type loader struct {
	ConfigFile *string `kernel:"flag,station,Station loader to use"`
	Config     *Stations
}

func (c *loader) Stations() *Stations {
	return c.Config
}

func (c *loader) Accept(v Visitor) error {
	return v.VisitStations(c.Config)
}

func (c *loader) Start() error {
	if c.ConfigFile == nil || *c.ConfigFile == "" {
		return errors.New("-station required")
	}

	m := Stations(make(map[string]*Station))
	c.Config = &m

	err := io.NewReader().
		Yaml(c.Config).
		Open(*c.ConfigFile)

	// Once loaded ensure the structure is intact and ids are set up
	if err == nil {
		err = c.Config.Init()
	}

	return err
}
