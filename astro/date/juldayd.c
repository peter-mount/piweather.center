
#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include <time.h>
#include "astro/time.h"

double astro_julday_d(int year, int month, double day) {
    double A, B;
    double Y = (double) year;
    double M = (double) month;
    if (month <= 2) {
        Y -= 1.0;
        M += 12.0;
    }
    if (year > 1582 || (year == 1582 && (month > 10 || (month == 10 && day > 5.0)))) {
        A = trunc(Y / 100.0);
        B = 2.0 - A + trunc(A / 4.0);
    } else {
        B = 0;
    }
    return trunc(DAYS_YEAR * (Y + 4716.0)) + trunc(30.6001 * (M + 1)) + day + B - 1524.5;
}
