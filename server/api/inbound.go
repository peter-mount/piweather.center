package api

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"github.com/peter-mount/piweather.center/util/template"
	"path"
	"sort"
	"strings"
	"sync"
)

// Inbound presents an api under /api/inbound which will accept data from an
// HTTP request and pass it onto a Task.
//
// Examples of this are:
//   - Stations submitting data in the Ecowitt protocol
//   - Stations submitting data in the Weather Underground protocol
type Inbound struct {
	Rest      *rest.Server      `kernel:"inject"`
	Templates *template.Manager `kernel:"inject"`
	endPoints map[string][]*endpoint
	mutex     sync.Mutex
}

type endpoint struct {
	endpoint EndpointEntry
	task     task.Task
}

type Endpoints struct {
	Category  string
	Endpoints []EndpointEntry
}

// EndpointEntry contains info exposed to the templates
type EndpointEntry struct {
	Category string
	Id       string
	Name     string
	Endpoint string
	Method   string
	Protocol string
}

// Identical returns true of both entries have the same Endpoint and Method
func (a EndpointEntry) Identical(b EndpointEntry) bool {
	return a.Endpoint == b.Endpoint && strings.ToUpper(a.Method) == strings.ToUpper(b.Method)
}

func (s *Inbound) PostInit() error {
	s.endPoints = make(map[string][]*endpoint)
	return nil
}

func (s *Inbound) Start() error {
	s.Rest.Do("/status/endpoints", s.showEndpoints).Methods("GET")
	return nil
}

// RegisterEndpoint registers a new endpoint under /api/inbound which will be sent to a task
func (s *Inbound) RegisterEndpoint(category, pathName, id, name, method, protocol string, t task.Task) error {
	// Sanitize category and pathName
	category = strings.ToLower(strings.TrimSpace(category))

	pathName = path.Clean(pathName)
	if pathName == "." || pathName == "/" {
		return fmt.Errorf("ecowitt path invalid")
	}

	pathName = path.Join("/api", category, pathName)

	e := &endpoint{
		task: t,
		endpoint: EndpointEntry{
			Category: category,
			Id:       id,
			Name:     name,
			Endpoint: pathName,
			Protocol: protocol,
			Method:   method,
		},
	}

	if err := s.addEndpoint(e); err != nil {
		return err
	}

	s.Rest.Do(e.endpoint.Endpoint, e.task).Methods(e.endpoint.Method)

	return nil
}

func (s *Inbound) addEndpoint(e *endpoint) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cat, exists := s.endPoints[e.endpoint.Category]
	if exists {
		// Verify that the path and method are not already in use by another task
		for _, ep := range cat {
			if e.endpoint.Identical(ep.endpoint) {
				return fmt.Errorf("path %s:%s already in use", e.endpoint.Method, e.endpoint.Endpoint)
			}
		}
	}

	s.endPoints[e.endpoint.Category] = append(s.endPoints[e.endpoint.Category], e)
	return nil
}

func (s *Inbound) getEndpoints() []Endpoints {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var r []Endpoints
	for cat, ep := range s.endPoints {
		l := Endpoints{Category: cat}
		for _, e := range ep {
			l.Endpoints = append(l.Endpoints, e.endpoint)
		}
		sort.SliceStable(l.Endpoints, func(i, j int) bool {
			return l.Endpoints[i].Endpoint < l.Endpoints[j].Endpoint
		})
		r = append(r, l)
	}
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Category < r[j].Category
	})
	return r
}

func (s *Inbound) showEndpoints(ctx context.Context) error {
	return s.Templates.Render(ctx, "info/endpoints.html", map[string]interface{}{
		"endpoints":  s.getEndpoints(),
		"navSection": "Status",
		"navLink":    "System",
	})
}
