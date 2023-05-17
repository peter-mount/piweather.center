package view

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/util/template"
	_ "github.com/peter-mount/piweather.center/weather/measurment"
	"github.com/peter-mount/piweather.center/weather/value"
	"sort"
)

type Units struct {
	Rest       *rest.Server      `kernel:"inject"`
	Templates  *template.Manager `kernel:"inject"`
	categories []string
	units      map[string][]*value.Unit
}

func (s *Units) Start() error {
	// Populate unit map grouped by category
	s.units = make(map[string][]*value.Unit)
	for _, u := range value.GetUnits() {
		s.units[u.Category()] = append(s.units[u.Category()], u)
	}

	// Build category list
	for c, _ := range s.units {
		s.categories = append(s.categories, c)
	}
	sort.SliceStable(s.categories, func(i, j int) bool {
		return s.categories[i] < s.categories[j]
	})

	// Sort each group
	for _, u := range s.units {
		sort.SliceStable(u, func(i, j int) bool {
			return u[i].ID() < u[j].ID()
		})
	}

	s.Rest.Do("/info/units", s.showUnits).Methods("GET")

	return nil
}

func (s *Units) showUnits(ctx context.Context) error {
	return s.Templates.Render(ctx, "info/units.html", map[string]interface{}{
		"navSection": "Status",
		"navLink":    "Units",
		"cats":       s.categories,
		"units":      s.units,
	})
}
