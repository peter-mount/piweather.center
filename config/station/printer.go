package station

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/peter-mount/go-script/errors"
	"github.com/peter-mount/piweather.center/config/util/command"
	"github.com/peter-mount/piweather.center/config/util/location"
	"github.com/peter-mount/piweather.center/config/util/time"
	"github.com/peter-mount/piweather.center/config/util/units"
	"strings"
)

var (
	printVisitor = NewBuilder[*printState]().
		Command(printCommand).
		Container(printContainer).
		CronTab(printCron).
		Dashboard(printDashboard).
		Gauge(printGauge).
		Location(printLocation).
		Metric(printMetric).
		Station(printStation).
		Task(printTask).
		TaskCondition(printTaskCondition).
		Tasks(printTasks).
		Text(printText).
		TimeZone(printTimeZone).
		Unit(printUnit).
		Value(printValue).
		Build()
)

func newPrintState() *printState {
	root := &printNode{}
	return &printState{
		indentSize: 4,
		root:       root,
		current:    root,
	}
}

type printState struct {
	indentSize int        // Side of indent
	root       *printNode // Root of tree
	current    *printNode // Current node
}

type printNode struct {
	parent   *printNode
	children []*printNode
	indent   int
	header   []string
	body     []string
	footer   []string
}

func (s *printState) String() string {
	l := s.root.append([]string{})

	return strings.Join(l, "\n")
}

func (n *printNode) append(l []string) []string {
	indent := strings.Repeat(" ", n.indent)

	l = n.appendImpl(indent, l, n.body)

	for _, child := range n.children {
		l = n.appendImpl(indent, l, child.header)
		l = child.append(l)
		l = n.appendImpl(indent, l, child.footer)
	}

	return l
}

func (n *printNode) appendImpl(indent string, l []string, content []string) []string {
	for _, e := range content {
		l = append(l, indent+e)
	}
	return l
}

func (s *printState) Start() *printState {
	n := &printNode{
		parent: s.current,
		indent: s.current.indent + s.indentSize,
	}

	s.current.children = append(s.current.children, n)
	s.current = n
	return s
}

func (s *printState) End() *printState {
	if s.current.parent != nil {
		s.current = s.current.parent
	}
	return s
}

func (s *printState) EndError(p lexer.Position, err error) error {
	s.End()

	if err == nil {
		return errors.VisitorStop
	}

	return errors.Error(p, err)
}

func appendf(l []string, f string, a ...any) []string {
	if len(a) == 0 {
		return append(l, f)
	}
	return append(l, fmt.Sprintf(f, a...))
}

func (s *printState) AppendPos(p lexer.Position) *printState {
	return s.AppendHead("").AppendHead("// %s", p.String())
}

func (s *printState) AppendComponent(d *Component) *printState {
	if d != nil {
		if d.Title != "" {
			s.AppendBody("title %q", d.Title)
		}

		if d.Class != "" {
			s.AppendBody("class %q", d.Class)
		}

		if d.Style != "" {
			s.AppendBody("style %q", d.Style)
		}
	}
	return s
}

func (s *printState) AppendHead(f string, a ...any) *printState {
	s.current.header = appendf(s.current.header, f, a...)
	return s
}

func (s *printState) AppendBody(f string, a ...any) *printState {
	s.current.body = appendf(s.current.body, f, a...)
	return s
}

func (s *printState) AppendFooter(f string, a ...any) *printState {
	s.current.footer = appendf(s.current.footer, f, a...)
	return s
}

func (s *printState) Append(f string, a ...any) *printState {
	if len(s.current.header) == 0 {
		return s.AppendHead(f, a...)
	}

	st := f
	if len(a) > 0 {
		st = fmt.Sprintf(f, a...)
	}

	l := len(s.current.header) - 1
	e := s.current.header[l]
	if e != "" {
		e = e + " "
	}
	s.current.header[l] = e + st
	return s
}

func printCron(v Visitor[*printState], d time.CronTab) error {
	return v.Get().
		Start().
		AppendHead("%q", d.Definition()).
		EndError(d.Position(), nil)
}

func printLocation(v Visitor[*printState], d *location.Location) error {
	st := v.Get()

	if d.Notes == "" {
		st.Start().
			AppendHead("location( %q %q %q %.0f )", d.Name, d.Latitude, d.Longitude, d.Altitude).
			End()

	} else {
		st.Start().
			AppendHead("location( %q %q %q %.0f", d.Name, d.Latitude, d.Longitude, d.Altitude).
			AppendBody("note %s ", d.Notes).
			AppendFooter(")").
			End()
	}

	return errors.VisitorStop
}

func printTimeZone(v Visitor[*printState], d *time.TimeZone) error {
	st := v.Get()

	st.Start().
		AppendHead("timezone( %q )", d.TimeZone).
		End()

	return errors.VisitorStop
}

func printUnit(v Visitor[*printState], d *units.Unit) error {
	st := v.Get()

	st.Start().
		AppendHead("unit %q", d.Using).
		End()

	return errors.VisitorStop
}

func printCommand(v Visitor[*printState], d command.Command) error {
	// create command line with any " escaped
	var args []string
	args = append(args, fmt.Sprintf("%q", d.Command()))
	for _, arg := range d.Args() {
		args = append(args, fmt.Sprintf("%q", arg))
	}

	v.Get().
		Start().
		AppendHead(strings.Join(args, " ")).
		End()

	return errors.VisitorStop
}
