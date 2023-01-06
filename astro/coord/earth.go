package coord

import (
	"encoding/xml"
	"github.com/peter-mount/nre-feeds/util"
	util2 "github.com/peter-mount/piweather.center/astro/util"
	"github.com/soniakeys/meeus/v3/globe"
)

// LatLong represents a location on the Earth's surface.
//
// Be aware that here Longitude is East positive, West negative, which is normal for geography/gis.
// Some calculations however use the opposite with West positive.

type LatLong struct {
	Coord    globe.Coord
	Altitude float64
	Name     string
	Notes    string
}

func (r *LatLong) String() string {
	return util2.String(r)
}

func (r *LatLong) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	return util.NewXmlBuilder(encoder, start).
		AddAttributeIfSet(xml.Name{Local: "name"}, r.Name).
		AddFloatAttribute(xml.Name{Local: "lat"}, r.Coord.Lat.Deg()).
		AddFloatAttribute(xml.Name{Local: "lon"}, r.Coord.Lon.Deg()).
		AddFloatAttribute(xml.Name{Local: "alt"}, r.Altitude).
		ElementIf(r.Notes != "",
			xml.Name{Local: "notes"},
			func(builder *util.XmlBuilder) error {
				builder.AddCharData(r.Notes)
				return nil
			}).
		Build()
}
