# sun/ephemeris.c
#
# Draw a graph of the sun across the year, including times of sun rise, set and twilight.
#
# It does this by running through the entire year calculating the Suns ephemeris for each day
# and then plotting the times for Astronomical, Nautical, Civil twilight and Sun rise/Set.
#
# Issues:
#
# This is a bit convoluted because currently due to issue go-script#2 we cannot change existing
# values in a map, so we have to pass parameters in a long form.
# Once that issue is fixed then this can be made simpler when plotting the twilight/rise/set times.
#

main() {

    // The year to plot
    year := 2023

    // Location of London, UK
    location := geo.LatLong(51.5, -8/60.0, 0)
    timeZone := time.LoadLocation("Europe/London")

    // Alternate location: New York, New York, USA
//    location := geo.LatLong(40.7128, -74.0060, 0)
//    timeZone := time.LoadLocation("US/Eastern")

    // The start time in julian day's
    // Note: use time.Date so we go from the required timeZone
    jd0 := astroTime.FromTime( astroTime.LocalMidnight( time.Date( year, 1, 1, 0, 0, 0, 0, timeZone) ) )

    // Number of days in this year by counting julian days from the start
    // of next year
    jd1 := astroTime.FromTime( astroTime.LocalMidnight( time.Date( year+1, 1, 1, 0, 0, 0, 0, timeZone) ) )
    days := jd1 - jd0

    hourScale   := 20       // Pixels horizontally per hour
    dayScale    := 5        // Pixels vertically per day
    w := 24*hourScale
    h := (days-1)*dayScale

    // Calculate the ephemeris
    data := newArray()
    {
        jd := jd0
        for i:=0; i<days; i=i+1 {
            if jd.Time().Weekday() == 0 {
                fmt.Print("\rCalculating Ephemeris ",jd.Time().Format("2006 Jan 02"))
            }
            t := value.BasicTime(jd.Time(), location.Coord(), 0.0)
            data=append( data, calculator.SolarEphemeris( t ) )
            jd=jd.Add(1)
        }
        fmt.Println()
    }

    // Graphics context with final image size
    ctx := animGraphic.NewSizedContext(w+100,h+100)

    try( ctx ) {
        gc := ctx.Gc()

        image.Fill( ctx, colour.Colour("white") )

        gc.SetStrokeColor( colour.Colour("black") )
        gc.BeginPath()
        animGraphic.Rectangle(gc,50,50,w,h)
        gc.Stroke()

        // Draw hour lines
        try (ctx) {
            gc.BeginPath()
            for hr:=0;hr<24;hr=hr+1 {
                x = 50 + (hr * hourScale)
                gc.MoveTo(x,50)
                gc.LineTo(x,50+h)
            }
            gc.SetStrokeColor( colour.Colour("lightGrey"))
            gc.Stroke()

            gc.BeginPath()
            for hr:=0;hr<24;hr=hr+3 {
                x = 50 + (hr * hourScale)
                gc.MoveTo(x,50)
                gc.LineTo(x,50+h)
            }
            gc.SetStrokeColor( colour.Colour("grey"))
            gc.Stroke()

        }

        // Draw day lines
        try (ctx) {
            gc.BeginPath()
            for doy, _ := range data {
                y = 50 + (doy * dayScale)
                gc.MoveTo(50,y)
                gc.LineTo(50+(24*hourScale),y)
            }
            gc.SetStrokeColor( colour.Colour("#aaaaaa"))
            gc.Stroke()
        }

        // Draw week lines
        try (ctx) {
            jd = jd0
            gc.BeginPath()
            for doy:=0; doy<=days; doy=doy+1 {
                if jd.Time().Weekday() == 0 {
                    y = 50 + (doy * dayScale)
                    gc.MoveTo(50,y)
                    gc.LineTo(50+(24*hourScale),y)
                }
                jd=jd.Add(1)
            }
            gc.SetStrokeColor( colour.Colour("#666666"))
            gc.Stroke()
        }

        // Draw month lines
        try (ctx) {
            gc.BeginPath()
            for m:=1; m<=12; m=m+1 {
                doy := astroTime.FromDate( year, m, 1, 0, 0, 0) - jd0
                y = 50 + (doy * dayScale)
                gc.MoveTo(50,y)
                gc.LineTo(50+(24*hourScale),y)
            }
            gc.SetStrokeColor( colour.Colour("#333333"))
            gc.Stroke()
        }

        // Hour and date labels
        try( ctx ) {

            //gc.SetStrokeColor(colour.Colour("blue"))
            gc.SetFillColor(colour.Colour("black"))

            animGraphic.SetFont( gc, "luxi 14 mono bold" )
            animUtil.DrawString(gc, 50 + (12*hourScale), 40-14-16, "Daylight & Twilight for %d", year)

            animGraphic.SetFont( gc, "luxi 10 mono" )
            animUtil.DrawString(gc, 50 + (12*hourScale), 40-14, "Local Time of Day")

            animUtil.DrawString(gc, 50 + (12*hourScale), 50 + h + 30,
                "Lat %.4f Long %.4f %s",
                location.Latitude.Deg(),
                location.Longitude.Deg(),
                timeZone.String() )

            animGraphic.SetFont( gc, "luxi 8 mono" )

            // Hour labels
            for hr:=0;hr<=24;hr=hr+3 {
                x = 50 + (hr * hourScale)
                gc.SetFillColor(colour.Colour("black"))
                animUtil.DrawString(gc, x, 40, "%d", hr)
            }

            lastTz := ""
            for jd:= jd0; jd < jd1; jd=jd.Add(1) {
                t := jd.Time()
                doy := jd - jd0
                y = 50 + (doy * dayScale)

                gc.SetFillColor(colour.Colour("black"))

                // At mid-month show the month name
                if t.Day() == 15 {
                    try(ctx) {
                        gc.Translate(20,y)
                        gc.Rotate( -math.Pi/2.0 )
                        animUtil.DrawString(gc, 0,0, t.Month().String())
                    }
                }

                // Date labels shown for each Sunday
                if t.Weekday()==0 {
                    animUtil.DrawString(gc, 35, y, "%d", t.Day())
                }

                // Show Timezone at start and at any changes during the year
                tz := t.In(timeZone).Format("MST")
                if tz != lastTz {
                    gc.BeginPath()
                    gc.MoveTo(50+w,y)
                    gc.LineTo(55+w,y)
                    gc.Stroke()
                    gc.FillStringAt( tz, 55 + w, y)
                    zone := t.In(timeZone).Zone()
                    gc.FillStringAt( fmt.Sprintf("UTC%+d", zone[1]/3600), 55+w, y+10)
                    lastTz = tz
                }
            }

            // Legend
            y := 50 + h + 15
            gc.SetFillColor( colour.Colour("black"))
            gc.FillStringAt("Astronomical",50,y)

            gc.SetFillColor( colour.Colour("blue"))
            gc.FillStringAt("Nautical",130,y)

            gc.SetFillColor( colour.Colour("red"))
            gc.FillStringAt("Civil Twilight",185,y)

            gc.SetFillColor( colour.Colour("darkGreen"))
            gc.FillStringAt("Sun Rise/Set",280,y)
        }

        try (ctx) {

            gc.BeginPath()
            valid := false
            lx := -1
            for doy, ephemeris := range data {
                r = plot(ctx, valid, doy, ephemeris.AstronomicalDawn, hourScale, dayScale, timeZone, lx)
                valid=r[0]
                lx=r[1]
            }
            valid = false
            lx := -1
            for doy, ephemeris := range data {
                r = plot(ctx, valid, doy, ephemeris.AstronomicalDusk, hourScale, dayScale, timeZone, lx)
                valid=r[0]
                lx=r[1]
            }
            gc.SetStrokeColor( colour.Colour("black") )
            gc.Stroke()

            gc.BeginPath()
            valid := false
            lx := -1
            for doy, ephemeris := range data {
                r = plot(ctx, valid, doy, ephemeris.NauticalDawn, hourScale, dayScale, timeZone, lx)
                valid=r[0]
                lx=r[1]
            }
            valid = false
            lx := -1
            for doy, ephemeris := range data {
                r = plot(ctx, valid, doy, ephemeris.NauticalDusk, hourScale, dayScale, timeZone, lx)
                valid=r[0]
                lx=r[1]
            }
            gc.SetStrokeColor( colour.Colour("blue") )
            gc.Stroke()

            gc.BeginPath()
            valid := false
            lx := -1
            for doy, ephemeris := range data {
                r = plot(ctx, valid, doy, ephemeris.CivilDawn, hourScale, dayScale, timeZone, lx)
                valid=r[0]
                lx=r[1]
            }

            valid = false
            lx := -1
            for doy, ephemeris := range data {
                r = plot(ctx, valid, doy, ephemeris.CivilDusk, hourScale, dayScale, timeZone, lx)
                valid=r[0]
                lx=r[1]
            }
            gc.SetStrokeColor( colour.Colour("red") )
            gc.Stroke()

            gc.BeginPath()
            valid := false
            lx := -1
            for doy, ephemeris := range data {
                r = plot(ctx, valid, doy, ephemeris.SunRise, hourScale, dayScale, timeZone, lx)
                valid=r[0]
                lx=r[1]
            }

            valid = false
            lx := -1
            for doy, ephemeris := range data {
                r = plot(ctx, valid, doy, ephemeris.SunSet, hourScale, dayScale, timeZone, lx)
                valid=r[0]
                lx=r[1]
            }
            gc.SetStrokeColor( colour.Colour("darkGreen") )
            gc.Stroke()
        }
    }

    try( f:=os.Create("sun-ephemeris.png") ) {
        png.Encode(f,ctx.Image())
    }
}

plot(ctx, valid, doy, altAz, hourScale, dayScale, timeZone, lx) {
    altAzValid := altAz.IsValid()
    x:=lx
    if altAzValid {
        tt := altAz.Time.In(timeZone)
        tm := math.Float(tt.Hour())+(math.Float(tt.Minute())/60.0)
        x = 50 + (tm * hourScale)
        y := 50 + (doy * dayScale)

        gc := ctx.Gc()
        if valid && lx > 0 && math.Abs(x-lx)<(3*hourScale) {
            gc.LineTo(x,y)
        } else {
            gc.MoveTo(x,y)
        }
    }
    return append(newArray(), altAzValid, x)
}
