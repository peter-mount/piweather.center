package bot

import (
	"errors"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/io"
	"github.com/peter-mount/piweather.center/weather/state"
	"path/filepath"
	"strings"
)

// getPost loads the yaml config and gets the named post to publish
func (t *Bot) getPost() error {
	t.posts = make(map[string]*Post)
	if err := io.NewReader().
		Yaml(&t.posts).
		Open(filepath.Join(*t.RootDir, "weatherbot.yaml")); err != nil {
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
func (t *Bot) createPostText() error {
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

// processRow processes a Row and returns it as a string
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
