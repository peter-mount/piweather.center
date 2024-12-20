// Script used to generate a sky map at the camera's location.
//
// This really works for cameras using fish eye lenses so they see most of the sky.
//
// Normal cameras have too small a field of view for this to be useful other than
// showing what's up rather than whats in the main camera view.

import (
    "github.com/peter-mount/go-anim/script/graph"
    "github.com/peter-mount/go-anim/script/util"
    "github.com/peter-mount/piweather.center/script/astro/chart"
)

// create the resources required by renderSkyMap
//
// cfg  The configuration
//
// requirements:
//
// cfg.location     The location of the camera on Earth
// cfg.black        background colour, usually color.Black
// cfg.horizonColour    Colour to plot the horizon (if used)
// cfg.horizonBorder    Colour for the horizon border (if used)
// cfg.milkyWay         Colour of the milky way (if used)
// cfg.constBorder      Colour of the constellation borders (if used)
// cfg.constLine        Colour of the constellation outlines (if used)
// cfg.magLimit         Magnitude limit for stars. 3, 4 or 5 for light pollution, 6 for visible, 99 for everything

createSkyMap(cfg) {
    width := cfg.auxViewW
    chart := chart.NewHorizon( cfg.location, util.Rect(0,0,width,width).Rect() )
    chart.Background().SetFillStroke(cfg.black)

    // Horizon with the border colour being optional
    if mapContains(cfg,"horizonColour", "horizonBorder") {
        hz := chart.Horizon().SetFill(cfg.horizonColour)
        if mapContains(cfg,"horizonBorder") hz.SetStroke(cfg.horizonBorder)
    }

    // Optional layers included if their colours are defined
    if mapContains(cfg,"milkyWay") chart.Feature("milkyway").SetFillStroke(cfg.milkyWay)

    if mapContains(cfg,"constBorder") chart.Feature("const.border").SetStroke(cfg.constBorder)

    if mapContains(cfg,"constLine") chart.Feature("const.line").SetStroke(cfg.constLine)

    // Star layer with magnitude limit applied if defined
    ybsc := chart.YaleBrightStarCatalog()
    if mapContains(cfg,"magLimit") ybsc.MagLimit(cfg.magLimit)

    // Add to the config our chart and it's own context
    cfg.map = chart
    cfg.mapCtx = graph.NewSizedContext(width,width)
}

// render the sky map
//
// cfg      Configuration
// jd       Julian Day Number of the moment to display
//
// requirements:
//
// cfg.map      The Chart containing the layers to be shown
// cfg.mapCtx   GraphicsContext specific to the map
//
// components in cfg.layout:
//
// "auxView"    Image component to display the map
//
renderSkyMap( cfg, jd ) {

    // This will speed things up by reducing lookup times
    mapCtx := cfg.mapCtx

    try( mapCtx ) {
        gc := mapCtx.Gc()
        cfg.map.JD(jd)
        cfg.map.Draw(gc)
    }

    cfg.layout.Get("auxView").SetImage( mapCtx.Image() )
}
