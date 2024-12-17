// Example script to plot a stereographic chart centered on the constellation of Orion.

main() {
    black := colour.Colour("black")
    white := colour.Colour("white")

    // Create the context we will render into
    ctx := animGraphic.NewSizedContext(900,900)

    // Start a Horizon chart for the specified location and date
    bounds := ctx.Image().Bounds()
    chart := chart.NewStereographic( angle.RAFromHour(5.5).Angle(), angle.AngleFromDeg(0.0), bounds.Dx(), bounds )

    // Set the background to black
    chart.Background().SetFillStroke(black)

    chart.Feature("milkyway").SetFillStroke(colour.Grey(17))
    chart.Feature("const.border").SetStroke(colour.Colour("#444444"))
    chart.Feature("const.line").SetStroke(colour.Colour("#555555"))
    chart.YaleBrightStarCatalog()

    // Finally draw the image
    try (gc := ctx.Gc() ) {
        chart.Draw(gc)
    }

    fileName:="/home/peter/test-spherical.png"
    try( f:=os.Create(fileName) ) {
        render.Encoder(fileName).Encode(f,ctx.Image())
    }
}
