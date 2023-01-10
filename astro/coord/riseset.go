package coord

import (
	"encoding/xml"
	"github.com/peter-mount/nre-feeds/util"
	util2 "github.com/peter-mount/piweather.center/astro/util"
	"github.com/soniakeys/unit"
)

type RiseSet struct {
	Circumpolar bool      `json:"circumpolar,omitempty" xml:"circumpolar,omitempty,attr" yaml:"circumpolar,omitempty"`
	Rise        unit.Time `json:"rise,omitempty" xml:"rise,omitempty,attr" yaml:"rise,omitempty"`
	Transit     unit.Time `json:"transit,omitempty" xml:"transit,omitempty,attr" yaml:"transit,omitempty"`
	Set         unit.Time `json:"set,omitempty" xml:"set,omitempty,attr" yaml:"set,omitempty"`
}

func (r *RiseSet) String() string {
	return util2.String(r)
}

func (r *RiseSet) Equals(b *RiseSet) bool {
	if r == nil {
		return b == nil
	}
	if r.Circumpolar || b.Circumpolar {
		return r.Circumpolar == b.Circumpolar
	}

	return int(r.Rise.Sec()) == int(b.Rise.Sec()) &&
		int(r.Transit.Sec()) == int(b.Transit.Sec()) &&
		int(r.Set.Sec()) == int(b.Set.Sec())
}

func (r *RiseSet) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	b := util.NewXmlBuilder(encoder, start)

	if r.Circumpolar {
		b.AddBoolAttribute(xml.Name{Local: "circumpolar"}, r.Circumpolar)
	} else {
		b.AddAttribute(xml.Name{Local: "rise"}, util2.HourDMSString(r.Rise)).
			AddAttribute(xml.Name{Local: "transit"}, util2.HourDMSString(r.Transit)).
			AddAttribute(xml.Name{Local: "set"}, util2.HourDMSString(r.Set))
	}
	return b.Build()
}
