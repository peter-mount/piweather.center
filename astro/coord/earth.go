package coord

import (
	"encoding/xml"
	"github.com/peter-mount/nre-feeds/util"
	util2 "github.com/peter-mount/piweather.center/astro/util"
	"github.com/peter-mount/piweather.center/weather/value"
	"github.com/soniakeys/meeus/v3/globe"
	"github.com/soniakeys/unit"
	"time"
)

// LatLong represents a location on the Earth's surface.
//
// Be aware that here Longitude is East positive, West negative, which is normal for geography/gis.
// Some calculations however use the opposite with West positive.

type LatLong struct {
	Longitude unit.Angle `xml:"longitude,attr"`
	Latitude  unit.Angle `xml:"latitude,attr"`
	Altitude  float64    `xml:"altitude,attr"`
	Name      string     `xml:"name,attr,omitempty"`
	Notes     string     `xml:",chardata"`
	coord     globe.Coord
}

func (r *LatLong) String() string {
	return util2.String(r)
}

// Coord returns a globe.Coord from the meeus library.
//
// Note: in the meeus library (& book) it uses West positive for longitudes
// whereas we use East positive, so this function is the best way to get
// the coordinates correct.
func (r *LatLong) Coord() *globe.Coord {
	if r.coord.Lat == 0 && r.coord.Lon == 0 {
		r.coord.Lon = -r.Longitude
		r.coord.Lat = r.Latitude
	}
	return &r.coord
}

func (r *LatLong) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	return util.NewXmlBuilder(encoder, start).
		AddAttributeIfSet(xml.Name{Local: "name"}, r.Name).
		AddFloatAttribute(xml.Name{Local: "lon"}, r.Longitude.Deg()).
		AddFloatAttribute(xml.Name{Local: "lat"}, r.Latitude.Deg()).
		AddFloatAttribute(xml.Name{Local: "alt"}, r.Altitude).
		ElementIf(r.Notes != "",
			xml.Name{Local: "notes"},
			func(builder *util.XmlBuilder) error {
				builder.AddCharData(r.Notes)
				return nil
			}).
		Build()
}

func (r *LatLong) Time(t time.Time) value.Time {
	return value.BasicTime(t, r.Coord(), r.Altitude)
}
