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
)

include (
    "demos/timelapse/skycamcommon.c"
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

        // Privacy mask
        "privmask": readImage( "/home/peter/weather/cam2-privmask.png" ),
        // Sky mask for cloud detection
        "skymask": readImage( "/home/peter/weather/cam2-skymask.png" ),

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

    createLayout(cfg)

    createSkyMap(cfg)

    ctx := graph.New4k()
    //ctx := graph.New1080p().Scale(0.5,0.5)
    //ctx := graph.New720p().Scale(1/3.0,1/3.0)

    files := util.GetImageFiles(cfg.srcDir)
    frames := util.SequenceIn(15,files,cfg.timeZone)

    // Here we render just the first frame
    if frames.HasNext() {
        frame := frames.Next()
        renderFrame(ctx,cfg,frame)
        image.WriteImage(cfg.output, ctx.Image())
    }
}

createLayout(cfg) {
    cfg.layout = layout.New(image.Width4K,image.Height4K).
        RowContainer().
            ColScaleContainer(1/3.0,1/3.0,1/3.0).
                Font("luxi 32 mono bold").
                Fill( cfg.white ).
                Text("",cfg.title).End().
                Text("","ME15Weather").Align("center").End().
                Text("timeDisplay","%s").Align("right").End().
            End().
            ColScaleContainer(0.4,0.4,0.2).
                Image("skyCamera").Inset(10).End().
                Image("auxView").Inset(10).End().
                RowContainer().
                    Font("luxi 20 mono bold").
                    Fill( cfg.white ).
                    Text("cloudCover", "Cloud Cover %3.0f%% Sky %3.0f%% Obscured %3.0f%%").End().
                    Text("jab1", "JabberGhjkfjdf").End().
                    Text("jab2", "JabberGhjkfjdf").End().
                End().
            End().
        End().
    End().
    Build()
}