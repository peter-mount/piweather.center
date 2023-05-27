package station

import "context"

// Graph represents a graph of either an individual Reading
// or a custom one combining multiple Reading's.
type Graph struct {
	// Title on top of graph (optional)
	Title           string `json:"title,omitempty" xml:"title,attr,omitempty" yaml:"title,omitempty"`
	Line            *Line  `json:"line,omitempty" xml:"line,attr,omitempty" yaml:"line,omitempty"`
	Path            string `json:"-" xml:"-" yaml:"-"`
	reading         *Reading
	calculatedValue *CalculatedValue
}

func GraphFromContext(ctx context.Context) *Graph {
	return ctx.Value("Graph").(*Graph)
}

func (g *Graph) WithContext(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, "Graph", g), nil
}

func (g *Graph) Accept(v Visitor) error {
	return v.VisitGraph(g)
}

func (g *Graph) GetMinMax() (*float64, *float64) {
	if g.Line != nil {
		return g.Line.Min, g.Line.Max
	}
	return nil, nil
}

func (g *Graph) GetReading() *Reading                 { return g.reading }
func (g *Graph) GetCalculatedValue() *CalculatedValue { return g.calculatedValue }
func (g *Graph) Sensor() Sensor {
	if g.reading != nil {
		return g.reading
	}
	return g.calculatedValue
}
