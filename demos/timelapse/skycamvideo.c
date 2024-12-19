// Script which takes a set of images taken over time and generate a 4K video of those images.
//
// Each image is placed at the top left of the frame.
// A keogram is below this image showing the sky conditions over time.
//
// On the right is a cloud cover view for data during daytime, and a map of the sky during nighttime.
//

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

main() {

    // The width of each column on the top part of the frame
    // e.g. this contains the camera image, cloud cover and sky map
    topColWidth := image.Width4K/3
    // This is the width of each entry, reduced so there's a gap between them
    topColCellWidth := topColWidth-20

    cfg := map(
        // Set this to the directory containing the images
        "srcDir": "/home/peter/weather/cam2",

        // The output video name
        "output": "/home/peter/test-video.mp4",

        // Location of London, UK
        "location": geo.LatLong(51.5, -8/60.0, 0),
        "timeZone": time.LoadLocation("Europe/London"),

        "title": "Example sky camera timelapse",


        "black": colour.Colour("black"),
        "white": colour.Colour("white"),

        // Overall background colour of the video
        "background": colour.Colour("black"),

        // Width and position of the sky camera view - the left 40% of the frame
        "skyWidth": topColCellWidth,
        "skyX": 10,
        "skyY": 60,
        // usable image is 2656x2154 but as we should keep it square then limit it to
        // part of the frame with the most sky visible
        "skyBounds": util.Rect(2656-2154,0,2656,2154).Rect(),

        // cloud config
        "cloudX": (image.Width4K-topColCellWidth)/2,
        //"cloudX": image.Width4K-topColCellWidth,
        "cloudY": 60,
        "cloudWidth": topColCellWidth,

        // Position of the cloud coverage or skymap view - the right 30% of the frame
        "auxViewX": (image.Width4K-topColCellWidth)/2,
        //"auxViewX": image.Width4K-topColCellWidth,
        "auxViewY": 60,
        "auxViewW": topColCellWidth,

        // skyMap config
        "horizonColour": colour.Colour("black"), // "#00320033"
        "horizonBorder": colour.Colour("white"),
        "milkyWay": colour.Grey(17),
        "constBorder": nil, // colour.Colour("#0000aa",
        "constLine": colour.Colour("#0000aa"),
        "magLimit": 99
    )

    createSkyMap(cfg)

    //ctx := graph.New4k()
    //ctx := graph.New1080p().Scale(0.5,0.5)
    ctx := graph.New720p().Scale(1/3.0,1/3.0)

    files := util.GetImageFiles(cfg.srcDir)
    frames := util.SequenceIn(15,files,cfg.timeZone)
    frameCount := frames.Size()
    fmt.Printf("Rendering %d frames\n", frameCount)

    try( out := render.New( cfg.output, 25 ) ) {
        for i,frame := range frames {
//            fmt.Printf( "\rRendering %d/%d %.0f%% %s ",
//                i, frameCount,
//                (100.0*i)/frameCount,
//                frame.Time.Format("2006-01-02 15:04:05")
//            )
            renderFrame(ctx,cfg,frame)
            out.WriteImage(ctx.Image())

//            if i>100 {
//                break
//            }
        }
    }

    fmt.Println("\nRender complete")
}

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
        // usable image is 2656x2154 but as we should keep it square then limit
        //img = image.Crop(img, cfg.skyBounds)
        img = image.Resize(math.Min(cfg.skyWidth,cfg.skyBounds.Dx()),0,img,"nearestNeighbour")
        frame.SetImage(img)
    }
    skyImage := frame.Image()

    try( ctx ) {
        gc=ctx.Gc()

        image.Fill(ctx,cfg.background)

        // The title
        try( ctx ) {
            gc.SetFillColor( cfg.white )
            graph.SetFont( gc, "luxi 32 mono bold" )
            util.DrawStringLeft(gc, 10, 23,cfg.title )

            util.DrawStringRight( gc,
                image.Width4K - 10, 23,
                "%s",
                srcTime.Format(time.RFC1123))
        }

        try( ctx ) {
            gc.Translate(cfg.skyX,cfg.skyY)
            gc.DrawImage(skyImage)
        }

        if sun.GetHorizontal().Alt >= 0 {
            renderClouds(ctx,cfg,skyImage)
        } else {
            renderSkyMap(ctx,cfg,jd)
        }
    }
}

renderClouds(ctx, cfg, srcImg) {
    // Run the cloud filter generating the statistics and an image
    filter := cloud.FilterNoMask().Limit( 0.7 )
    img := image.FilterNew(filter.Filter(),srcImg)
    coverage := filter.Coverage()

    // Render the results
    //
    // NB: If the totals rendered are >100% this is down to %f in Sprintf always
    // rounding up if the fraction is >=0.5 so although it appears wrong internally it's
    // correct.
    try( ctx ) {
        gc := ctx.Gc()
        gc.Translate(cfg.cloudX, cfg.cloudY)
        gc.DrawImage( img )

        gc.SetFillColor( cfg.white )
        graph.SetFont( gc, "luxi 20 mono bold" )
        util.DrawStringLeft(gc,
            0, cfg.cloudWidth+20,
            "Cloud Cover %3.0f%% Sky %3.0f%% Obscured %3.0f%%",
            coverage.Cloud, coverage.Sky, coverage.Obscured
        )
    }
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

    try( ctx ) {
        gc := ctx.Gc()
        gc.Translate(cfg.auxViewX, cfg.auxViewY)
        gc.DrawImage( mapCtx.Image() )
    }
}
