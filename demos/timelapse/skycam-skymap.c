// Script used to generate a sky map at the camera's location.
//
// This really works for cameras using fish eye lenses so they see most of the sky.
//
// Normal cameras have too small a field of view for this to be useful other than
// showing what's up rather than whats in the main camera view.

import (
    "github.com/peter-mount/go-anim/script/graph"
    "github.com/peter-mount/go-anim/script/image"
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

createSkyMap(cfg, bounds) {
    bounds = util.Rect(0,0, cfg.mapW, cfg.mapH).Rect()
    cfg.mapBounds = bounds
    //chart := chart.NewHorizon( cfg.location, bounds )
    chart := chart.NewAngular( cfg.location, bounds, math.Min(bounds.Dx(), bounds.Dy())/2, 80 )

    chart.Background().SetFillStroke(cfg.mapBackground)

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

    cfg.ephemLayer = chart.EphemerisDayLayer()

    // Add to the config our chart and it's own context
    cfg.map = chart
}

// render the sky map
//
// cfg      Configuration
// jd       Julian Day Number of the moment to display
// srcImg   Source image from the camera
// ephem    Ephemeris for solar system objects
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
renderSkyMap( cfg, jd, srcImg, ephem) {

    if !mapContains(cfg, "map" ) {
        createSkyMap(cfg, srcImg.Bounds())
    }

    cfg.ephemLayer.SetEphemeris(ephem)

    bounds := srcImg.Bounds()
    mapCtx := graph.NewImageContext(image.Duplicate(srcImg))

    try( mapCtx ) {
        gc := mapCtx.Gc()
        cfg.map.JD(jd)

        gc.Translate( cfg.mapX, cfg.mapY )
        cfg.map.Draw(gc)
    }

    // Now draw this image over the source using skymask to show just
    // the stars actually visible in the camera
    mapImg := image.DrawMask( mapCtx.Image(), cfg.skymask, srcImg )

    auxView := cfg.layout.Get("auxView")
    auxView.SetImage( mapImg )
}
