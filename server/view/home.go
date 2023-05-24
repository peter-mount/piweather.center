package view

import (
	"bytes"
	"context"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/piweather.center/io"
	"github.com/peter-mount/piweather.center/server/api"
	"github.com/peter-mount/piweather.center/server/store"
	"github.com/peter-mount/piweather.center/util/template"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type Home struct {
	Rest       *rest.Server         `kernel:"inject"`
	Templates  *template.Manager    `kernel:"inject"`
	Store      *store.Store         `kernel:"inject"`
	ApiInbound *api.EndpointManager `kernel:"inject"`
	meta       *Meta
	lastUpdate time.Time
	mutex      sync.Mutex
}

type Meta struct {
	StartTime time.Time
	OSName    string // OS name
	OSVersion string // OS Version
	OSUptime  string // OS uptime
}

func (m *Meta) Uptime() time.Duration {
	return time.Now().Sub(m.StartTime)
}

func (s *Home) Start() error {
	s.meta = &Meta{
		StartTime: time.Now(),
	}
	_ = io.NewReader().
		ForEachLine(func(line string) error {
			l := strings.SplitN(line, "=", 2)
			if len(l) == 2 {
				v := strings.Trim(l[1], "\"")
				switch l[0] {
				case "VERSION":
					s.meta.OSVersion = v
				case "NAME":
					s.meta.OSName = v
				}
			}
			return nil
		}).
		Open("/etc/os-release")

	s.Rest.Do("/", s.showHome).Methods("GET")
	s.Rest.Do("/status", s.showStatus).Methods("GET")
	s.Rest.Do("/status/system", s.showSystem).Methods("GET")

	return nil
}

func (s *Home) uptime() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if time.Now().Sub(s.lastUpdate) < time.Minute {
		return
	}

	var buf bytes.Buffer
	cmd := exec.Command("uptime", "-p")
	cmd.Stdout = &buf
	err := cmd.Run()
	if err == nil {
		s.meta.OSUptime = buf.String()
	}

	s.lastUpdate = time.Now()
}

func (s *Home) showHome(ctx context.Context) error {
	return s.Templates.Render(ctx, "home.html", map[string]interface{}{
		"navSection": "Home",
		"navLink":    "Home",
	})
}

func (s *Home) showStatus(ctx context.Context) error {
	return s.Templates.Render(ctx, "info/status.html", map[string]interface{}{
		"navSection": "Status",
		"navLink":    "Status",
	})
}

func (s *Home) showSystem(ctx context.Context) error {
	s.uptime()
	return s.Templates.Render(ctx, "info/system.html", map[string]interface{}{
		"meta":       s.meta,
		"navSection": "Status",
		"navLink":    "System",
	})
}
