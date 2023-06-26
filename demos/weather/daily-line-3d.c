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

    // a Reducer reduces recorded data into fixed periods of time.
    // So here, with underlying readings every 20 seconds, a reducer
    // with 5 minutes will reduce the number of points by a factor of 15
    reducer := weather.ReducerMinutes(5)

    // Size of area
    w := 1440 // Also number of minutes in 24 hours
    h := 300

    days := 31      // Number of days to plot
    zScale := 5.0   // Scale of zAxis

    // Coordinates of origin
    x0 := 100 + (days * zScale)
    y0 := 100 + h + (days * zScale)
    x1 := x0 + w
    y1 := y0 - h

    // Time of start of today in julian.Day's
    t0 := astroTime.StartOfToday()

    // Extract the data
    data := newArray()
    try( store := weather.NewStore( "/home/peter/tmp/weather/data" ) ) {
        t := t0.Add(-days)
        for i:=0; i<days; i=i+1 {
            // Load the data for the day contained by t
            store.Load(t.Time())

            // Get the sensor readings
            readings:= store.Get(sensor)

            // reduce the readings so we have a uniform data set.
            //
            // The following reductions are available:
            //
            // Min the minimum value within the reduction period
            // Max the maximum value within the reduction period
            // Sum the sum of all values in the reduction period
            // Mean the Sum within the reduction period / the number of entries in that period.
            readings = reducer.Mean(readings)

            // Append the reduced readings to the array ready for plotting
            data=append(data,readings)
            t=t.Add(1)
        }
    }

    // range of Y axes TODO make this dependent on the data
    yMin := 999999.0
    yMax := -999999.0
    for _,readings := range data {
        for _,reading := range readings {
            v:=reading.Value.Float()
            yMin = math.Min(yMin,v)
            yMax = math.Max(yMax,v)
        }
    }
    fmt.Printf("Data range: %.3f ... %.3f\n",yMin,yMax)
    // Calculate scale, handle /by-zero
    dy := (yMax-yMin)
    if dy == 0.0 {dy = 1}
    yScale := h / dy

    // Graphics context with final image size
    ctx := animGraphic.NewSizedContext(x0+w+100,y0+100)

    try( ctx ) {
        gc := ctx.Gc()

        image.Fill( ctx, colour.Colour("white") )

        // Draw background axes
        try(ctx) {
            gc.SetStrokeColor( colour.Colour("black") )

            zOffset := zScale*(days-1)

            gc.BeginPath()
            // y-axis
            gc.MoveTo(x0-zOffset,y0-zOffset)
            gc.LineTo(x0-zOffset,y1-zOffset)
            gc.LineTo(x1-zOffset,y1-zOffset)
            gc.LineTo(x1-zOffset,y0-zOffset)
            // z-axis
            gc.MoveTo(x1,y1)
            gc.LineTo(x1-zOffset,y1-zOffset)
            gc.MoveTo(x1,y0)
            gc.LineTo(x1-zOffset,y0-zOffset)
            gc.MoveTo(x1,y1)
            gc.LineTo(x1,y0)
            gc.Stroke()
        }

        try(ctx) {
            gc.SetFillColor( colour.Colour("white") )
            gc.SetStrokeColor( colour.Colour("red") )

            for day,readings:= range data {
                // Time of start of day in time.Time
                tDay := t0.Add(-days+day).Time()

                // z-axis offset
                zOffset := zScale*(days-day-1)

                gc.BeginPath()
                px0:=0              // first point
                py0:=0
                px1:=0              // last point's x
                py1:=y0-zOffset     // Y of origin at z
                for i, reading := range readings {
                    x := reading.Time.Sub(tDay).Minutes()
                    y := (reading.Value-yMin) * yScale
                    x = x + x0 - zOffset
                    y = y0 - y - zOffset
                    gc.LineTo(x,y)

                    px1 =x
                    if i==0 {
                        px0=x
                        py0=y
                    }
                }
                gc.LineTo(px1,py1)
                gc.LineTo(px0,py1)
                gc.LineTo(px0,py0)
                gc.FillStroke()
            }
        }

        // Draw left hand side as a solid shape
        try(ctx) {
            gc.SetFillColor( colour.Colour("white") )
            gc.SetStrokeColor( colour.Colour("red") )

            zOffset := zScale*(days-1)
            gc.MoveTo(x0-zOffset,y0-zOffset)

            for day,readings:= range data {
                zOffset := zScale*(days-day-1)
                if len(readings) > 0 {
                    reading := readings[0]
                    y := (reading.Value-yMin) * yScale
                    x := x0 - zOffset
                    y = y0 - y - zOffset
                    gc.LineTo(x,y)
                }
            }
            gc.LineTo(x0,y0)
            zOffset := zScale*(days-1)
            gc.LineTo(x0-zOffset,y0-zOffset)
            gc.FillStroke()
        }

        // Draw foreground axes
        try(ctx) {
            gc.SetStrokeColor( colour.Colour("black") )

            zOffset := zScale*(days-1)

            gc.BeginPath()
            // y-axis
            gc.MoveTo(x0,y0)
            gc.LineTo(x0,y1)
            // x-axis
            gc.MoveTo(x0,y0)
            gc.LineTo(x1,y0)
            gc.MoveTo(x0,y1)
            gc.LineTo(x1,y1)
            // z-axis
            gc.MoveTo(x0,y0)
            gc.LineTo(x0-zOffset,y0-zOffset)
            gc.MoveTo(x0,y1)
            gc.LineTo(x0-zOffset,y1-zOffset)
            gc.Stroke()
        }

    }

    try( f:=os.Create("daily-line-2d.png") ) {
        png.Encode(f,ctx.Image())
    }
}
