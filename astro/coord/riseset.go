package coord

import (
	"encoding/xml"
	"github.com/peter-mount/nre-feeds/util"
	util2 "github.com/peter-mount/piweather.center/astro/util"
	util3 "github.com/peter-mount/piweather.center/util/strings"
	"github.com/soniakeys/unit"
	"time"
)

type RiseSet struct {
	Circumpolar bool          `json:"circumpolar,omitempty" xml:"circumpolar,omitempty,attr" yaml:"circumpolar,omitempty"`
	Rise        unit.Time     `json:"rise,omitempty" xml:"rise,omitempty,attr" yaml:"rise,omitempty"`
	Transit     unit.Time     `json:"transit,omitempty" xml:"transit,omitempty,attr" yaml:"transit,omitempty"`
	Set         unit.Time     `json:"set,omitempty" xml:"set,omitempty,attr" yaml:"set,omitempty"`
	Duration    time.Duration `json:"duration,omitempty" xml:"duration,omitempty,attr" yaml:"duration,omitempty"`
	DayLength   unit.Time     `json:"dayLength,omitempty" xml:"dayLength,omitempty,attr" yaml:"dayLength,omitempty"`
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
		dur := r.Set - r.Rise
		if dur < 0 {
			dur += 86400.0
		}

		b.AddAttribute(xml.Name{Local: "rise"}, util3.HourDMSString(r.Rise)).
			AddAttribute(xml.Name{Local: "transit"}, util3.HourDMSString(r.Transit)).
			AddAttribute(xml.Name{Local: "set"}, util3.HourDMSString(r.Set)).
			AddFloatAttribute(xml.Name{Local: "duration"}, dur.Hour()).
			AddAttribute(xml.Name{Local: "dayLength"}, util3.HourDMSString(dur))
	}
	return b.Build()
}
