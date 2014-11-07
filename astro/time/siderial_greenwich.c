#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include <time.h>
#include "astro/time.h"
#include "astro/math.h"

// Polynomial constants for siderial time in degrees
static const double T0[] = {100.46061837, 36000.770053608, 0.000387933, -1 / 38710000.0};
// 4 terms in T0
#define T0n 4
/**
 * Siderial time at Greenwich at 0h
 * @param jd Julian day number
 * @return siderial time in decimal hours
 */
double astro_siderial_greenwich_0h(double jd) {
    double Z = astro_julian_0h(jd);
    double T = astro_julian_century(Z);
    double t = astro_polynomial(T, T0n, T0);
    return astro_range( t, 360.0 ) /15.0;
}

double astro_siderial_greenwich(double jd) {
    double F, Z;

    // Z is the JD at midnight
    F = modf(jd, &Z);
    Z += 0.5;

    double T = (Z - 2451545.0) / DAYS_CENTURY_D;

    double theta = 280.46061837 + (360.98564736629 * (jd - 2451545.0))+(0.000387933 * T * T) - (T * T * T / 38710000.0);
    theta /= 360.0;
    F = modf(theta, &Z);
    if (F < 0.0)
        F += 1.0;
    return F * 24.0;
}