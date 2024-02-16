package bot

import (
	"fmt"
	util2 "github.com/peter-mount/piweather.center/config/util"
	"github.com/peter-mount/piweather.center/store/ql/lang"
	"github.com/peter-mount/piweather.center/store/ql/parser"
	"github.com/peter-mount/piweather.center/util"
	"sort"
	"strings"
)

type Query struct {
	post       *Post
	columns    map[string]column
	colKeys    util.StringMap
	parser     util2.Parser[lang.Expression]
	Query      string
	queryRange string
}

var (
	ranges = util.StringMap{
		"yesterday":  `between "midnight" add "-24h" and "midnight"`,
		"today":      `between "midnight" and "now"`,
		"hour":       `between "row" truncate "1h" and "row" add "every"`,
		"trendLimit": `offset "-10m"`,
	}
)

type column struct {
	expression string // Expression used
	name       string // Name of column
}

// ParsePost parses a Post to form a query to issue to weatherdb
func ParsePost(post *Post) (*Query, error) {
	q := &Query{
		post:       post,
		columns:    make(map[string]column),
		colKeys:    util.NewStringMap(),
		parser:     parser.NewExpressionParser(),
		queryRange: post.Range,
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

	for _, w := range r.When {
		err := q.parseValue(w.Value)
		if err == nil {
			err = q.parseValue(w.LessThan)
		}
		if err == nil {
			err = q.parseValue(w.LessThanEqual)
		}
		if err == nil {
			err = q.parseValue(w.Equal)
		}
		if err == nil {
			err = q.parseValue(w.NotEqual)
		}
		if err == nil {
			err = q.parseValue(w.GreaterThanEqual)
		}
		if err == nil {
			err = q.parseValue(w.GreaterThan)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (q *Query) parseValue(v *Value) error {
	if v == nil || v.Query == "" {
		return nil
	}

	// Verify expression is valid
	_, err := q.parser.ParseString("", v.Query)
	if err != nil {
		return fmt.Errorf("error in %s: %s", v.Query, err.Error())
	}

	// Using a Set means we only include a query once
	if q.colKeys.Contains(v.Query) {
		// Just set the shared column name
		v.Col = q.colKeys.Get(v.Query)
	} else {
		// New entry
		colName := fmt.Sprintf(`col%03d`, len(q.colKeys))

		q.columns[colName] = column{
			expression: v.Query,
			name:       colName,
		}

		v.Col = colName

		q.colKeys.Add(v.Query, colName)
	}
	return nil
}

func (q *Query) generateQuery() string {
	var a []string

	if q.queryRange == "" {
		q.queryRange = `BETWEEN "now" TRUNCATE "10m" ADD "-10m" AND "now" ADD "1m" EVERY "10m"`
	}
	a = append(a, q.queryRange)

	var d []string
	for k, v := range ranges {
		if len(d) > 0 {
			d = append(d, ",\n")
		}
		d = append(d, `  "`, k, `" AS `, v)
	}
	a = append(a, "DECLARE", strings.Join(d, ""))

	var columns []string
	for k, _ := range q.columns {
		columns = append(columns, k)
	}
	sort.SliceStable(columns, func(i, j int) bool {
		return columns[i] < columns[j]
	})

	var s []string
	for _, k := range columns {
		if len(s) > 0 {
			s = append(s, ",\n")
		}
		c := q.columns[k]
		s = append(s, "  ", c.expression, ` AS "`, c.name, `"`)
	}
	a = append(a, "SELECT", strings.Join(s, ""))

	return strings.Join(a, "\n")
}
