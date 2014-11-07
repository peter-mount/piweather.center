/**
 * Renderer which takes the captured image and applies a R/B filter to identify clouds from the sky
 * 
 */

#include <gd.h>
#include <stdlib.h>
#include <math.h>
#include "camera/camera.h"
#include "renderers/annotation.h"
#include "webserver/webserver.h"
#include "clouds/clouds.h"
#include "imagefilter/filter.h"

/*
 * These are the Aviation units for Oktas.
 * 
 * http://en.wikipedia.org/wiki/Okta
 */
static const char *OKTA[] = {
    // 0 Sky clear
    "SKC",
    // 1..2 Few
    "FEW", "FEW",
    // 3..4 Scattered
    "SCT", "SCT",
    // 5..7 Broken
    "BKN", "BKN", "BKN",
    // 8 Overcast
    "OVC",
    // Not defined in aviation but in case we support 9 & prevent a segfault, Obscured
    "OBS"
};

void clouds_render(struct image_renderer *r, CAMERA camera) {
    // No camera then do nothing
    if (!camera->enabled)
        return;

    struct cloud_renderer *state = (struct cloud_renderer *) r;

    double rblim = state->rblim;
    int cloud = 0, total = 0;

    // The source image in SD (640x480)
    gdImagePtr srcImage = imagerenderer_getSDImage(&camera->renderers);
    // Flag to indicate if we destroy srcImage, i.e. if it's the one from
    // the renderer above then we must not destroy it
    int freeImage = 0;

    // Apply mask
    gdImagePtr img = NULL;
    if (state->mask) {
        // Create the mask image
        img = imagefilter_apply_mask(srcImage, state->mask);

        if (img) {
            // Use mask as the source image
            srcImage = img;
            freeImage = 1;
        }
    }

    // Detect the sun
    if (state->sunLimit) {
        img = imagefilter_sun(srcImage, state->sunLimit);

        if (img) {
            if (freeImage)
                gdImageDestroy(srcImage);
            srcImage = img;
            freeImage = 1;
        }
    }

    // Identify the clouds
    gdImagePtr clouds = imagefilter_rb_ratio(srcImage, rblim, state->cloud, state->sky, state->black, &total, &cloud);

    // Free memory of the masked image as no longer needed
    // but never free the one from the renderer
    if (freeImage)
        gdImageDestroy(srcImage);


    // Generate it's thumbnail
    gdImagePtr thumb = imagefilter_createThumbnail(clouds, IMAGE_THUMB_WIDTH, IMAGE_THUMB_HEIGHT);

    // Make them public
    imagerenderer_createJpegResponse(&camera->renderers, clouds, 90, state->response);
    imagerenderer_createJpegResponse(&camera->renderers, thumb, 90, state->thumb);

    // Free memory
    gdImageDestroy(thumb);
    gdImageDestroy(clouds);

    // Calculate the percentage cloud cover
    double cloud_pc = 0.0;
    if (total > 0)
        cloud_pc = 100.0 * (double) cloud / (double) total;

    // According to the MetOffice, Okta is measured in eights, so its int(pc/12.5)
    // http://www.metoffice.gov.uk/climatechange/science/monitoring/ukcp09/faq.html#faq1.11
    int okta;

    // total is number of sky/cloud pixels, not black so check to see if it's
    // at least 1/8 of the image size
    int lowLim = clouds->sx * clouds->sy / 8;
    if (total > lowLim) {
        // We have at least 1/8 of sky in calculation so calculate okta
        okta = round(cloud_pc / 12.5);
    } else {
        // Presume sky is obscured so okta 9
        okta = 9;
    }

    // sanity check on the range
    if (okta < 0)
        okta = 0;
    else if (okta > 9)
        okta = 9;

    // Update the sensor
    state->cloud_sensor->cloudcoverage = cloud_pc;

    // Now log. As we are logging outside of the sensor loop we must set the time here
    time(&state->cloud_sensor->sensor.last_update);
    state->okta_sensor->last_update = state->cloud_sensor->sensor.last_update;
    sensor_log(&state->cloud_sensor->sensor, (int) round(cloud_pc * 10.0), "Cloud %0.1f%%", cloud_pc);
    sensor_log(state->okta_sensor, okta, "Okta %d %s", okta, OKTA[okta]);
}

