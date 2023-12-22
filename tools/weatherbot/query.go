package bot

import (
	"fmt"
	"github.com/peter-mount/piweather.center/util"
	"sort"
	"strings"
)

type Query struct {
	post    *Post
	ranges  util.StringSet
	columns util.StringMap
	colKeys util.StringSet
	Query   string
}

var (
	ranges = util.StringMap{
		"today":      `between "midnight" and "now"`,
		"hour":       `between "row" truncate "1h" and "row" add "every"`,
		"trendLimit": `offset "-10m"`,
	}
)

// ParsePost parses a Post to form a query to issue to weatherdb
func ParsePost(post *Post) (*Query, error) {
	q := &Query{
		post:    post,
		ranges:  util.NewStringSet(),
		columns: util.NewStringMap(),
		colKeys: util.NewStringSet(),
	}
	err := q.parsePost()
	if err == nil {
		q.Query = q.generateQuery()
	}
	return q, err
}

func (q *Query) parsePost() error {
	for _, t := range q.post.Threads {
		if err := q.parseThread(t); err != nil {
			return err
		}
	}
	return nil
}

func (q *Query) parseThread(t *Thread) error {
	for _, r := range t.Table {
		if err := q.parseRow(r); err != nil {
			return err
		}
	}
	return nil
}

func (q *Query) parseRow(r *Row) error {
	for _, v := range r.Values {
		if err := q.parseValue(v); err != nil {
			return err
		}
	}
	return nil
}

func (q *Query) parseValue(v Value) error {
	var a []string

	switch {
	case v.Type != "" && v.Sensor != "":
		a = append(a, v.Type, "(", v.Sensor, ")")
	case v.Type != "":
		a = append(a, v.Type, "()")
	case v.Sensor != "":
		a = append(a, v.Sensor)
	default:
		// No type or sensor then ignore the value
		return nil
	}

	if v.Range != "" {
		a = append(a, ` USING "`, v.Range, `"`)
		q.ranges.Add(v.Range)
	}

	if v.Unit.Unit != "" {
		a = append(a, ` UNIT "`, v.Unit.Unit, `"`)
	}

	l := strings.Join(a, "")
	if q.colKeys.Add(l) {
		colName := fmt.Sprintf(`col%03d`, len(q.colKeys))

		// Using a Set means we only include a query once
		q.columns.Add(colName, l+` AS "`+colName+`"`)
	}
	return nil
}

func (q *Query) generateQuery() string {
	var a []string
	// TODO range
	//a = append(a, `BETWEEN "midnight" AND "now"`)
	a = append(a,
		`BETWEEN "now" TRUNCATE "10m" ADD "-10m" AND "now" ADD "1m"`,
		`EVERY "10m"`)

	if !q.ranges.IsEmpty() {
		var d []string
		for _, k := range q.ranges.Entries() {
			if len(d) > 0 {
				d = append(d, ",\n")
			}
			d = append(d, `  "`, k, `" AS `, ranges[k])
		}
		a = append(a, "DECLARE", strings.Join(d, ""))
	}

	columns := q.columns.Keys()
	sort.SliceStable(columns, func(i, j int) bool {
		return columns[i] < columns[j]
	})
	var s []string
	for _, k := range columns {
		if len(s) > 0 {
			s = append(s, ",\n")
		}
		s = append(s, "  ", q.columns.Get(k))
	}
	a = append(a, "SELECT", strings.Join(s, ""))

	return strings.Join(a, "\n")
}
