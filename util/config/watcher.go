package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/peter-mount/go-kernel/v2/log"
	"os"
	"path/filepath"
	"strings"
)

type updater struct {
	factory      Factory
	updater      Updater
	unmarshaller Unmarshaller
}

// Factory creates a struct to pass to an Unmarshaller then an Updater
type Factory func() any

// Updater is a function that will be notified of changes
type Updater func(fsnotify.Event, any) error

// Unmarshaller will be called to unmarshal a file when it changes.
// Usually this is yaml.Unmarshal but it can be a custom format
type Unmarshaller func([]byte, interface{}) error

func (m *manager) WatchDirectory(d string, f Factory, u Updater, um Unmarshaller) {
	d = filepath.Clean(d) + pathSep
	m.add(d, f, u, um)
}

func (m *manager) ReadAndWatch(n string, f Factory, u Updater, um Unmarshaller) error {
	// Add the Updater
	m.add(n, f, u, um)

	// Implicitly read the file here
	fn := filepath.Join(m.EtcDir(), n)
	b, err := os.ReadFile(fn)

	if err == nil || os.IsNotExist(err) {
		o := f()

		// We don't want to pass NotExist error but only unmarshal if we have read the file
		if err == nil {
			err = um(b, o)
		} else {
			err = nil
		}

		if err == nil {
			err = u(fsnotify.Event{Name: fn}, o)
		}

	}
	return err
}

func (m *manager) run() {
	for {
		select {
		case event := <-m.watcher.Events:
			log.Printf("event: %d %q", event.Op, event.Name)
			if event.Op&fsnotify.Write == fsnotify.Write {
				m.notify(event)
			}
		case err := <-m.watcher.Errors:
			log.Printf("Config: watcher error: %v", err)
		}
	}
}

func (m *manager) add(n string, f Factory, u Updater, um Unmarshaller) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.watching[n] = append(m.watching[n], &updater{
		factory:      f,
		updater:      u,
		unmarshaller: um,
	})

	if len(m.watching[n]) == 1 {
		p := filepath.Join(m.EtcDir(), n)
		err := m.watcher.Add(p)
		if err == nil {
			log.Printf("Config: watching %q", n)
		} else {
			log.Printf("Config: watching %q error %v", p, err)
		}
	}
}

func (m *manager) get(n string) []*updater {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for {
		u, ok := m.watching[n]
		if ok {
			return u
		}

		n = filepath.Dir(n)
		if n == "." || n == m.EtcDir() {
			break
		}
	}

	return nil
}

func (m *manager) notify(event fsnotify.Event) {
	n := strings.TrimPrefix(event.Name, m.EtcDir()+pathSep)
	watchers := m.get(n)
	if len(watchers) > 0 {
		switch event.Op {
		// Create or Write, unmarshal new contents and notify
		case fsnotify.Create, fsnotify.Write:
			log.Printf("Config: update %q", event.Name)

			// Read the file just once
			b, err := os.ReadFile(event.Name)
			if err != nil {
				log.Printf("Config: read failed %q %v", event.Name, err)
				return
			}

			for _, w := range watchers {
				// Unmarshal to the required object for this specific watcher
				o := w.factory()
				err = w.unmarshaller(b, o)

				if err == nil {
					// Notify watcher of the update
					err = w.updater(event, o)
				}
				if err != nil {
					// Log but continue to next watcher
					log.Printf("Config: notify error %q %v", event.Name, err)
				}
			}

		// Remove or Rename, notify with nil value
		case fsnotify.Remove, fsnotify.Rename:
			for _, w := range watchers {
				_ = w.updater(event, nil)
			}

		default:
			// Ignore for now
		}
	}
}
