package station

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
	"strings"
)

func NewParser() util.Parser[Stations] {
	return util.NewParser[Stations](nil, nil, stationInit)
}

var (
	initVisitor = NewBuilder[*initState]().
		Axis(initAxis).
		Calculation(initCalculation).
		Container(initContainer).
		CronTab(initCronTab).
		Dashboard(initDashboard).
		Ephemeris(initEphemeris).
		EphemerisTarget(initEphemerisTarget).
		EphemerisTargetOption(initEphemerisTargetOption).
		Gauge(initGauge).
		Http(initHttp).
		I2C(initI2c).
		Location(initLocation).
		Metric(initMetric).
		MetricExpression(initMetricExpression).
		MetricPattern(initMetricPattern).
		TimeZone(initTimeZone).
		Sensor(initSensor).
		Serial(initSerial).
		SourceParameter(initSourceParameter).
		SourceWithin(initSourceWithin).
		SourcePath(initSourcePath).
		Station(initStation).
		StationEntryList(initStationEntryList).
		Stations(initStations).
		Unit(initUnit).
		Value(initValue).
		Build()
)

func stationInit(q *Stations, err error) (*Stations, error) {

	if err == nil {
		err = initVisitor.Clone().
			Set(&initState{}).
			Stations(q)
	}

	return q, err
}

type initState struct {
	stationId        string                      // copy of the stationId being processed
	stationPrefix    string                      // stationId + "."
	sensorPrefix     string                      // sensorId + "."
	stationIds       map[string]*Station         // map of Stations, for id uniqueness
	calculations     map[string]lexer.Position   // map of calculations within a Station, for target uniqueness
	dashboards       map[string]lexer.Position   // map of Dashboards within a Station, for id uniqueness
	sensors          map[string]lexer.Position   // map of Sensors within a station, for sensorPrefix uniqueness
	sensorParameters map[string]*SourceParameter // Map of SourceParameter's used to ensure target metrics are unique
	sourcePath       []string                    // Prefix for source path, used with SourceWithin
	ephemeris        *Ephemeris                  // Ephemeris being scanned
	ephemerisTarget  *EphemerisTarget            // EphemerisTarget being scanned
}

func (s *initState) prefixMetric(m string) string {
	return s.stationPrefix + s.sensorPrefix + m
}

func initLocation(v Visitor[*initState], d *location.Location) error {
	s := v.Get()

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

func initCronTab(_ Visitor[*initState], d *time.CronTab) error {
	return errors.Error(d.Pos, d.Init())
}

func initTimeZone(_ Visitor[*initState], d *time.TimeZone) error {
	return errors.Error(d.Pos, d.Init())
}

func initUnit(_ Visitor[*initState], d *units.Unit) error {
	return errors.Error(d.Pos, d.Init())
}
