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
    // is this the firstFrame? if so then we need to render it twice
    // to ensure the layout is correct
    firstFrame := !mapContains(cfg,"layoutCompleted")

    // Get time from the file name
    srcTime := frame.Time
    jd := calendar.FromTime(srcTime)
    tm := value.BasicTime(srcTime,cfg.location.Coord(),0)

    layout := cfg.layout

    // Calculate the solar system
    ephem := calculator.SolarSystem(tm)

    // Calculate the sun - we need the altitude to know when to show the clouds or
    // the sky map.
    sun := ephem.GetByName("Sun")
    sunLimit := 0 // Show clouds whilst the Sun is above the horizon
    altAz := sun.GetHorizontal()
    sunAlt := altAz.Alt
    layout.Get("sunAltAz").Args( sunAlt, altAz.Az )
    layout.Get("sunDist").Args( sun.GetDistance() )
    layout.Get("sunTime").Args( sun.GetLightTime() )

    moon := ephem.GetByName("Moon")
    altAz = moon.GetHorizontal()
    layout.Get("moonAltAz").Args( altAz.Alt, altAz.Az )
    layout.Get("moonDist").Args( moon.GetDistance() )
    layout.Get("moonTime").Args( moon.GetLightTime() )

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

        // render clouds whilst the sun is above sunLimit
        // however we also do this for the firstFrame regardless as we
        // need to configure the Image component with the correct size
        if firstFrame || sunAlt >= sunLimit {
            renderClouds( cfg, skyImage )
        }

        if !firstFrame && sunAlt < sunLimit {
            renderSkyMap( cfg, jd, skyImage, ephem )
        }

        // Layout on the first frame only
        if firstFrame cfg.layout.Layout(gc)

        // These must be done after we are laid out as they need to know their sizes
        cfg.layout.Get("keogram").Sample(skyImage)

        cfg.layout.Draw(gc)

        // If first frame then render it a second time.
        // This allows us to be certain the first frame is laid out correctly
        // specifically the Image components
        if firstFrame {
            cfg.layoutCompleted = true

            // We also need a copy of the bounds for auxView as right now it's got the cloud view
            cfg.skymapBounds := cfg.layout.Get("auxView").InsetBounds()

            renderFrame(ctx,cfg,frame)
        }

    }
}
