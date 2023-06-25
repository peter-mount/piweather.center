#
# Calculate the times of Sunrise and Sunset for a location on Earth
# using the Meeus method to get the approximate times
#

main() {

    // Location of London, UK
    location := geo.LatLong(52.5, 0.5, 0)

    // Julian Day to calculate from
    day:= astroTime.StartOfToday()

    eq := astro.Sun.ApparentEquatorial(day)
    rs := eq.RiseSet( location.Coord(), day.Apparent0UT(), astro.Angle.AngleFromMin(-50) )

    fmt.Printf("       Date %s\n", day.Time().Format(time.RFC822Z))
    fmt.Printf(" Julian Day %.3f\n", day)
    fmt.Printf(" Equatorial %.6f\t%.6f\n", eq.Alpha.Hour(), eq.Delta.Deg() )
    if rs.Circumpolar {
        fmt.Printf("Circumpolar\n")
    }else {
        fmt.Printf("       Rise %v UTC\n", astroTime.HourDMSString(rs.Rise))
        fmt.Printf("    Transit %v UTC\n", astroTime.HourDMSString(rs.Transit))
        fmt.Printf("        Set %v UTC\n", astroTime.HourDMSString(rs.Set))
    }
}