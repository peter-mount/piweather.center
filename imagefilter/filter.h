/* 
 * File:   filter.h
 * Author: Peter T Mount
 *
 * Created on April 20, 2014, 9:33 AM
 */

#ifndef FILTER_H
#define	FILTER_H

#include <gd.h>
#include <stdlib.h>
#include <stdint.h>

typedef enum {
    IMAGE_CHANNEL_RED,
    IMAGE_CHANNEL_GREEN,
    IMAGE_CHANNEL_BLUE
} IMAGE_CHANNEL_T;

// Histograms
int *imagefilter_histogram(gdImagePtr img, IMAGE_CHANNEL_T channel);

// Generic functions
extern gdImagePtr imagefilter_createThumbnail(gdImagePtr srcImg, int w, int h);
extern gdImagePtr imagefilter_readFile( char *n );
extern void imagefilter_writeFile( gdImagePtr img, char *n );

// Image filter/transformations
extern gdImagePtr imagefilter_add(gdImagePtr img1, gdImagePtr img2);
extern gdImagePtr imagefilter_difference(gdImagePtr img1, gdImagePtr img2);
extern gdImagePtr imagefilter_intersect(gdImagePtr img1, gdImagePtr img2);
extern gdImagePtr imagefilter_merge(gdImagePtr img1, gdImagePtr img2);
extern gdImagePtr imagefilter_subtract(gdImagePtr img1, gdImagePtr img2);

// Masking filters
extern gdImagePtr imagefilter_apply_mask(gdImagePtr src, gdImagePtr mask);
extern gdImagePtr imagefilter_sun(gdImagePtr srcImage, int lim);

// ==================================================
// R/B Filter, use for cloud detection

// A default value for rblim
#define DEFAULT_RBLIM 0.84

extern gdImagePtr imagefilter_rb_ratio(gdImagePtr srcImage,
        double rblim,
        uint32_t cloud, uint32_t sky, uint32_t black,
        int *pixTotal,
        int *pixCloud);

#endif	/* FILTER_H */

