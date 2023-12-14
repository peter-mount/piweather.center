package lang

import (
	"github.com/alecthomas/participle/v2/lexer"
	"strings"
)

type Query struct {
	Pos lexer.Position

	QueryRange *QueryRange `parser:"@@"`
	Select     []*Select   `parser:"( @@ )+"`
}

func (a *Query) Accept(v Visitor) error {
	return v.Query(a)
}

func (a *Query) String() string {
	qp := queryPrinter{}
	_ = NewBuilder().
		QueryRange(qp.queryRange).
		Time(qp.time).
		Duration(qp.duration).
		Select(qp.selectStatement).
		SelectExpression(qp.selectExpression).
		AliasedExpression(qp.aliasedExpression).
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

func (qp *queryPrinter) time(v Visitor, b *Time) error {
	qp.save()

	qp.appendString(b.Def)

	if b.Add != nil {
		qp.append("ADD")
		_ = v.Duration(b.Add)
	}

	if b.Truncate != nil {
		qp.append("TRUNCATE")
		_ = v.Duration(b.Truncate)
	}

	qp.restore()
	return VisitorStop
}

func (qp *queryPrinter) duration(_ Visitor, b *Duration) error {
	qp.appendString(b.Def)
	return nil
}

func (qp *queryPrinter) selectStatement(v Visitor, b *Select) error {
	qp.append("SELECT")
	return nil
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

	if b.Offset != nil {
		qp.append("OFFSET")
		_ = v.Duration(b.Offset)
	}

	if b.As != "" {
		qp.append("AS")
		qp.appendString(b.As)
	}

	qp.restore()
	return VisitorStop
}

func (qp *queryPrinter) metric(v Visitor, b *Metric) error {
	qp.append(b.Name)
	return nil
}

func (qp *queryPrinter) function(v Visitor, b *Function) error {
	qp.save()
	switch {
	case b.TimeOf:
		qp.append("TIMEOF")
	default:
		qp.append(b.Name)
	}
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
