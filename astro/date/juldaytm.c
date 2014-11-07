#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include <time.h>
#include "astro/time.h"

double astro_julday_tm(struct tm *tm) {
    return astro_julday(tm->tm_year + 1900, tm->tm_mon + 1, tm->tm_mday, tm->tm_hour, tm->tm_min, tm->tm_sec);
}
