package rtl433

import (
	"github.com/peter-mount/go-kernel/v2/log"
	io2 "github.com/peter-mount/piweather.center/util/io"
	"os"
	"os/exec"
	"sync"
)

// Frequency handles everything needed for a specific frequency
type Frequency struct {
	mutex     sync.Mutex
	frequency string
	listeners []*Listener
}

func (f *Frequency) AddListener(listener *Listener) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.listeners = append(f.listeners, listener)
}

func (f *Frequency) getListeners() []*Listener {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	return f.listeners
}

func (f *Frequency) Process(m *Message) {
	for _, l := range f.getListeners() {
		l.Accept(m)
	}
}

func (f *Frequency) start() {
	log.Printf("Starting RTL433 frequency %s", f.frequency)

	cmd := exec.Command("rtl_433", "-f", f.frequency, "-F", "json")

	r, err := cmd.StdoutPipe()

	if err == nil {
		err = cmd.Start()
	}

	if err == nil {
		err = io2.NewReader().
			ForEachLine(func(line string) error {
				log.Println(f.frequency, line)

				m, err := NewMessage([]byte(line))
				if err == nil {
					f.Process(m)
				}
				return nil
			}).Read(r)
	}

	if err == nil {
		log.Println(f.frequency, "rtl_433 waiting")
		err = cmd.Wait()
	}

	if err != nil {
		log.Println(f.frequency, err)
		os.Exit(1)
	}
}
