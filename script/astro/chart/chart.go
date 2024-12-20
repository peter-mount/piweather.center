package chart

import (
	"errors"
	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/astro/catalogue"
	"github.com/peter-mount/piweather.center/astro/chart"
	"github.com/peter-mount/piweather.center/astro/chart/render"
	"github.com/peter-mount/piweather.center/astro/coord"
	"github.com/peter-mount/piweather.center/astro/julian"
	"github.com/peter-mount/piweather.center/astro/sidereal"
	"github.com/soniakeys/unit"
	"image"
	"time"
)

func init() {
	packages.RegisterPackage(&Package{})
}

type Package struct {
}

type Chart struct {
	rootLayer   chart.Layers
	chartLayer  chart.Layers
	transformer coord.CoordinateTransformer
	projection0 chart.Projection
	projection  chart.Projection
	background  chart.ConfigurableLayer
	horizon     chart.ConfigurableLayer
	manager     *catalogue.Manager
}

func (_ Package) NewStereographic(cx, cy unit.Angle, R float64, bounds image.Rectangle) *Chart {
	projection := chart.NewStereographicProjection(cx, cy, R, bounds)

	c := &Chart{
		chartLayer: chart.NewLayers(),
		projection: projection,
		manager:    &catalogue.Manager{},
	}

	c.background = chart.FloodFillLayer(c.projection)

	c.rootLayer = chart.NewLayers().
		// We need to flip both axes when plotting equatorial coordinates
		SetProjection(c.projection).
		Flip(true, true).
		// Core layers
		Add(c.background).
		Add(c.chartLayer)

	return c
}

func (_ Package) NewHorizon(loc *coord.LatLong, bounds image.Rectangle) *Chart {
	jd := julian.FromTime(time.Now())

	c := &Chart{
		chartLayer:  chart.NewLayers(),
		transformer: coord.NewCoordinateTransformer(loc.Latitude, loc.Longitude).Sidereal(sidereal.FromJD(jd)),
		projection0: chart.NewStereographicProjection(0, unit.AngleFromDeg(90), float64(bounds.Dx())/4.0, bounds),
		manager:     &catalogue.Manager{},
	}

	c.projection = c.projection0.Transform(func(p chart.Point) chart.Point {
		A, h := c.transformer.EqToHz(p.X.RA(), p.Y)
		f := A.Deg() + 180.0
		for f < 0.0 {
			f = f + 360
		}
		for f >= 360 {
			f = f - 360
		}
		return chart.Point{X: unit.AngleFromDeg(f), Y: h}
	})

	c.background = chart.FloodFillLayer(c.projection0)
	c.horizon = chart.HorizonLayer(c.projection0)

	c.rootLayer = chart.NewLayers().
		// Use the root projection, but we need to flip the x-axis
		SetProjection(c.projection0).
		Flip(true, false).
		// Core layers
		Add(c.background).
		Add(c.chartLayer).
		Add(c.horizon)

	return c
}

func (c *Chart) Draw(ctx draw2d.GraphicContext) {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Printf("Panic %v\n", err1)
		}
	}()
	c.rootLayer.Draw(ctx)
}

func (c *Chart) JD(jd julian.Day) (*Chart, error) {
	if c.transformer == nil {
		return nil, errors.New("transformer not set")
	}
	c.transformer.Sidereal(sidereal.FromJD(jd))
	return c, nil
}

// Background returns the background layer, or an error if not available
func (c *Chart) Background() (chart.ConfigurableLayer, error) {
	if c.background == nil {
		return nil, errors.New("background not set")
	}
	return c.background, nil
}

// Horizon returns the horizon layer, or an error if not available
func (c *Chart) Horizon() (chart.ConfigurableLayer, error) {
	if c.horizon == nil {
		return nil, errors.New("horizon not set")
	}
	return c.horizon, nil
}

// Projection returns the projection in use
func (c *Chart) Projection() chart.Projection {
	return c.projection
}

// Layer returns the main layer for rendering
func (c *Chart) Layer() chart.Layers {
	return c.chartLayer
}

func (c *Chart) Feature(n string) (chart.ConfigurableLayer, error) {
	f, err := c.manager.Feature(n)
	if err != nil {
		return nil, err
	}
	l := f.GetLayerAll(c.Projection())
	c.Layer().Add(l)
	return l, nil
}

func (c *Chart) YaleBrightStarCatalog() (chart.ConfigurableLayer, error) {
	f, err := c.manager.YaleBrightStarCatalog()
	if err != nil {
		return nil, err
	}
	l := f.NewLayer(render.BrightnessPixelStarRenderer, c.Projection())
	c.Layer().Add(l)
	return l, nil
}
