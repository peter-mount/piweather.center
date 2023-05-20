package graph

import (
	"fmt"
	svg "github.com/ajstarks/svgo/float"
	"github.com/peter-mount/piweather.center/weather/value"
	"io"
	"strings"
)

func Number(f float64) string {
	s := fmt.Sprintf("%.2f", f)
	switch {
	// xx.00 -> xx
	case strings.HasSuffix(s, ".00"):
		return s[:len(s)-3]
	// xx.x0 -> xx.x
	case strings.HasSuffix(s, "0"):
		return s[:len(s)-1]
	// xx.xx
	default:
		return s
	}
}

func Text(canvas *svg.SVG, x, y, rot float64, class, text string, arg ...interface{}) {
	if len(arg) > 0 {
		text = fmt.Sprintf(text, arg...)
	}
	text = strings.TrimSpace(text)

	if text != "" {
		s := []string{"<text x=\"", Number(x), "\" y=\"", Number(y), "\""}

		if value.NotEqual(rot, 0) {
			s = append(s, " transform=\"rotate(", Number(rot), ",", Number(x), ",", Number(y), ")\"")
		}

		if class != "" {
			s = append(s, " class=\"", class, "\"")
		}

		s = append(s, "><![CDATA[", text, "]]></text>")

		_, _ = io.WriteString(canvas.Writer, strings.Join(s, ""))
	}
}
