/*
 * Calculate the Modified Julian Day Number (MJD) from a date
 */

#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include "astro/time.h"

double astro_mjd(int year, int month, int day, int h, int m, int s) {
    return astro_julday(year, month, day, h, m, s) - MJD0;
}

double astro_mjd_d(int year, int month, double day) {
    return astro_julday_d(year, month, day) - MJD0;
}

double astro_mjd_tm(struct tm *tm) {
    return astro_julday_tm(tm) - MJD0;
}

double astro_mjd_time(time_t *time) {
    return astro_julday_time(time) - MJD0;
}