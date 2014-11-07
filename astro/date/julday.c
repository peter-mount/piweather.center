#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include <time.h>
#include "astro/time.h"

double astro_julday(int year, int month, int day, int h, int m, int s) {
    double D = (double) ((h * SECONDS_HOUR_I)+(m * SECONDS_MINUTE_I) + s);
    D = (double) day + (D / SECONDS_DAY_D);
    return astro_julday_d(year, month, D);
}
