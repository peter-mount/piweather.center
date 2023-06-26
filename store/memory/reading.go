package memory

import (
	"github.com/peter-mount/piweather.center/weather/value"
	"strconv"
	"strings"
	"time"
)

type Reading struct {
	Name  string
	Value value.Value
	Time  time.Time
}

func (r *Reading) String() string {
	return strings.Join([]string{
		r.Name,
		strconv.FormatFloat(r.Value.Float(), 'f', 3, 64),
		strconv.Itoa(int(r.Time.UTC().Unix())),
	}, " ")
}
