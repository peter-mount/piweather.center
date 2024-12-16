package catalogue

import (
	"github.com/llgcode/draw2d"
	"github.com/peter-mount/piweather.center/astro/chart"
	"sort"
)

type CatalogLayer interface {
	chart.ConfigurableLayer
	chart.ProjectionLayer

	// BrightestFirst orders stars brightest first.
	// For certain plots this allows for faint stars to be plotted over brighter ones.
	// e.g. Printed Star Atlases use this with star size based on Magnitude
	BrightestFirst()

	// FaintestFirst orders faint stars first.
	// Used for plots where stars are point objects and the Magnitude is shown by the intensity
	// of each pixel.
	//
	// This allows for bright stars not to be obscured by fainter ones.
	FaintestFirst()
}

type catalogLayer struct {
	chart.BaseLayer
	chart.BaseProjectionLayer
	stars    []Star
	renderer StarRenderer
}

// Star within a CatalogLayer which has been projected onto a chart.
type Star struct {
	X, Y, Mag float64
}

// StarRenderer renders a star
type StarRenderer func(draw2d.GraphicContext, Star)

// NewCatalogLayer creates a new CatalogLayer based on a Catalog.
//
// The entries of the Catalog will be filtered so the layer only contains entries that would be plotted.
//
// The default order of the entries are that in the Catalog. Use BrightestFirst or FaintestFirst to order them
// by magnitude.
func NewCatalogLayer(catalog *Catalog, renderer StarRenderer, proj chart.Projection) CatalogLayer {
	l := &catalogLayer{renderer: renderer}
	l.BaseLayer.Drawable = l.draw
	l.BaseProjectionLayer.SetProjection(proj)

	_ = catalog.ForEach(l.add)

	return l
}

func (l *catalogLayer) BrightestFirst() {
	sort.SliceStable(l.stars, func(i, j int) bool {
		return l.stars[i].Mag < l.stars[j].Mag
	})
}

func (l *catalogLayer) FaintestFirst() {
	sort.SliceStable(l.stars, func(i, j int) bool {
		return l.stars[i].Mag > l.stars[j].Mag
	})
}

func (l *catalogLayer) add(e Entry) error {
	x, y := l.Project(e.RA().Angle(), e.Dec())
	if l.Contains(x, y) {
		l.stars = append(l.stars, Star{X: x, Y: y, Mag: e.Mag()})
	}
	return nil
}

func (l *catalogLayer) draw(gc draw2d.GraphicContext) {
	for _, s := range l.stars {
		l.renderer(gc, s)
	}
}
