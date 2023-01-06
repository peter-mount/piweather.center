package coord

import (
	"encoding/xml"
	"github.com/peter-mount/piweather.center/astro/julian"
)

type Orbit struct {
	TrueAnomaly       float64 `json:"trueAnomaly" xml:"trueAnomaly,attr" yaml:"trueAnomaly"`
	RadiusVector      float64 `json:"radiusVector" xml:"radiusVector,attr" yaml:"radiusVector"`
	ObliquityEcliptic float64 `json:"obliquityEcliptic" xml:"obliquityEcliptic,attr" yaml:"obliquityEcliptic"`
	TrueLongitude     float64 `json:"trueLongitude" xml:"trueLongitude,attr" yaml:"trueLongitude"`
	ApparentLongitude float64 `json:"apparentLongitude" xml:"apparentLongitude,attr" yaml:"apparentLongitude"`
}

type Sun struct {
	XMLName    xml.Name   `xml:"sun" json:"-" yaml:"-"`
	Date       julian.Day `json:"jd" xml:"jd,attr" yaml:"jd"`
	Orbit      Orbit      `json:"orbit" xml:"orbit" yaml:"orbit"`
	Equatorial Equatorial `json:"equatorial" xml:"equatorial" yaml:"equatorial"`
}

func (s *Sun) String() string {
	if s == nil {
		return "nil"
	}
	b, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}
