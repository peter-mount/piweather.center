package util

func Fsum(v ...float64) float64 {
	var r float64
	for _, e := range v {
		r = r + e
	}
	return r
}
