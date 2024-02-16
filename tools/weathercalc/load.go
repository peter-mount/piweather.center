package weathercalc

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/config/calc"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	"strings"
)

// load a metric's value on startup if they specify getting the value from the db
func (calc *Calculator) loadFromDB(c *calc.Calculation) error {
	b := c.Load

	if *calc.DBServer == "" {
		return participle.Errorf(b.Pos, "DB load requested but no server defined")
	}

	// Form the query
	var q []string

	switch b.When {
	case "today":
		q = append(q,
			`BETWEEN "midnight" AND "now" ADD "24h"`,
			`EVERY "24h"`)

	case "hour":
		q = append(q,
			`BETWEEN "now" TRUNCATE "1h" AND "now" ADD "1h"`,
			`EVERY "1h"`)

	case "minute":
		q = append(q,
			`BETWEEN "now" ADD "-1m" AND "now" TRUNCATE "1m"`,
			`EVERY "1m"`)

	default:
		return participle.Errorf(b.Pos, "Unsupported when %q in exec", b.When)
	}

	useFirst := ""
	if c.UseFirst != nil && c.UseFirst.Metric != nil {
		useFirst = "LAST(" + c.UseFirst.Metric.Name + ")"
	}
	q = append(q, "LIMIT 1", `SELECT`, `TIMEOF(`+useFirst+`),`, b.With)

	query := strings.Join(q, " ")
	log.Printf("DB: %s", query)

	cl := client.Client{Url: *calc.DBServer}
	res, err := cl.Query(query)
	if err != nil {
		return participle.Errorf(b.Pos, "%s", err.Error())
	}

	for _, t := range res.Table {
		for _, r := range t.Rows {
			if len(*r) < 2 || !(*r)[1].Value.IsValid() {
				log.Printf("no data returned for %q", b.With)
			} else {
				e0 := (*r)[0]
				e1 := (*r)[1]
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
