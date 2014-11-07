#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include "astro/math.h"

double astro_polynomial(double t, int termc, const double *termv) {
    double T = 1.0;
    double v = 0.0;
    int i;
    for (i = 0; i < termc; i++) {
        v += *termv++ * T;
        T *= t;
    }
    return v;
}
