package station

import (
	"github.com/peter-mount/piweather.center/store/api"
	"sort"
	"sync"
)

type Component struct {
	mutex            sync.Mutex
	dashboard        *Dashboard                       // Dashboard containing this component
	id               string                           // component ID
	metricNames      []string                         // metric ID's in this component
	metricNamesIndex map[string][]ComponentEntryIndex // map of ID's keyed by metric ID
	sorted           bool                             // true for multiview
}

func newComponent(id string, dashboard *Dashboard) *Component {
	return &Component{
		id:               id,
		dashboard:        dashboard,
		metricNamesIndex: make(map[string][]ComponentEntryIndex),
	}
}

type ComponentEntryIndex struct {
	Index  int
	Suffix string
	Metric string
}

func (e *Component) Sorted() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.sorted = true
}

// Id of the component being managed
func (e *Component) Id() string { return e.id }

func (e *Component) AcceptMetric(m api.Metric) bool {
	if e.dashboard.AcceptMetric(m.Metric) {
		e.mutex.Lock()
		defer e.mutex.Unlock()
		_, exists := e.metricNamesIndex[m.Metric]
		return exists
	}

	return false
}

func (e *Component) NumMetrics(metricId string) int {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	mni, exists := e.metricNamesIndex[metricId]
	if exists {
		return len(mni)
	}
	return 0
}

func (e *Component) GetMetrics(metricId string) []ComponentEntryIndex {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	mni, exists := e.metricNamesIndex[metricId]
	if exists {
		return append([]ComponentEntryIndex{}, mni...)
	}

	return nil
}

// AddMetric adds a new StateEntryIndex for idx,suffix.
// Returns true if a new entry was created, false if it already existed
func (e *Component) AddMetric(idx int, suffix, metricId string) bool {
	r := e.addMetric(idx, suffix, metricId)
	if r {
		e.dashboard.linkComponent(metricId, e)
	}
	return r
}

func (e *Component) addMetric(idx int, suffix, metricId string) bool {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	// Look for an existing entry
	mni, exists := e.metricNamesIndex[metricId]
	if exists {
		for _, m := range mni {
			if m.Index == idx && m.Suffix == suffix {
				return false
			}
		}
	}

	// Add new index
	mni = append(mni, ComponentEntryIndex{
		Index:  idx,
		Suffix: suffix,
		Metric: metricId,
	})

	// Add metric name if it doesn't exist
	found := false
	for _, m := range e.metricNames {
		if m == metricId {
			found = true
			break
		}
	}
	if !found {
		e.metricNames = append(e.metricNames, metricId)
	}

	// Update the map with the new slice
	e.metricNamesIndex[metricId] = mni

	if e.sorted {
		e.sortMetrics()
	}

	return true
}

// RequiresMetric returns true if this StateEntry requires a metricId
func (e *Component) RequiresMetric(metricId string) bool {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	_, exists := e.metricNamesIndex[metricId]
	return exists
}

// GetMetricNames returns the metric ID's in the sequence they were added to the State.
// The order is fixed unless SortMetrics is called.
func (e *Component) GetMetricNames() []string {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	return append([]string{}, e.metricNames...)
}

// SortMetrics so that they are in alphabetical order.
// This is used by certain components, e.g. MultiValue, where the index of a metric
// is not defined by the sequence they are added to the State
func (e *Component) SortMetrics() {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.sortMetrics()
}

func (e *Component) sortMetrics() {

	sort.SliceStable(e.metricNames, func(i, j int) bool {
		return e.metricNames[i] < e.metricNames[j]
	})

	// Update the StateEntryIndex with the new index
	for i, name := range e.metricNames {
		mni, _ := e.metricNamesIndex[name]
		for j, v := range mni {
			v.Index = i
			mni[j] = v
		}
		e.metricNamesIndex[name] = mni
	}
}

// Submit a metric to it's components
//
// Note: idx is the index required, or -1 for all
func (e *Component) Submit(r *Response, c ResponseComponent, idx int, m api.Metric) {
	if r != nil {
		metrics := e.GetMetrics(m.Metric)
		for _, v := range metrics {
			if v.Index == idx || idx < 0 {
				r.SetComponent(c, v, m)
			}
		}
	}
}
