package store

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/piweather.center/station"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/state"
	"github.com/peter-mount/piweather.center/weather/value"
	"sort"
	"sync"
	"time"
)

// State manages an in memory copy of the ID data
type State struct {
	Store    *Store            `kernel:"inject"`
	Config   station.Config    `kernel:"inject"`
	Cron     *cron.CronService `kernel:"inject"`
	mutex    sync.Mutex
	stations map[string]*state.Station
}

func (s *State) Start() error {
	s.stations = make(map[string]*state.Station)

	if _, err := s.Cron.AddTask("58 * * * * ?", s.updateStations); err != nil {
		return err
	}

	// Trigger a scan so we have some data.
	_ = s.updateStations(nil)

	return nil
}

func (s *State) GetStation(id string) *state.Station {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.stations[id]
}

func (s *State) updateStations(_ context.Context) error {
	for _, stationDef := range *s.Config.Stations() {
		s.updateStation(stationDef)
	}
	return nil
}

// updateStation calculates a Station of Measurements for a specific Station
func (s *State) updateStation(station *station.Station) {
	now := time.Now().Truncate(time.Minute)

	stn := &state.Station{
		ID: station.ID,
		Meta: state.Meta{
			Name:       station.Name,
			Units:      make(map[string]state.Unit),
			Time:       now,
			Minute10:   now.Add(-10 * time.Minute),
			Previous10: now.Add(-20 * time.Minute),
			Hour:       now.Add(-time.Hour),
			Hour24:     now.Add(-24 * time.Hour),
			Today:      time2.LocalMidnight(now),
		},
	}

	for _, sensor := range station.Sensors {
		for _, reading := range sensor.Readings {
			stn.AddMeasurement(s.getMeasurement(stn, reading.ID))
		}
		for _, calc := range sensor.Calculations {
			stn.AddMeasurement(s.getMeasurement(stn, calc.ID))
		}
	}

	// Sort the measurements by id
	sort.SliceStable(stn.Measurements, func(i, j int) bool {
		return stn.Measurements[i].ID < stn.Measurements[j].ID
	})

	// Populate units in meta
	for _, m := range stn.Measurements {
		if _, exists := stn.Meta.Units[m.Unit]; !exists {
			if u, ok := value.GetUnit(m.Unit); ok {
				stn.Meta.Units[m.Unit] = state.Unit{
					ID:   m.Unit,
					Name: u.Name(),
					Unit: u.Unit(),
				}
			}
		}
	}

	// Update map, lock for shortest amount of time
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stations[station.ID] = stn
}

// getMeasurement calculates the ID of a specific reading id
func (s *State) getMeasurement(stn *state.Station, id string) *state.Measurement {

	meta := stn.Meta

	r := s.Store.GetHistoryBetween(id, meta.Hour24, meta.Time)
	if r == nil {
		return nil
	}

	m := &state.Measurement{ID: id}

	// Populate from the data
	r.ForEach(func(i int, t time.Time, value value.Value) {
		if value.IsValid() {
			// Set unit on first valid entry
			if m.Unit == "" {
				m.Unit = value.Unit().ID()
			}

			f := value.Float()

			if t.After(meta.Hour24) {
				m.Hour24 = m.Hour24.Include(f)
			}
			if t.After(meta.Today) {
				m.Today = m.Today.Include(f)
			}
			if t.After(meta.Hour) {
				m.Hour = m.Hour.Include(f)
			}

			if t.After(meta.Minute10) {
				m.Current = state.Point{
					Value: state.RoundedFloat(f),
					Time:  t,
				}
				m.Current10 = m.Current10.Include(f)
			}
			if t.After(meta.Previous10) && t.Before(meta.Minute10) {
				m.Previous = state.Point{
					Value: state.RoundedFloat(f),
					Time:  t,
				}
				m.Previous10 = m.Previous10.Include(f)
			}

			if t.After(m.Time) {
				m.Time = t
			}
		}
	})

	// Calculate trends
	m.Trends = state.Trends{
		From:    m.Previous.Time,
		To:      m.Current.Time,
		Current: state.TrendFrom(float64(m.Previous.Value), float64(m.Current.Value)),
		Min:     state.TrendFrom(float64(m.Previous10.Min), float64(m.Current10.Min)),
		Max:     state.TrendFrom(float64(m.Previous10.Max), float64(m.Current10.Max)),
		Mean:    state.TrendFrom(float64(m.Previous10.Mean), float64(m.Current10.Mean)),
	}

	return m
}
