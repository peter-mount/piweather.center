package angle

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/soniakeys/unit"
)

func init() {
	packages.RegisterPackage(&Angle{})
}

type Angle struct{}

func (_ Angle) AngleFromDeg(d float64) unit.Angle {
	return unit.AngleFromDeg(d)
}

func (_ Angle) AngleFromMin(d float64) unit.Angle {
	return unit.AngleFromMin(d)
}

func (_ Angle) AngleFromSec(d float64) unit.Angle {
	return unit.AngleFromSec(d)
}

func (_ Angle) NewAngle(neg byte, d, m int, s float64) unit.Angle {
	return unit.NewAngle(neg, d, m, s)
}

func (_ Angle) HourAngleFromHour(d float64) unit.HourAngle {
	return unit.HourAngleFromHour(d)
}

func (_ Angle) HourAngleFromMin(d float64) unit.HourAngle {
	return unit.HourAngleFromMin(d)
}

func (_ Angle) HourAngleFromSec(d float64) unit.HourAngle {
	return unit.HourAngleFromSec(d)
}

func (_ Angle) NewHourAngle(neg byte, d, m int, s float64) unit.HourAngle {
	return unit.NewHourAngle(neg, d, m, s)
}

func (_ Angle) RAFromRad(d float64) unit.RA {
	return unit.RAFromRad(d)
}

func (_ Angle) RAFromDeg(d float64) unit.RA {
	return unit.RAFromDeg(d)
}

func (_ Angle) RAFromHour(d float64) unit.RA {
	return unit.RAFromHour(d)
}

func (_ Angle) RAFromMin(d float64) unit.RA {
	return unit.RAFromMin(d)
}

func (_ Angle) RAFromSec(d float64) unit.RA {
	return unit.RAFromSec(d)
}

func (_ Angle) NewRA(d, m int, s float64) unit.RA {
	return unit.NewRA(d, m, s)
}
