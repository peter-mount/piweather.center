package station

import (
	"github.com/peter-mount/piweather.center/config/station"
	"sync"
)

type Dashboard struct {
	mutex      sync.Mutex
	station    *Station                // Containing Station
	dashboard  *station.Dashboard      // Associated Dashboard
	uid        string                  // dashboard uid used by front end to indicate a refresh
	cronSeq    int                     // cron sequence
	cronId     int                     // cron id
	components map[string]*Component   // Component indexed by component id
	metrics    map[string][]*Component // map of Component's that use a metric
}

func newDashboard(station *Station, dashboard *station.Dashboard) *Dashboard {
	return &Dashboard{
		station:    station,
		dashboard:  dashboard,
		uid:        station.uid,
		components: make(map[string]*Component),
		metrics:    make(map[string][]*Component),
	}
}

func (d *Dashboard) Station() *Station {
	return d.station
}

func (d *Dashboard) AcceptMetric(id string) bool {
	return d.station.AcceptMetric(id)
}

func (d *Dashboard) Dashboard() *station.Dashboard {
	return d.dashboard
}

func (d *Dashboard) GetComponent(id string) *Component {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.components[id]
}

func (d *Dashboard) GetOrCreateComponent(id string) *Component {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	c, exists := d.components[id]
	if !exists {
		c = newComponent(id, d)
		d.components[id] = c
		d.metrics[id] = append(d.metrics[id], c)
	}
	return c
}

func (d *Dashboard) linkComponent(metric string, c *Component) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	m, exists := d.metrics[metric]

	// Look for an existing mapping & ignore if found
	if exists {
		for _, e := range m {
			if e == c {
				return
			}
		}
	}

	d.metrics[metric] = append(m, c)
}

func (d *Dashboard) GetComponentsByMetric(id string) []*Component {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	r, exists := d.metrics[id]
	if exists {
		r = append([]*Component{}, r...)
	}
	return r
}

func (d *Dashboard) GetType() string {
	return "dashboard"
}

func (d *Dashboard) Definition() any {
	return d.Dashboard()
}

func (d *Dashboard) GetData() map[string]interface{} {
	st := d.Station()
	dash := d.Dashboard()
	return map[string]interface{}{
		"stationId": st.Station().Name,
		"dash":      dash.Name,
		"board":     d,
		"js":        nil,
	}
}
