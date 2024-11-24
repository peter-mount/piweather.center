package publisher

import (
	"encoding/json"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/sensors/reading"
	"github.com/peter-mount/piweather.center/store/api"
	"strings"
)

// logPublisher is a Publisher which will log the Reading as JSON to the log
func logPublisher(r *reading.Reading) error {
	if log.IsVerbose() {
		b, err := json.Marshal(r)
		if err != nil {
			return err
		}
		log.Println(string(b))
	}
	return nil
}

// filterEmptyReadings wraps a Publisher which will only be invoked when the provided Publisher when
// the passed Reading contains values.
func filterEmptyReadings(p Publisher) Publisher {
	if p == nil {
		return nil
	}

	return func(r *reading.Reading) error {
		if !r.IsEmpty() {
			return p(r)
		}
		return nil
	}
}

// setId sets the Results ID before passing to the
func setId(id string) Publisher {
	return func(r *reading.Reading) error {
		if r != nil {
			r.ID = id
		}
		return nil
	}
}

type DBPublisher interface {
	PublishMetric(api.Metric) error
}

func dbPublisher(p DBPublisher) Publisher {
	return func(r *reading.Reading) error {
		for key, value := range r.Readings {
			metric := api.NewMetric(
				strings.Join([]string{r.ID, key}, "."),
				r.Time,
				value,
			)
			err := p.PublishMetric(metric)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
