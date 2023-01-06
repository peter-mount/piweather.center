package sun

import (
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/util"
	"math"
)

func Low(jd julian.Day) coord.Sun {

	T := jd.CenturiesJ2k()
	L0 := util.DegRange(util.Polynomial(T, 280.46645, 36000.76983, 0.0003032))
	M := util.DegRange(util.Polynomial(T, 357.52910, 35999.05030, -0.0001559, -0.00000048))
	e := util.Polynomial(T, 0.016708617, -0.000042037, -0.0000001236)

	Mr := util.Deg2Rad(M)

	C := (util.Polynomial(T, 1.914600, -0.004817, -0.000014) * math.Sin(Mr)) +
		(util.Polynomial(T, 0.019993, -0.000101) * math.Sin(2*Mr)) +
		(0.000290 * math.Sin(3*Mr))

	sun := coord.Sun{
		Date: jd,
		Orbit: coord.Orbit{
			TrueLongitude:     L0 + C,
			TrueAnomaly:       M + C,
			ObliquityEcliptic: Obliquity(T),
		},
	}

	sun.Orbit.RadiusVector = (1.000001018 * (1 - (e * e))) / (1 + (e * math.Cos(util.Deg2Rad(sun.Orbit.TrueAnomaly))))

	ohm := util.Deg2Rad(util.Polynomial(T, 125.04, -1934.136))
	sun.Orbit.ApparentLongitude = sun.Orbit.TrueLongitude - 0.00569 - (0.00478 * math.Sin(ohm))

	L0r := util.Deg2Rad(sun.Orbit.ApparentLongitude)
	obliquity := util.Deg2Rad(sun.Orbit.ObliquityEcliptic + (0.00256 * math.Cos(ohm)))

	alpha := util.Rad2Deg(math.Atan2(math.Cos(obliquity)*math.Sin(L0r), math.Cos(L0r)))
	delta := util.Rad2Deg(math.Asin(math.Sin(obliquity) * math.Sin(L0r)))

	sun.Equatorial = coord.New(util.DegRange(alpha), delta)
	return sun
}
