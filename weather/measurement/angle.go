package measurement

import (
	"github.com/peter-mount/piweather.center/util/strings"
	"github.com/peter-mount/piweather.center/weather/value"
	"math"
)

func init() {
	Radian = value.NewUnit("Radian", "Radians", " rad", 4, nil)
	Degree = value.NewUnit("Degree", "Degrees", "°", 3, nil)
	ArcMinute = value.NewUnit("ArcMinute", "Arc Minute", "'", 3, nil)
	ArcSecond = value.NewUnit("ArcSecond", "Arc Second", "\"", 3, nil)
	Gradian = value.NewUnit("Gradian", "Gradian", " grad", 3, nil)
	HourAngle = value.NewBoundedUnitF("HourAngle", "Hour Angle", " ha", 3, -12+value.EqualityError, 12.0, nil)
	Turn = value.NewUnit("Turn", "Turn", " turn", 6, nil)

	// RA is 0...24 but 24 is not a valid value hence we set the limit to 24- the equality error
	RA = value.NewBoundedUnitF("RA", "Right Ascension", "", 4, 0.0, 24.0-value.EqualityError, func(f float64) string {
		return strings.DegDMSStringExt(f, false, "", "", 2, 1)
	})

	Declination = value.NewBoundedUnitF("Dec", "Declination", "", 4, -90.0, 90.0, func(f float64) string {
		return strings.DegDMSStringExt(f, false, "+", "-", 2, 1)
	})

	// Longitude is -180...180 but -180 is not a valid value hence we set the limit to -180+ the equality error.
	// Also, this is East positive, West negative
	Longitude = value.NewBoundedUnitF("Longitude", "Longitude", "", 4, -180.0-value.EqualityError, 180.0, func(f float64) string {
		return strings.DegDMSStringExt(f, false, "E", "W", 2, 1)
	})

	Latitude = value.NewBoundedUnitF("Latitude", "Latitude", "", 4, -90.0, 90.0, func(f float64) string {
		return strings.DegDMSStringExt(f, false, "N", "S", 2, 1)
	})

	// Azimuth is -180 < az <= 180 with West positive
	Azimuth = value.NewBoundedUnitF("Azimuth", "Azimuth", "°", 4, -180.0-value.EqualityError, 180.0, func(f float64) string {
		return strings.DegDMSStringExt(f, false, "W", "E", 2, 1)
	})

	// Turn is the default unit
	value.NewBasicBiTransform(Turn, Degree, 360)
	value.NewBasicBiTransform(Turn, Radian, 2.0*math.Pi)
	value.NewBasicBiTransform(Turn, ArcMinute, 360*60)
	value.NewBasicBiTransform(Turn, ArcSecond, 360*3600)
	value.NewBasicBiTransform(Turn, Gradian, 400)
	value.NewBasicBiTransform(Turn, RA, 24)
	value.NewBasicBiTransform(Turn, Declination, 360)

	value.NewBasicBiTransform(Turn, Latitude, 360)

	// HourAngle is -12 < ha <= 12
	value.NewBiTransform(Turn, HourAngle,
		value.BasicTransform(24).Then(degreeToHourAngle),
		value.Of(hourAngleToDegree).Then(value.BasicInverseTransform(24)))

	// Turn<->Longitude is same as Turn->Degree->Longitude and Longitude->Degree->Turn
	// We have to do it this way as Longitude is -180 < lon <= 180
	value.NewBiTransform(Turn, Longitude,
		value.BasicTransform(360).Then(degreeToLongitude),
		value.Of(longitudeToDegree).Then(value.BasicInverseTransform(360)))

	value.NewBiTransform(Turn, Azimuth,
		value.BasicTransform(360).Then(degreeToAzimuth),
		value.Of(azimuthToDegree).Then(value.BasicInverseTransform(360)))

	// Common transforms to save on going via Turn
	value.NewBasicBiTransform(Degree, Radian, math.Pi/180.0)
	value.NewBasicBiTransform(Degree, ArcMinute, 60.0)
	value.NewBasicBiTransform(ArcMinute, ArcSecond, 60.0)
	value.NewBasicBiTransform(Degree, ArcSecond, 3600.0)
	value.NewBasicBiTransform(HourAngle, Degree, 15.0)
	value.NewBasicBiTransform(RA, Degree, 15.0)
	value.NewBiTransform(Degree, Latitude, value.NopTransformer, value.NopTransformer)
	value.NewBiTransform(Degree, Longitude, degreeToLongitude, longitudeToDegree)

	value.NewBiTransform(Degree, Azimuth, degreeToAzimuth, azimuthToDegree)

	// Ensure all others exist
	Angle = value.NewGroup("Angle", Turn, Radian, Degree, ArcMinute, ArcSecond, Gradian, HourAngle, RA, Declination)
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
	// HourAngle is the distance, west positive, in hours
	HourAngle *value.Unit
	Turn      *value.Unit
	// RA Right Ascension of an object. This is a value between 0...24 and is formatted in hours:minutes:seconds
	RA *value.Unit
	// Declination of an object. This is a value between -90...90 and is formatted in +degrees:minutes:seconds
	Declination *value.Unit
	Latitude    *value.Unit
	Longitude   *value.Unit
	// Azimuth measured westwards from the South.
	//
	// It should be noted that Navigators and Meteorologists count the compass direction, or azimuth,
	// from the North (0), East (90), South (180) and West (270).
	// But Astronomers measure the Azimuth from the South, because Hour Angles are measured from the South.
	//
	// Ref: Jean Meeus, c12 p87, Astronomical Algorithms, 1st edition 1991
	//
	// Ref: William Chauvenet, p20, A Manual of Spherical and Practical Astronomy, 5th edition 1891
	Azimuth *value.Unit
)

func AngleRoundDown(v value.Value) value.Value {
	if v.IsValid() {
		if err := Angle.AssertValue(v); err != nil {
			return value.Value{}
		}

		if value.LessThan(v.Float(), 1.0) {
			switch v.Unit() {
			case Degree:
				return AngleRoundDown(v.AsOrInvalid(ArcMinute))
			case ArcMinute:
				return AngleRoundDown(v.AsOrInvalid(ArcSecond))
			}
		}
	}

	return v
}

func degreeToHourAngle(f float64) (float64, error) {
	if value.GreaterThan(f, 12) {
		return f - 24, nil
	}
	return f, nil
}

func hourAngleToDegree(f float64) (float64, error) {
	if value.IsNegative(f) {
		return f + 24, nil
	}
	return f, nil
}

func degreeToLongitude(f float64) (float64, error) {
	if value.GreaterThan(f, 180) {
		f = f - 360
	}
	return f, nil
}

func longitudeToDegree(f float64) (float64, error) {
	if value.IsNegative(f) {
		f = f + 360
	}
	return f, nil
}

func degreeToAzimuth(f float64) (float64, error) {
	f = f - 180.0

	if value.LessThanEqual(f, -180+value.EqualityError) {
		f = f + 360.0
	}

	return f, nil
}

func azimuthToDegree(f float64) (float64, error) {
	f = f + 180

	if value.GreaterThanEqual(f, 360-value.EqualityError) {
		f = f - 360.0
	}

	return f, nil
}
