package station

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/command"
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Visitor[T any] interface {
	Axis(*Axis) error
	Calculation(*Calculation) error
	Command(command.Command) error
	Component(*Component) error
	ComponentList(*ComponentList) error
	ComponentListEntry(*ComponentListEntry) error
	Container(*Container) error
	CronTab(time.CronTab) error
	Current(*Current) error
	Dashboard(*Dashboard) error
	Ephemeris(*Ephemeris) error
	EphemerisSchedule(*EphemerisSchedule) error
	EphemerisTarget(*EphemerisTarget) error
	EphemerisTargetOption(*EphemerisTargetOption) error
	Expression(*Expression) error
	ExpressionAtom(*ExpressionAtom) error
	ExpressionLevel1(b *ExpressionLevel1) error
	ExpressionLevel2(b *ExpressionLevel2) error
	ExpressionLevel3(b *ExpressionLevel3) error
	ExpressionLevel4(b *ExpressionLevel4) error
	ExpressionLevel5(b *ExpressionLevel5) error
	Function(*Function) error
	Gauge(*Gauge) error
	Http(*Http) error
	HttpFormat(*HttpFormat) error
	I2C(*I2C) error
	Load(*Load) error
	Location(*location.Location) error
	LocationExpression(*LocationExpression) error
	Metric(*Metric) error
	MetricExpression(*MetricExpression) error
	MetricList(*MetricList) error
	MetricPattern(*MetricPattern) error
	MultiValue(*MultiValue) error
	Publisher(*Publisher) error
	Rtl433(*Rtl433) error
	Sensor(*Sensor) error
	Serial(*Serial) error
	SourceParameter(*SourceParameter) error
	SourceParameterList(*SourceParameterList) error
	SourceParameterListEntry(*SourceParameterListEntry) error
	SourcePath(*SourcePath) error
	SourceWithin(d *SourceWithin) error
	Station(*Station) error
	StationEntry(*StationEntry) error
	StationEntryList(*StationEntryList) error
	Stations(*Stations) error
	Task(*Task) error
	TaskCondition(*TaskCondition) error
	Tasks(*Tasks) error
	Text(*Text) error
	TimeZone(*time.TimeZone) error
	Unit(*units.Unit) error
	UseFirst(*UseFirst) error
	Value(*Value) error

	// Get returns the attached value T
	Get() T
	// Set the attached value T
	Set(T) Visitor[T]
	// Clone the visitor - used when you need to pass an alternate data object in multiple goroutines.
	// Note: The attached value T is not included in the clone
	Clone() Visitor[T]
}

type visitor[T any] struct {
	common[T]
	data T
}

type common[T any] struct {
	axis                     func(Visitor[T], *Axis) error
	calculation              func(Visitor[T], *Calculation) error
	command                  func(Visitor[T], command.Command) error
	component                func(Visitor[T], *Component) error
	componentList            func(Visitor[T], *ComponentList) error
	componentListEntry       func(Visitor[T], *ComponentListEntry) error
	container                func(Visitor[T], *Container) error
	crontab                  func(Visitor[T], time.CronTab) error
	current                  func(Visitor[T], *Current) error
	dashboard                func(Visitor[T], *Dashboard) error
	ephemeris                func(Visitor[T], *Ephemeris) error
	ephemerisSchedule        func(Visitor[T], *EphemerisSchedule) error
	ephemerisTarget          func(Visitor[T], *EphemerisTarget) error
	ephemerisTargetOption    func(Visitor[T], *EphemerisTargetOption) error
	expression               func(Visitor[T], *Expression) error
	expressionAtom           func(Visitor[T], *ExpressionAtom) error
	expressionLevel1         func(Visitor[T], *ExpressionLevel1) error
	expressionLevel2         func(Visitor[T], *ExpressionLevel2) error
	expressionLevel3         func(Visitor[T], *ExpressionLevel3) error
	expressionLevel4         func(Visitor[T], *ExpressionLevel4) error
	expressionLevel5         func(Visitor[T], *ExpressionLevel5) error
	function                 func(Visitor[T], *Function) error
	gauge                    func(Visitor[T], *Gauge) error
	http                     func(Visitor[T], *Http) error
	httpFormat               func(Visitor[T], *HttpFormat) error
	i2c                      func(Visitor[T], *I2C) error
	load                     func(Visitor[T], *Load) error
	location                 func(Visitor[T], *location.Location) error
	locationExpression       func(Visitor[T], *LocationExpression) error
	metric                   func(Visitor[T], *Metric) error
	metricExpression         func(Visitor[T], *MetricExpression) error
	metricList               func(Visitor[T], *MetricList) error
	metricPattern            func(Visitor[T], *MetricPattern) error
	multiValue               func(Visitor[T], *MultiValue) error
	publisher                func(Visitor[T], *Publisher) error
	rtl433                   func(Visitor[T], *Rtl433) error
	sensor                   func(Visitor[T], *Sensor) error
	serial                   func(Visitor[T], *Serial) error
	sourceParameter          func(Visitor[T], *SourceParameter) error
	sourceParameterList      func(Visitor[T], *SourceParameterList) error
	sourceParameterListEntry func(Visitor[T], *SourceParameterListEntry) error
	sourcePath               func(Visitor[T], *SourcePath) error
	sourceWithin             func(Visitor[T], *SourceWithin) error
	station                  func(Visitor[T], *Station) error
	stationEntry             func(Visitor[T], *StationEntry) error
	stationEntryList         func(Visitor[T], *StationEntryList) error
	stations                 func(Visitor[T], *Stations) error
	task                     func(Visitor[T], *Task) error
	taskCondition            func(Visitor[T], *TaskCondition) error
	tasks                    func(Visitor[T], *Tasks) error
	text                     func(Visitor[T], *Text) error
	timeZone                 func(Visitor[T], *time.TimeZone) error
	unit                     func(Visitor[T], *units.Unit) error
	useFirst                 func(Visitor[T], *UseFirst) error
	value                    func(Visitor[T], *Value) error
}

func (c *visitor[T]) Get() T {
	return c.data
}

func (c *visitor[T]) Set(data T) Visitor[T] {
	c.data = data
	return c
}

func (c *visitor[T]) Clone() Visitor[T] {
	return &visitor[T]{common: c.common}
}

func (c *visitor[T]) Command(d command.Command) error {
	var err error
	if d != nil {
		if c.command != nil {
			err = c.command(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Position(), err)
	}
	return err
}

func (c *visitor[T]) CronTab(d time.CronTab) error {
	var err error
	if d != nil {
		if c.crontab != nil {
			err = c.crontab(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Position(), err)
	}
	return err
}

func (c *visitor[T]) Location(d *location.Location) error {
	var err error
	if d != nil && c.location != nil {
		err = c.location(c, d)
		if errors.IsVisitorStop(err) {
			return nil
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) TimeZone(d *time.TimeZone) error {
	var err error
	if d != nil && c.timeZone != nil {
		err = c.timeZone(c, d)
		if errors.IsVisitorStop(err) {
			return nil
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Unit(d *units.Unit) error {
	var err error
	if d != nil {
		if c.unit != nil {
			err = c.unit(c, d)
			if errors.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}
