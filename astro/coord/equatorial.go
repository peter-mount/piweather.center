package coord

import (
	"encoding/xml"
	"github.com/peter-mount/nre-feeds/util"
	util2 "github.com/peter-mount/piweather.center/astro/util"
	"github.com/soniakeys/meeus/v3/coord"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/meeus/v3/rise"
	"github.com/soniakeys/unit"
)

type Equatorial struct {
	Alpha    unit.RA    `json:"alpha" xml:"alpha,attr" yaml:"alpha"`
	Delta    unit.Angle `json:"delta" xml:"delta,attr" yaml:"delta"`
	RA       string     `json:"ra" xml:"ra,attr" yaml:"ra"`
	Dec      string     `json:"dec" xml:"dec,attr" yaml:"dec"`
	Diameter unit.Angle `json:"diameter,omitempty" xml:"diameter,attr,omitempty" yaml:"diameter,omitempty"`
}

func (e *Equatorial) Equals(b *Equatorial) bool {
	if e == nil {
		return b == nil
	}
	return int(e.Alpha.Sec()) == int(b.Alpha.Sec()) &&
		int(e.Delta.Sec()) == int(b.Delta.Sec())
}

func New(alpha unit.RA, delta unit.Angle) Equatorial {
	e := Equatorial{Alpha: alpha, Delta: delta}
	e.RA = util2.HourDMSString(e.Alpha.Time())
	e.Dec = util2.DegDMSStringExt(e.Delta.Deg(), true, "N", "S")
	return e
}

func NewFromEq(e coord.Equatorial) Equatorial {
	return New(e.RA, e.Dec)
}

func (e *Equatorial) Equatorial() coord.Equatorial {
	return coord.Equatorial{
		RA:  e.Alpha,
		Dec: e.Delta,
	}
}

func (e *Equatorial) RiseSet(p *globe.Coord, th0 unit.Time, h0 unit.Angle) RiseSet {
	tRise, tTransit, tSet, err := rise.ApproxTimes(*p, h0, th0, e.Alpha, e.Delta)
	if err != nil {
		return RiseSet{Circumpolar: true}
	}

	return RiseSet{Transit: tTransit, Rise: tRise, Set: tSet}
}

func (e *Equatorial) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	return util.NewXmlBuilder(encoder, start).
		AddFloatAttribute(xml.Name{Local: "alpha"}, e.Alpha.Hour()).
		AddFloatAttribute(xml.Name{Local: "delta"}, e.Delta.Deg()).
		AddAttribute(xml.Name{Local: "ra"}, e.RA).
		AddAttribute(xml.Name{Local: "dec"}, e.Dec).
		Build()
}
