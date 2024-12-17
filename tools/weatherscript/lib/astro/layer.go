package astro

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/astro/catalogue"
	"github.com/peter-mount/piweather.center/astro/chart"
	"github.com/peter-mount/piweather.center/astro/chart/render"
)

func init() {
	packages.Register("chartLayer", &Layer{
		manager: &catalogue.Manager{},
	})
}

type Layer struct {
	manager *catalogue.Manager
	ybsc    *catalogue.Catalog
}

func (_ Layer) NewLayers() chart.Layers {
	return chart.NewLayers()
}

func (l Layer) FeatureLayer(n string, proj chart.Projection) (chart.ConfigurableLayer, error) {
	f, err := l.manager.Feature(n)
	if err != nil {
		return nil, err
	}
	return f.GetLayerAll(proj), nil
}

func (_ Layer) FloodFillLayer(proj chart.Projection) chart.ConfigurableLayer {
	return chart.FloodFillLayer(proj)
}

func (_ Layer) HorizonLayer(proj chart.Projection) chart.ConfigurableLayer {
	return chart.HorizonLayer(proj)
}

func (l Layer) YaleBrightStarCatalogLayer(proj chart.Projection) (chart.ConfigurableLayer, error) {
	cat, err := l.manager.YaleBrightStarCatalog()
	if err != nil {
		return nil, err
	}
	// FIXME declare renderer
	return cat.NewLayer(render.PixelStarsRenderer, proj), nil
}
