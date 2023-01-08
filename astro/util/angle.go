package util

import (
	"errors"
	"fmt"
	"github.com/soniakeys/unit"
	"math"
	"strconv"
	"strings"
)

const (
	toRad = math.Pi / 180.0
	toDeg = 180.0 / math.Pi
)

// Deg2Rad converts an angle in degrees to radians
func Deg2Rad(d float64) float64 {
	return d * toRad
}

// Rad2Deg converts an angle in radians to degrees
func Rad2Deg(r float64) float64 {
	return r * toDeg
}

func DegRange(d float64) float64 {
	for d < 0.0 {
		d = d + 360.0
	}
	for d >= 360.0 {
		d = d - 360.0
	}
	return d
}

func DegDMS(d float64) (int, int, float64) {
	deg, fDeg := math.Modf(d)
	m, s := math.Modf(fDeg * 60)
	s = s * 60.0
	return int(deg), int(m), s
}

func DegDMSString(d float64, sign bool) string {
	return DegDMSStringExt(d, sign, "+", "-")
}

func DegDMSStringExt(d float64, sign bool, p, m string) string {
	var s string
	deg, min, sec := DegDMS(math.Abs(d))
	if sign {
		s = fmt.Sprintf("%02d:%02d:%02d", deg, min, int(math.Round(sec)))
		if d < 0.0 {
			s = m + s
		} else {
			s = p + s
		}
	} else {
		s = fmt.Sprintf("%3d:%02d:%02d", deg, min, int(math.Round(sec)))
	}
	return strings.TrimSpace(s)
}

func HourDMSString(t unit.Time) string {
	return HourDMSStringExt(t.Hour())
}

func HourDMSStringExt(d float64) string {
	deg, min, sec := DegDMS(math.Abs(d))
	return fmt.Sprintf("%02d:%02d:%02d", deg, min, int(math.Round(sec)))
}

var (
	noAngle = errors.New("no angle")
)

func ParseAngle(s string) (unit.Angle, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, noAngle
	}

	var neg bool
	switch s[0] {
	case '-', 'S', 's':
		neg = true
		s = s[1:]
	case '+', 'N', 'n':
		neg = false
		s = s[1:]
	}

	if s == "" {
		return 0, noAngle
	}

	var err error
	var angle float64
	if strings.Contains(s, ":") {
		angle, err = parseAngleDMS(strings.SplitN(s, ":", 3))
	} else {
		angle, err = strconv.ParseFloat(s, 64)
	}

	if err != nil {
		return 0, err
	}

	if neg {
		angle = -angle
	}
	return unit.AngleFromDeg(angle), nil
}

func parseAngleDMS(s []string) (float64, error) {
	angle := 0.0
	multiplier := 1.0
	for _, v := range s {
		v = strings.TrimSpace(v)
		if v != "" {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return 0, err
			}
			angle = angle + (math.Abs(f) * multiplier)
			multiplier = multiplier / 60.0
		}
	}

	return angle, nil
}
