package station

import (
	"github.com/peter-mount/go-kernel/v2/cron"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/store/api"
	"sort"
	"strings"
	"sync"
)

// Stations holds a map of State indexed by stationId and dashboardId
type Stations struct {
	Cron              *cron.CronService `kernel:"inject"`
	mutex             sync.Mutex
	stations          map[string]*Station            // map of State's
	newStationVisitor station.Visitor[*visitorState] // Visitor to add a new/updated station
	notifyVisitor     station.Visitor[*visitorState] // Visitor to handle new metrics
	loadVisitor       station.Visitor[*visitorState] // Visitor to handle new metrics
}

func (s *Stations) Start() error {
	s.stations = make(map[string]*Station)

	// Visitor to add a station & it's dashboards to this instance
	s.newStationVisitor = station.NewBuilder[*visitorState]().
		Dashboard(addDashboard).
		Gauge(addGauge).
		MultiValue(addMultiValue).
		Station(addStation).
		Stations(visitStations).
		Value(addValue).
		Build()

	// Visitor to add metrics to the station and to create Response's if a Dashboard is live
	s.notifyVisitor = station.NewBuilder[*visitorState]().
		Dashboard(notifyDashboard).
		Gauge(visitGauge).
		MultiValue(visitMultiValue).
		Station(visitStationFilterMetric).
		Stations(visitStations).
		Value(visitValue).
		Build()

	// Visitor used to load metrics. Identical to notifyVisitor but without creating responses
	s.loadVisitor = station.NewBuilder[*visitorState]().
		Dashboard(visitDashboard).
		Gauge(visitGauge).
		MultiValue(visitMultiValue).
		Station(visitStationFilterMetric).
		Stations(visitStations).
		Value(visitValue).
		Build()

	return nil
}

func (s *Stations) GetStation(id string) *Station {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	st := s.stations[id]
	if st != nil {
		// Ensure that the station points back to us
		st.stations = s
	}
	return st
}

// AddStations adds all stations, replacing any existing ones.
// This function is usually called at startup rather than when updating
// stations later.
func (s *Stations) AddStations(stations *station.Stations) {
	_ = s.newStationVisitor.Clone().
		Set(newVisitorState(s)).
		Stations(stations)
}

// AddStation adds a new station, replacing any existing one
func (s *Stations) AddStation(station *station.Station) {
	_ = s.newStationVisitor.Clone().
		Set(newVisitorState(s)).
		Station(station)
}

func (s *Stations) addStation(station *Station) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stations[station.station.Name] = station
}

// RemoveStation removes the named station
func (s *Stations) RemoveStation(stationId string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.stations, stationId)
}

func (s *Stations) Notify(m api.Metric) []*Response {
	if m.IsValid() {
		name := strings.SplitN(m.Metric, ".", 2)
		st := s.GetStation(name[0])
		if st != nil {
			state := newVisitorState(s)
			state.metric = m

			_ = s.notifyVisitor.Clone().
				Set(state).
				Station(st.Station())

			// Return the build responses, one per Dashboard
			return state.responses
		}
	}

	return nil
}

func (s *Stations) Load(metrics []api.Metric) {

	sort.SliceStable(metrics, func(i, j int) bool {
		return metrics[i].Metric < metrics[j].Metric
	})

	state := newVisitorState(s)
	v := s.loadVisitor.Clone().Set(state)

	var st *Station
	lastName := ""

	for _, m := range metrics {
		if m.IsValid() {
			name := strings.SplitN(m.Metric, ".", 2)[0]
			if name != lastName {
				lastName = name
				st = s.GetStation(name)
			}
			if st != nil {
				state.metric = m
				_ = v.Station(st.Station())
			}
		}
	}
}
