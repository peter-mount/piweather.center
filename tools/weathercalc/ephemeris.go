package weathercalc

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
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
	_, err := calc.Cron.AddFunc(schedule.Every.Definition, func() {
		calc.calculateEphemeris(stn, ephemeris, schedule)
	})
	return err
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
	// TODO implement

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
		err = util.VisitorStop
	}

	return err
}
