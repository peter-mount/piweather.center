package chart

import (
	"github.com/llgcode/draw2d"
	"github.com/soniakeys/unit"
	"image"
	"image/color"
)

// Layer represents a drawable layer
type Layer interface {
	Draw(draw2d.GraphicContext)
}

// ConfigurableLayer is a Layer which can be configured with properties required for drawing.
// Specifically the Fill, Stroke and LineWidth parameters
type ConfigurableLayer interface {
	Layer
	SetFill(fill color.Color) ConfigurableLayer
	SetLineWidth(lineWidth float64) ConfigurableLayer
	SetStroke(stroke color.Color) ConfigurableLayer
}

// BaseLayer implements all of ConfigurableLayer except for Draw().
//
// Most ConfigurableLayer implementations use this as their base.
type BaseLayer struct {
	stroke    color.Color
	fill      color.Color
	lineWidth float64
	Drawable  Drawable
}

func (b *BaseLayer) SetFill(fill color.Color) ConfigurableLayer {
	b.fill = fill
	return b
}

func (b *BaseLayer) SetLineWidth(lineWidth float64) ConfigurableLayer {
	b.lineWidth = lineWidth
	return b
}

func (b *BaseLayer) SetStroke(stroke color.Color) ConfigurableLayer {
	b.stroke = stroke
	return b
}

func (b *BaseLayer) Draw(ctx draw2d.GraphicContext) {
	b.DrawLayer(ctx, b.Drawable)
}

// DrawLayer should be called by the parent Draw() function. It will call the passed Drawable
// with any configuration applied to the GraphicContext.
//
// The GraphicContext will be restored to its previous state when drawing is completed.
func (b *BaseLayer) DrawLayer(gc draw2d.GraphicContext, drawable Drawable) {
	if drawable == nil {
		panic("drawable must not be nil")
	}

	gc.Save()
	defer gc.Restore()

	if b.stroke != nil {
		gc.SetStrokeColor(b.stroke)
	}

	if b.fill != nil {
		gc.SetFillColor(b.fill)
	}

	if b.lineWidth > 0 {
		gc.SetLineWidth(b.lineWidth)
	}

	drawable(gc)
}

// Drawable is a function that can draw to a context.
type Drawable func(draw2d.GraphicContext)

// NewDrawableLayer returns a ConfigurableLayer which will use the provided Drawable for it's content.
func NewDrawableLayer(drawable Drawable) ConfigurableLayer {
	return &BaseLayer{Drawable: drawable}
}

type ProjectionLayer interface {
	// Projection being used
	Projection() Projection

	// GetCenter returns the center coordinates
	GetCenter() (float64, float64)

	// Contains returns true if the point is within the bounds of the chart (including the buffer area)
	Contains(x, y float64) bool

	// Correct moves (x, y) to be from the center of the plot, and adjusts the axes to account for
	// the original x positive to the left and y upwards to Image x positive to the right and y downwards
	//
	// Note: This is for use when using Projection directly. If the Project() function is used then
	// the result has already been corrected so no need to use this function.
	Correct(x, y float64) (float64, float64)

	// Project the coordinates (x,y) to image coordinates.
	//
	// Note: These coordinates have been corrected already so no need to use Correct()
	Project(x, y unit.Angle) (float64, float64)
}

type BaseProjectionLayer struct {
	projection Projection
	bounds     image.Rectangle
	cx, cy     float64
}

func (p *BaseProjectionLayer) SetProjection(proj Projection) {
	p.projection = proj
	// cx & cy based on original bounds
	bounds := proj.Bounds()
	p.cx = float64(bounds.Dx()) / 2.0
	p.cy = float64(bounds.Dy()) / 2.0
	// bounds expanded by 50 pixels, so we still plot lines crossing the boundary
	p.bounds = bounds.Inset(-50)
}

func (p *BaseProjectionLayer) Projection() Projection {
	return p.projection
}

func (p *BaseProjectionLayer) GetCenter() (float64, float64) {
	return p.cx, p.cy
}

func (p *BaseProjectionLayer) Contains(x, y float64) bool {
	return image.Pt(int(x), int(y)).In(p.bounds)
}

func (p *BaseProjectionLayer) Correct(x, y float64) (float64, float64) {
	return p.cx - x, p.cy - y
}

func (p *BaseProjectionLayer) Project(x, y unit.Angle) (float64, float64) {
	px, py := p.projection.Project(x, y)
	return p.Correct(px, py)
}
