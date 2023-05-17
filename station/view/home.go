package view

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/util/template"
)

type Home struct {
	Rest      *rest.Server      `kernel:"inject"`
	Templates *template.Manager `kernel:"inject"`
}

func (s *Home) Start() error {

	s.Rest.Do("/", s.showHome).Methods("GET")

	return nil
}

func (s *Home) showHome(ctx context.Context) error {
	return s.Templates.Render(ctx, "home.html", map[string]interface{}{
		"navSection": "Home",
		"navLink":    "Home",
	})
}
