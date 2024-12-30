package weathercalc

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/config/station"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	"strings"
)

// load a metric's value on startup if they specify getting the value from the db
func (calc *Calculator) loadFromDB(c *station.Calculation) error {
	b := c.Load

	if *calc.DBServer == "" {
		return participle.Errorf(b.Pos, "DB load requested but no server defined")
	}

	// Form the query
	var q []string

	switch b.When {
	case "today":
		q = append(q,
			`between "midnight" and "now" add "24h"`,
			`every "24h"`)

	case "hour":
		q = append(q,
			`between "now" truncate "1h" and "now" add "1h"`,
			`every "1h"`)

	case "minute":
		q = append(q,
			`between "now" add "-1m" and "now" truncate "1m"`,
			`every "1m"`)

	default:
		return participle.Errorf(b.Pos, "Unsupported when %q in exec", b.When)
	}

	useFirst := ""
	if c.UseFirst != nil && c.UseFirst.Metric != nil {
		useFirst = "last(" + c.UseFirst.Metric.Name + ")"
	}
	q = append(q, "limit 1", `select`, `timeof(`+useFirst+`),`, b.With)

	query := strings.Join(q, " ")
	log.Printf("DB: %s", query)

	cl := client.Client{Url: *calc.DBServer, Internal: true}
	res, err := cl.Query(query)
	if err != nil {
		return participle.Errorf(b.Pos, "%s", err.Error())
	}

	for _, t := range res.Table {
		for _, r := range t.Rows {
			if r.Size() < 2 || r.Cell(1).Value.IsValid() {
				log.Printf("no data returned for %q", b.With)
			} else {
				e0 := r.Cell(0)
				e1 := r.Cell(1)
				calc.Latest.Set(c.Target, record.Record{Time: e0.Time, Value: e1.Value})

				if err = calc.DatabaseBroker.PublishMetric(api.Metric{
					Metric: c.Target,
					Time:   e0.Time,
					Unit:   e1.Value.Unit().ID(),
					Value:  e1.Value.Float(),
				}); err != nil {
					log.Printf("post %q failed", c.Target)
				}
			}
			return nil
		}
	}

	return nil
}
