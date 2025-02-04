package station

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/piweather.center/config/util/command"
	"github.com/peter-mount/piweather.center/config/util/location"
	time2 "github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
	"reflect"
	"strings"
)

// ParseAndInclude will set the Pos field within a script's structure to a single specific value.
// This is used when we generate code based on part of a script and the
// generated code then reports any errors based on the original
// and not the generated code.
//
// e.g. calculate from will generate calculations saving boilerplate
func (s *initState) ParseAndInclude(pos lexer.Position, script []string) error {
	// Add station preamble based on the station being processed before the
	// generated script
	lines := []string{fmt.Sprintf("station( %q", s.station.Name)}
	if s.station.Location != nil {
		lines = append(lines, fmt.Sprintf("location(%q %q %q %.0f)",
			s.station.Location.Name,
			s.station.Location.Latitude,
			s.station.Location.Longitude,
			s.station.Location.Altitude,
		))
	}

	if s.station.TimeZone != nil {
		lines = append(lines, fmt.Sprintf("timezone(%q)", s.station.TimeZone.TimeZone))
	}

	lines = append(lines, script...)
	lines = append(lines, ")")

	newStations, err := s.parser.ParseString(pos.Filename, strings.Join(lines, "\n"))

	if err == nil {
		// Ensure entries have the Position passed to us
		err = pseudoVisitor.Clone().
			Set(reflect.ValueOf(pos)).
			Stations(newStations)
	}

	if err == nil {
		// Add any new entries to newEntries ready to merge post init
		err = pseudoAddNewStationEntryVisitor.Clone().
			Set(s).
			Stations(newStations)
	}

	return err
}

var (
	pseudoVisitor = NewBuilder[reflect.Value]().
			Axis(pseudoSetPositionStruct[Axis]).
			Calculation(pseudoSetPositionStruct[Calculation]).
			CalculateFrom(pseudoSetPositionStruct[CalculateFrom]).
			Command(pseudoSetPositionInterface[command.Command]).
			Component(pseudoSetPositionStruct[Component]).
			ComponentList(pseudoSetPositionStruct[ComponentList]).
			ComponentListEntry(pseudoSetPositionStruct[ComponentListEntry]).
			Container(pseudoSetPositionStruct[Container]).
			CronTab(pseudoSetPositionInterface[time2.CronTab]).
			Current(pseudoSetPositionStruct[Current]).
			Dashboard(pseudoSetPositionStruct[Dashboard]).
			Ephemeris(pseudoSetPositionStruct[Ephemeris]).
			EphemerisSchedule(pseudoSetPositionStruct[EphemerisSchedule]).
			EphemerisTarget(pseudoSetPositionStruct[EphemerisTarget]).
			EphemerisTargetOption(pseudoSetPositionStruct[EphemerisTargetOption]).
			Expression(pseudoSetPositionStruct[Expression]).
			ExpressionAtom(pseudoSetPositionStruct[ExpressionAtom]).
			ExpressionLevel1(pseudoSetPositionStruct[ExpressionLevel1]).
			ExpressionLevel2(pseudoSetPositionStruct[ExpressionLevel2]).
			ExpressionLevel3(pseudoSetPositionStruct[ExpressionLevel3]).
			ExpressionLevel4(pseudoSetPositionStruct[ExpressionLevel4]).
			ExpressionLevel5(pseudoSetPositionStruct[ExpressionLevel5]).
			Function(pseudoSetPositionStruct[Function]).
			Gauge(pseudoSetPositionStruct[Gauge]).
			Http(pseudoSetPositionStruct[Http]).
			HttpFormat(pseudoSetPositionStruct[HttpFormat]).
			I2C(pseudoSetPositionStruct[I2C]).
			Load(pseudoSetPositionStruct[Load]).
			Location(pseudoSetPositionStruct[location.Location]).
			LocationExpression(pseudoSetPositionStruct[LocationExpression]).
			Metric(pseudoSetPositionStruct[Metric]).
			MetricExpression(pseudoSetPositionStruct[MetricExpression]).
			MetricList(pseudoSetPositionStruct[MetricList]).
			MetricPattern(pseudoSetPositionStruct[MetricPattern]).
			MultiValue(pseudoSetPositionStruct[MultiValue]).
			Publisher(pseudoSetPositionStruct[Publisher]).
			Sensor(pseudoSetPositionStruct[Sensor]).
			Serial(pseudoSetPositionStruct[Serial]).
			SourceParameter(pseudoSetPositionStruct[SourceParameter]).
			SourceParameterList(pseudoSetPositionStruct[SourceParameterList]).
			SourceParameterListEntry(pseudoSetPositionStruct[SourceParameterListEntry]).
			SourcePath(pseudoSetPositionStruct[SourcePath]).
			SourceWithin(pseudoSetPositionStruct[SourceWithin]).
			Station(pseudoSetPositionStruct[Station]).
			StationEntry(pseudoSetPositionStruct[StationEntry]).
			StationEntryList(pseudoSetPositionStruct[StationEntryList]).
			Stations(pseudoSetPositionStruct[Stations]).
			Task(pseudoSetPositionStruct[Task]).
			TaskCondition(pseudoSetPositionStruct[TaskCondition]).
			Tasks(pseudoSetPositionStruct[Tasks]).
			Text(pseudoSetPositionStruct[Text]).
			TimeZone(pseudoSetPositionStruct[time2.TimeZone]).
			Unit(pseudoSetPositionStruct[units.Unit]).
			UseFirst(pseudoSetPositionStruct[UseFirst]).
			Value(pseudoSetPositionStruct[Value]).
			Build()

	pseudoAddNewStationEntryVisitor = NewBuilder[*initState]().
		//
		Calculation(func(v Visitor[*initState], d *Calculation) error {
			return v.Get().
				addStationEntry(&StationEntry{
					Pos:         d.Pos,
					Calculation: d,
				})
		}).
		Dashboard(func(v Visitor[*initState], d *Dashboard) error {
			return v.Get().
				addStationEntry(&StationEntry{
					Pos:       d.Pos,
					Dashboard: d,
				})
		}).
		Ephemeris(func(v Visitor[*initState], d *Ephemeris) error {
			return v.Get().
				addStationEntry(&StationEntry{
					Pos:       d.Pos,
					Ephemeris: d,
				})
		}).
		Sensor(func(v Visitor[*initState], d *Sensor) error {
			return v.Get().
				addStationEntry(&StationEntry{
					Pos:    d.Pos,
					Sensor: d,
				})
		}).
		Tasks(func(v Visitor[*initState], d *Tasks) error {
			return v.Get().
				addStationEntry(&StationEntry{
					Pos:   d.Pos,
					Tasks: d,
				})
		}).
		Build()
)

// pseudoSetPositionStruct handles setting Pos against a Struct
func pseudoSetPositionStruct[T any](v Visitor[reflect.Value], d *T) error {
	return pseudoSetPositionImpl(v, d)
}

// pseudoSetPositionInterface handles setting Pos against an Interface.
// Note: The underlying struct implementing the interface must have Pos as a field.
func pseudoSetPositionInterface[T any](v Visitor[reflect.Value], d T) error {
	return pseudoSetPositionImpl(v, d)
}

func pseudoSetPositionImpl(v Visitor[reflect.Value], d any) error {
	tv := reflect.ValueOf(d)
	t := tv.Type()
	if t == nil {
		return nil
	}

	val := tv.Elem().FieldByName("Pos")
	if val.IsValid() && val.CanSet() {
		val.Set(v.Get())
	}
	return nil
}
