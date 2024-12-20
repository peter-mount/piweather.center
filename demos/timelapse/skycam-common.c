// Common functionality for the skycam scripts

import (
    "github.com/peter-mount/go-anim/script/image"
    "github.com/peter-mount/piweather.center/astro/calculator"
    "github.com/peter-mount/piweather.center/script/astro/calendar"
    "github.com/peter-mount/piweather.center/script/weather/value"
)

include (
    "demos/timelapse/skycam-clouds.c"
    "demos/timelapse/skycam-skymap.c"
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

//        if sun.GetHorizontal().Alt >= 0 {
//            renderClouds( cfg, skyImage )
//        } else {
            renderSkyMap( cfg, jd )
//        }

        cfg.layout.Layout(gc)

        cfg.layout.Get("keogram").Sample(skyImage)

        cfg.layout.Draw(gc)
    }
}
