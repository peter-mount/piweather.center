package rtl433

import (
	"strings"
	"sync"
)

// RTL433 is a pseudo bus which uses the rtl_433 utility to receive decoded
// messages from sensors using an SDR (Software Defined Modem).
//
// There are multiple "buses" available, each bus identified by the central
// frequency.
//
// Each "bus" then has it's own instance of rtl_433 running, hence this requires
// an SDR for each frequency.
type RTL433 struct {
	mutex sync.Mutex
	buses map[string]*Frequency
}

func (s *RTL433) PostInit() error {
	s.buses = make(map[string]*Frequency)
	return nil
}

func (s *RTL433) Start() error {
	for _, f := range s.getFrequencies() {
		go s.getFrequency(f).start()
	}
	return nil
}

func (s *RTL433) Stop() {}

func (s *RTL433) getFrequencies() []string {
	var f []string
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for k, _ := range s.buses {
		f = append(f, k)
	}
	return f
}

func GetFrequency(f string) string {
	f = strings.TrimSpace(f)
	if f == "" {
		return "433.92M"
	}
	return f
}

func (s *RTL433) getFrequency(f string) *Frequency {
	f = GetFrequency(f)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.buses[f]
}

func (s *RTL433) getOrCreateFrequency(f string) *Frequency {
	f = GetFrequency(f)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	freq, exists := s.buses[f]
	if !exists {
		freq = &Frequency{
			frequency: f,
		}
		s.buses[f] = freq
	}
	return freq
}

func (s *RTL433) AddListener(freq string, l *Listener) {
	s.getOrCreateFrequency(freq).AddListener(l)
}
