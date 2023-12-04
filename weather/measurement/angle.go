package measurement

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

func init() {
	Radian = value.NewUnit("Radian", "Radians", " rad", 4)
	Degree = value.NewUnit("Degree", "Degrees", "°", 1)
	ArcMinute = value.NewUnit("ArcMinute", "Arc Minute", "'", 3)
	ArcSecond = value.NewUnit("ArcSecond", "Arc Second", "\"", 3)
	Gradian = value.NewUnit("Gradian", "Gradian", " grad", 3)
	HourAngle = value.NewUnit("HourAngle", "Hour Angle", " ha", 3)
	Turn = value.NewUnit("Turn", "Turn", " turn", 6)

	// Turn is the default unit
	value.NewBasicBiTransform(Turn, Degree, 360)
	value.NewBasicBiTransform(Turn, Radian, 2.0*math.Pi)
	value.NewBasicBiTransform(Turn, ArcMinute, 360*60)
	value.NewBasicBiTransform(Turn, ArcSecond, 360*3600)
	value.NewBasicBiTransform(Turn, Gradian, 400)
	value.NewBasicBiTransform(Turn, HourAngle, 24)

	// Common transforms to save on going via Turn
	value.NewBasicBiTransform(Degree, Radian, math.Pi/180.0)
	value.NewBasicBiTransform(Degree, ArcMinute, 60.0)
	value.NewBasicBiTransform(ArcMinute, ArcSecond, 60.0)
	value.NewBasicBiTransform(Degree, ArcSecond, 3600.0)
	value.NewBasicBiTransform(HourAngle, Degree, 15.0)

	// Ensure all others exist
	Angle = value.NewGroup("Angle", Turn, Radian, Degree, ArcMinute, ArcSecond, Gradian, HourAngle)
}

var (
	// Angle value.group of all angular value.Unit's
	Angle *value.Group
	// Radian is determined by the circumference of a circle that is equal in
	// length to the radius of the circle (n = 2π = 6.283...). It is the angle
	// subtended by an arc of a circle that has the same length as the circle's
	// radius.
	//
	// The symbol for radian is rad. One turn is 2π radians, and one radian is
	// 180°/π, or about 57.2958 degrees.
	//
	// Often, particularly in mathematical texts, one radian is assumed to equal
	// one, resulting in the unit rad being omitted.
	// The radian is used in virtually all mathematical work beyond simple
	// practical geometry, due, for example, to the pleasing and "natural"
	// properties that the trigonometric functions display when their arguments
	// are in radians. The radian is the (derived) unit of angular measurement
	// in the SI.
	Radian *value.Unit
	// Degree denoted by a small superscript circle (°), is 1/360 of a turn,
	// so one turn is 360°.
	//
	// One advantage of this old sexagesimal subunit is that many angles common
	// in simple geometry are measured as a whole number of degrees.
	// Fractions of a degree may be written in normal decimal notation
	// (e.g. 3.5° for three and a half degrees),
	// but the "minute" and "second" sexagesimal subunits of the
	// "degree–minute–second" system are also in use, especially for
	// geographical coordinates and in astronomy and ballistics (n = 360)
	Degree *value.Unit
	// ArcMinute The minute of arc (or MOA, arcminute, or just minute)
	// is 1/60 of a degree = 1/21,600 turn.
	// It is denoted by a single prime ( ′ ).
	//
	// For example, 3° 30′ is equal to 3 × 60 + 30 = 210 minutes
	// or 3 + 30/60 = 3.5 degrees.
	//
	// A mixed format with decimal fractions is also sometimes used,
	// e.g. 3° 5.72′ = 3 + 5.72/60 degrees.
	//
	// A nautical mile was historically defined as an arcminute along
	// a great circle of the Earth. (n = 21,600).
	ArcMinute *value.Unit
	ArcSecond *value.Unit
	Gradian   *value.Unit
	HourAngle *value.Unit
	Turn      *value.Unit
)
