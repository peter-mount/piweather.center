#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include <time.h>
#include "astro/time.h"

double astro_julday_time(time_t *time) {
    struct tm tm;
    gmtime_r(time, &tm);
    return astro_julday_tm(&tm);
}
