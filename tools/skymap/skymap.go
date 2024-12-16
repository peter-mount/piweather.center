package skymap

import (
	"flag"
	"github.com/peter-mount/go-build/application"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/astro/catalogue"
	io2 "github.com/peter-mount/piweather.center/util/io"
)

type Skymap struct {
	overview     *string `kernel:"flag,skymap-overview,Generate overview map"`
	sphericalMap *string `kernel:"flag,spherical,Generate spherical map"`
	catalog      *catalogue.Catalog
}

func (s *Skymap) Start() error {
	log.Println(version.Version)

	// Load the YBSC catalog
	s.catalog = &catalogue.Catalog{}
	if err := io2.NewReader(s.catalog.Read).
		Decompress().
		Open(application.FileName(application.STATIC, "bsc5.bin")); err != nil {
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
