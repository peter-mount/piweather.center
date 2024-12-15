package station

import (
	"github.com/peter-mount/piweather.center/config/util/command"
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
)

type Builder[T any] interface {
	Axis(func(Visitor[T], *Axis) error) Builder[T]
	Calculation(func(Visitor[T], *Calculation) error) Builder[T]
	Command(f func(Visitor[T], command.Command) error) Builder[T]
	Component(func(Visitor[T], *Component) error) Builder[T]
	ComponentList(func(Visitor[T], *ComponentList) error) Builder[T]
	ComponentListEntry(func(Visitor[T], *ComponentListEntry) error) Builder[T]
	Container(func(Visitor[T], *Container) error) Builder[T]
	CronTab(func(Visitor[T], time.CronTab) error) Builder[T]
	Current(func(Visitor[T], *Current) error) Builder[T]
	Dashboard(func(Visitor[T], *Dashboard) error) Builder[T]
	Ephemeris(func(Visitor[T], *Ephemeris) error) Builder[T]
	EphemerisSchedule(func(Visitor[T], *EphemerisSchedule) error) Builder[T]
	EphemerisTarget(func(Visitor[T], *EphemerisTarget) error) Builder[T]
	EphemerisTargetOption(func(Visitor[T], *EphemerisTargetOption) error) Builder[T]
	Expression(func(Visitor[T], *Expression) error) Builder[T]
	ExpressionAtom(f func(Visitor[T], *ExpressionAtom) error) Builder[T]
	ExpressionLevel1(f func(Visitor[T], *ExpressionLevel1) error) Builder[T]
	ExpressionLevel2(f func(Visitor[T], *ExpressionLevel2) error) Builder[T]
	ExpressionLevel3(f func(Visitor[T], *ExpressionLevel3) error) Builder[T]
	ExpressionLevel4(f func(Visitor[T], *ExpressionLevel4) error) Builder[T]
	ExpressionLevel5(f func(Visitor[T], *ExpressionLevel5) error) Builder[T]
	Function(func(Visitor[T], *Function) error) Builder[T]
	Gauge(func(Visitor[T], *Gauge) error) Builder[T]
	Http(func(Visitor[T], *Http) error) Builder[T]
	HttpFormat(func(Visitor[T], *HttpFormat) error) Builder[T]
	I2C(func(Visitor[T], *I2C) error) Builder[T]
	Load(func(Visitor[T], *Load) error) Builder[T]
	Location(func(Visitor[T], *location.Location) error) Builder[T]
	LocationExpression(func(Visitor[T], *LocationExpression) error) Builder[T]
	Metric(func(Visitor[T], *Metric) error) Builder[T]
	MetricExpression(func(Visitor[T], *MetricExpression) error) Builder[T]
	MetricList(func(Visitor[T], *MetricList) error) Builder[T]
	MetricPattern(func(Visitor[T], *MetricPattern) error) Builder[T]
	MultiValue(func(Visitor[T], *MultiValue) error) Builder[T]
	Publisher(func(Visitor[T], *Publisher) error) Builder[T]
	Sensor(func(Visitor[T], *Sensor) error) Builder[T]
	Serial(func(Visitor[T], *Serial) error) Builder[T]
	SourceParameter(func(Visitor[T], *SourceParameter) error) Builder[T]
	SourceParameterList(func(Visitor[T], *SourceParameterList) error) Builder[T]
	SourceParameterListEntry(func(Visitor[T], *SourceParameterListEntry) error) Builder[T]
	SourcePath(func(Visitor[T], *SourcePath) error) Builder[T]
	SourceWithin(func(Visitor[T], *SourceWithin) error) Builder[T]
	Station(func(Visitor[T], *Station) error) Builder[T]
	StationEntry(func(Visitor[T], *StationEntry) error) Builder[T]
	StationEntryList(func(Visitor[T], *StationEntryList) error) Builder[T]
	Stations(func(Visitor[T], *Stations) error) Builder[T]
	Task(func(Visitor[T], *Task) error) Builder[T]
	TaskCondition(func(Visitor[T], *TaskCondition) error) Builder[T]
	Tasks(func(Visitor[T], *Tasks) error) Builder[T]
	Text(func(Visitor[T], *Text) error) Builder[T]
	TimeZone(func(Visitor[T], *time.TimeZone) error) Builder[T]
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

func (b *builder[T]) Command(f func(Visitor[T], command.Command) error) Builder[T] {
	b.command = f
	return b
}

func (b *builder[T]) CronTab(f func(Visitor[T], time.CronTab) error) Builder[T] {
	b.crontab = f
	return b
}

func (b *builder[T]) Location(f func(Visitor[T], *location.Location) error) Builder[T] {
	b.location = f
	return b
}

func (b *builder[T]) TimeZone(f func(Visitor[T], *time.TimeZone) error) Builder[T] {
	b.timeZone = f
	return b
}

func (b *builder[T]) Unit(f func(Visitor[T], *units.Unit) error) Builder[T] {
	b.unit = f
	return b
}
