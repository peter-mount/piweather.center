package station

import (
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Builder[T any] interface {
	Axis(func(Visitor[T], *Axis) error) Builder[T]
	Calculation(func(Visitor[T], *Calculation) error) Builder[T]
	CalculationList(func(Visitor[T], *CalculationList) error) Builder[T]
	Component(func(Visitor[T], *Component) error) Builder[T]
	ComponentList(func(Visitor[T], *ComponentList) error) Builder[T]
	ComponentListEntry(func(Visitor[T], *ComponentListEntry) error) Builder[T]
	Container(func(Visitor[T], *Container) error) Builder[T]
	CronTab(func(Visitor[T], *time.CronTab) error) Builder[T]
	Current(func(Visitor[T], *Current) error) Builder[T]
	Dashboard(func(Visitor[T], *Dashboard) error) Builder[T]
	DashboardList(func(Visitor[T], *DashboardList) error) Builder[T]
	Expression(func(Visitor[T], *Expression) error) Builder[T]
	Forecast(func(Visitor[T], *Forecast) error) Builder[T]
	Function(func(Visitor[T], *Function) error) Builder[T]
	Gauge(func(Visitor[T], *Gauge) error) Builder[T]
	Load(func(Visitor[T], *Load) error) Builder[T]
	Location(func(Visitor[T], *location.Location) error) Builder[T]
	LocationExpression(func(Visitor[T], *LocationExpression) error) Builder[T]
	Metric(func(Visitor[T], *Metric) error) Builder[T]
	MetricExpression(func(Visitor[T], *MetricExpression) error) Builder[T]
	MetricList(func(Visitor[T], *MetricList) error) Builder[T]
	MetricPattern(func(Visitor[T], *MetricPattern) error) Builder[T]
	MultiValue(func(Visitor[T], *MultiValue) error) Builder[T]
	Station(func(Visitor[T], *Station) error) Builder[T]
	Stations(func(Visitor[T], *Stations) error) Builder[T]
	Text(func(Visitor[T], *Text) error) Builder[T]
	Unit(func(Visitor[T], *units.Unit) error) Builder[T]
	UseFirst(func(Visitor[T], *UseFirst) error) Builder[T]
	Value(func(Visitor[T], *Value) error) Builder[T]
	Build() Visitor[T]
}

type builder[T any] struct {
	common[T]
}

func NewBuilder[T any]() Builder[T] {
	return &builder[T]{}
}

func (b *builder[T]) Build() Visitor[T] {
	return &visitor[T]{common: b.common}
}

func (b *builder[T]) Axis(f func(Visitor[T], *Axis) error) Builder[T] {
	b.axis = f
	return b
}

func (b *builder[T]) Calculation(f func(Visitor[T], *Calculation) error) Builder[T] {
	b.calculation = f
	return b
}

func (b *builder[T]) CalculationList(f func(Visitor[T], *CalculationList) error) Builder[T] {
	b.calculationList = f
	return b
}

func (b *builder[T]) Component(f func(Visitor[T], *Component) error) Builder[T] {
	b.component = f
	return b
}

func (b *builder[T]) ComponentList(f func(Visitor[T], *ComponentList) error) Builder[T] {
	b.componentList = f
	return b
}

func (b *builder[T]) ComponentListEntry(f func(Visitor[T], *ComponentListEntry) error) Builder[T] {
	b.componentListEntry = f
	return b
}

func (b *builder[T]) Container(f func(Visitor[T], *Container) error) Builder[T] {
	b.container = f
	return b
}

func (b *builder[T]) CronTab(f func(Visitor[T], *time.CronTab) error) Builder[T] {
	b.crontab = f
	return b
}

func (b *builder[T]) Current(f func(Visitor[T], *Current) error) Builder[T] {
	b.current = f
	return b
}

func (b *builder[T]) Dashboard(f func(Visitor[T], *Dashboard) error) Builder[T] {
	b.dashboard = f
	return b
}

func (b *builder[T]) DashboardList(f func(Visitor[T], *DashboardList) error) Builder[T] {
	b.dashboardList = f
	return b
}

func (b *builder[T]) Expression(f func(Visitor[T], *Expression) error) Builder[T] {
	b.expression = f
	return b
}

func (b *builder[T]) Forecast(f func(Visitor[T], *Forecast) error) Builder[T] {
	b.forecast = f
	return b
}

func (b *builder[T]) Function(f func(Visitor[T], *Function) error) Builder[T] {
	b.function = f
	return b
}

func (b *builder[T]) Gauge(f func(Visitor[T], *Gauge) error) Builder[T] {
	b.gauge = f
	return b
}

func (b *builder[T]) Load(f func(Visitor[T], *Load) error) Builder[T] {
	b.load = f
	return b
}

func (b *builder[T]) Location(f func(Visitor[T], *location.Location) error) Builder[T] {
	b.location = f
	return b
}

func (b *builder[T]) LocationExpression(f func(Visitor[T], *LocationExpression) error) Builder[T] {
	b.locationExpression = f
	return b
}

func (b *builder[T]) Metric(f func(Visitor[T], *Metric) error) Builder[T] {
	b.metric = f
	return b
}

func (b *builder[T]) MetricExpression(f func(Visitor[T], *MetricExpression) error) Builder[T] {
	b.metricExpression = f
	return b
}

func (b *builder[T]) MetricList(f func(Visitor[T], *MetricList) error) Builder[T] {
	b.metricList = f
	return b
}

func (b *builder[T]) MetricPattern(f func(Visitor[T], *MetricPattern) error) Builder[T] {
	b.metricPattern = f
	return b
}

func (b *builder[T]) MultiValue(f func(Visitor[T], *MultiValue) error) Builder[T] {
	b.multiValue = f
	return b
}

func (b *builder[T]) Station(f func(Visitor[T], *Station) error) Builder[T] {
	b.station = f
	return b
}

func (b *builder[T]) Stations(f func(Visitor[T], *Stations) error) Builder[T] {
	b.stations = f
	return b
}

func (b *builder[T]) Text(f func(Visitor[T], *Text) error) Builder[T] {
	b.text = f
	return b
}

func (b *builder[T]) Unit(f func(Visitor[T], *units.Unit) error) Builder[T] {
	b.unit = f
	return b
}

func (b *builder[T]) UseFirst(f func(Visitor[T], *UseFirst) error) Builder[T] {
	b.useFirst = f
	return b
}

func (b *builder[T]) Value(f func(Visitor[T], *Value) error) Builder[T] {
	b.value = f
	return b
}
