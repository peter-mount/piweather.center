package html

func (e *Element) Svg() *Element {
	return e.Element("svg").
		Attr("xmlns", "http://www.w3.org/2000/svg").
		Attr("version", "1.1")
}

func (e *Element) ViewBox(x, y, width, height int) *Element {
	return e.Attr("viewBox", "%d %d %d %d", x, y, width, height)
}

func (e *Element) R(v int) *Element {
	return e.AttrInt("r", v)
}

func (e *Element) X(v int) *Element {
	return e.AttrInt("x", v)
}

func (e *Element) Y(v int) *Element {
	return e.AttrInt("y", v)
}

func (e *Element) X1(v int) *Element {
	return e.AttrInt("x1", v)
}

func (e *Element) Y1(v int) *Element {
	return e.AttrInt("y1", v)
}

func (e *Element) X2(v int) *Element {
	return e.AttrInt("x2", v)
}

func (e *Element) Y2(v int) *Element {
	return e.AttrInt("y2", v)
}

func (e *Element) ClipPath() *Element {
	return e.Element("clipPath")
}

func (e *Element) CX(v int) *Element {
	return e.AttrInt("cx", v)
}

func (e *Element) CY(v int) *Element {
	return e.AttrInt("cy", v)
}

func (e *Element) DX(v int) *Element {
	return e.AttrInt("dx", v)
}

func (e *Element) DY(v int) *Element {
	return e.AttrInt("dy", v)
}

func (e *Element) Width(v int) *Element {
	return e.AttrInt("width", v)
}

func (e *Element) Height(v int) *Element {
	return e.AttrInt("height", v)
}

func (e *Element) Fill(v string, a ...interface{}) *Element {
	return e.Attr("fill", v, a...)
}

func (e *Element) Stroke(v string, a ...interface{}) *Element {
	return e.Attr("stroke", v, a...)
}

func (e *Element) StrokeWidth(v string, a ...interface{}) *Element {
	return e.Attr("stroke-width", v, a...)
}

// SvgText is the text Element, cannot use Text() as that adds plain text
func (e *Element) SvgText() *Element {
	return e.Element("text")
}

func (e *Element) TSpan() *Element {
	return e.Element("tspan")
}

func (e *Element) Circle() *Element {
	return e.Element("circle")
}

func (e *Element) G() *Element {
	return e.Element("g")
}

func (e *Element) Polygon() *Element {
	return e.Element("polygon")
}

func (e *Element) Polyline() *Element {
	return e.Element("polyline")
}

func (e *Element) Point(x, y int) *Element {
	e.points = append(e.points, Point{X: x, Y: y})
	return e
}

func (e *Element) Rect() *Element {
	return e.Element("rect")
}

func (e *Element) Line() *Element {
	return e.Element("line")
}
