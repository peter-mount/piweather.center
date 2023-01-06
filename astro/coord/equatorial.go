package coord

import (
	"encoding/xml"
	"github.com/peter-mount/nre-feeds/util"
	util2 "github.com/peter-mount/piweather.center/astro/util"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/meeus/v3/rise"
	"github.com/soniakeys/unit"
)

type Equatorial struct {
	Alpha unit.RA    `json:"alpha" xml:"alpha,attr" yaml:"alpha"`
	Delta unit.Angle `json:"delta" xml:"delta,attr" yaml:"delta"`
	RA    string     `json:"ra" xml:"ra,attr" yaml:"ra"`
	Dec   string     `json:"dec" xml:"dec,attr" yaml:"dec"`
}

func New(alpha unit.RA, delta unit.Angle) Equatorial {
	e := Equatorial{Alpha: alpha, Delta: delta}
	e.RA = util2.HourDMSString(e.Alpha.Time())
	e.Dec = util2.DegDMSStringExt(e.Delta.Deg(), true, "N", "S")
	return e
}

func (e *Equatorial) RiseSet(p globe.Coord, th0 unit.Time, h0 unit.Angle) RiseSet {
	rise, transit, set, err := rise.ApproxTimes(p, h0, th0, e.Alpha, e.Delta)
	if err != nil {
		return RiseSet{Circumpolar: true}
	}

	return RiseSet{Transit: transit, Rise: rise, Set: set}
}

func (e *Equatorial) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	return util.NewXmlBuilder(encoder, start).
		AddFloatAttribute(xml.Name{Local: "alpha"}, e.Alpha.Hour()).
		AddFloatAttribute(xml.Name{Local: "delta"}, e.Delta.Deg()).
		AddAttribute(xml.Name{Local: "ra"}, e.RA).
		AddAttribute(xml.Name{Local: "dec"}, e.Dec).
		Build()
}
