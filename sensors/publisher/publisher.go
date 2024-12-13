package publisher

import (
	"github.com/peter-mount/piweather.center/sensors/reading"
)

// Publisher processes Reading's
type Publisher func(*reading.Reading) error

// Then will return a Publisher which will call this one then, if no error a subsequent Publisher
func (a Publisher) Then(b Publisher) Publisher {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(r *reading.Reading) error {
		err := a(r)
		if err == nil {
			err = b(r)
		}
		return err
	}
}

// Do will invoke a Publisher but will do nothing if the Publisher is nil
func (a Publisher) Do(r *reading.Reading) error {
	if a != nil {
		return a(r)
	}
	return nil
}

// Of will return a Publisher which will invoke each Publisher in sequence
func Of(pubs ...Publisher) Publisher {
	var a Publisher

	for _, b := range pubs {
		a = a.Then(b)
	}

	return a
}
