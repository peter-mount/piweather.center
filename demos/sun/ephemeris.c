# sun/ephemeris.c
#
# Draw a graph of the sun across the year, including times of sun rise, set and twilight.
#
#

main() {

    // The year to plot
    year := 2023

    // The start time in julian day's
    jd := astroTime.FromDate( year, 1, 1, 0, 0, 0)

    // Location of London, UK
    location := geo.LatLong(51.5, -8/60.0, 0)

    // Number of days in this year
    jd1 := astroTime.FromDate( year+1, 1, 1, 0, 0, 0)
    days := jd1-jd
    fmt.Println("Year",year,"days",days)

    hourScale   := 20       // Pixels horizontally per hour
    dayScale    := 5        // Pixels vertically per day
    w := 24*hourScale
    h := (days-1)*dayScale

    // Calculate the ephemeris
    data := newArray()
    for i:=0; i<days; i=i+1 {
        if i % 15 == 0 {
            fmt.Print("\rCalculating Ephemeris ",jd.Time().Format("2006 Jan 02"))
        }
        t := value.BasicTime(jd.Time(), location.Coord(), 0.0)
        data=append( data, calculator.SolarEphemeris( t ) )
        jd=jd.Add(1)
    }
    fmt.Println()

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
            gc.SetStrokeColor( colour.Colour("lightGrey"))
            gc.Stroke()
        }


        try (ctx) {

            gc.BeginPath()
            valid := false
            for doy, ephemeris := range data {
                valid = plot(ctx, valid, doy, ephemeris.AstronomicalDawn, hourScale, dayScale)
            }
            valid = false
            for doy, ephemeris := range data {
                valid = plot(ctx, valid, doy, ephemeris.AstronomicalDusk, hourScale, dayScale)
            }
            gc.SetStrokeColor( colour.Colour("black") )
            gc.Stroke()

            gc.BeginPath()
            valid := false
            for doy, ephemeris := range data {
                valid = plot(ctx, valid, doy, ephemeris.NauticalDawn, hourScale, dayScale)
            }
            valid = false
            for doy, ephemeris := range data {
                valid = plot(ctx, valid, doy, ephemeris.NauticalDusk, hourScale, dayScale)
            }
            gc.SetStrokeColor( colour.Colour("blue") )
            gc.Stroke()

            gc.BeginPath()
            valid := false
            for doy, ephemeris := range data {
                valid = plot(ctx, valid, doy, ephemeris.CivilDawn, hourScale, dayScale)
            }

            valid = false
            for doy, ephemeris := range data {
                valid = plot(ctx, valid, doy, ephemeris.CivilDusk, hourScale, dayScale)
            }
            gc.SetStrokeColor( colour.Colour("red") )
            gc.Stroke()

            gc.BeginPath()
            valid := false
            for doy, ephemeris := range data {
                valid = plot(ctx, valid, doy, ephemeris.SunRise, hourScale, dayScale)
            }

            valid = false
            for doy, ephemeris := range data {
                valid = plot(ctx, valid, doy, ephemeris.SunSet, hourScale, dayScale)
            }
            gc.SetStrokeColor( colour.Colour("darkGreen") )
            gc.Stroke()
        }
    }

    try( f:=os.Create("sun-ephemeris.png") ) {
        png.Encode(f,ctx.Image())
    }
}

plot(ctx, valid, doy, altAz, hourScale, dayScale) {
    altAzValid := altAz.IsValid()
    if altAzValid {
        tm := math.Float(altAz.Time.Hour())+(math.Float(altAz.Time.Minute())/60.0)
        x := 50 + (tm * hourScale)
        y := 50 + (doy * dayScale)

        gc := ctx.Gc()
        if valid {
            gc.LineTo(x,y)
        } else {
            gc.MoveTo(x,y)
        }
    }
    return altAzValid
}
