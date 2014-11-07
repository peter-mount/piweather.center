#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include "astro/math.h"

double astro_range(double v, double max) {
    double i;
    double f = modf(v / max, &i);
    if (f < 0.0)
        f += 1.0;
    return f*max;
}

double astro_range_360(double v) {
    return astro_range(v,360.0);
}

double astro_range_180(double v) {
    double c = astro_range_360(v);
    if(c>180.0)
        c=c-360.0;
    return c;
}