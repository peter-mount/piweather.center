/* 
 * File:   clouds.h
 * Author: Peter T Mount
 *
 * Created on April 14, 2014, 3:27 PM
 */

#ifndef CLOUDS_H
#define	CLOUDS_H

#include <stdint.h>
#include <pthread.h>
#include <gd.h>
#include "lib/bytebuffer.h"
#include "webserver/webserver.h"
#include "renderers/imagerenderer.h"
#include "sensors/sensors.h"

struct cloudcoverage_sensor {
    struct sensor sensor;
    // The percentage of cloud
    double cloudcoverage;
};

struct cloud_renderer {
    struct image_renderer renderer;
    // Colour for the sky, default #0000ff
    uint32_t sky;
    // Colour for clouds, default #ffffff
    uint32_t cloud;
    // Colour for anything else, default #000000
    uint32_t black;
    // R/B ratio above which we detect cloud
    double rblim;
    // If set the limit to use in detecting the sun
    int sunLimit;
    // =================
    // Internal use only
    // =================
    // Mask image (optional)
    gdImagePtr mask;
    // The 640x480 image showing the coverage
    const char *response;
    // Thumbnail of image
    const char *thumb;
    // The sensor
    struct cloudcoverage_sensor *cloud_sensor;
    // virtual sensor which converts percentage cover into Okta's
    struct sensor *okta_sensor;
};

extern void clouds_render_postinit(struct image_renderer *r, CAMERA camera);
extern void clouds_render(struct image_renderer *r, CAMERA camera);

#endif	/* CLOUDS_H */

