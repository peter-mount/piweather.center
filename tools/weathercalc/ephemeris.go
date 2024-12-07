package weathercalc

import (
	"github.com/peter-mount/piweather.center/astro/api"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	api2 "github.com/peter-mount/piweather.center/store/api"
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

	case targetType.IsPlanet():
		// TODO implement

	default:
		// Other, so just ignore for now
	}

	if err == nil && result != nil {
		for _, opt := range d.Options {
			val := result.Value(opt.TargetType())
			if val.IsValid() {
				t := st.time.Time()
				metric := api2.Metric{
					Metric:    opt.As,
					Time:      t,
					Unit:      val.Unit().ID(),
					Value:     val.Float(),
					Formatted: val.String(),
					Unix:      t.Unix(),
				}
				if metric.IsValid() {
					err = st.calc.DatabaseBroker.PublishMetric(metric)
					if err != nil {
						break
					}
				}

			}
		}
	}

	if err == nil {
		err = util.VisitorStop
	}

	return err
}
