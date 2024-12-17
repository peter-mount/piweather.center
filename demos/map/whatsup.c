// Example script to plot a chart based on what's visible in the sky at a specific time
// and location on the Earth.

main() {
    // Location of London, UK
    location := geo.LatLong(51.5, -8/60.0, 0)
    timeZone := time.LoadLocation("Europe/London")

    black := colour.Colour("black")
    white := colour.Colour("white")
    horizonColour := colour.Colour("darkgreen")

    jd := astroTime.FromTime( time.Now())
    fmt.Printf("Generating map for %v\n", jd)

    // Create the context we will render into
    ctx := animGraphic.NewSizedContext(900,900)

    // Start a Horizon chart for the specified location and date
    chart := chart.NewHorizon( location, ctx.Image().Bounds() ).JD(jd)

    // Set the background to black
    chart.Background().SetFillStroke(black)

    // The horizon colour
    chart.Horizon().SetFillStroke(colour.Colour("#00320033"))

    //chart.Feature("milkyway").SetFillStroke(colour.Grey(17))

    //chart.Feature("const.border").SetStroke(colour.Colour("#0000aa"))

    chart.Feature("const.line").SetStroke(colour.Colour("#0000aa"))

    chart.YaleBrightStarCatalog()//.MagLimit(6)

    // Finally draw the image
    try (gc := ctx.Gc() ) {
        chart.Draw(gc)
    }

    fileName:="/home/peter/test-horizon.png"
    try( f:=os.Create(fileName) ) {
        render.Encoder(fileName).Encode(f,ctx.Image())
    }
}
