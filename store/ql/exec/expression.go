package exec

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-script/errors"
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/store/ql/functions"
	"time"
)

func expression(v lang2.Visitor[*Executor], s *lang2.Expression) error {
	ex := v.Get()

	var err error

	// If offset defined, temporarily adjust the current time by that offset
	if s.Using != "" || s.Modifier != nil {
		ex.save()
		defer ex.restore()

		// Resolve the modifier if we are declaring using
		mod := s.Modifier
		if s.Using != "" {
			uDef, exists := ex.using[s.Using]
			if !exists {
				// Should not happen as we checked before running
				return errors.Errorf(s.Pos, "panic: %q missing", s.Using)
			}
			mod = uDef.Modifier
		}

		for _, e := range mod {
			if err == nil {
				err = v.ExpressionModifier(e)
			}
		}
	}

	if err == nil {
		switch {
		case s.Metric != nil:
			err = v.Metric(s.Metric)
		case s.Function != nil:
			err = v.Function(s.Function)
		}
	}

	if err != nil {
		return err
	}

	return errors.VisitorStop
}

func expressionModifier(v lang2.Visitor[*Executor], s *lang2.ExpressionModifier) error {
	ex := v.Get()

	var err error

	if s.Offset != nil {
		ex.time = ex.time.Add(s.Offset.Duration(ex.timeRange.Every))
	}

	if s.Range != nil {
		if s.Range.IsRow() {
			panic("range not implemented") // FIXME
			//err = s.Range.SetTime(ex.time, ex.timeRange.Every, v)
		}

		r := s.Range.Range()
		ex.time = r.From
		ex.timeRange.Every = r.Duration()
	}

	return err
}

func aliasedExpression(v lang2.Visitor[*Executor], s *lang2.AliasedExpression) error {
	ex := v.Get()

	// If a group is present then jump into that
	if s.Group != nil {
		ex.save()
		defer ex.restore()
		ex.inGroup = true

		if err := v.AliasedGroup(s.Group); err != nil {
			return err
		}
		return errors.VisitorStop
	}

	// Call summarize first
	if err := v.Summarize(s.Summarize); err != nil {
		return err
	}

	ex.resetStack()

	err := v.Expression(s.Expression)

	val, ok := ex.Pop()

	// If invalid but have values attached then get the last value in the Set.
	// Required with metrics without an aggregation function around them
	if !val.IsTime && !val.Value.IsValid() && len(val.Values) > 0 {
		val = functions.InitialLast(val)
	}

	switch {
	case err != nil:
		log.Println(err)
		ex.row.AddNull()

	case !ok,
		val.IsNull():
		ex.row.AddNull()

	case val.IsTime:
		ex.row.AddDynamic(val.Time, val.Time.Format(time.RFC3339))

	default:
		col := ex.table.Columns[ex.row.Size()]
		val1, err := col.Transform(val.Value)
		if err != nil {
			return err
		}
		// Add to table
		ex.row.AddValue(val.Time, val1)

		// Add to summary if it's being summarized
		if ex.summary.IsAggregated(ex.selectColumn) {
			ex.summary.Set(ex.selectColumn, val1)
		}
	}

	// Now we are done with this column, increment for the next one
	ex.selectColumn++

	return errors.VisitorStop
}
