// Common functionality for the skycam scripts

import (
    "github.com/peter-mount/go-anim/script/colour"
    "github.com/peter-mount/go-anim/script/graph"
    "github.com/peter-mount/go-anim/script/image"
    "github.com/peter-mount/go-anim/script/render"
    "github.com/peter-mount/go-anim/script/util"
    "github.com/peter-mount/piweather.center/astro/calculator"
    "github.com/peter-mount/piweather.center/script/astro/calendar"
    "github.com/peter-mount/piweather.center/script/astro/chart"
    "github.com/peter-mount/piweather.center/script/astro/geo"
    "github.com/peter-mount/piweather.center/script/weather/cloud"
    "github.com/peter-mount/piweather.center/script/weather/value"
)

// read an image, but on failure will write which file failed to the console
readImage(srcName) {
    try {
        return image.ReadImage(srcName)
    } catch( e ) {
        fmt.Printf("\nFailed to read %q: %v\n", srcName, e)
        throw(e)
    }
}

renderFrame(ctx,cfg,frame) {
    // Get time from the file name
    srcTime := frame.Time
    jd := calendar.FromTime(srcTime)
    tm := value.BasicTime(srcTime,cfg.location.Coord(),0)

    sun := calculator.CalculateSun(tm)

    // Get the sky camera image, caching it as necessary
    if frame.RequiresImage() {
        img:=readImage(frame.Source)
        // Apply privacy mask where black pixels are removed
        img=image.Mask(img, cfg.privmask)
        // Now auto-crop it, removing any black borders introduced by privmask
        img=image.AutoCrop(img)
        // Cache the image in the frame
        frame.SetImage(img)
    }
    skyImage := frame.Image()

    try( ctx ) {
        gc=ctx.Gc()
        image.Fill(ctx,cfg.background)

        cfg.layout.Get("timeDisplay").Args(srcTime.Format(time.RFC1123))
        cfg.layout.Get("skyCamera").SetImage(skyImage)

        if sun.GetHorizontal().Alt >= 0 {
            renderClouds(ctx,cfg,skyImage)
        } else {
//            renderSkyMap(ctx,cfg,jd)
        }

        cfg.layout.Layout(gc)
        cfg.layout.Draw(gc)
    }
}

renderClouds(ctx, cfg, srcImg) {
    // Run the cloud filter generating the statistics and an image
    filter := cloud.Filter( cfg.privmask, cfg.skymask ).Limit( 0.7 )
    img := image.FilterNew(filter.Filter(),srcImg)
    coverage := filter.Coverage()

    // Render the results
    //
    // NB: If the totals rendered are >100% this is down to %f in Sprintf always
    // rounding up if the fraction is >=0.5 so although it appears wrong internally it's
    // correct.
    cfg.layout.Get("auxView").SetImage(img)
    cfg.layout.Get("cloudCover").Args(coverage.Cloud, coverage.Sky, coverage.Obscured)
}

createSkyMap(cfg) {
    width := cfg.auxViewW
    chart := chart.NewHorizon( cfg.location, util.Rect(0,0,width,width).Rect() )
    chart.Background().SetFillStroke(cfg.black)
    chart.Horizon().SetFill(cfg.horizonColour).SetStroke(cfg.horizonBorder)

    chart.Feature("milkyway").SetFillStroke(cfg.milkyWay)
//  chart.Feature("const.border").SetStroke(cfg.constBorder)
    chart.Feature("const.line").SetStroke(cfg.constLine)

    chart.YaleBrightStarCatalog().MagLimit(cfg.magLimit)

    // Add to the config our chart and it's own context
    cfg.map = chart
    cfg.mapCtx = graph.NewSizedContext(width,width)
}

renderSkyMap(ctx,cfg,jd) {

    // This will speed things up by reducing lookup times
    mapCtx := cfg.mapCtx

    try( mapCtx ) {
        gc := mapCtx.Gc()
        cfg.map.JD(jd)
        cfg.map.Draw(gc)
    }

    cfg.layout.Get("auxView").SetImage( mapCtx.Image() )
}
