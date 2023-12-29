package model

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
)

func init() {
	f := func() Instance { return &Query{} }
	Register("wind-rose", f)
}

// Query represents a custom db query
type Query struct {
	Component `yaml:",inline"`
	Label     string      `yaml:"label"`           // optional label
	Query     string      `yaml:"query,omitempty"` // Query to get values
	Result    *api.Result `yaml:"-"`               // Cache of last query result
}

func (q *Query) Init(url string) {
	if q.Query != "" {
		c := client.Client{Url: url}
		r, err := c.Query(q.Query)
		if err == nil {
			q.Result = r
		} else {
			log.Printf("Query %q %q = %v", q.ID, q.Query, err)
		}
	}
}
