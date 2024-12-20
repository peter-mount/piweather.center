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
    "demos/timelapse/skycam-common.c"
    "demos/timelapse/skycam-layout.c"
)

main() {

    //ctx := graph.New4k()
    //ctx := graph.New1080p().Scale(0.5,0.5)
    ctx := graph.New720p().Scale(1/3.0,1/3.0)

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
        "output": "/home/peter/test-video.mp4",

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
        // Position of map overlaid on the camera image.
        // This can take some trial and error, use skycamframe.c to render a single frame until
        // you get this just right
        "mapX": 115,
        "mapY": 30,
        "mapW": 785,
        "mapH": 730,
        // Map background
        "mapBackground": colour.Colour("#00000000"),
        //"milkyWay": colour.Grey(17),
        //"constBorder": colour.Colour("#0000aa"),
        "constLine": colour.Colour("#0000aa"),
        "magLimit": 5,
        // Horizon - horizonColour is required, border is optional
        //"horizonBorder": colour.Colour("white"),
        "horizonColour": colour.Colour("#00320033") // "#00320033"
    )

    createLayout(cfg)

    files := util.GetImageFiles(cfg.srcDir)
    frames := util.SequenceIn(15,files,cfg.timeZone)

    frameCount := frames.Size()
    fmt.Printf("Rendering %d frames\n", frameCount)

    try( out := render.New( cfg.output, 25 ) ) {
        for i,frame := range frames {
            renderFrame(ctx,cfg,frame)
            out.WriteImage(ctx.Image())

            //if i>100 { break }
        }
    }

    fmt.Println("\nRender complete")
}
