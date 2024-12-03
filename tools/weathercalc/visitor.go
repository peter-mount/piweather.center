package weathercalc

import (
	station2 "github.com/peter-mount/piweather.center/config/station"
	"log"
)

type calcState struct {
	calc    *Calculator
	station *station2.Station
	c       *station2.Calculation
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

		log.Printf("Target %q", c.Target)

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

func addMetric(v station2.Visitor[*calcState], c *station2.Metric) error {
	s := v.Get()
	log.Printf("Metric %q", c.Name)
	s.calc.addMetric(c.Name, s.c, s.station)
	return nil
}

func visitStation(v station2.Visitor[*calcState], c *station2.Station) error {
	v.Get().station = c

	// Calculator doesn't need the following so prune to save memory
	c.Dashboards = nil

	return nil
}
