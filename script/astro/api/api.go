package api

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/astro/api"
)

func init() {
	packages.RegisterPackage(&Package{
		EquatorialName:    int(api.EquatorialName),
		EquatorialRA:      int(api.EquatorialRA),
		EquatorialDec:     int(api.EquatorialDec),
		Distance:          int(api.Distance),
		LightTime:         int(api.LightTime),
		DistanceSun:       int(api.DistanceSun),
		SemiDiameter:      int(api.SemiDiameter),
		HorizonAltitude:   int(api.HorizonAltitude),
		HorizonAzimuth:    int(api.HorizonAzimuth),
		HorizonBearing:    int(api.HorizonBearing),
		EclipticLatitude:  int(api.EclipticLatitude),
		EclipticLongitude: int(api.EclipticLongitude),
		GalacticLatitude:  int(api.GalacticLatitude),
		GalacticLongitude: int(api.GalacticLongitude),
	})
}

type Package struct {
	EquatorialName    int
	EquatorialRA      int
	EquatorialDec     int
	Distance          int
	LightTime         int
	DistanceSun       int
	SemiDiameter      int
	HorizonAltitude   int
	HorizonAzimuth    int
	HorizonBearing    int
	EclipticLatitude  int
	EclipticLongitude int
	GalacticLatitude  int
	GalacticLongitude int
}
