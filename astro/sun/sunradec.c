
#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include <time.h>
#include "astro/time.h"
#include "astro/math.h"

/**
 * Calculates the Sun's equatorial coordinates RA, Dec and also its distance
 * 
 * @param jd Julian Day Number
 * @param RA Right Ascension
 * @param dec Declination
 * @param r Distance
 */
void astro_sunRADec(double jd, double *RA, double *dec, double *r) {
    double lon, obl_ecl, x, y, z;

    // Compute Sun's ecliptical coordinates
    astro_sunpos(jd, &lon, r);

    // Compute ecliptic rectangular coordinates (z=0)
    x = *r * cosd(lon);
    y = *r * sind(lon);

    // Compute obliquity of ecliptic (inclination of Earth's axis)
    obl_ecl = 23.4393 - 3.563E-7 * astro_days_since_2000(jd);

    // Convert to equatorial rectangular coordinates - x is unchanged
    z = y * sind(obl_ecl);
    y = y * cosd(obl_ecl);

    // Convert to spherical coordinates
    *RA = atan2d(y, x);
    *dec = atan2d(z, sqrt(x * x + y * y));
}
