#
# Calculate the times of Sunrise and Sunset for a location on Earth
# using the Meeus method.
#

main() {

    // Location of London, UK
    location := geo.LatLong(52.5, 0.5, 0)

    // Julian Day to calculate from
    day:= astroTime.FromDate2(2023,6,24,0,0,0)

    eq := astro.Sun.ApparentEquatorial(day)
    rs := eq.RiseSet( location.Coord(), day.Apparent0UT(), astro.Angle.AngleFromMin(-50) )

    fmt.Printf(" Julian Day %.3f\n", day)
    fmt.Printf(" Equatorial %.6f\t%.6f\n", eq.Alpha.Hour(), eq.Delta.Deg() )
    fmt.Printf("Circumpolar %v\n", rs.Rise)
    fmt.Printf("       Rise %v\n", rs.Rise)
    fmt.Printf("    Transit %v\n", rs.Rise)
    fmt.Printf("        Set %v\n", rs.Rise)
    fmt.Printf("   Duration %v\n", rs.Rise)
    fmt.Printf(" Day Length %v\n", rs.Rise)
}