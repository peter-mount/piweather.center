package station

import (
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Builder[T any] interface {
	Component(func(Visitor[T], *Component) error) Builder[T]
	ComponentList(func(Visitor[T], *ComponentList) error) Builder[T]
	ComponentListEntry(func(Visitor[T], *ComponentListEntry) error) Builder[T]
	Container(func(Visitor[T], *Container) error) Builder[T]
	CronTab(func(Visitor[T], *time.CronTab) error) Builder[T]
	Dashboard(func(Visitor[T], *Dashboard) error) Builder[T]
	DashboardList(func(Visitor[T], *DashboardList) error) Builder[T]
	Gauge(func(Visitor[T], *Gauge) error) Builder[T]
	Location(func(Visitor[T], *location.Location) error) Builder[T]
	Metric(func(Visitor[T], *Metric) error) Builder[T]
	MetricList(func(Visitor[T], *MetricList) error) Builder[T]
	MetricPattern(func(Visitor[T], *MetricPattern) error) Builder[T]
	MultiValue(func(Visitor[T], *MultiValue) error) Builder[T]
	Station(func(Visitor[T], *Station) error) Builder[T]
	Stations(func(Visitor[T], *Stations) error) Builder[T]
	Unit(func(Visitor[T], *units.Unit) error) Builder[T]
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

func (b *builder[T]) Dashboard(f func(Visitor[T], *Dashboard) error) Builder[T] {
	b.dashboard = f
	return b
}

func (b *builder[T]) DashboardList(f func(Visitor[T], *DashboardList) error) Builder[T] {
	b.dashboardList = f
	return b
}

func (b *builder[T]) Gauge(f func(Visitor[T], *Gauge) error) Builder[T] {
	b.gauge = f
	return b
}

func (b *builder[T]) Location(f func(Visitor[T], *location.Location) error) Builder[T] {
	b.location = f
	return b
}

func (b *builder[T]) Metric(f func(Visitor[T], *Metric) error) Builder[T] {
	b.metric = f
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

func (b *builder[T]) Unit(f func(Visitor[T], *units.Unit) error) Builder[T] {
	b.unit = f
	return b
}

func (b *builder[T]) Value(f func(Visitor[T], *Value) error) Builder[T] {
	b.value = f
	return b
}
