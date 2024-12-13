package forecast

type Hemisphere bool

const (
	NorthernHemisphere Hemisphere = false
	SouthernHemisphere Hemisphere = true
)

func HemisphereFor(latitude float64) Hemisphere {
	if latitude < 0 {
		return SouthernHemisphere
	}
	return NorthernHemisphere
}
