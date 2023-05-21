package svg

import (
	"fmt"
	"github.com/peter-mount/piweather.center/weather/value"
	"strings"
)

type Transformer struct {
	s []string
}

func Transform() *Transformer {
	return &Transformer{}
}

func (t *Transformer) String() string {
	return strings.Join(t.s, " ")
}

func (t *Transformer) Attr() string {
	return Attr("transform", t.String())
}

func (t *Transformer) add(f string, a ...interface{}) *Transformer {
	t.s = append(t.s, fmt.Sprintf(f, a...))
	return t
}

func (t *Transformer) Rotate(rot, x, y float64) *Transformer {
	if value.NotEqual(rot, 0) {
		return t.add("rotate(%s,%s,%s)", Number(rot), Number(x), Number(y))
	}
	return t
}
