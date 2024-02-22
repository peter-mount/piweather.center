package model

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
)

func init() {
	f := func() Instance { return &Query{} }
	Register("wind-rose", f)
	Register("wind-rose-line", f)
	Register("wind-rose-simple", f)
}

// Query represents a custom db query
type Query struct {
	Component `yaml:",inline"`
	Label     string      `yaml:"label"`           // optional label
	Query     string      `yaml:"query,omitempty"` // Query to get values
	Result    *api.Result `yaml:"-"`               // Cache of last query result
}

func (q *Query) Init(url string) {
	log.Printf("Query %q %q = %q", q.ID, q.Query, url)
	if q.Query != "" {
		c := client.Client{Url: url}
		r, err := c.Query(q.Query)
		if err == nil {
			q.Result = r
			log.Printf("Query %q %q = %v", q.ID, q.Query, q.Result.WindRose)
		} else {
			log.Printf("Query %q %q = %v", q.ID, q.Query, err)
		}
	}
}
