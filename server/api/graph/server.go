package graph

import (
	"context"
	"github.com/peter-mount/go-kernel/v2/util/task"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"path"
	"time"
)

type Generator func(start, end time.Time, ctx context.Context) error

type GeneratorFactory func(path string, g Generator) (string, string, task.Task)

func ServeLatest(p string, g Generator) (string, string, task.Task) {
	return path.Join(p, "latest.svg"),
		" latest",
		func(ctx context.Context) error {
			now := time.Now()
			start := now.Truncate(time.Hour)
			end := start.Add(time.Hour)
			return g(start, end, ctx)
		}
}

// ServeHour returns a Task that will call a Generator with the start, end times
// being the start and end of the current hour.
func ServeHour(p string, g Generator) (string, string, task.Task) {
	return path.Join(p, "hour.svg"),
		" for current hour",
		func(ctx context.Context) error {
			now := time.Now()
			start := now.Truncate(time.Hour)
			end := start.Add(time.Hour)
			return g(start, end, ctx)
		}
}

// ServeDay returns a Task that will call a Generator with the end being the end
// of the current hour. The start will be 24 hours before the end.
//
// e.g. If it's currently 09:10 then the range will be from 10:00 yesterday until
// 10:00 today.
func ServeDay(p string, g Generator) (string, string, task.Task) {
	return path.Join(p, "day.svg"),
		" for last 24 hours",
		func(ctx context.Context) error {
			// End is end of current hour
			end := time.Now().Truncate(time.Hour).Add(time.Hour)
			start := end.Add(-24 * time.Hour)
			return g(start, end, ctx)
		}
}

// ServeToday returns a Task that will call a Generator with the start being
// local midnight and the end 24 hours later.
//
// e.g. If it's currently 09:10 then start will be 00:00 that morning and end
// will be 00:00 the following evening.
func ServeToday(p string, g Generator) (string, string, task.Task) {
	return path.Join(p, "today.svg"),
		" since midnight",
		func(ctx context.Context) error {
			now := time.Now()
			// Start at beginning of the current local day
			start := time2.LocalMidnight(now)
			end := start.Add(time.Hour * 24)
			return g(start, end, ctx)
		}
}
