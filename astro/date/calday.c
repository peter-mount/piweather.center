/*
 * Calculates the calendar date from a Julian Day Number
 */

#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include <string.h>
#include <stdbool.h>
#include <time.h>
#include "astro/time.h"

// Value of K for a Bissextile or common year used in calculating the day of year
#define K_BISSEXTILE 1
#define K_COMMON 2

static const char *TIMEZONE = "UTC";

/**
 * Calculates the calendar day from the supplied julian day number and places the result into the supplied struct tm.
 * 
 * Note: The result is always in UTC
 * 
 * @param tm struct tm to write the result
 * @param jd Julian Day Number
 */
void astro_calday(struct tm *tm, double jd) {
    double A, B, C, D, E, F, Z, DAY, tod;
    int month, year, k, doy;
    div_t d;
    bool gregorian;

    // Clear tm & set our timezone
    memset(tm, 0, sizeof (struct tm));
    tm->tm_gmtoff = 0;
    tm->tm_zone = TIMEZONE;

    // Calendar date from JD
    F = modf(jd + 0.5, &Z);
    if (Z < 2299161.0) {
        A = Z;
        gregorian = false;
    } else {
        double alpha = trunc((Z - 1867216.25) / DAYS_JULIAN_CENTURY_D);
        A = Z + 1.0 + alpha - trunc(alpha / 4.0);
        gregorian = true;
    }

    B = A + 1524.0;
    C = trunc((B - 122.1) / DAYS_YEAR);
    D = trunc(DAYS_YEAR * C);
    E = trunc((B - D) / 30.6001);

    // Month
    if (E < 14.0)
        month = (int) E - 1;
    else
        month = (int) E - 13;

    // tm_mon has January as 0 not 1
    tm->tm_mon = month - 1;

    // Year
    if (month > 2)
        year = (int) C - 4716;
    else
        year = (int) C - 4715;

    // tm_year has 1900 as year 0
    tm->tm_year = year - 1900;

    // Day of month
    tod = modf(B - D - trunc(30.6001 * E) + F, &DAY);
    tm->tm_mday = (int) DAY;

    // Time of day
    d = div((int) round(tod * SECONDS_DAY_D), SECONDS_HOUR_I);
    tm->tm_hour = d.quot;
    d = div(d.rem, SECONDS_MINUTE_I);
    tm->tm_min = d.quot;
    tm->tm_sec = d.rem;

    // Day of week
    d = div((int) Z + 1, 7);
    tm->tm_wday = d.rem;

    // Day of year
    if (gregorian) {
        // Gregorian calendar
        if (year % 400 == 0)
            k = K_BISSEXTILE;
        else if (year % 100 == 0)
            k = K_COMMON;
        else if (year % 4 == 0)
            k = K_BISSEXTILE;
        else
            k = K_COMMON;
    } else {
        // Julian calendar
        if (year % 4 == 0)
            k = K_BISSEXTILE;
        else
            k = K_COMMON;
    }
    doy = ((275 * month) / 9) - k * ((month + 9) / 12) + tm->tm_mday - 30;
    // tm_yday has Jan 1 as 0 not 1
    tm->tm_yday = doy - 1;
}
