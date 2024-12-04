package station

import (
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/config/util/location"
	time2 "github.com/peter-mount/piweather.center/config/util/time"
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
		CalculationList(initCalculationList).
		Container(initContainer).
		CronTab(initCronTab).
		Dashboard(initDashboard).
		DashboardList(initDashboardList).
		Gauge(initGauge).
		I2C(initI2c).
		Location(initLocation).
		Metric(initMetric).
		MetricExpression(initMetricExpression).
		MetricPattern(initMetricPattern).
		Sensor(initSensor).
		SensorList(initSensorList).
		Serial(initSerial).
		Station(initStation).
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
	stationId     string                  // copy of the stationId being processed
	stationPrefix string                  // stationId + "."
	sensorPrefix  string                  // sensorId + "."
	stationIds    map[string]*Station     // map of Stations
	calculations  map[string]*Calculation // map of calculations within a Station
	dashboards    map[string]*Dashboard   // map of Dashboards within a Station
	sensors       map[string]*Sensor      // map of Sensors within a station
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

func initCronTab(_ Visitor[*initState], d *time2.CronTab) error {
	return errors.Error(d.Pos, d.Init())
}

func initUnit(_ Visitor[*initState], d *units.Unit) error {
	return errors.Error(d.Pos, d.Init())
}
