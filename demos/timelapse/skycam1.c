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
    "github.com/peter-mount/piweather.center/script/astro/calendar"
    "github.com/peter-mount/piweather.center/script/astro/chart"
    "github.com/peter-mount/piweather.center/script/astro/geo"
    "github.com/peter-mount/piweather.center/script/weather/cloud"
)

main() {

    cfg := map(
        // Set this to the directory containing the images
        "srcDir": "/home/peter/weather/cam",

        // The output video name
        "output": "/home/peter/test-video.png",

        // Location of London, UK
        "location": geo.LatLong(51.5, -8/60.0, 0),
        "timeZone": time.LoadLocation("Europe/London"),

        "title": "Example sky camera timelapse",


        "black": colour.Colour("black"),
        "white": colour.Colour("white"),

        // Overall background colour of the video
        "background": colour.Colour("black"),

        // Width and position of the sky camera view - the left 40% of the frame
        "skyWidth": image.Width4K*0.4,
        "skyX": 10,
        "skyY": 60,
        // usable image is 2656x2154 but as we should keep it square then limit it to
        // part of the frame with the most sky visible
        "skyBounds": util.Rect(2656-2154,0,2656,2154).Rect(),

        // Position of the cloud coverage or skymap view - the right 30% of the frame
        "auxViewX": image.Width4K-(image.Width4K*0.3),
        "auxViewY": 60,
        "auxViewW": (image.Width4K*0.3)-10,
        "auxViewH": (image.Width4K*0.3)-10,

        // cloud config
        "cloudX": image.Width4K-(image.Width4K*0.6)+20,
        "cloudY": 60,
        "cloudWidth": (image.Width4K*0.3)-20,
        "cloudHeight": (image.Width4K*0.3)-20,

        // skyMap config
        "horizonColour": colour.Colour("black"), // "#00320033"
        "horizonBorder": colour.Colour("white"),
        "milkyWay": nil, // colour.Grey(17),
        "constBorder": nil, // colour.Colour("#0000aa",
        "constLine": colour.Colour("#0000aa"),
        "magLimit": 99
    )

    createSkyMap(cfg)

    ctx := graph.New4k()

    renderFrame(ctx,cfg,"20230810162632.png")

    image.WriteImage(cfg.output, ctx.Image() )
}

renderFrame(ctx,cfg,srcName) {
    // Get time from the file name
    srcTime := util.TimeFromFileNameIn(srcName,cfg.timeZone)

    // Get the sky camera image
    skyImage := image.ReadImage(cfg.srcDir+"/"+ srcName)
    // usable image is 2656x2154 but as we should keep it square then limit
    skyImage = image.Crop(skyImage, cfg.skyBounds)
    skyImage = image.Resize(math.Min(cfg.skyWidth,cfg.skyBounds.Dx()),0,skyImage,"nearestNeighbour")

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

        renderClouds(ctx,cfg,skyImage)

        renderSkyMap(ctx,cfg,calendar.FromTime(srcTime))
    }
}

renderClouds(ctx, cfg, srcImg) {
    // Get a copy of the source resized to fit the output
    img := image.Resize( cfg.cloudWidth, cfg.cloudHeight, srcImg, "")

    // Run the cloud filter generating the statistics and
    filter := cloud.FilterNoMask().Limit( 0.7 )
    image.FilterOver(filter.Filter(),img)
    coverage := filter.Coverage()

    // Render the results
    try( ctx ) {
        gc := ctx.Gc()
        gc.Translate(cfg.cloudX, cfg.cloudY)
        gc.DrawImage( img )

        gc.SetFillColor( cfg.white )
        graph.SetFont( gc, "luxi 20 mono bold" )
        util.DrawStringLeft(gc,
            0, cfg.cloudHeight+20,
            "Cloud Cover %3.0f%%",
            coverage.Cloud*100
        )
    }
}

createSkyMap(cfg) {
    width := math.Min(cfg.auxViewW,cfg.auxViewH)
    chart := chart.NewHorizon( cfg.location, util.Rect(0,0,width,width).Rect() )
    chart.Background().SetFillStroke(cfg.black)
    chart.Horizon().SetFill(cfg.horizonColour).SetStroke(cfg.horizonBorder)

//    if cfg.milkyWay != nil {
//        chart.Feature("milkyway").SetFillStroke(cfg.milkyWay)
//    }

//    if cfg.constBorder != nil {
//        chart.Feature("const.border").SetStroke(cfg.constBorder)
//    }

//    if cfg.constLine != nil {
        chart.Feature("const.line").SetStroke(cfg.constLine)
//    }

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
