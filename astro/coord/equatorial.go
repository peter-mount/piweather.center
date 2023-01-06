package coord

import (
	"github.com/peter-mount/piweather.center/astro/util"
)

type Equatorial struct {
	Alpha float64 `json:"alpha" xml:"alpha,attr" yaml:"alpha"`
	Delta float64 `json:"delta" xml:"delta,attr" yaml:"delta"`
	RA    string  `json:"ra" xml:"ra,attr" yaml:"ra"`
	Dec   string  `json:"dec" xml:"dec,attr" yaml:"dec"`
}

func New(alpha, delta float64) Equatorial {
	e := Equatorial{
		Alpha: util.DegRange(alpha),
		Delta: delta,
	}

	e.RA = util.DegDMSString(e.Alpha/15.0, false)
	e.Dec = util.DegDMSStringExt(e.Delta, true, "N", "S")

	return e
}
