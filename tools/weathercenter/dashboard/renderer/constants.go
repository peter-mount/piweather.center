package renderer

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
)

const (
	D                = "d"
	DominantBaseline = "dominant-baseline"
	FontSize         = "font-size"
	ID               = "id"
	TextAnchor       = "text-anchor"
	Transform        = "transform"

	Black     = "black"
	Middle    = "middle"
	None      = "none"
	White     = "white"
	Red       = "red"
	Blue      = "blue"
	Green     = "green"
	Yellow    = "yellow"
	Purple    = "purple"
	Cyan      = "cyan"
	LightBlue = "lightblue"
)

func fix(f float64) string {
	s := fmt.Sprintf("%.4f", f)
	s = strings.TrimRight(s, "0")
	return strings.TrimSuffix(s, ".")
}

func Translate(x, y float64) string {
	if value.IsZero(x) && value.IsZero(y) {
		return ""
	}
	return fmt.Sprintf("translate(%s,%s)", fix(x), fix(y))
}

func Rotate(d float64) string {
	s := fix(d)
	if s != "" {
		return fmt.Sprintf("rotate(%s)", fix(d))
	}
	return ""
}
