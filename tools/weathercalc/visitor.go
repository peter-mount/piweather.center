package weathercalc

import (
	"github.com/peter-mount/go-script/errors"
	station2 "github.com/peter-mount/piweather.center/config/station"
)

type calcState struct {
	calc      *Calculator
	station   *station2.Station
	c         *station2.Calculation
	ephemeris *station2.Ephemeris
}

func visitCalculation(v station2.Visitor[*calcState], c *station2.Calculation) error {
	v.Get().c = c
	return nil
}

func addCalculation(v station2.Visitor[*calcState], c *station2.Calculation) error {
	err := visitCalculation(v, c)
	if err == nil {
		st := v.Get()
		calc := st.calc

		// RESET EVERY cron definition
		if c.ResetEvery != nil {
			if _, err := calc.Cron.AddFunc(c.ResetEvery.Definition, func() {
				calc.Latest.Remove(c.Target)
				if c.Load != nil {
					//_ = calc.loadFromDB(c)
				}
				calc.calculateTarget(c.Target)
			}); err != nil {
				return err
			}
		}

		// Every definition
		if c.Every != nil {
			if _, err := calc.Cron.AddFunc(c.Every.Definition, func() {
				if c.Load != nil {
					//_ = calc.loadFromDB(c)
				}
				calc.calculateTarget(c.Target)
			}); err != nil {
				return err
			}
		}

		// If the target still has no Calculation registered then create it.
		// This will happen when a calculation is defined that doesn't
		// reference any metrics. e.g. SolarAltitude which uses just location and time
		if calc.getCalculationByTarget(c.Target) == nil {
			calc.addCalculationByTarget(NewCalculation(c, st.station))
		}

		if c.Load != nil {
			//if err := calc.loadFromDB(c); err != nil {
			//	return err
			//}
		}
	}
	return err
}

func visitEphemeris(v station2.Visitor[*calcState], c *station2.Ephemeris) error {
	v.Get().ephemeris = c
	return nil
}

func addEphemeris(v station2.Visitor[*calcState], d *station2.Ephemeris) error {
	err := visitEphemeris(v, d)

	if err == nil {
		st := v.Get()
		calc := st.calc

		for _, c := range d.Schedules {
			err = calc.addEphemeris(st.station, d, c)
			if err != nil {
				err = errors.Error(c.Pos, err)
				break
			}
		}

	}

	return err
}

func addMetric(v station2.Visitor[*calcState], c *station2.Metric) error {
	s := v.Get()
	s.calc.addMetric(c.Name, s.c, s.station)
	return nil
}

func visitStation(v station2.Visitor[*calcState], c *station2.Station) error {
	v.Get().station = c
	return nil
}
