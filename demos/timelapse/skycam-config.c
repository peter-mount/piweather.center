
import (
    "github.com/peter-mount/go-anim/script/colour"
    "github.com/peter-mount/go-anim/script/image"
)


config() {

    // The width of each column on the top part of the frame
    // e.g. this contains the camera image, cloud cover and sky map
    topColWidth := image.Width4K/3
    // This is the width of each entry, reduced so there's a gap between them
    topColCellWidth := topColWidth-20

    return map(
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
}