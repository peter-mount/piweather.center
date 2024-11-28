package station

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/units"
	util2 "github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
)

func NewParser() util.Parser[Station] {
	return util.NewParser[Station](nil, nil, stationInit)
}

func stationInit(q *Station, err error) (*Station, error) {

	if err == nil {
		s := &state{
			dashboards: make(map[string]*Dashboard),
		}

		err = NewBuilder[*state]().
			Container(s.container).
			Dashboard(s.dashboard).
			Gauge(s.gauge).
			Location(s.location).
			Metric(s.metric).
			MetricPattern(s.metricPattern).
			Station(s.station).
			Unit(s.unit).
			Value(s.value).
			Build().
			Set(s).
			Station(q)
	}

	return q, err
}

type state struct {
	stationId     string // copy of the stationId
	stationPrefix string // stationId + "."
	sensorPrefix  string // sensorId + "."
	dashboards    map[string]*Dashboard
}

func (s *state) station(_ Visitor[*state], d *Station) error {
	var err error

	// Enforce lower case name
	d.Name = strings.ToLower(strings.TrimSpace(d.Name))

	if d.Name == "" {
		err = errors.Errorf(d.Pos, "station id is required")
	}

	if err == nil && strings.ContainsAny(d.Name, ". _") {
		err = errors.Errorf(d.Pos, "station id must not contain '.', '_' or spaces")
	}

	if err == nil {
		s.stationId = d.Name
		s.stationPrefix = s.stationId + "."

		// Ensure we have a Location
		if d.Location == nil {
			d.Location = &location.Location{Pos: d.Pos}
		}
	}

	return errors.Error(d.Pos, err)
}

func (s *state) location(_ Visitor[*state], d *location.Location) error {
	var err error

	d.Name = strings.TrimSpace(d.Name)
	d.Longitude = strings.TrimSpace(d.Longitude)
	d.Latitude = strings.TrimSpace(d.Latitude)
	d.Notes = strings.TrimSpace(d.Notes)

	if d.Name == "" {
		d.Name = s.stationId
	}

	if d.Longitude == "" && d.Latitude == "" {
		// set to Null Island
		d.Longitude = "0.0"
		d.Latitude = "0.0"
	}

	if d.Longitude == "" || d.Latitude == "" {
		err = errors.Errorf(d.Pos, "both latitude AND longitude are required")
	}

	if err == nil {
		err = d.Init()
	}

	return errors.Error(d.Pos, err)
}

func (s *state) container(_ Visitor[*state], d *Container) error {
	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}
	return nil
}

func (s *state) dashboard(_ Visitor[*state], d *Dashboard) error {
	var err error

	// Enforce lower case name
	d.Name = strings.ToLower(strings.TrimSpace(d.Name))
	if d.Name == "" {
		err = errors.Errorf(d.Pos, "dashboard name is required")
	}
	if err == nil && strings.ContainsAny(d.Name, ". _") {
		err = errors.Errorf(d.Pos, "dashboard name must not contain '.', '_' or spaces")
	}

	// Check name is unique
	if err == nil {
		if e, exists := s.dashboards[d.Name]; exists {
			err = errors.Errorf(d.Pos, "dashboard %q already exists at %s", d.Name, e.Pos.String())
		}
	}

	if err == nil {
		s.dashboards[d.Name] = d

		// Ensure Component exists, require by templates
		if d.Component == nil {
			d.Component = &Component{}
		}
	}

	return errors.Error(d.Pos, err)
}

func (s *state) multiValue(_ Visitor[*state], d *MultiValue) error {
	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}

	return nil
}

func (s *state) value(_ Visitor[*state], d *Value) error {
	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}
	return nil
}

func (s *state) gauge(_ Visitor[*state], d *Gauge) error {
	var err error

	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}

	// ensure min < max
	if value.GreaterThan(d.Min, d.Max) {
		d.Min, d.Max = d.Max, d.Min
	}

	// default values
	if value.IsZero(d.Min) && value.IsZero(d.Max) {
		// Default to 0...100
		d.Min, d.Max = 0.0, 100.0
	}

	if d.Ticks == 0 {
		// Default to 10 ticks
		d.Ticks = 10
	}

	// verify state
	switch {
	case value.Equal(d.Min, d.Max):
		err = errors.Errorf(d.Pos, "Min and Max must not be the same")

	case d.Ticks < 0:
		err = errors.Errorf(d.Pos, "Ticks %d is invalid", d.Ticks)

	case d.Metrics == nil || len(d.Metrics.Metrics) == 0:
		// We must have at least 1 metric for gauges
		err = errors.Errorf(d.Pos, "No metrics provided for Gauge")
	}

	return errors.Error(d.Pos, err)
}

func (s *state) metric(_ Visitor[*state], d *Metric) error {
	var err error

	// enforce metrics to be lower case
	d.Name = strings.ToLower(strings.TrimSpace(d.Name))

	if d.Name == "" {
		err = errors.Errorf(d.Pos, "metric name is required")
	}

	if err == nil && strings.ContainsAny(d.Name, " ") {
		err = errors.Errorf(d.Pos, "metric name must not include spaces")
	}

	// Prefix with the stationId & sensorId to become a full metric id
	d.Name = s.stationPrefix + s.sensorPrefix + d.Name

	return errors.Error(d.Pos, err)
}

func (s *state) metricPattern(_ Visitor[*state], d *MetricPattern) error {
	var err error

	t, p := util2.ParsePatternType(d.Pattern)

	if strings.ContainsAny(p, " *") {
		err = errors.Errorf(d.Pos, "pattern must not include '*' or spaces")
	}

	// Disallow equality as that makes no sense for this component
	if err == nil && t == util2.PatternEquals {
		err = errors.Errorf(d.Pos, "No wildcard provided")
	}

	// For MetricPattern we want "" as an alias for "*"
	if err == nil && t == util2.PatternNone {
		t = util2.PatternAll
	}

	if err == nil {
		d.Pattern = strings.ToLower(p)
		d.Type = t

		d.Prefix = s.stationPrefix + s.sensorPrefix
	}

	return err
}

func (s *state) unit(_ Visitor[*state], d *units.Unit) error {
	return errors.Error(d.Pos, d.Init())
}
