package weather

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/io"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/store/memory"
	"time"
)

func init() {
	packages.Register("weather", &Weather{})
}

type Weather struct {
	stations *station.Stations
}

func (w *Weather) LoadConfig(fileName string) error {
	if w.stations == nil {
		config := station.Stations(make(map[string]*station.Station))

		if err := io.NewReader().
			Yaml(&config).
			Open(fileName); err != nil {
			return err
		}

		if err := config.Init(); err != nil {
			return err
		}

		w.stations = &config
	}
	return nil
}

func (w *Weather) NewStore(dir string) (*memory.Store, error) {
	return memory.New(dir, w.stations)
}

func (w *Weather) Reducer(period time.Duration) *memory.Reducer {
	return memory.NewReducer(period)
}

func (w *Weather) ReducerMinutes(minutes int) *memory.Reducer {
	return memory.NewReducerMins(minutes)
}
