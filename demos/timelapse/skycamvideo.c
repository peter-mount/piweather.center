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
    "demos/timelapse/skycam-config.c"
    "demos/timelapse/skycam-common.c"
    "demos/timelapse/skycam-layout.c"
)

main() {

    //ctx := graph.New4k()
    //ctx := graph.New1080p().Scale(0.5,0.5)
    ctx := graph.New720p().Scale(1/3.0,1/3.0)

    cfg := config()
    // Set this to the directory containing the images
    cfg.srcDir = "/home/peter/weather/cam4"

    // Privacy mask
    cfg.privmask = readImage( "/home/peter/weather/cam4-privmask.png" )
    // Sky mask for cloud detection
    cfg.skymask = readImage( "/home/peter/weather/cam4-skymask.png" )

    // The output video name
    cfg.output = "/home/peter/test-video.mp4"

    // Location of London, UK
    cfg.location = geo.LatLong(51.5, -8/60.0, 0)
    cfg.timeZone = time.LoadLocation("Europe/London")

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
