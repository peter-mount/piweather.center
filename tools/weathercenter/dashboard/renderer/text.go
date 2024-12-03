package renderer

import (
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/forecast"
	"github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

func Text(v station.Visitor[*State], d *station.Text) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			s.Builder().TextNbsp(d.Text).End()
			return nil
		})

	if err != nil {
		return err
	}

	return util.VisitorStop
}

func Forecast(v station.Visitor[*State], d *station.Forecast) error {
	err := v.Get().
		Component(v, d, d.Component, func(s *State) error {
			var txt string

			stn := s.Dashboard().Station()
			loc := stn.Station().Location

			temp, _ := stn.GetMetric(d.Temperature.Name)
			pressure, _ := stn.GetMetric(d.Pressure.Name)
			windDir, _ := stn.GetMetric(d.WindDirection.Name)

			if temp.IsValid() && pressure.IsValid() && windDir.IsValid() {
				t := time2.LatestTime(temp.Time, pressure.Time, windDir.Time)

				t1, _ := temp.ToValue()
				p1, _ := pressure.ToValue()
				wd, _ := windDir.ToValue()

				_, str := calcForecast(t, t1, p1, wd, loc)

				if str != "" {
					txt = str

					// Hack, for now notify a dummy metric for the forecast until calculator handle it
					//if st := s.Dashboard().Station().Stations(); st != nil {
					//	go func() {
					//		time.Sleep(time.Second)
					//		log.Printf("Pub forecast")
					//		st.Notify(api.Metric{
					//			Metric:    "home.ecowitt.forecast",
					//			Time:      t,
					//			Unit:      forecast.Zambretti.ID(),
					//			Value:     float64(zam),
					//			Formatted: zam.String(),
					//			Unix:      t.Unix(),
					//		})
					//	}()
					//}
				}
			}
			s.Builder().Span().TextNbsp(txt).End().End()
			return nil
		})

	if err != nil {
		return err
	}

	return util.VisitorStop
}

func calcForecast(t time.Time, t1, p1, wd value.Value, loc *location.Location) (forecast.ZambrettiSeverity, string) {
	h := forecast.HemisphereFor(loc.LatLong().Latitude.Deg())

	p0, err := measurement.PressureMeanSeaLevel(p1, t1, measurement.Meters.Value(loc.Altitude))
	if err != nil {
		return 0, ""
	}

	zam, err := forecast.CalculateZambrettiForecast(t, h, p0, p1, wd)
	if err != nil {
		return 0, ""
	}
	return zam, zam.String()
}
