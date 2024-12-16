package skymap

import (
	"flag"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/astro/catalogue"
)

type Skymap struct {
	overview     *string            `kernel:"flag,skymap-overview,Generate overview map"`
	sphericalMap *string            `kernel:"flag,spherical,Generate spherical map"`
	Manager      *catalogue.Manager `kernel:"inject"`
	catalog      *catalogue.Catalog
}

func (s *Skymap) Start() error {
	log.Println(version.Version)

	var err error

	// Load the YBSC catalog
	s.catalog, err = s.Manager.YaleBrightStarCatalog()
	if err != nil {
		return err
	}

	done := false

	if *s.overview != "" {
		done = true
		if err := s.renderOverview(); err != nil {
			return err
		}
	}

	if *s.sphericalMap != "" {
		done = true
		if err := s.spherical(); err != nil {
			return err
		}
	}

	if !done {
		flag.PrintDefaults()
	}
	return nil
}
