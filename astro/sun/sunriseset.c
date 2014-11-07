
#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include <time.h>
#include <stdio.h>
#include "astro/time.h"
#include "astro/math.h"
#include "astro/observatory.h"
#include "astro/sun.h"

/**
 * Calculates the time of sunrise and sunset and the length of the day
 * 
 * @param jd Julian Day Number, 1801-2099 only
 * @param lon Longitude, East positive
 * @param lat Latitude, North positive
 * @param altit Altitude sun must cross for rise/set
 * @param upper_limb 1 upper limb, 0 center
 * @param trise time rise
 * @param tset time set
 * @return 0 sun rises/sets, 1 sun above horizon all day, 0 sun never rises
 */
static void sunriset(double jd, RISE_SET *rs, OBSERVATORY *obs, double altit, int upper_limb) {
    // Days since 2000 Jan 0.0 (negative before)
    double d;
    // Obliquity (inclination) of Earth's axis
    double obl_ecl;
    // Solar distance, astronomical units
    double sr;
    // True solar longitude
    double slon;
    // Sine of Sun's declination
    double sin_sdecl;
    // Cosine of Sun's declination
    double cos_sdecl;
    // Sun's Right Ascension
    double sRA;
    // Sun's declination
    double sdec;
    // Sun's apparent radius
    double sradius;
    // Diurnal arc
    double t;
    // Time when Sun is at south
    double tsouth;
    // Local sidereal time
    double sidtime;
    double gmsidtime;
    double cost;

    // Compute d of 12h local mean solar time
    d = astro_days_since_2000(jd) + 0.5 - obs->longitude / 360.0;

    // Compute the local sidereal time of this moment
    gmsidtime = 15.0 * astro_siderial_greenwich_0h(jd);
    sidtime = astro_range_360(gmsidtime + 180.0 + obs->longitude);

    // Compute Sun's RA, Decl and distance at this moment
    astro_sunRADec(jd, &sRA, &sdec, &sr);

    // Compute time when Sun is at south - in hours UT
    tsouth = 12.0 - astro_range_180(sidtime - sRA) / 15.0;

    // Compute obliquity of ecliptic (inclination of Earth's axis)
    obl_ecl = 23.4393 - 3.563E-7 * d;

    // Compute Sun's ecliptic longitude and distance
    astro_sunpos(jd, &slon, &sr);

    // Compute sine and cosine of Sun's declination
    sin_sdecl = sind(obl_ecl) * sind(slon);
    cos_sdecl = sqrt(1.0 - sin_sdecl * sin_sdecl);

    // Compute the Sun's apparent radius in degrees
    sradius = 0.2666 / sr;

    // Do correction to upper limb, if necessary
    if (upper_limb)
        altit -= sradius;

    // Compute the diurnal arc that the Sun traverses to reach
    // the specified altitude altit:
    cost = (sind(altit) - sind(obs->latitude) * sind(sdec)) / (cosd(obs->latitude) * cosd(sdec));
    if (cost >= 1.0) {
        // Sun always below altit
        rs->type = RISE_SET_NEVER_RISES;
        rs->length = 0.0;
        t = 0.0;
    } else if (cost <= -1.0) {
        // Sun always above altit
        rs->type = RISE_SET_NEVER_SETS;
        rs->length = 24.0;
        t = 12.0;
    } else {
        // The diurnal arc, hours
        rs->type = RISE_SET_NORMAL;
        rs->length = (2.0 / 15.0) * acosd(cost);
        t = acosd(cost) / 15.0;
    }

    // Store rise and set times - in hours UT
    rs->rise = tsouth - t;
    rs->set = tsouth + t;
}

void astro_sunriseset(double jd, OBSERVATORY *obs, SOLAR_EPHEMERIS *se) {
    se->jd = jd;
    sunriset(jd, &se->day, obs, SUN_DAY_ALTITUDE, 1);
    sunriset(jd, &se->civil, obs, SUN_CIVIL_ALTITUDE, 0);
    sunriset(jd, &se->nautical, obs, SUN_NAUTICAL_ALTITUDE, 0);
    sunriset(jd, &se->astronomical, obs, SUN_ASTRONOMICAL_ALTITUDE, 0);
}