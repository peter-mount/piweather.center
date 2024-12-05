package query

import (
	"flag"
	"fmt"
	"github.com/peter-mount/piweather.center/store/client"
	"net/http"
	"strings"
)

// Query is a utility to perform queries against the DB.
//
// -query-db http://192.168.1.137:9001
// -query 'between "2024-11-01" and "2024-12-01" add "1h" every "24h"
// select timeof(last(home.ecowitt.temp)),home.ecowitt.temp,min(home.ecowitt.temp),max(home.ecowitt.temp)'
type Query struct {
	Query *bool   `kernel:"flag,query,Query the DB"`
	DBUrl *string `kernel:"flag,query-db,DB url,http://127.0.0.1:9001"`
}

func (q *Query) Run() error {
	if *q.Query {
		for _, query := range flag.Args() {
			if err := q.query(query); err != nil {
				return err
			}
		}
	}
	return nil
}

func (q *Query) query(query string) error {
	c := client.Client{Url: *q.DBUrl}
	r, err := c.Query(query)
	if err != nil {
		return err
	}

	if r.Status != http.StatusOK {
		return fmt.Errorf("query returned %d: %q", r.Status, r.Message)
	}

	if len(r.Table) > 0 {
		var s []string
		for _, t := range r.Table {
			s = t.String(s)
		}
		fmt.Println(strings.Join(s, "\n"))
	}

	return nil
}
