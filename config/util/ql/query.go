package ql

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type Query struct {
	Pos        lexer.Position
	QueryRange *QueryRange       `parser:"@@"`
	Using      *UsingDefinitions `parser:"(@@)?"`
	Histogram  []*Histogram      `parser:"( ( @@ )+"`
	WindRose   []*WindRose       `parser:"| ( @@ )+"`
	Limit      int               `parser:"| ( 'limit' @Int )?"`
	Select     []*Select         `parser:"  ( @@ )+ )"`
}

/*
func (a *Query) String() string {
	qp := queryPrinter{}
	_ = NewBuilder().
		Query(qp.query).
		QueryRange(qp.queryRange).
		UsingDefinitions(qp.usingDefinitions).
		Time(qp.time).
		Duration(qp.duration).
		Histogram(qp.histogram).
		Select(qp.selectStatement).
		WindRose(qp.windRose).
		SelectExpression(qp.selectExpression).
		AliasedExpression(qp.aliasedExpression).
		Expression(qp.expression).
		ExpressionModifier(qp.expressionModifier).
		Metric(qp.metric).
		Function(qp.function).
		Build().
		Query(a)
	return strings.Join(qp.s, "\n")
}

type queryPrinter struct {
	queryPrinterChild
}

type queryPrinterChild struct {
	s     []string
	child *queryPrinterChild
}

func (qp *queryPrinter) save() {
	child := qp.queryPrinterChild
	qp.queryPrinterChild = queryPrinterChild{
		s:     nil,
		child: &child,
	}
}

func (qp *queryPrinter) restore(seps ...string) {
	sep := " "
	if len(seps) > 0 {
		sep = seps[0]
	}
	s := strings.Join(qp.s, sep)
	qp.queryPrinterChild = *qp.child
	qp.append(s)
}

func (qp *queryPrinter) append(s ...string) {
	qp.s = append(qp.s, strings.Join(s, " "))
}

func (qp *queryPrinter) appendString(s string) {
	qp.s = append(qp.s, "\""+s+"\"")
}

func (qp *queryPrinter) pop(f func()) string {
	f()
	l := len(qp.s)
	s := qp.s[l-1]
	qp.s = qp.s[:l-1]
	return s
}

func (qp *queryPrinter) popTime(v Visitor, t *Time) string {
	return qp.pop(func() {
		_ = v.Time(t)
	})
}

func (qp *queryPrinter) popDuration(v Visitor, d *Duration) string {
	return qp.pop(func() {
		_ = v.Duration(d)
	})
}

func (qp *queryPrinter) query(_ Visitor, b *Query) error {
	if b.Limit > 0 {
		qp.append("LIMIT", strconv.Itoa(b.Limit))
	}
	return nil
}

func (qp *queryPrinter) queryRange(v Visitor, b *QueryRange) error {
	qp.save()
	switch {
	case b.At != nil:
		qp.append("AT")
		_ = v.Time(b.At)

	case b.From != nil && b.For != nil:
		qp.append("FROM")
		_ = v.Time(b.From)
		qp.append("FOR")
		_ = v.Duration(b.For)

	case b.Start != nil && b.End != nil:
		qp.append("BETWEEN")
		_ = v.Time(b.Start)
		qp.append("AND")
		_ = v.Time(b.End)
	}
	qp.restore()

	if b.Every != nil {
		qp.save()
		qp.append("EVERY")
		_ = v.Duration(b.Every)
		qp.restore()
	}

	return VisitorStop
}

func (qp *queryPrinter) usingDefinitions(v Visitor, b *UsingDefinitions) error {
	qp.append("DECLARE")
	qp.save()
	for _, e := range b.Defs {
		qp.save()
		qp.append("      ")
		qp.appendString(e.Name)
		qp.append("AS")
		for _, m := range e.Modifier {
			_ = v.ExpressionModifier(m)
		}
		qp.restore(" ")
	}
	qp.restore(",\n")
	return VisitorStop
}

func (qp *queryPrinter) time(v Visitor, b *Time) error {
	qp.save()

	qp.appendString(b.Def)

	for _, e := range b.Expression {
		switch {
		case e.Add != nil:
			qp.append("ADD")
			_ = v.Duration(e.Add)

		case e.Truncate != nil:
			qp.append("TRUNCATE")
			_ = v.Duration(e.Truncate)
		}
	}

	qp.restore()
	return VisitorStop
}

func (qp *queryPrinter) duration(_ Visitor, b *Duration) error {
	qp.appendString(b.Def)
	return nil
}

func (qp *queryPrinter) histogram(v Visitor, b *Histogram) error {
	qp.append("HISTOGRAM")

	if b.Expression != nil {
		if err := v.AliasedExpression(b.Expression); err != nil {
			return err
		}
	}

	return VisitorStop
}

func (qp *queryPrinter) selectStatement(v Visitor, b *Select) error {
	qp.append("SELECT")

	if b.Expression != nil {
		if err := v.SelectExpression(b.Expression); err != nil {
			return err
		}
	}

	if b.Limit > 0 {
		qp.append("    LIMIT", strconv.Itoa(b.Limit))
	}

	return VisitorStop
}

func (qp *queryPrinter) windRose(v Visitor, b *WindRose) error {
	qp.append("WINDROSE")

	qp.save()

	qp.save()
	qp.append("       ")
	err := v.Expression(b.Degrees)
	qp.restore("")

	if err == nil {
		err = v.Expression(b.Speed)
	}

	if err != nil {
		return err
	}

	qp.restore(",\n       ")

	qp.append("    AS")

	qp.save()
	for _, e := range b.Options {
		qp.save()
		qp.append("       ")
		switch {
		case e.Rose:
			qp.append("ROSE")

		case e.Count:
			qp.append("COUNT")

		case e.Max:
			qp.append("MAX")
		}
		qp.restore("")
	}
	qp.restore(",\n")

	return VisitorStop
}

func (qp *queryPrinter) selectExpression(v Visitor, b *SelectExpression) error {
	qp.save()
	for _, e := range b.Expressions {
		qp.save()
		qp.append("       ")
		_ = v.AliasedExpression(e)
		qp.restore("")
	}
	qp.restore(",\n")
	return VisitorStop
}

func (qp *queryPrinter) aliasedExpression(v Visitor, b *AliasedExpression) error {
	qp.save()

	_ = v.Expression(b.Expression)

	if b.Unit != "" {
		qp.append("UNIT")
		qp.appendString(b.Unit)
	}

	if b.As != "" {
		qp.append("AS")
		qp.appendString(b.As)
	}

	qp.restore()
	return VisitorStop
}

func (qp *queryPrinter) expression(v Visitor, b *Expression) error {
	qp.save()

	switch {
	case b.Metric != nil:
		_ = v.Metric(b.Metric)

	case b.Function != nil:
		_ = v.Function(b.Function)
	}

	if b.Using != "" {
		qp.append("USING")
		qp.appendString(b.Using)
	} else {
		for _, e := range b.Modifier {
			_ = v.ExpressionModifier(e)
		}
	}

	qp.restore()
	return VisitorStop
}

func (qp *queryPrinter) expressionModifier(v Visitor, b *ExpressionModifier) error {
	_ = v.QueryRange(b.Range)

	if b.Offset != nil {
		qp.append("OFFSET")
		_ = v.Duration(b.Offset)
	}

	return VisitorStop
}

func (qp *queryPrinter) metric(v Visitor, b *Metric) error {
	qp.append(b.Name)
	return nil
}

func (qp *queryPrinter) function(v Visitor, b *Function) error {
	qp.save()
	qp.append(b.Name)
	qp.append("(")

	qp.save()
	for _, e := range b.Expressions {
		_ = v.Expression(e)
	}
	qp.restore(", ")

	qp.append(")")
	qp.restore("")
	return VisitorStop
}
*/
