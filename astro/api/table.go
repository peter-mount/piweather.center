package api

import (
	"github.com/peter-mount/piweather.center/store/api"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"time"
)

func (e *ephemeris) Table(options EphemerisOption) *api.Table {
	caser := cases.Title(language.English)

	//optionNames := options.OptionNames()

	t := api.NewTable()

	addColumns(t, options)

	optionList := options.Split()
	e.ForEach(func(day EphemerisDay) {
		day.ForEach(func(result EphemerisResult) {
			r := t.NewRow()
			tm := result.Time()
			r.AddString(tm, result.Time().Format(time.RFC3339)).
				AddString(tm, caser.String(result.Name()))
			for _, opt := range optionList {
				if v := result.Value(opt); v.IsValid() {
					//r.AddString(tm, v.String())
					r.AddValue(tm, v)
				} else {
					r.AddNull()
				}
			}
		})
	})
	return t
}

func addColumns(t *api.Table, options EphemerisOption) {
	caser := cases.Title(language.English)

	opts := options.Split()

	t.AddColumnGroup("", 2).
		AddColumn(&api.Column{Name: "Date", Type: api.ColumnRight}).
		AddColumn(&api.Column{Name: "Name", Type: api.ColumnLeft})

	var groupKeys []EphemerisOption
	groups := make(map[EphemerisOption]int)
	for _, column := range opts {
		if n, exists := column.Name(); exists {
			s := strings.SplitN(n, ".", 2)
			n = s[len(s)-1]
			t.AddColumn(&api.Column{Name: caser.String(n)})

			group := column.Group()
			if g, exists := groups[group]; exists {
				groups[group] = g + 1
			} else {
				groups[group] = 1
				groupKeys = append(groupKeys, group)
			}
		}
	}

	for _, group := range groupKeys {
		if l, exists := groups[group]; exists {
			s := ""
			if group.IsGroup() {
				s = group.String()
			}
			t.AddColumnGroup(s, l)
		}
	}
}
