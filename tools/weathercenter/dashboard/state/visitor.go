package state

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/store/api"
	"strconv"
	"strings"
)

type AcceptMetric interface {
	AcceptMetric(api.Metric) bool
}

type visitorState struct {
	stations  *Stations  // Link to root Stations instance
	station   *Station   // Station being processed
	dashboard *Dashboard // Dashboard being processed
	component *Component // Component at this point in time
	metric    api.Metric // The metric being processed
	response  *Response  // The response being built
}

func newVisitorState(s *Stations) *visitorState {
	return &visitorState{stations: s}
}

func addStation(v station.Visitor[*visitorState], d *station.Station) error {
	se := newStation(d)

	// initial id is the instance uid appended with the sha1 sum of the dashboard's bytes
	var id []byte
	id = append(id, uid...)

	// Note: we cannot use id=append(id,sha1.Sum(b)...) here as it returns [20]byte instead of
	// []byte, and go does not allow that hence we manually append the bytes
	sum := d.GetChecksum()
	for i := 0; i < len(sum); i++ {
		id = append(id, sum[i])
	}

	id = compress(id)

	se.uid = string(encode(id))

	st := v.Get()
	st.station = se

	st.stations.addStation(se)

	log.Printf("Added Station %q", d.Name)

	return nil
}

func addDashboard(v station.Visitor[*visitorState], d *station.Dashboard) error {
	st := v.Get()
	de := newDashboard(st.station, d)
	st.dashboard = de
	st.station.addDashboard(de)

	var useCron bool
	if d.Update != nil {
		id, err := st.stations.Cron.AddFunc(d.Update.Definition, func() {
			//d.Init(*s.Server.DBServer)
			// Make a new Uid so client refreshes
			st.dashboard.cronSeq++
			uid := strings.Split(st.dashboard.uid, "-")
			st.dashboard.uid = uid[0] + "-" + strconv.Itoa(st.dashboard.cronSeq)
		})
		if err == nil {
			st.dashboard.cronId = int(id)
			useCron = true
			log.Printf("Cron: Adding %q %d", d.Name, st.dashboard.cronId)
		}
	}
	st.station.updateCron(st.dashboard, useCron)

	log.Printf("Added Dashboard \"%s/%s\"", st.station.station.Name, d.Name)

	return nil
}

func addMultiValue(v station.Visitor[*visitorState], d *station.MultiValue) error {
	st := v.Get()
	dash := st.dashboard

	if dash != nil && d.Pattern != nil {
		comp := dash.GetOrCreateComponent(d.GetID())

		// Prepopulate if we have metrics already available
		metrics := st.station.AcceptMetrics(d.Pattern)
		if len(metrics) > 0 {
			for i, m := range metrics {
				if dash.AcceptMetric(m) {
					comp.AddMetric(i, "", m)

					if d.Time {
						comp.AddMetric(i, "T", m)
					}
				}
			}

			comp.SortMetrics()
		}
	}

	return util.VisitorStop
}

func addGauge(v station.Visitor[*visitorState], d *station.Gauge) error {
	return addMetricListImpl(v, d.GetID(), d.Metrics)
}

func addValue(v station.Visitor[*visitorState], d *station.Value) error {
	return addMetricListImpl(v, d.GetID(), d.Metrics)
}

func addMetricListImpl(v station.Visitor[*visitorState], id string, d *station.MetricList) error {
	st := v.Get()
	dash := st.dashboard

	if dash != nil && d != nil {
		comp := dash.GetOrCreateComponent(id)

		for i, m := range d.Metrics {
			if dash.AcceptMetric(m.Name) {
				comp.AddMetric(i, "", m.Name)
			}
		}
	}

	return util.VisitorStop
}

func visitDashboard(v station.Visitor[*visitorState], d *station.Dashboard) error {
	st := v.Get()

	// If somehow we have no station set then stop processing the Dashboard
	if st.station == nil {
		st.dashboard = nil
		return util.VisitorStop
	}

	st.dashboard = st.station.GetDashboard(d.Name)
	if st.dashboard == nil {
		return util.VisitorStop
	}

	return nil
}

// visitStation sets the station within the context
func visitStation(v station.Visitor[*visitorState], d *station.Station) error {
	st := v.Get()

	st.station = st.stations.GetStation(d.Name)

	// If no station or the station then stop processing it
	if st.station == nil {
		return util.VisitorStop
	}

	// Set the response uid which is unique to the station (or stations loaded from the same file)
	if st.response != nil {
		st.response.Uid = st.station.uid
	}

	return nil
}

// visitStationFilterMetric calls visitStation but then will stop processing
// the station if the metric being passed is not accepted by the station.
func visitStationFilterMetric(v station.Visitor[*visitorState], d *station.Station) error {
	err := visitStation(v, d)
	if err == nil {
		st := v.Get()
		m := st.metric

		// If the station does not accept the metric then stop processing it
		if !m.IsValid() || !st.station.AcceptMetric(m.Metric) {
			err = util.VisitorStop
		}
	}

	return err
}

func visitStations(v station.Visitor[*visitorState], _ *station.Stations) error {
	st := v.Get()
	st.station = nil
	st.dashboard = nil
	return nil
}

func visitMultiValue(v station.Visitor[*visitorState], d *station.MultiValue) error {
	st := v.Get()
	resp := st.response
	if resp != nil {
		dash := st.dashboard

		m := st.metric

		if dash != nil && d.Pattern != nil && m.IsValid() {
			comp := dash.GetComponent(d.GetID())
			if comp != nil {
				st.component = comp

				if d.AcceptMetric(m) {
					// TODO check if component has the metric or needs refresh if new
					comp.Submit(st.response, d, m)
					st.station.SetMetric(m)
				}
			}
		}
	}

	return util.VisitorStop
}

func visitGauge(v station.Visitor[*visitorState], d *station.Gauge) error {
	return visitComponentImpl(v, d, d.Metrics)
}

func visitValue(v station.Visitor[*visitorState], d *station.Value) error {
	return visitComponentImpl(v, d, d.Metrics)
}

func visitComponentImpl(v station.Visitor[*visitorState], d ResponseComponent, l *station.MetricList) error {
	st := v.Get()
	dash := st.dashboard
	m := st.metric

	if dash != nil && l != nil && m.IsValid() {
		comp := dash.GetComponent(d.GetID())
		if comp != nil {
			for _, e := range l.Metrics {
				if e.AcceptMetric(m) {
					comp.Submit(st.response, d, m)
					st.station.SetMetric(m)
				}
			}
			return nil
		}
	}

	return util.VisitorStop
}
