package station

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/store/api"
	cron2 "gopkg.in/robfig/cron.v2"
	"sort"
	"strings"
	"sync"
)

type Station struct {
	mutex        sync.Mutex
	stations     *Stations               // Stations containing this station
	station      *station.Station        // Stad *Dashboardtion associated with this instance
	metricPrefix string                  // metric prefix to limit metrics to this station
	metrics      map[string]api.Metric   // map of current Metric values for this Station
	dashboards   map[string]*Dashboard   // map of Dashboards
	calculations map[string]*Calculation // map of Calculations
	uid          string                  // uid of the original Stations file loaded
	cronIds      map[string]int          // Map of cron ids
}

func newStation(s *station.Station) *Station {
	return &Station{
		station:      s,
		metricPrefix: s.Name + ".",
		metrics:      make(map[string]api.Metric),
		dashboards:   make(map[string]*Dashboard),
		calculations: make(map[string]*Calculation),
		cronIds:      make(map[string]int),
	}
}

func (s *Station) updateCron(d *Dashboard, useCron bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	n := d.dashboard.Name
	if oid, exists := s.cronIds[n]; exists {
		delete(s.cronIds, n)
		defer func() {
			log.Printf("Cron: Removing %q %d", n, oid)
			s.stations.Cron.Remove(cron2.EntryID(oid))
		}()
	}
	if useCron {
		// record ID only if we are creating a new one.
		// Note: cronId can be 0 so we can't do this blindly,
		//hence the bool indicating we have actually created one
		s.cronIds[n] = d.cronId
	}
}

func (s *Station) Station() *station.Station {
	return s.station
}

func (s *Station) Stations() *Stations {
	return s.stations
}

func (s *Station) GetUid() string {
	return s.uid
}

func (s *Station) addDashboard(d *Dashboard) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.dashboards[d.dashboard.Name] = d
}

func (s *Station) addCalculation(d *Calculation) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.calculations[d.calculation.Target] = d
}

// GetMetric returns the current value of a Metric
func (s *Station) GetMetric(n string) (api.Metric, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	m, exists := s.metrics[n]
	return m, exists
}

// SetMetric sets a new value for a Metric.
// Normally this function is not called directly but via the relevant components
func (s *Station) SetMetric(newMetric api.Metric) {
	if s.AcceptMetric(newMetric.Metric) {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		oldMetric, exists := s.metrics[newMetric.Metric]
		if !exists || newMetric.IsNewerThan(oldMetric) {
			s.metrics[newMetric.Metric] = newMetric
			//log.Printf("SetMetric: Added new metric %q %.3f", newMetric.Metric, newMetric.Value)
		}
	}
}

func (s *Station) AcceptMetric(id string) bool {
	return strings.HasPrefix(id, s.metricPrefix)
}

// AcceptMetrics returns a list of existing metrics which match the AcceptMetric interface
// accepts
func (s *Station) AcceptMetrics(a AcceptMetric) []string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var r []string
	var m api.Metric

	for k, _ := range s.metrics {
		m.Metric = k
		if a.AcceptMetric(m) {
			r = append(r, k)
		}
	}

	sort.SliceStable(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return r
}

func (s *Station) GetDashboard(id string) *Dashboard {
	if s == nil {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.dashboards[id]
}

func (s *Station) GetCalculation(target string) *Calculation {
	if s == nil {
		return nil
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.calculations[target]
}
