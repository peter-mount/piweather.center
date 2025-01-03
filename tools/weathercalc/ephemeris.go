package weathercalc

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/config/station"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"gopkg.in/robfig/cron.v2"
	"time"
)

type ephemerisCalculator struct {
	calc      *Calculator
	station   *station.Station
	ephemeris *station.Ephemeris
	schedule  *station.EphemerisSchedule
	time      value.Time
	data      api.EphemerisResult
}

var (
	ephemerisVisitor = station.NewBuilder[*ephemerisCalculator]().
		EphemerisTarget(calculateEphemerisTarget).
		Build()
)

func (calc *Calculator) addEphemeris(stn *station.Station, ephemeris *station.Ephemeris, schedule *station.EphemerisSchedule) error {
	task := func() {
		calc.calculateEphemeris(stn, ephemeris, schedule)
	}

	_ = calc.Cron.Schedule(schedule.Every.Schedule(), cron.FuncJob(task))

	// On start calculation
	if schedule.OnStartup {
		go task()
	}

	return nil
}

func (calc *Calculator) calculateEphemeris(stn *station.Station, ephemeris *station.Ephemeris, schedule *station.EphemerisSchedule) {
	tm := stn.Location.Time().
		SetTime(time.Now().
			In(stn.TimeZone.Location()))

	_ = ephemerisVisitor.Clone().
		Set(&ephemerisCalculator{
			calc:      calc,
			station:   stn,
			ephemeris: ephemeris,
			schedule:  schedule,
			time:      tm,
		}).
		EphemerisSchedule(schedule)
}

func calculateEphemerisTarget(v station.Visitor[*ephemerisCalculator], d *station.EphemerisTarget) error {
	var err error

	st := v.Get()

	var result api.EphemerisResult

	targetType := d.GetTarget()
	switch {
	case targetType.IsSun():
		result, err = st.calc.Astro.CalculateSun(st.time)

	case targetType == station.EphemerisTargetMoon:
		result, err = st.calc.Astro.CalculateMoon(st.time)

	case targetType.IsPlanet():
		result, err = st.calc.Astro.CalculatePlanet(targetType, st.time)

	default:
		// Other, so just ignore for now
	}

	if err == nil && result != nil {
		for _, metric := range result.ToMetrics(st.ephemeris.Target+"."+d.As, d.GetEphemerisOption()) {
			log.Printf("Ephem %q %s %s\n", metric.Metric, metric.Formatted, metric.Time.Format(time2.RFC3339))
			err = st.calc.DatabaseBroker.PublishMetric(metric)
			if err != nil {
				break
			}
		}
	}

	if err == nil {
		err = errors.VisitorStop
	}

	return err
}
