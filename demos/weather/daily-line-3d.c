# daily-line-2d.c
#
# This will accept a date and, from the stations archives plot a line graph
# of a specific sensor.
#

main() {
    // Load the weather station config
    weather.LoadConfig( "workConfig/station.yaml" )

    title := "A229 Pollution 2.5µm"
    yAxisLabel := "Pollution µg/m³"
    // Sensor reading to plot
    //sensor:= "home.ecowitt.temp"
    sensor:= "home.drive.pm2_5"

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
            fmt.Printf("\rLoading %s", t.Time().Format("2006 Jan 02"))
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
            readings = reducer.Max(readings)

            // Append the reduced readings to the array ready for plotting
            data=append(data,readings)
            t=t.Add(1)
        }
        fmt.Println()
    }

    // range of Y axes TODO make this dependent on the data
    yMin := 999999.0
    yMax := -999999.0
    for _,readings := range data {
        for _,reading := range readings {
            v:=reading.Value
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

        try(ctx) {
            gc.SetFillColor(colour.Colour("black"))

            // Main title
            animGraphic.SetFont( gc, "luxi 18 mono bold" )
            animUtil.DrawString(gc, x0 + (w/2), 40-14-12, title)

            animGraphic.SetFont( gc, "luxi 14 mono bold" )
            animUtil.DrawString(gc, x0 + (w/2), 45,
             "For the period %s to %s",
             t0.Add(-days).Time().Format("2006 Jan 02"),
             t0.Add(-1).Time().Format("2006 Jan 02")
             )

            // y-axis
            try(ctx) {
                zOffset := zScale*(days-1)
                x0L := x0 - zOffset
                y0L := y0 - zOffset
                x1 := x0 + w
                y1 := y0 + w

                animGraphic.SetFont( gc, "luxi 10 mono bold" )
                yStep := (yMax-yMin)/10
                gc.BeginPath()
                for i:=0; i<=10; i=i+1 {
                    y := h*i/10.0
                    gc.MoveTo(x0L, y0L-y)
                    gc.LineTo(x0L-10, y0L-y)
                    gc.MoveTo(x1, y1-y)
                    gc.LineTo(x1-10, y1-y)
                    s := fmt.Sprintf("%.0f",i*yStep)
                    gc.FillStringAt(s,x0L-5-(10*len(s)),y0L-y+5)
                }
                gc.Stroke()

                // y-axis label
                try(ctx) {
                    animGraphic.SetFont( gc, "luxi 14 mono bold" )
                    sMin := len(fmt.Sprintf("%.0f",yMin))
                    sMax := len(fmt.Sprintf("%.0f",yMax))
                    gc.Translate(x0L-20-(10*math.Max(sMin,sMax)),y0L-(h/2))
                    gc.Rotate( -math.Pi/2.0 )
                    animUtil.DrawString(gc, 0,0, yAxisLabel)
                }
            }

            // x-axis
            try(ctx) {
                animGraphic.SetFont( gc, "luxi 10 mono bold" )
                gc.BeginPath()
                for hr:=0; hr<=24; hr=hr+1 {
                    x := x0+(hr*60)
                    animUtil.DrawString(gc, x,y0+14, "%d", hr)
                    gc.MoveTo(x,y0)
                    gc.LineTo(x,y0+10)
                }
                gc.Stroke()

                animGraphic.SetFont( gc, "luxi 14 mono bold" )
                animUtil.DrawString(gc, x0 + (w/2), y0+35, "Local Time")
            }

            // z-axis
            try(ctx) {
                animGraphic.SetFont( gc, "luxi 10 mono bold" )
                gc.BeginPath()
                for day,_ := range data {
                    // Time of start of day in time.Time
                    tDay := t0.Add(-days+day).Time()
                    // Label every Sunday
                    if tDay.Weekday() == 0 {
                        zOffset := (days-day)*zScale
                        x := x0-zOffset
                        y := y0-zOffset
                        gc.MoveTo(x,y)
                        gc.LineTo(x-10,y)

                        s := tDay.Format("Jan 02")
                        gc.FillStringAt(s,x-10-(10*len(s)),y+5)
                    }
                }
                gc.Stroke()
            }
        }

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
