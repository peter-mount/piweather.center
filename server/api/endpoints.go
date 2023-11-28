package api

import (
	"context"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/rest"
	"github.com/peter-mount/go-kernel/v2/util/task"
	"sort"
	"strings"
	"sync"
	"time"
)

// EndpointManager acts as a frontend to the Rest server and provides a common
// interface for both inbound feeds (using http or message brokers) as well as
// for svg and image sources.
type EndpointManager struct {
	Rest      *rest.Server `kernel:"inject"`
	endPoints map[string][]*endpoint
	mutex     sync.Mutex
}

type endpoint struct {
	endpoint EndpointEntry
	task     task.Task
}

func (e *endpoint) invoke(ctx context.Context) error {
	e.endpoint.LastCall = time.Now()
	e.endpoint.NumCalls++
	return e.task.Do(ctx)
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
	LastCall time.Time
	NumCalls int64
}

// Identical returns true of both entries have the same Endpoint and Method
func (a EndpointEntry) Identical(b EndpointEntry) bool {
	return a.Endpoint == b.Endpoint && strings.ToUpper(a.Method) == strings.ToUpper(b.Method)
}

func (s *EndpointManager) PostInit() error {
	s.endPoints = make(map[string][]*endpoint)
	return nil
}

// RegisterHttpEndpoint registers a new endpoint with the system webserver which will be sent to a task.
func (s *EndpointManager) RegisterHttpEndpoint(category, pathName, id, name, method, protocol string, t task.Task) error {
	e, _, err := s.registerEndpoint(category, pathName, id, name, method, protocol, t)
	if err == nil {
		s.Rest.Do(e.endpoint.Endpoint, e.invoke).Methods(e.endpoint.Method)
	}
	return err
}

// RegisterEndpoint registers a new endpoint will be sent to a task.
// This doesn't do much other than count the number of times the task is run and show it
// on the status page.
// This is usually used for mqtt and amqp queues
func (s *EndpointManager) RegisterEndpoint(category, pathName, id, name, method, protocol string, t task.Task) (task.Task, error) {
	_, rt, err := s.registerEndpoint(category, pathName, id, name, method, protocol, t)
	return rt, err
}

func (s *EndpointManager) registerEndpoint(category, pathName, id, name, method, protocol string, t task.Task) (*endpoint, task.Task, error) {
	// Sanitize category and pathName
	category = strings.ToLower(strings.TrimSpace(category))

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
		return nil, nil, err
	}

	return e, e.invoke, nil
}

func (s *EndpointManager) addEndpoint(e *endpoint) error {
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

func (s *EndpointManager) getEndpointSlice() []Endpoints {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var r []Endpoints
	for cat, ep := range s.endPoints {
		l := Endpoints{Category: cat}
		for _, e := range ep {
			l.Endpoints = append(l.Endpoints, e.endpoint)
		}
		r = append(r, l)
	}
	return r
}

func (s *EndpointManager) getEndpoints() []Endpoints {
	r := s.getEndpointSlice()

	// Sort by category
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Category < r[j].Category
	})

	// Sort each category by ID, then by endpoint length then endpoint
	for _, cat := range r {
		sort.SliceStable(cat.Endpoints, func(i, j int) bool {
			a, b := cat.Endpoints[i], cat.Endpoints[j]
			if a.Id != b.Id {
				return a.Id < b.Id
			}
			as := strings.Split(a.Endpoint, "/")
			bs := strings.Split(b.Endpoint, "/")

			if len(as) != len(bs) {
				return len(as) < len(bs)
			}

			return a.Endpoint < b.Endpoint
		})
	}

	return r
}
