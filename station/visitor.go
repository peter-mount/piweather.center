package station

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
	stations     *Stations                       // Link to root Stations instance
	station      *Station                        // Station being processed
	dashboard    *Dashboard                      // Dashboard being processed
	calculation  *Calculation                    // Calculation being processed
	metric       api.Metric                      // The metric being processed
	idSeq        int                             // ID sequence within a dashboard
	response     *Response                       // The response currently being built within a dashboard
	responses    []*Response                     // Slice of built responses
	loadOption   LoadOption                      // If set defines how LoadDirectory operates
	calculations map[string]*station.Calculation // Map of calculations, used to ensure
}

func (s *visitorState) nextId() string {
	s.idSeq++
	return string(encode([]byte{uid[0], uid[1], byte(s.idSeq & 0xff), byte((s.idSeq >> 8) & 0xff)}))
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

	// Remove children we do not require
	if st.loadOption.Not(DashboardOption) {
		d.Dashboards = nil
	}
	if st.loadOption.Not(CalculationOption) {
		d.Calculations = nil
	}

	log.Printf("Added Station %q", d.Name)

	return nil
}

func addCalculation(v station.Visitor[*visitorState], d *station.Calculation) error {
	st := v.Get()
	calc := newCalculation(st.station, d)
	st.calculation = calc
	st.station.addCalculation(calc)

	log.Printf("Added calculation %q", d.Target)

	return util.VisitorStop
}

func addDashboard(v station.Visitor[*visitorState], d *station.Dashboard) error {
	st := v.Get()
	de := newDashboard(st.station, d)
	st.dashboard = de
	st.idSeq = 0
	st.station.addDashboard(de)

	var useCron bool
	if d.Update != nil {
		id, err := st.stations.Cron.AddFunc(d.Update.Definition, func() {
			// TODO check we need to update the UID here?
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

// visitDashboard will stop processing if the dashboard does not exist
func visitDashboard(v station.Visitor[*visitorState], d *station.Dashboard) error {
	st := v.Get()

	// If somehow we have no station set then stop processing the Dashboard
	if st.station == nil {
		st.dashboard = nil
		return util.VisitorStop
	}

	// Stop here if the dashboard isn't registered - should never occur
	st.dashboard = st.station.GetDashboard(d.Name)
	if st.dashboard == nil {
		return util.VisitorStop
	}

	return nil
}

// notifyDashboard calls visitDashboard but will then create a Response for this dashboard if the dashboard is live
func notifyDashboard(v station.Visitor[*visitorState], d *station.Dashboard) error {
	err := visitDashboard(v, d)

	// The dashboard is live so build a response
	if err == nil && d.Live {
		st := v.Get()

		st.response = &Response{
			Station:   st.station.Station().Name,
			Dashboard: d.Name,
			// Set the response uid which is unique to the station (or stations loaded from the same file)
			Uid: st.dashboard.uid,
		}

		err = v.ComponentListEntry(d.Components)
		if err == nil || util.IsVisitorStop(err) {
			// Add the response to responses and stop here as we've already visited the component
			if st.response.IsValid() {
				st.responses = append(st.responses, st.response)
			}
			st.response = nil

			// Stop processing the dashboard as we have done that here
			err = util.VisitorStop
		}
	}

	return err
}

func addMetric(v station.Visitor[*visitorState], d *station.Metric) error {
	st := v.Get()
	calc := st.calculation
	if calc != nil {
		calc.AddMetric(d.Name)
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
		if !st.station.AcceptMetric(m.Metric) {
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

func addMultiValue(v station.Visitor[*visitorState], d *station.MultiValue) error {
	st := v.Get()
	dash := st.dashboard

	if dash != nil && d.Pattern != nil {
		d.Component.ID = st.nextId()
		comp := dash.GetOrCreateComponent(d.GetID())
		comp.Sorted()
	}

	return util.VisitorStop
}

func visitMultiValue(v station.Visitor[*visitorState], d *station.MultiValue) error {
	st := v.Get()
	resp := st.response
	if resp != nil {
		dash := st.dashboard

		m := st.metric

		if dash != nil && d.Pattern != nil && m.IsValid() {
			comp := dash.GetComponent(d.GetID())
			if comp != nil && d.AcceptMetric(m) {
				// Submit and if the metric is not present then add it
				mn := m.Metric
				cei := comp.GetMetrics(mn)
				if len(cei) == 0 {
					comp.AddMetric(0, "", mn)
					if d.Time {
						comp.AddMetric(0, "T", mn)
					}
					// force a page refresh
					st.response.Uid = ""
				}

				comp.Submit(st.response, d, -1, m)

				st.station.SetMetric(m)
			}
		}
	}

	return util.VisitorStop
}

func addGauge(v station.Visitor[*visitorState], d *station.Gauge) error {
	d.Component.ID = v.Get().nextId()
	return addMetricListImpl(v, d.GetID(), d.Metrics)
}

func addValue(v station.Visitor[*visitorState], d *station.Value) error {
	d.Component.ID = v.Get().nextId()
	return addMetricListImpl(v, d.GetID(), d.Metrics)
}

func visitGauge(v station.Visitor[*visitorState], d *station.Gauge) error {
	return visitMetricListImpl(v, d, d.Metrics)
}

func visitValue(v station.Visitor[*visitorState], d *station.Value) error {
	return visitMetricListImpl(v, d, d.Metrics)
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

func visitMetricListImpl(v station.Visitor[*visitorState], d ResponseComponent, l *station.MetricList) error {
	st := v.Get()
	dash := st.dashboard
	m := st.metric

	if dash != nil && l != nil && m.IsValid() {
		comp := dash.GetComponent(d.GetID())
		if comp != nil {
			for idx, e := range l.Metrics {
				if e.AcceptMetric(m) {
					// set the metric in the DB
					st.station.SetMetric(m)

					// Now check for different unit
					if metric, exists := st.station.GetMetric(m.Metric); exists {
						val, _ := metric.ToValue()
						val, _ = e.Convert(val)
						comp.Submit(st.response, d, idx, api.Metric{
							Metric:    m.Metric,
							Time:      m.Time,
							Unit:      val.Unit().ID(),
							Value:     val.Float(),
							Formatted: val.String(),
							Unix:      m.Unix,
						})
					}
				}
			}
			return nil
		}
	}

	return util.VisitorStop
}
