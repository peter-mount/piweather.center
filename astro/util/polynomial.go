package util

func Polynomial(t float64, v ...float64) float64 {
	r := 0.0
	b := 1.0
	for _, e := range v {
		r = r + (e * b)
		b = b * t
	}
	return r
}
