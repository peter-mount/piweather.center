package bot

import (
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/weather/state"
	"strings"
)

func (t *Bot) Run() error {
	err := t.createPost()
	if err != nil {
		return err
	}

	return nil
}

func (t *Bot) createPost() error {
	for _, thread := range t.post.Threads {
		var str []string

		for _, row := range thread.Table {
			s, err := t.processRow(row)
			if err != nil {
				return err
			}
			str = append(str, s)
		}

		text := thread.Prefix + strings.Join(str, "\n") + thread.Suffix

		// Replace errors when data is unavailable with "N/A"
		text = t.cleanup.ReplaceAllString(text, "N/A")

		if *t.Test {
			log.Printf("Post length: %d\n%s\n", len(text), text)
		}
	}

	return nil
}

func (t *Bot) getMeasurement(id string) *state.Measurement {
	for _, m := range t.station.Measurements {
		if m.ID == id {
			return m
		}
	}
	return nil
}

func (t *Bot) processRow(row *Row) (string, error) {
	// Blank row
	if row.Format == "" {
		return "", nil
	}

	var a []interface{}

	for _, value := range row.Values {
		var v interface{}
		switch value.Type {
		case ValueTime:
			v = t.station.Meta.Time

		case ValueStationName:
			v = t.station.Meta.Name

		default:
			m := t.getMeasurement(value.Sensor)
			if m == nil {
				return fmt.Sprintf("sensor %q missing", value.Sensor), nil
			}

			var f state.RoundedFloat
			switch value.Type {
			case "", ValueLatest:
				f = m.Current.Value

			case ValueTrend:
				v = m.Trends.Current.Char

			case ValueMin:
				f = m.Minute10.Min

			case ValueMax:
				f = m.Minute10.Max

			case ValueMean:
				f = m.Minute10.Mean

			default:
				return "", fmt.Errorf("unsupported type %q for sensor %q", value.Type, value.Sensor)
			}

			if v == nil {
				if value.Factor != 0.0 {
					f = f * state.RoundedFloat(value.Factor)
				}
				v = f
			}
		}
		a = append(a, v)
	}

	return fmt.Sprintf(row.Format, a...), nil
}
