package util

import (
	"math"
)

func FdayToHMS(day float64) (int, int, int, int) {
	d, fdy := math.Modf(day)
	fdy = fdy * 24.0
	h, fdy := math.Modf(fdy)
	fdy = fdy * 60.0
	m, fdy := math.Modf(fdy)
	fdy = fdy * 24.0
	s := math.Round(fdy)
	return int(d), int(h), int(m), int(s)
}

func HMSToFday(d, h, m, s int) float64 {
	return float64(d) + ((float64(h) + (float64(m) / 60.0) + (float64(s) / 3600.0)) / 24.0)

}
