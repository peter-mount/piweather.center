package view

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/util/template"
	_ "github.com/peter-mount/piweather.center/weather/measurment"
	"github.com/peter-mount/piweather.center/weather/value"
)

type Units struct {
	Rest      *rest.Server      `kernel:"inject"`
	Templates *template.Manager `kernel:"inject"`
}

func (s *Units) Start() error {
	s.Rest.Do("/info/units", s.showUnits).Methods("GET")
	s.Rest.Do("/info/transforms", s.showTransforms).Methods("GET")

	return nil
}

func (s *Units) showUnits(ctx context.Context) error {
	return s.Templates.Render(ctx, "info/units.html", map[string]interface{}{
		"navSection": "Status",
		"navLink":    "Units",
		"groups":     value.GetGroups(),
	})
}

func (s *Units) showTransforms(ctx context.Context) error {
	return s.Templates.Render(ctx, "info/transforms.html", map[string]interface{}{
		"navSection": "Status",
		"navLink":    "Transforms",
		"groups":     value.GetGroups(),
		"transforms": value.GetTransforms(),
	})
}
