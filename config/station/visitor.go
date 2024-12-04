package station

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Visitor[T any] interface {
	Axis(*Axis) error
	Calculation(*Calculation) error
	CalculationList(*CalculationList) error
	Component(*Component) error
	ComponentList(*ComponentList) error
	ComponentListEntry(*ComponentListEntry) error
	Container(*Container) error
	CronTab(*time.CronTab) error
	Current(*Current) error
	Dashboard(*Dashboard) error
	DashboardList(*DashboardList) error
	Expression(*Expression) error
	Function(*Function) error
	I2C(d *I2C) error
	Gauge(*Gauge) error
	Http(*Http) error
	HttpFormat(*HttpFormat) error
	Load(*Load) error
	Location(*location.Location) error
	LocationExpression(*LocationExpression) error
	Metric(*Metric) error
	MetricExpression(*MetricExpression) error
	MetricList(*MetricList) error
	MetricPattern(*MetricPattern) error
	MultiValue(*MultiValue) error
	Publisher(*Publisher) error
	Sensor(*Sensor) error
	SensorList(*SensorList) error
	Serial(*Serial) error
	Station(*Station) error
	Stations(*Stations) error
	Text(*Text) error
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
	axis               func(Visitor[T], *Axis) error
	calculation        func(Visitor[T], *Calculation) error
	calculationList    func(Visitor[T], *CalculationList) error
	component          func(Visitor[T], *Component) error
	componentList      func(Visitor[T], *ComponentList) error
	componentListEntry func(Visitor[T], *ComponentListEntry) error
	container          func(Visitor[T], *Container) error
	crontab            func(Visitor[T], *time.CronTab) error
	current            func(Visitor[T], *Current) error
	dashboard          func(Visitor[T], *Dashboard) error
	dashboardList      func(Visitor[T], *DashboardList) error
	expression         func(Visitor[T], *Expression) error
	function           func(Visitor[T], *Function) error
	gauge              func(Visitor[T], *Gauge) error
	http               func(Visitor[T], *Http) error
	httpFormat         func(Visitor[T], *HttpFormat) error
	i2c                func(Visitor[T], *I2C) error
	load               func(Visitor[T], *Load) error
	location           func(Visitor[T], *location.Location) error
	locationExpression func(Visitor[T], *LocationExpression) error
	metric             func(Visitor[T], *Metric) error
	metricExpression   func(Visitor[T], *MetricExpression) error
	metricList         func(Visitor[T], *MetricList) error
	metricPattern      func(Visitor[T], *MetricPattern) error
	multiValue         func(Visitor[T], *MultiValue) error
	publisher          func(Visitor[T], *Publisher) error
	sensor             func(Visitor[T], *Sensor) error
	sensorList         func(Visitor[T], *SensorList) error
	serial             func(Visitor[T], *Serial) error
	station            func(Visitor[T], *Station) error
	stations           func(Visitor[T], *Stations) error
	text               func(Visitor[T], *Text) error
	unit               func(Visitor[T], *units.Unit) error
	useFirst           func(Visitor[T], *UseFirst) error
	value              func(Visitor[T], *Value) error
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

func (c *visitor[T]) CronTab(d *time.CronTab) error {
	var err error
	if d != nil {
		if c.crontab != nil {
			err = c.crontab(c, d)
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}

func (c *visitor[T]) Location(d *location.Location) error {
	var err error
	if d != nil && c.location != nil {
		err = c.location(c, d)
		if util.IsVisitorStop(err) {
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
			if util.IsVisitorStop(err) {
				return nil
			}
		}

		err = errors.Error(d.Pos, err)
	}
	return err
}
