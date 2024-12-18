package utils

import (
	"github.com/alecthomas/participle/v2"
	"github.com/peter-mount/go-script/errors"
	lang2 "github.com/peter-mount/piweather.center/config/ql"
	"github.com/peter-mount/piweather.center/store/api"
	"strings"
)

type colResolver struct {
	visitor   lang2.Visitor[*colResolver]
	table     *api.Table
	lastGroup *api.ColumnGroup
	path      []string
}

var (
	colResolverVisitor = lang2.NewBuilder[*colResolver]().
		AliasedGroup(crAliasedGroup).
		AliasedExpression(crAliasedExpression).
		Function(crFunction).
		Metric(crMetric).
		Select(crSelect).
		Build()
)

func (r *colResolver) append(s ...string) {
	r.path = append(r.path, s...)
}

func ResolveColumns(table *api.Table, f *lang2.SelectExpression) error {
	return colResolverVisitor.Clone().
		Set(&colResolver{table: table}).
		SelectExpression(f)
}

func crSelect(v lang2.Visitor[*colResolver], f *lang2.Select) error {
	if err := v.SelectExpression(f.Expression); err != nil {
		return err
	}
	return errors.VisitorStop
}

func crAliasedGroup(v lang2.Visitor[*colResolver], f *lang2.AliasedGroup) error {
	var err error

	s := v.Get()

	if s.lastGroup != nil {
		err = participle.Errorf(f.Pos, "Cannot nest groups")
	}

	if err == nil {
		defer func() { s.lastGroup = nil }()

		s.table.AddColumnGroup(f.Name, 0)
		s.lastGroup = s.table.LastColumnGroup()

		err = v.SelectExpression(f.Expressions)
	}

	if err != nil {
		return errors.Error(f.Pos, err)
	}
	return errors.VisitorStop
}

func crAliasedExpression(v lang2.Visitor[*colResolver], f *lang2.AliasedExpression) error {
	var err error

	r := v.Get()

	switch {
	case f.Group != nil:
		err = v.AliasedGroup(f.Group)

	case f.As != "":
		r.table.AddColumn(&api.Column{Name: f.As})
		if r.lastGroup != nil {
			r.lastGroup.Width++
		}

	default:
		r.path = nil
		err = v.Expression(f.Expression)
		if err == nil {
			n := strings.Join(r.path, "")
			r.table.AddColumn(&api.Column{Name: n})

			if r.lastGroup != nil {
				r.lastGroup.Width++
			}
		}
	}

	if err != nil {
		return err
	}

	return errors.VisitorStop
}

func crFunction(v lang2.Visitor[*colResolver], f *lang2.Function) error {
	r := v.Get()

	r.append(f.Name, "(")
	for i, e := range f.Expressions {
		if i > 0 {
			r.append(",")
		}
		if err := v.Expression(e); err != nil {
			return err
		}
	}
	r.append(")")
	return errors.VisitorStop
}

func crMetric(v lang2.Visitor[*colResolver], f *lang2.Metric) error {
	v.Get().append(f.Name)
	return nil
}
