# month-chart.c
#
# Generate a chart of a metric for a specific month
#

daily3d(config) {

    query := fmt.Sprintf("between %q and %q every %q table select timeof(last(%s)), max(%s)",
        config.startDate, config.endDate, config.every, config.metric, config.metric
    )

    fmt.Println(query)
    try( result := weatherdb.Query( config.dbUrl, query)) {
        if result.Status != 200 {
            fmt.Printf("Got %d\n%s\n", result.Status, result.Message )
            return null
        }
        if len(result.Table)==0 || len(result.Table[0].Rows)==0 {
            fmt.Println("No results returned")
            return null
        }

        // range of Y axes
        fmt.Println("Getting min/max values")
        yMin := 999999.0
        yMax := -999999.0
        rMax := 0
        for _,row := range result.Table[0].Rows {
            rMax = math.Max(rMax,len(row))
            for _,v := range row {
                yMin = math.Min(yMin,v)
                yMax = math.Max(yMax,v)
            }
        }

        xScale := config.w / rMax
        xAxisScale := config.w / 25 // 24 hours but we need to account for 0

        // Number of days to plot
        days := len(result.Table[0].Rows)

        fmt.Printf("Data range: %.3f ... %.3f, rows %d\n",yMin,yMax, days)

        // Scale of zAxis
        zScale := 5.0

        // Coordinates of origin
        x0 := 100 + (days * zScale)
        y0 := 100 + config.h + (days * zScale)
        x1 := x0 + (24*xAxisScale)
        y1 := y0 - config.h

        // Calculate scale, handle /by-zero
        dy := (yMax-yMin)
        if dy == 0.0 {dy = 1}
        yScale := config.h / dy

        // Graphics context with final image size
        ctx := animGraphic.NewSizedContext(x0+config.w,y0+50)

        try( ctx ) {
            gc := ctx.Gc()

            image.Fill( ctx, colour.Colour(config.background) )

            // Draw axes
            try(ctx) {
                cx := x0 - (zScale*days) + (config.w/2)

                gc.SetFillColor(colour.Colour(config.axesColour))

                // Main title
                animGraphic.SetFont( gc, "luxi 18 mono bold" )
                animUtil.DrawString(gc, cx, 40-14-12, config.title)

                animGraphic.SetFont( gc, "luxi 14 mono bold" )
                animUtil.DrawString(gc, cx, 45, "For the period %s to %s", config.startDate, config.endDate)

                // y-axis
                try(ctx) {
                    zOffset := zScale*days
                    x0L := x0 - zOffset
                    y0L := y0 - zOffset
                    x1 := x0 + config.w
                    y1 := y0 + config.w

                    animGraphic.SetFont( gc, "luxi 10 mono bold" )
                    yStep := (yMax-yMin)/10
                    gc.BeginPath()
                    for i:=0; i<=10; i=i+1 {
                        y := config.h*i/10.0
                        gc.MoveTo(x0L, y0L-y)
                        gc.LineTo(x0L-10, y0L-y)
                        gc.MoveTo(x1, y1-y)
                        gc.LineTo(x1-10, y1-y)
                        s := fmt.Sprintf("%.1f",i*yStep)
                        gc.FillStringAt(s,x0L-5-(10*len(s)),y0L-y+5)
                    }
                    gc.Stroke()

                    // y-axis label
                    try(ctx) {
                        animGraphic.SetFont( gc, "luxi 14 mono bold" )
                        sMin := len(fmt.Sprintf("%.1f",yMin))
                        sMax := len(fmt.Sprintf("%.1f",yMax))
                        gc.Translate(x0L-20-(10*math.Max(sMin,sMax)),y0L-(config.h/2))
                        gc.Rotate( -math.Pi/2.0 )
                        animUtil.DrawString(gc, 0,0, config.yAxisLabel)
                    }
                }

                // x-axis
                try(ctx) {
                    animGraphic.SetFont( gc, "luxi 10 mono bold" )
                    gc.BeginPath()
                    for hr:=0; hr<=24; hr=hr+1 {
                        x := x0+(hr*xAxisScale)
                        animUtil.DrawString(gc, x,y0+14, "%d", hr)
                        gc.MoveTo(x,y0)
                        gc.LineTo(x,y0+10)
                    }
                    gc.Stroke()

                    animGraphic.SetFont( gc, "luxi 14 mono bold" )
                    animUtil.DrawString(gc, x0 + (config.w/2), y0+35, "Local Time")
                }

                // z-axis
                try(ctx) {
                    animGraphic.SetFont( gc, "luxi 10 mono bold" )
                    gc.BeginPath()
                    for day,row := range result.Table[0].Rows {
                        // Time of start of day in time.Time
                        tDay := row[0].Time
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
                gc.SetStrokeColor( colour.Colour(config.axesColour) )

                zOffset := zScale*days

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

            // Plot Data
            try( ctx ) {
                gc.SetFillColor( colour.Colour(config.fillColour) )
                gc.SetStrokeColor( colour.Colour(config.strokeColour) )

                gc.BeginPath()

                for rowNum,row := range result.Table[0].Rows {
                    gc.BeginPath()
                    zOffset := zScale * (days-rowNum)
                    px0:=0              // first point
                    py0:=0
                    px1:=0              // last point's x
                    py1:=y0-zOffset     // Y of origin at z

                    tM := row[0].Time
                    for vN,v := range row {
                        if vN>0 {
                            y := (v - yMin) * yScale
                            x = ((vN - 1)*xScale) + x0 - zOffset
                            y = y0 - y - zOffset
                            gc.LineTo(x,y)

                            px1 =x
                            if vN==1 {
                                px0=x
                                py0=y
                            }
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
                gc.SetFillColor( colour.Colour(config.fillColour) )
                gc.SetStrokeColor( colour.Colour(config.strokeColour) )

                zOffset := zScale*(days-1)
                gc.MoveTo(x0-zOffset,y0-zOffset)

                for day,row := range result.Table[0].Rows {
                    zOffset := zScale*(days-day-1)
                    if len(row) > 1 {
                        reading := row[1]
                        y := (reading-yMin) * yScale
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
                gc.SetStrokeColor( colour.Colour(config.axesColour) )

                zOffset := zScale*days

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

        return ctx.Image()
    } catch(e) {
        fmt.Println("Failed to query db:", e)
    }

}
