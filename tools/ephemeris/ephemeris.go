package ephemeris

import (
	"context"
	"flag"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/ephemeris"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/util"
)

type Ephemeris struct {
	worker    task.Queue           `kernel:"worker"`
	name      *string              `kernel:"flag,name,Name of ephemeris"`
	site      *string              `kernel:"flag,site,simple location definition"`
	step      *float64             `kernel:"flag,s,Step size in days,1.0"`
	stdout    *bool                `kernel:"flag,d,Dump xml to stdout"`
	riseSet   *bool                `kernel:"flag,rs,Include Rise/Transit/Set times in output"`
	sun       *bool                `kernel:"flag,sun,Include the Sun"`
	ephemeris *ephemeris.Ephemeris // Ephemeris to generate
}

func (e *Ephemeris) Start() error {
	e.ephemeris = &ephemeris.Ephemeris{
		Name: *e.name,
	}

	// Take a set of dates from the command line and add to the Range
	for _, arg := range flag.Args() {
		d, err := julian.Parse(arg)
		if err != nil {
			return err
		}
		e.ephemeris = e.ephemeris.Include(d)
	}

	// If the range is invalid, e.g. no dates provided,
	// use the current system time
	if !e.ephemeris.Range.Valid() {
		e.ephemeris.Include(julian.StartOfToday())
	}

	if *e.site != "" {
		ll, err := coord.Parse(*e.site)
		if err != nil {
			return err
		}
		e.ephemeris.Meta.LatLong = ll
	}

	if *e.sun {
		e.worker.AddTask(e.calcSun)
	}

	if *e.riseSet {
		e.worker.AddTask(e.calcRiseSet)
	}
	if *e.stdout {
		e.worker.AddTask(e.print)
	}

	return nil
}

func (e *Ephemeris) print(_ context.Context) error {
	fmt.Println(util.String(e.ephemeris))
	return nil
}

func (e *Ephemeris) calcRiseSet(_ context.Context) error {
	return e.ephemeris.CalculateRiseSetTimes()
}

func (e *Ephemeris) calcSun(_ context.Context) error {
	return e.ephemeris.Generate(*e.step, &ephemeris.SunProvider{})
}
