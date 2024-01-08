package weathercalc

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/store/file/record"
	"github.com/peter-mount/piweather.center/tools/weathercalc/lang"
	"strings"
)

// load a metric's value on startup if they specify getting the value from the db
func (calc *Calculator) loadFromDB(c *lang.Calculation) error {
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

	q = append(q, "LIMIT 1", `SELECT`, b.With)

	query := strings.Join(q, " ")
	log.Printf("DB: %s", query)

	cl := client.Client{Url: *calc.DBServer}
	res, err := cl.Query(query)
	if err != nil {
		return participle.Errorf(b.Pos, "%s", err.Error())
	}

	for _, t := range res.Table {
		for _, r := range t.Rows {
			if len(*r) == 0 || !(*r)[0].Value.IsValid() {
				log.Printf("no data returned for %q", b.With)
			} else {
				e := (*r)[0]
				calc.Latest.Append(c.Target, record.Record{Time: e.Time, Value: e.Value})

				if err = calc.DatabaseBroker.PublishMetric(api.Metric{
					Metric: c.Target,
					Time:   e.Time,
					Unit:   e.Value.Unit().ID(),
					Value:  e.Value.Float(),
				}); err != nil {
					log.Printf("post %q failed", c.Target)
				}
			}
			return nil
		}
	}

	return nil
}
