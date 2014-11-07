/* 
 * File:   math.h
 * Author: Peter T Mount
 *
 * Created on April 27, 2014, 6:12 PM
 */

#ifndef MATH_H
#define	MATH_H

#ifndef PI
#define PI        3.1415926535897932384
#endif

#define RADEG     ( 180.0 / PI )
#define DEGRAD    ( PI / 180.0 )

/* The trigonometric functions in degrees */

#define sind(x)  sin((x)*DEGRAD)
#define cosd(x)  cos((x)*DEGRAD)
#define tand(x)  tan((x)*DEGRAD)

#define atand(x)    (RADEG*atan(x))
#define asind(x)    (RADEG*asin(x))
#define acosd(x)    (RADEG*acos(x))
#define atan2d(y,x) (RADEG*atan2(y,x))

// Calculate polynomial in termv to termc terms
extern double astro_polynomial(double t, int termc, const double *termv);

// Convert v into the range 0<=v<max
extern double astro_range(double v, double max);
extern double astro_range_360(double v);
extern double astro_range_180(double v);

#endif	/* MATH_H */

