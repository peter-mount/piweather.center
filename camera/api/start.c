
#include <string.h>
#include "camera/camera.h"
#include "lib/config.h"
#include "lib/list.h"

extern int verbose;

/**
 * Start all cameras.
 * 
 * NOTE: Later starting a camera will have the option of running during specific
 * hours of the day, i.e. daytime only.
 * 
 * For now it's the current run all the time mode.
 * 
 * @return 
 */
void camera_start() {
    struct Node *n = list_getHead(&cameras);
    while (list_isNode(n)) {
        CAMERA camera = (CAMERA) n;
        n = n->n_succ;

        if (camera->enabled) {

            if (verbose > 1)
                fprintf(stderr, "Running with interval %ds\n", camera->timelapse);

            // Let the camera start itself
            if (camera->registry->start)
                camera->registry->start(camera);

            // Start the thread if needed
/*
            if (!camera->thread)
                pthread_create(&camera->thread, NULL, camera->loop, (void *) camera);
*/
        }
    }
}
