package strings

import (
	"errors"
	"github.com/soniakeys/unit"
	"math"
	"strconv"
	"strings"
)

var (
	noAngle = errors.New("no angle")
)

func ParseAngle(s string) (unit.Angle, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, noAngle
	}

	var neg bool
	// check for prefix
	switch s[0] {
	case '-', 'S', 's', 'W', 'w':
		neg = true
		s = s[1:]
	case '+', 'N', 'n', 'E', 'e':
		neg = false
		s = s[1:]
	}

	if s == "" {
		return 0, noAngle
	}

	// Check for suffix
	sl := len(s) - 1
	switch s[sl] {
	case 'S', 's', 'W', 'w':
		neg = true
		s = s[:sl]
	case 'N', 'n', 'E', 'e':
		neg = false
		s = s[:sl]
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
