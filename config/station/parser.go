package station

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
	util2 "github.com/peter-mount/piweather.center/util"
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

		err = NewBuilder[state]().
			Container(s.container).
			Dashboard(s.dashboard).
			Metric(s.metric).
			MetricPattern(s.metricPattern).
			Station(s.station).
			Value(s.value).
			Build().
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

func (s *state) station(_ Visitor[state], d *Station) error {
	// Enforce lower case name
	d.Name = strings.ToLower(strings.TrimSpace(d.Name))

	if d.Name == "" {
		return errors.Errorf(d.Pos, "station id is required")
	}

	if strings.ContainsAny(d.Name, ". _") {
		return errors.Errorf(d.Pos, "station id must not contain '.', '_' or spaces")
	}

	s.stationId = d.Name
	s.stationPrefix = s.stationId + "."

	// Ensure we have a Location, so set to Null Island
	if d.Location == nil {
		d.Location = &location.Location{
			Pos:       d.Pos,
			Name:      d.Name,
			Latitude:  "0.0",
			Longitude: "0.0",
			Altitude:  0,
		}
	}

	return nil
}

func (s *state) container(_ Visitor[state], d *Container) error {
	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}
	return nil
}

func (s *state) dashboard(_ Visitor[state], d *Dashboard) error {
	// Enforce lower case name
	d.Name = strings.ToLower(strings.TrimSpace(d.Name))
	if d.Name == "" {
		return errors.Errorf(d.Pos, "dashboard name is required")
	}
	if strings.ContainsAny(d.Name, ". _") {
		return errors.Errorf(d.Pos, "dashboard name must not contain '.', '_' or spaces")
	}

	// Check name is unique
	if e, exists := s.dashboards[d.Name]; exists {
		return errors.Errorf(d.Pos, "dashboard %q already exists at %s", d.Name, e.Pos.String())
	}
	s.dashboards[d.Name] = d

	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}

	return nil
}

func (s *state) multivalue(_ Visitor[state], d *MultiValue) error {
	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}

	return nil
}

func (s *state) value(_ Visitor[state], d *Value) error {
	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}
	return nil
}

func (s *state) metric(_ Visitor[state], d *Metric) error {
	// enforce metrics to be lower case
	d.Name = strings.ToLower(strings.TrimSpace(d.Name))

	if d.Name == "" {
		return errors.Errorf(d.Pos, "metric name is required")
	}

	if strings.ContainsAny(d.Name, " _") {
		return errors.Errorf(d.Pos, "metric name must not include '_' or spaces")
	}

	// Prefix with the stationId & sensorId to become a full metric id
	d.Name = s.stationPrefix + s.sensorPrefix + d.Name

	return nil
}

func (s *state) metricPattern(_ Visitor[state], d *MetricPattern) error {
	var err error

	t, p := util2.ParsePatternType(d.Pattern)

	if strings.ContainsAny(p, " _*") {
		err = errors.Errorf(d.Pos, "pattern must not include '_', '*' or spaces")
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
