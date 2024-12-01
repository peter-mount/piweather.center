package html

import (
	"fmt"
	strings2 "github.com/peter-mount/go-kernel/v2/util/strings"
	"strings"
)

// Element represents an HTML Element.
type Element struct {
	name     string
	class    strings2.StringSlice
	attrs    []Attr
	parent   *Element
	children []*Element
	text     string
	points   []Point
}

type Attr struct {
	Name  string
	Value string
}

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func printfOptimized(format string, a []interface{}) string {
	s := format
	if len(a) > 0 {
		s = fmt.Sprintf(s, a...)
	}
	return s
}

// Class adds a css class name to the Element.
// This takes a format string like fmt.Sprintf()
func (e *Element) Class(c string, a ...interface{}) *Element {
	s := printfOptimized(c, a)
	if s != "" {
		e.class = append(e.class, s)
	}
	return e
}

func (e *Element) Attr(name, format string, args ...interface{}) *Element {
	s := printfOptimized(format, args)
	if s != "" {
		e.attrs = append(e.attrs, Attr{
			Name:  name,
			Value: s,
		})
	}
	return e
}

func (e *Element) AttrInt(name string, val int) *Element {
	return e.Attr(name, "%d", val)
}

// String converts the Element & any children to a string
func (e *Element) String() string {
	var a []string

	if e.text != "" {
		a = append(a, e.text)
	} else {
		if e.name != "" {
			a = append(a, "<", e.name)

			if !e.class.IsEmpty() {
				a = append(a, e.class.Join2(" class=\"", "\"", " "))
			}

			if len(e.attrs) > 0 {
				for _, attr := range e.attrs {
					a = append(a, fmt.Sprintf(" %s=%q", attr.Name, attr.Value))
				}
			}

			if len(e.points) > 0 {
				var r []string
				for _, p := range e.points {
					r = append(r, p.String())
				}
				a = append(a, " points=\"", strings.Join(r, " "), "\"")
			}

			a = append(a, ">")
		}

		for _, child := range e.children {
			a = append(a, child.String())
		}

		if e.name != "" {
			a = append(a, "</", e.name, ">")
		}
	}

	return strings.Join(a, "")
}

// Builder creates a new Element builder
func Builder() *Element {
	return &Element{}
}

// End terminates the current Element, returning its parent.
// For the root it return's the current root Element.
func (e *Element) End() *Element {
	if e != nil && e.parent != nil {
		return e.parent
	}
	return e
}

// Exec invokes a function returning the return value from it.
// It's used to embed common html sequences into a builder.
func (e *Element) Exec(f func(*Element) *Element) *Element {
	return f(e)
}

// If calls a function if a condition is true.
// The returned instance will be the returned instance from that function or the current Element if the condition was false.
func (e *Element) If(p bool, f func(*Element) *Element) *Element {
	if p {
		return f(e)
	}
	return e
}

// RootElement returns the root Element of the current document.
func (e *Element) RootElement() *Element {
	n := e
	for n.parent != nil {
		n = n.parent
	}
	return n
}

// Textf appends a Text Element based on a fmt.Sprintf() formatted string.
func (e *Element) Textf(f string, a ...interface{}) *Element {
	return e.Text(printfOptimized(f, a))
}

// Text appends each provided string as a text Element.
func (e *Element) Text(s ...string) *Element {
	ne := e
	for _, a := range s {
		ne := ne.Element("")
		ne.text = a
	}
	return ne
}

// TextNbsp appends each provided string as a text Element.
func (e *Element) TextNbsp(s ...string) *Element {
	for i, a := range s {
		if a == "" {
			s[i] = "&nbsp;"
		}
	}
	return e.Text(s...)
}

type SequenceHandler func(int, *Element) *Element

// Sequence calls a function for each integer between start and end inclusively.
// If start > end then the sequence will count down from start until end.
func (e *Element) Sequence(start, end int, f SequenceHandler) *Element {
	ne := e
	if start < end {
		for i := start; i <= end; i++ {
			ne = f(i, ne)
		}
	} else {
		for i := start; i >= end; i-- {
			ne = f(i, ne)
		}
	}
	return ne
}

func (e *Element) Element(n string) *Element {
	child := &Element{
		name:   n,
		parent: e,
	}
	e.children = append(e.children, child)
	return child
}

func (e *Element) Html() *Element {
	return e.Element("html")
}

func (e *Element) Head() *Element {
	return e.Element("head")
}

func (e *Element) Title() *Element {
	return e.Element("title")
}

func (e *Element) Body() *Element {
	return e.Element("body")
}

func (e *Element) H1() *Element {
	return e.Element("h1")
}

func (e *Element) A() *Element {
	return e.Element("a")
}

// Div Element. This is open so End() must be called to terminate it.
func (e *Element) Div() *Element {
	return e.Element("div")
}

// Span Element. This is open so End() must be called to terminate it.
func (e *Element) Span() *Element {
	return e.Element("span")
}

// OL Element. This is open so End() must be called to terminate it.
func (e *Element) OL() *Element {
	return e.Element("ol")
}

// LI Element. This is open so End() must be called to terminate it.
func (e *Element) LI() *Element {
	return e.Element("li")
}

// Sub Element. This is open so End() must be called to terminate it.
func (e *Element) Sub() *Element {
	return e.Element("sub")
}

// Sup Element. This is open so End() must be called to terminate it.
func (e *Element) Sup() *Element {
	return e.Element("sup")
}

func (e *Element) Style() *Element {
	return e.Element("style")
}

func (e *Element) Table() *Element {
	return e.Element("table")
}

func (e *Element) THead() *Element {
	return e.Element("thead")
}

func (e *Element) TBody() *Element {
	return e.Element("tbody")
}

func (e *Element) TR() *Element {
	return e.Element("tr")
}

func (e *Element) TD() *Element {
	return e.Element("td")
}

func (e *Element) TH() *Element {
	return e.Element("th")
}
