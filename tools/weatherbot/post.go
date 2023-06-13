package bot

import (
	"errors"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-mastodon"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/state"
	value2 "github.com/peter-mount/piweather.center/weather/value"
	"strings"
)

// getPost loads the yaml config and gets the named post to publish
func (t *Bot) getPost() error {
	t.posts = make(map[string]*Post)
	if err := t.ConfigManager.ReadYaml("weatherbot.yaml", &t.posts); err != nil {
		return err
	}

	// Lookup post, show available posts & exit if not found
	t.post = t.posts[*t.Post]
	if *t.Post == "" || t.post == nil {
		a := append([]string{}, "Available posts:")
		for k, e := range t.posts {
			a = append(a, fmt.Sprintf("%s: %s", k, e.Name))
		}
		return errors.New(strings.Join(a, "\n"))
	}

	return nil
}

// createPostText takes the post and generates the Mastodon post text
func (t *Bot) postText() error {
	for tid, thread := range t.post.Threads {
		var str []string

		for _, row := range thread.Table {
			s, err := t.processRow(row)
			if err != nil {
				return err
			}
			str = append(str, s)
		}

		text := thread.Prefix + strings.Join(str, " ") + thread.Suffix
		text = strings.ReplaceAll(text, " \n", "\n")
		text = strings.ReplaceAll(text, "\n ", "\n")
		text = strings.ReplaceAll(text, "  ", " ")

		if *t.Test {
			log.Printf("---- thread %d length %d\n%s\n---- thread %d end\n\n", tid, len(text), text, tid)
		} else {
			if err := t.postToMastodon(mastodon.PostStatus{
				IdempotencyKey: "",
				Text:           text,
				InReplyTo:      0,
				Visibility:     mastodon.VisibilityUnlisted,
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

// processRow processes a Row and returns it as a string
func (t *Bot) processRow(row *Row) (string, error) {
	// filter the row
	if len(row.When) > 0 {
		pass := false
		for _, when := range row.When {
			b, err := t.when(when)
			if err != nil {
				return "", err
			}
			pass = pass || b
		}

		if !pass {
			return "", nil
		}
	}

	// Blank row
	if row.Format == "" {
		return "", nil
	}

	var a []interface{}

	for _, value := range row.Values {
		v, err := t.GetValue(value)
		if err != nil {
			return "", err
		}
		a = append(a, v)
	}

	return util.Sprintf(row.Format, a...), nil
}

func (t *Bot) when(when When) (bool, error) {
	v, err := t.GetValue(when.Value)
	if err != nil {
		return false, err
	}

	if val, ok := v.(value2.Value); ok {
		switch {
		case when.LessThan != nil:
			return t.compare(val, *when.LessThan, value2.LessThan)

		case when.LessThanEqual != nil:
			return t.compare(val, *when.LessThanEqual, value2.LessThanEqual)

		case when.Equal != nil:
			return t.compare(val, *when.Equal, value2.Equal)

		case when.NotEqual != nil:
			return t.compare(val, *when.NotEqual, value2.NotEqual)

		case when.GreaterThanEqual != nil:
			return t.compare(val, *when.GreaterThanEqual, value2.GreaterThanEqual)

		case when.GreaterThan != nil:
			return t.compare(val, *when.GreaterThan, value2.GreaterThan)

		}
	}

	return false, nil
}

func (t *Bot) compare(val1 value2.Value, value Value, comp value2.Comparator) (bool, error) {
	v, err := t.GetValue(value)
	if err != nil {
		return false, err
	}
	if val2, ok := v.(value2.Value); ok {
		return val1.Compare(val2, comp)
	}
	return false, nil
}

func (t *Bot) GetValue(value Value) (interface{}, error) {
	if value.Value != nil {
		f := *value.Value
		if value.Unit.Unit != "" {
			u, ok := value2.GetUnit(value.Unit.Unit)
			if !ok {
				return f, nil
			}
			return value.GetValue(u.Value(f))
		}
	}

	var v interface{}
	switch value.Type {
	case ValueTime:
		return t.station.Meta.Time.Format("2006 Jan 02 15:04 MST"), nil

	case ValueStationName:
		return t.station.Meta.Name, nil

	default:
		m := t.getMeasurement(value.Sensor)
		if m == nil {
			return fmt.Sprintf("sensor %q missing", value.Sensor), nil
		}

		valueRange := m.Current10
		switch value.Range {
		case RangeCurrent, "":
			valueRange = m.Current10
		case RangePrevious:
			valueRange = m.Previous10
		case RangeHour:
			valueRange = m.Hour
		case RangeHour24:
			valueRange = m.Hour24
		case RangeToday:
			valueRange = m.Today
		}

		var f state.RoundedFloat
		switch value.Type {
		case "", ValueLatest:
			f = m.Current.Value

		case ValuePrevious:
			f = m.Previous.Value

		case ValueTrend:
			v = m.Trends.Current.Char

		case ValueMin:
			f = valueRange.Min

		case ValueMax:
			f = valueRange.Max

		case ValueMean:
			f = valueRange.Mean

		case ValueTotal:
			f = valueRange.Total

		case ValueCount:
			f = state.RoundedFloat(valueRange.Count)

		default:
			return "", fmt.Errorf("unsupported type %q for sensor %q", value.Type, value.Sensor)
		}

		// v not set (see ValueTrend) then use f
		if v == nil {
			// Apply Factor if present
			if value.Factor != 0.0 {
				f = f * state.RoundedFloat(value.Factor)
			}

			// Handle units, src is from Measurement, dest from Value
			if src, srcOk := value2.GetUnit(m.Unit); srcOk {
				v0 := src.Value(float64(f))
				if v1, err := value.GetValue(v0); err != nil {
					// Cannot transform between requested units
					v = err.Error()
				} else {

					v = v1
				}
			} else {
				v = f
			}
		}

		// Handle "state.RoundedFloat=0" error when using %f
		if rf, ok := v.(state.RoundedFloat); ok {
			v = float64(rf)
		}

		return v, nil
	}
}
