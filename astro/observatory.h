/* 
 * File:   observatory.h
 * Author: Peter T Mount
 *
 * Created on April 29, 2014, 2:55 PM
 */

#ifndef OBSERVATORY_H
#define	OBSERVATORY_H

typedef struct {
    // Name of Observatory
    char *name;
    // Longitude, east positive
    double longitude;
    // Latitude, north positive
    double latitude;
    // Altitude, metres
    double altitude;
} OBSERVATORY;

#endif	/* OBSERVATORY_H */

