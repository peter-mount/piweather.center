/* 
 * File:   sun.h
 * Author: Peter T Mount
 *
 * Created on April 29, 2014, 1:16 PM
 */

#ifndef SUN_H
#define	SUN_H

#include "astro/observatory.h"

// The type of RISE_SET

typedef enum {
    // A normal day, sun rises & sets
    RISE_SET_NORMAL,
    // The sun never rises
    RISE_SET_NEVER_RISES,
    // The sun never sets
    RISE_SET_NEVER_SETS
} RISE_SET_T;

typedef struct {
    RISE_SET_T type;
    // Time of rising in UT
    double rise;
    // Time of setting in UT
    double set;
    // The length in hours when above the horizon
    double length;
} RISE_SET;

typedef struct {
    // Julian day number, 1801-2099 only
    double jd;
    // When the sun's upper limb is 35 arc minutes below the horizon or higher
    RISE_SET day;
    // Civil twilight, when sun center is 6 degrees below horizon or higher
    RISE_SET civil;
    // Nautical twilight, when sun center is 12 degrees below horizon
    // or higher
    RISE_SET nautical;
    // Astronomical twilight, when sun center is 18 degrees below horizon
    // or higher
    RISE_SET astronomical;
} SOLAR_EPHEMERIS;

// Altitude of sun upper lib below horizon for it to start the day
#define SUN_DAY_ALTITUDE (-35.0 / 60.0)
// Sun centre altitude for civil twilight
#define SUN_CIVIL_ALTITUDE -6.0
// Sun centre altitude for nautical twilight
#define SUN_NAUTICAL_ALTITUDE -12.0
// Sun centre altitude for astronomical twilight
#define SUN_ASTRONOMICAL_ALTITUDE -18.0

extern void astro_sunpos(double d, double *lon, double *r);
extern void astro_sunRADec(double d, double *RA, double *dec, double *r);

extern void astro_day_length(double jd, OBSERVATORY *obs, SOLAR_EPHEMERIS *srs);
extern void astro_sunriseset(double jd, OBSERVATORY *obs, SOLAR_EPHEMERIS *srs);

#endif	/* SUN_H */

