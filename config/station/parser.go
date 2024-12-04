package station

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
	time2 "github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
	util2 "github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
	"time"
)

func NewParser() util.Parser[Stations] {
	return util.NewParser[Stations](nil, nil, stationInit)
}

func stationInit(q *Stations, err error) (*Stations, error) {

	if err == nil {
		s := &state{
			dashboards: make(map[string]*Dashboard),
		}

		err = NewBuilder[*state]().
			Axis(s.axis).
			Calculation(s.calculation).
			CalculationList(s.calculationList).
			Container(s.container).
			CronTab(s.cronTab).
			Dashboard(s.dashboard).
			Gauge(s.gauge).
			I2C(s.i2c).
			Location(s.location).
			Metric(s.metric).
			MetricExpression(s.metricExpression).
			MetricPattern(s.metricPattern).
			Sensor(s.sensor).
			SensorList(s.sensorList).
			Serial(s.serial).
			Station(s.station).
			Stations(s.stations).
			Unit(s.unit).
			Value(s.value).
			Build().
			Set(s).
			Stations(q)
	}

	return q, err
}

type state struct {
	stationId     string                  // copy of the stationId being processed
	stationPrefix string                  // stationId + "."
	sensorPrefix  string                  // sensorId + "."
	stationIds    map[string]*Station     // map of Stations
	calculations  map[string]*Calculation // map of calculations within a Station
	dashboards    map[string]*Dashboard   // map of Dashboards within a Station
	sensors       map[string]*Sensor      // map of Sensors within a station
}

func (s *state) prefixMetric(m string) string {
	return s.stationPrefix + s.sensorPrefix + m
}

func (s *state) stations(v Visitor[*state], d *Stations) error {
	s.stationIds = make(map[string]*Station)
	return nil
}

func (s *state) station(_ Visitor[*state], d *Station) error {
	s.calculations = make(map[string]*Calculation)
	s.dashboards = make(map[string]*Dashboard)

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

		// Ensure stationId is unique
		err = assertStationUnique(&s.stationIds, d)
	}

	return errors.Error(d.Pos, err)
}

func assertStationUnique(m *map[string]*Station, s *Station) error {
	n := strings.ToLower(s.Name)
	if e, exists := (*m)[n]; exists {
		return errors.Errorf(s.Pos, "station %q already defined at %s", s.Name, e.Pos)
	}
	(*m)[n] = s
	return nil
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

func (s *state) calculation(_ Visitor[*state], d *Calculation) error {
	target := strings.ToLower(d.Target)

	if e, exists := s.calculations[target]; exists {
		return participle.Errorf(d.Pos, "calculation for %q already defined at %s", d.Target, e.Pos.String())
	}

	d.Target = s.prefixMetric(target)
	s.calculations[target] = d
	return nil
}

func (s *state) calculationList(_ Visitor[*state], d *CalculationList) error {
	// sensorPrefix is not used for calculations
	s.sensorPrefix = ""
	if s.stationPrefix == "" {
		// should never occur
		return errors.Errorf(d.Pos, "stationPrefix not defined")
	}
	return nil
}

func (s *state) container(_ Visitor[*state], d *Container) error {
	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}
	// Ensure we have an entry present so we don't need to check this in templates
	if d.Components == nil {
		d.Components = &ComponentList{}
	}

	return nil
}

func (s *state) cronTab(_ Visitor[*state], d *time2.CronTab) error {
	return d.Init()
}

func (s *state) dashboard(_ Visitor[*state], d *Dashboard) error {
	var err error

	// sensorPrefix is not used for dashboards
	s.sensorPrefix = ""
	if s.stationPrefix == "" {
		// should never occur
		err = errors.Errorf(d.Pos, "stationPrefix not defined")
	}

	// Enforce lower case name
	d.Name = strings.ToLower(strings.TrimSpace(d.Name))
	if err == nil && d.Name == "" {
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
		// Ensure we have an entry present so we don't need to check this in templates
		if d.Components == nil {
			d.Components = &ComponentListEntry{}
		}

		s.dashboards[d.Name] = d

		// Ensure Component exists, require by templates
		if d.Component == nil {
			d.Component = &Component{}
		}
	}

	return errors.Error(d.Pos, err)
}

func (s *state) initFunction(_ Visitor[*state], l *Function) error {
	l.Name = strings.ToLower(l.Name)

	if !value.CalculatorExists(l.Name) {
		return participle.Errorf(l.Pos, "function %q is undefined", l.Name)
	}

	return nil
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

func (s *state) axis(_ Visitor[*state], d *Axis) error {
	var err error

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

	// verify station
	switch {
	case value.Equal(d.Min, d.Max):
		err = errors.Errorf(d.Pos, "Min and Max must not be the same")

	case d.Ticks < 0:
		err = errors.Errorf(d.Pos, "Ticks %d is invalid", d.Ticks)
	}

	return errors.Error(d.Pos, err)
}

func (s *state) gauge(_ Visitor[*state], d *Gauge) error {
	var err error

	// Ensure Component exists, require by templates
	if d.Component == nil {
		d.Component = &Component{}
	}

	if d.Metrics == nil || len(d.Metrics.Metrics) == 0 {
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
	d.Name = s.prefixMetric(d.Name)

	return errors.Error(d.Pos, err)
}

func (s *state) metricExpression(_ Visitor[*state], d *MetricExpression) error {
	var err error

	if d.Offset != "" {
		d.offset, err = time.ParseDuration(d.Offset)
	}

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
		d.patternType = t
		d.prefix = s.prefixMetric("")
	}

	return err
}

func (s *state) unit(_ Visitor[*state], d *units.Unit) error {
	return errors.Error(d.Pos, d.Init())
}

func (s *state) i2c(_ Visitor[*state], d *I2C) error {
	if d.Bus < 1 || d.Device < 1 {
		return participle.Errorf(d.Pos, "invalid i2c address, got (%d:%d)", d.Bus, d.Device)
	}
	return nil
}

func (s *state) sensorList(_ Visitor[*state], d *SensorList) error {
	s.sensors = make(map[string]*Sensor)
	return nil
}

func (s *state) sensor(_ Visitor[*state], d *Sensor) error {
	// Should never occur
	if d.Target == nil {
		return participle.Errorf(d.Pos, "target is required")
	}

	if d.Target.Unit != nil {
		return participle.Errorf(d.Target.Unit.Pos, "unit is invalid as a target for sensors")
	}

	d.Device = strings.TrimSpace(d.Device)
	if d.Device == "" {
		return participle.Errorf(d.Pos, "no device defined")
	}

	// Check Target is unique within the station
	if e, exists := s.sensors[d.Target.Name]; exists {
		return participle.Errorf(d.Pos, "sensor %q already defined at %s", d.Target.Name, e.Pos)
	}
	s.sensors[d.Target.Name] = d

	return nil
}

func (s *state) serial(_ Visitor[*state], d *Serial) error {
	d.Port = strings.TrimSpace(d.Port)
	if d.Port == "" {
		return participle.Errorf(d.Pos, "no serial port defined")
	}

	// TODO define a common list of baud rates here?
	if d.BaudRate < 300 {
		return participle.Errorf(d.Pos, "Invalid baud rate")
	}

	return nil
}
