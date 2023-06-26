# daily-line-2d.c
#
# This will accept a date and, from the stations archives plot a line graph
# of a specific sensor.
#

main() {
    // Load the weather station config
    weather.LoadConfig( "workConfig/station.yaml" )

    // Sensor reading to plot
    sensor:= "home.ecowitt.temp"

    // Size of area
    w := 1440 // Also number of minutes in 24 hours
    h := 300

    // Coordinates of origin
    x0 := 100
    y0 := 100 + h

    ctx := animGraphic.NewSizedContext(w+200,h+200)

    // Create a memory store
    try( store := weather.NewStore( "/home/peter/tmp/weather/data" ) ) {

        t := astroTime.StartOfToday()
        store.Load(t.Time())

        // range of Y axes TODO make this dependent on the data
        yMin := 10.0
        yMax := 30.0
        yScale := 300.0 / (yMax-yMin)

        try( ctx ) {
            gc := ctx.Gc()

            image.Fill( ctx, colour.Colour("white") )

            // Draw axes
            gc.SetStrokeColor( colour.Colour("black") )
            gc.BeginPath()
            gc.MoveTo(x0,y0-h)
            gc.LineTo(x0,y0)
            gc.LineTo(x0+w,y0)
            gc.Stroke()

            gc.BeginPath()
            readings := store.Get(sensor)
            for i, reading := range readings {
                x := i + x0
                y := y0-((reading.Value-yMin) * yScale)
                //fmt.Printf("%4d %6.3f %6.3f\n",i,x,y)
                gc.LineTo(x,y)
            }
            gc.SetStrokeColor( colour.Colour("red") )
            gc.Stroke()
        }
    }

    try( f:=os.Create("daily-line-2d.png") ) {
        png.Encode(f,ctx.Image())
    }
}
