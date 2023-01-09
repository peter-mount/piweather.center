package coord

import (
	"errors"
	"github.com/peter-mount/piweather.center/astro/util"
	"strconv"
	"strings"
)

func Parse(s string) (LatLong, error) {
	a := strings.SplitN(s, ",", 4)
	l := len(a)
	if l < 2 {
		return LatLong{}, errors.New("syntax: long,lat[,alt,[name...]]")
	}

	ll := LatLong{}

	ang, err := util.ParseAngle(a[0])
	if err != nil {
		return ll, err
	}
	ll.Longitude = ang
	ll.coord.Lon = -ang

	ang, err = util.ParseAngle(a[1])
	if err != nil {
		return ll, err
	}
	ll.Latitude = ang
	ll.coord.Lat = ang

	if l > 2 {
		ll.Altitude, err = strconv.ParseFloat(a[2], 64)
		if err != nil {
			return ll, err
		}
	}

	if l > 3 {
		ll.Name = a[3]
	}

	return ll, nil
}
