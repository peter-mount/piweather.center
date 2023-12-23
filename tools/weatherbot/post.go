package bot

import (
	"errors"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-mastodon"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/util"
	"github.com/peter-mount/piweather.center/weather/value"
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

	for _, val := range row.Values {
		v, err := t.GetValue(val)
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

	if val, ok := v.(value.Value); ok {
		switch {
		case when.LessThan != nil:
			return t.compare(val, when.LessThan, value.LessThan)

		case when.LessThanEqual != nil:
			return t.compare(val, when.LessThanEqual, value.LessThanEqual)

		case when.Equal != nil:
			return t.compare(val, when.Equal, value.Equal)

		case when.NotEqual != nil:
			return t.compare(val, when.NotEqual, value.NotEqual)

		case when.GreaterThanEqual != nil:
			return t.compare(val, when.GreaterThanEqual, value.GreaterThanEqual)

		case when.GreaterThan != nil:
			return t.compare(val, when.GreaterThan, value.GreaterThan)

		}
	}

	return false, nil
}

func (t *Bot) compare(val1 value.Value, val *Value, comp value.Comparator) (bool, error) {
	v, err := t.GetValue(val)
	if err != nil {
		return false, err
	}
	if val2, ok := v.(value.Value); ok {
		return val1.Compare(val2, comp)
	}
	return false, nil
}

func (t *Bot) GetValue(val *Value) (interface{}, error) {
	if val.Value != nil {
		f := *val.Value
		if val.Unit.Unit != "" {
			u, ok := value.GetUnit(val.Unit.Unit)
			if !ok {
				return f, nil
			}
			return val.GetValue(u.Value(f))
		}
	}

	table := t.result.Table[0]
	cell := table.GetCell(val.Col, table.Rows[0])

	switch cell.Type {
	case api.CellNull:
		return "", nil
	case api.CellString:
		return cell.String, nil
	case api.CellNumeric:
		if cell.Value.IsValid() {
			return cell.Value, nil
		}
		return cell.Float, nil
	default:
		return "", nil
	}
}
