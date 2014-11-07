/**
 * The main core of the webcam server.
 * 
 * The main() function calls this code to start the system.
 * 
 */
#define WEBCAM_C

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "main.h"
#include "camera/camera.h"
#include "webserver/webserver.h"
#include "sensors/sensors.h"
#include "lib/charbuffer.h"
#include "lib/config.h"
#include "logger/logger.h"
#include "sensors/i2c/i2c.h"
#include "renderers/imagerenderer.h"
#include "weatherstation/weatherstation.h"
#include "astro/location.h"
#include "scheduler/scheduler.h"

// These are now mandatory so we define them here and register them in webcam_run() so they run first
extern struct image_renderer *create_annotatedrenderer();
extern struct image_renderer *create_rawrenderer();
extern struct image_renderer *create_thumbnailrenderer();

int webcam_run() {

#ifdef HAVE_CAMERA
    // Our mandatory renderers, ensures they are run first being registered last
    struct Node *n = cameras.l_head;
    while (list_isNode(n)) {
        CAMERA camera = (CAMERA) n;
        n = n->n_succ;
        // Put annotated last
        list_addTail(&camera->renderers.renderers, &create_annotatedrenderer()->node);
        // These go first
        list_addHead(&camera->renderers.renderers, &create_thumbnailrenderer()->node);
        list_addHead(&camera->renderers.renderers, &create_rawrenderer()->node);
    }
#endif

    // Now start the system up
    webserver_initialise(config);

#ifdef HAVE_CAMERA
    // Initialise the renderers
    n = cameras.l_head;
    while (list_isNode(n)) {
        CAMERA camera = (CAMERA) n;
        n = n->n_succ;
        imagerenderer_init(camera);
    }
#endif

    // Initialise the loggers
    logger_start();

    // Finish off configuring the webserver, default port etc
    webserver_set_defaults();

#ifdef HAVE_CAMERA
    n = cameras.l_head;
    while (list_isNode(n)) {
        CAMERA camera = (CAMERA) n;
        n = n->n_succ;
        imagerenderer_postinit(camera);
    }
#endif

    sensor_postinit();

    // The camera home page
    //create_homepage();

    // Start everything up
    webserver_start();

#ifdef HAVE_CAMERA
    camera_start();
#endif

    // Now the main loop, monitor for sensor updates
    //sensor_loop();
    while (1)
        sleep(60);

    // Shutdown - we never actually get here
    webserver_stop();
    logger_stop();

#ifdef HAVE_CAMERA
    camera_stop();
#endif

    return 0;
}