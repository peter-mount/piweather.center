
#include <string.h>
#include "camera/camera.h"
#include "lib/config.h"
#include "lib/list.h"

extern int verbose;

static void stop(CAMERA camera) {

    // Let the camera loop stop on the next iteration
/*
    camera->runLoop = 0;
*/

    // No camera then do nothing
    if (!camera->enabled)
        return;

    if (verbose > 1)
        fprintf(stderr, "Closing down\n");

    imagerenderer_stop(camera);

    // call stop hook, usually to allow the camera to interrupt a semaphore?
    if (camera->registry->stop)
        camera->registry->stop(camera);

    if (verbose > 1)
        fprintf(stderr, "Camera shutdown\n");
}

/**
 * Stop all cameras.
 * 
 * NOTE: Later starting a camera will have the option of running during specific
 * hours of the day, i.e. daytime only.
 * 
 * For now it's the current run all the time mode.
 * 
 * @return 
 */
void camera_stop() {
    struct Node *n = list_getHead(&cameras);
    while (list_isNode(n)) {
        CAMERA camera = (CAMERA) n;
        n = n->n_succ;
        stop(camera);
    }
}
