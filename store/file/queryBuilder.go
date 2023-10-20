package file

import (
	"github.com/peter-mount/go-kernel/v2/log"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"time"
)

const (
	dayDuration = 24 * time.Hour
)

func (s *store) Query(metric string) QueryBuilder {
	return &queryBuilder{
		store:  s,
		metric: metric,
	}
}

type QueryBuilder interface {
	// Today sets the search range to span from Local Midnight for 24 hours
	Today() QueryBuilder
	// TodayUTC sets the search range to span from Midnight UTC for 24 hours
	TodayUTC() QueryBuilder
	// DayFrom sets the search range to span for 24 hours from the specified time
	DayFrom(from time.Time) QueryBuilder
	// Between sets the time range of the query
	Between(from, to time.Time) QueryBuilder
	// Filter sets the filter to use. If one already exists then this will be or'ed with it.
	Filter(Filter) QueryBuilder
	// Build the final Query
	Build() Query
}

type queryBuilder struct {
	store  *store
	filter Filter
	metric string    // Metric to query
	start  time.Time // start time inclusive
	end    time.Time // end time exclusive
}

func (b *queryBuilder) Build() Query {
	var query Query

	// The scanner to cover the start and end date ranges
	query = &queryScanner{
		metric: b.metric,
		start:  b.start,
		end:    b.end.Truncate(dayDuration).Add(dayDuration),
		date:   b.start.Truncate(dayDuration),
		store:  b.store,
	}

	log.Printf("query from %s to %s", b.start, b.end)

	// As queryScanner works on entire day files, add a default filter to limit within
	// the required period. If a filter is defined in the builder then and it with the Between filter.
	filter := Between(b.start, b.end)
	if b.filter != nil {
		filter = filter.And(b.filter)
	}

	query = &filteredQuery{
		query:  query,
		filter: filter,
	}

	// TODO add aggregator here?

	return query
}

func (b *queryBuilder) Between(from, to time.Time) QueryBuilder {
	if from.After(to) {
		from, to = to, from
	}
	b.start = from.UTC()
	b.end = to.UTC()
	return b
}

func (b *queryBuilder) Today() QueryBuilder {
	return b.DayFrom(time2.LocalMidnight(time.Now()))
}

func (b *queryBuilder) TodayUTC() QueryBuilder {
	return b.DayFrom(time.Now().UTC())
}

func (b *queryBuilder) DayFrom(from time.Time) QueryBuilder {
	return b.Between(from, from.Add(24*time.Hour))
}

func (b *queryBuilder) Filter(f Filter) QueryBuilder {
	if b.filter == nil {
		b.filter = f
	} else {
		b.filter = b.filter.Or(f)
	}
	return b
}
