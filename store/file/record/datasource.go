package record

import (
	"github.com/peter-mount/piweather.center/util"
	time2 "github.com/peter-mount/piweather.center/util/time"
	"github.com/peter-mount/piweather.center/weather/value"
	"time"
)

func DataSource(rec []Record) (util.DataSource, error) {

	var unit *value.Unit
	var period time2.Period
	yRange := &value.Range{}
	if len(rec) > 0 {
		unit = rec[0].Value.Unit()

		for i := 0; i < len(rec); i++ {
			period = period.Add(rec[i].Time)
			if err := yRange.Add(rec[i].Value); err != nil {
				return nil, err
			}
		}
	}

	return &recordDataSource{
		rec:    rec,
		period: period,
		unit:   unit,
		yRange: yRange,
	}, nil
}

type recordDataSource struct {
	rec    []Record
	period time2.Period
	unit   *value.Unit
	yRange *value.Range
}

func (r *recordDataSource) Size() int {
	return len(r.rec)
}

func (r *recordDataSource) Get(i int) (time.Time, value.Value) {
	e := r.rec[i]
	if v, err := e.Value.As(r.unit); err == nil {
		return e.Time, v
	}
	return e.Time, e.Value
}

func (r *recordDataSource) Period() time2.Period {
	return r.period
}

func (r *recordDataSource) GetYRange() *value.Range {
	return r.yRange
}

func (r *recordDataSource) GetUnit() *value.Unit {
	return r.unit
}

func (r *recordDataSource) ForEach(f func(int, time.Time, value.Value)) {
	for i := 0; i < len(r.rec); i++ {
		t, v := r.Get(i)
		f(i, t, v)
	}
}
