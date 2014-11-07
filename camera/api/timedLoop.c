#include <stdlib.h>
#include <stdio.h>
#include "camera/camera.h"
#include "lib/list.h"

extern int verbose;

/**
 * A default implementation of a loop, this one captures a frame once every timelapse
 * and renders it.
 * 
 * @param arg
 * @return 
 */
void *camera_timedloop(void *arg) {
    CAMERA camera = (CAMERA) arg;

    time_t now, start;
    time_t timelapse = camera->timelapse;

    camera->runLoop = 1;
    while (camera->runLoop) {
        time(&start);

        // Capture a frame from the camera
        camera->registry->capture(camera);

        // Render it
        imagerenderer_render(camera);

        // Work out how long until the next trigger
        time(&now);
        start = start + timelapse;
        int delay = start - now;

        // Ensure we have a delay. If this triggers then timelapse is too short for
        // the amount of rendering required
        if (delay < 1) {
            fprintf(stderr, "WARNING: camera timelapse is too short for background processing, delay was %ds now using 1s\n", delay);
            delay = 1;
        }

        if (verbose > 1)
            fprintf(stderr, "Waiting %d\n", delay);
        sleep(delay);
        //vcos_sleep(delay*1000);
    }
}


