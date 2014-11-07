/* 
 * File:   time.h
 * Author: Peter T Mount
 *
 * Created on April 27, 2014, 9:39 AM
 */

#ifndef TIME_H
#define	TIME_H

#include <stdlib.h>
#include <sys/types.h>
#include <math.h>
#include <time.h>

// Number of seconds in a day
#define SECONDS_DAY_I 86400
#define SECONDS_DAY_D 86400.0

// Number of seconds in an hour
#define SECONDS_HOUR_I 3600
#define SECONDS_HOUR_D 3600.0

// Number of seconds in a minute
#define SECONDS_MINUTE_I 60
#define SECONDS_MINUTE_D 60.0

// Number of seconds in a siderial day
#define SECONDS_SIDERIALDAY_D 8640184.812866

// Days in a julian century
#define DAYS_JULIAN_CENTURY_D 36524.25

#define DAYS_YEAR 365.25
#define DAYS_CENTURY_I 36525
#define DAYS_CENTURY_D 36525.0

// Julian Day Number for 2000 Jan 1.0
#define JD_2000_I 2451545
#define JD_2000_D 2451545.0

// Julian Day Number for 2000 Jan 0.0
// used in some calculations, equivalent to 1999 Dec 31, 0h UT
#define JD_2000_JAN_0_0_I 2451544
#define JD_2000_JAN_0_0_D 2451544.0

// Days since 2000 Jan 0.0 from Julian Day Number
#define astro_days_since_2000(jd) ((jd)-JD_2000_JAN_0_0_D)

// Julian day jumber of day 0 in Modified Julian Day number
#define MJD0 2400000.5

extern void astro_calday(struct tm *tm, double jd);

extern double astro_julday(int year, int month, int day, int h, int m, int s);
extern double astro_julian_0h(double jd);
extern double astro_julian_century(double jd);
extern double astro_julday_d(int year, int month, double day);
extern double astro_julday_tm(struct tm *tm);
extern double astro_julday_time(time_t *time);

extern double astro_mjd(int year, int month, int day, int h, int m, int s);
extern double astro_mjd_d(int year, int month, double day);
extern double astro_mjd_tm(struct tm *tm);
extern double astro_mjd_time(time_t *time);

extern double astro_siderial_greenwich_0h(double jd);
extern double astro_siderial_greenwich(double jd);

#endif	/* TIME_H */

