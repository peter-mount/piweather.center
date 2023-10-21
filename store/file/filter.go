package file

import (
	"github.com/peter-mount/piweather.center/store/file/record"
	"time"
)

type Filter func(record.Record) bool

func Of(filters ...Filter) Filter {
	switch len(filters) {
	case 0:
		return False
	case 1:
		return filters[0]
	default:
		var f Filter
		for i, filter := range filters {
			if i == 0 {
				f = filter
			} else {
				f = f.Or(filter)
			}
		}
		return f
	}
}

func (a Filter) And(b Filter) Filter {
	return func(record record.Record) bool {
		return a(record) && b(record)
	}
}

func (a Filter) Or(b Filter) Filter {
	return func(record record.Record) bool {
		return a(record) || b(record)
	}
}

func True(_ record.Record) bool {
	return true
}

func False(_ record.Record) bool {
	return false
}

func Between(s, e time.Time) Filter {
	if s.After(e) {
		s, e = e, s
	}
	return func(record record.Record) bool {
		// don't use After(s) && Before(e) here as we want to match when equals on both as well
		return !record.Time.Before(s) && record.Time.Before(e)
	}
}

func After(t time.Time) Filter {
	return func(record record.Record) bool {
		// don't use After(t) here as we want to match when equals t as well
		return !record.Time.Before(t)
	}
}

func Before(t time.Time) Filter {
	return func(record record.Record) bool {
		// don't use After(t) here as we want to match when equals t as well
		return !record.Time.After(t)
	}
}
