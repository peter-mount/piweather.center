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
    "github.com/peter-mount/go-anim/script/layout"
    "github.com/peter-mount/go-anim/script/render"
    "github.com/peter-mount/go-anim/script/util"
    "github.com/peter-mount/piweather.center/script/astro/calendar"
    "github.com/peter-mount/piweather.center/script/astro/chart"
    "github.com/peter-mount/piweather.center/script/astro/geo"
    "github.com/peter-mount/piweather.center/script/weather/cloud"
    "github.com/peter-mount/piweather.center/script/weather/keogram"
)

include (
    "demos/timelapse/skycamcommon.c"
    "demos/timelapse/skycamlayout.c"
)

main() {

    ctx := graph.New4k()
    //ctx := graph.New1080p().Scale(0.5,0.5)
    //ctx := graph.New720p().Scale(1/3.0,1/3.0)

    // The width of each column on the top part of the frame
    // e.g. this contains the camera image, cloud cover and sky map
    topColWidth := image.Width4K/3
    // This is the width of each entry, reduced so there's a gap between them
    topColCellWidth := topColWidth-20

    cfg := map(
        // Set this to the directory containing the images
        "srcDir": "/home/peter/weather/cam4",

        // Privacy mask
        "privmask": readImage( "/home/peter/weather/cam4-privmask.png" ),
        // Sky mask for cloud detection
        "skymask": readImage( "/home/peter/weather/cam4-skymask.png" ),

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

    createLayout(cfg)

    createSkyMap(cfg)

    files := util.GetImageFiles(cfg.srcDir)
    frames := util.SequenceIn(15,files,cfg.timeZone)

    // Here we render just the first frame
    if frames.HasNext() {
        frame := frames.Next()
        renderFrame(ctx,cfg,frame)
        image.WriteImage(cfg.output, ctx.Image())
    }
}
