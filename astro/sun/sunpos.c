
#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include <time.h>
#include "astro/time.h"
#include "astro/math.h"

/**
 * Calculates the Sun's ecliptic longitude & distance at a given instant.
 * 
 * Note: The sun's ecliptic latitude is not calculated as it's very near to 0
 * 
 * @param jd Julian Day Number
 * @param lon Sun's ecliptic longitude
 * @param r solar distance
 */
void astro_sunpos(double jd, double *lon, double *r) {
    // Days since 2000 Jan 0.0
    double d = astro_days_since_2000(jd);
    // Mean anomaly of the Sun
    double M;
    // Mean longitude of perihelion
    double w;
    // Note: Sun's mean longitude = M + w
    // Eccentricity of Earth's orbit
    double e;
    // Eccentric anomaly
    double E;
    // x, y coordinates in orbit 
    double x, y;
    // True anomaly
    double v;

    // Compute mean elements
    M = astro_range(356.0470 + 0.9856002585 * d, 360.0);
    w = 282.9404 + 4.70935E-5 * d;
    e = 0.016709 - 1.151E-9 * d;

    // Compute true longitude and radius vector
    E = M + e * RADEG * sind(M) * (1.0 + e * cosd(M));
    x = cosd(E) - e;
    y = sqrt(1.0 - e * e) * sind(E);

    // Solar distance
    *r = sqrt(x * x + y * y);

    // True anomaly
    v = atan2d(y, x);

    // True solar longitude
    *lon = astro_range(v + w, 360.0);
}
