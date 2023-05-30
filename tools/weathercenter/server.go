package weathercenter

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-kernel/v2/rest"
	_ "github.com/peter-mount/piweather.center/server/menu"
	_ "github.com/peter-mount/piweather.center/server/view"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/util/template"
	"path/filepath"
)

// Server represents the primary service running the fully integrated weather station.
type Server struct {
	Rest      *rest.Server      `kernel:"inject"`
	Config    station.Config    `kernel:"inject"`
	Templates *template.Manager `kernel:"inject"`
}

func (s *Server) Start() error {
	// Static content to the webserver
	rootDir := filepath.Dir(s.Templates.GetRootDir())
	staticDir := filepath.Join(rootDir, "static")
	log.Printf("Static content: %s", staticDir)
	s.Rest.Static("/static", staticDir)

	return nil
}
