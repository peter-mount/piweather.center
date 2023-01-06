package sun

import "github.com/peter-mount/piweather.center/astro/util"

func Obliquity(t float64) float64 {
	return util.Polynomial(t, 23.43929111, -46.8150/3600.0, -0.00059/3600.0, 0.001813/3600.0)
}
