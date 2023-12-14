package api

import "sync"

// Listener provides a thread safe pool of Handlers which can listen for metrics,
// usually provided by a single data source like RabbitMQ.
//
// End use, create an instance using NewListener(), then call `go Run()` so that the
// listener runs in its own goroutine.
//
// Then add Handler's at any time and Notify() to send a Metric to all of them.
type Listener interface {
	// Add a Handler to the Listener
	Add(Handler)
	// Notify all Handler's of a Metric
	Notify(Metric)
	// Run the Listener.
	Run()
}

// Handler that will receive a Metric
type Handler func(Metric)

type listener struct {
	mutex        sync.Mutex
	running      bool
	listeners    []Handler
	listenerChan chan Metric
}

// NewListener creates a new Listener.
// Once created you need to call Run() in its own goroutine to cause it to run in the background.
func NewListener() Listener {
	return &listener{listenerChan: make(chan Metric)}
}

func (s *listener) Add(l Handler) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.listeners = append(s.listeners, l)
}

func (s *listener) getListeners() []Handler {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Make a copy so caller cannot modify the true slice outside the lock
	a := make([]Handler, len(s.listeners))
	copy(a, s.listeners)

	return a
}

func (s *listener) Notify(m Metric) {
	s.listenerChan <- m
}

func (s *listener) IsRunning() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.running
}

func (s *listener) setRunning(v bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.running = v
}

func (s *listener) Run() {
	if !s.IsRunning() {
		s.setRunning(true)
		defer s.setRunning(false)

		for {
			metric := <-s.listenerChan
			for _, l := range s.getListeners() {
				l(metric)
			}
		}
	}
}
