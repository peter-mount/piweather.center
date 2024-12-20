// This script provides methods to detect cloud cover and display as a separate view

import (
    "github.com/peter-mount/go-anim/script/image"
    "github.com/peter-mount/piweather.center/script/weather/cloud"
)

// render clouds from a source image
//
// cfg      Configuration
// srcImg   Source image from the camera
//
// requirements:
//
// cfg.privmask     The privacy mask - to limit what should not be shown
// cfg.skymask      The sky mask - to limit what is unobstructed sky
//
// components in cfg.layout:
//
// "auxView"    Image component to display the cloud image
// "cloudCover" Text component to write the cloud cover value to
//
renderClouds( cfg, srcImg ) {
    // Run the cloud filter generating the statistics and an image
    //
    // The limit 0.7 is the level used to determine clouds.
    // 0.7 seems reasonable, but originally I used 0.84
    filter := cloud.Filter( cfg.privmask, cfg.skymask ).Limit( 0.7 )
    img := image.FilterNew(filter.Filter(),srcImg)

    // The coverage statistics from the filter
    coverage := filter.Coverage()

    // Render the results
    //
    // NB: If the totals rendered are >100% this is down to %f in Sprintf always
    // rounding up if the fraction is >=0.5 so although it appears wrong internally it's
    // correct.
    cfg.layout.Get("auxView").SetImage(img)
    cfg.layout.Get("cloudCover").Args(coverage.Cloud, coverage.Sky, coverage.Obscured)
}
