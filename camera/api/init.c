/*
 * Initialise the camera system
 */

#include <string.h>
#include "camera/camera.h"
#include "lib/config.h"
#include "lib/list.h"
#include "imagefilter/filter.h"
#include "scheduler/scheduler.h"
#include "lib/string.h"

#include <stdio.h>

struct List cameras;
extern verbose;

static void *handler(void *arg) {
    CAMERA camera = (CAMERA) arg;

    if (verbose > 1)
        fprintf(stderr, "exec: %s\n", camera->cmd);

    system(camera->cmd);

    sleep(1);

    FILE *file = fopen(camera->image, "r");
    if (file) {

        if (verbose > 1)
            fprintf(stderr, "reading: %s\n", camera->image);

        bytebuffer_reset(&camera->imagedata);
        bytebuffer_read(&camera->imagedata, file);
        fclose(file);

        if (verbose > 1)
            fprintf(stderr, "rendering: %s\n", camera->image);

        imagerenderer_render(camera);

        if (verbose > 1)
            fprintf(stderr, "camera complete\n");
    } else if (verbose > 1)
        fprintf(stderr, "cannot find: %s\n", camera->image);
}

// Sets the default config common to all cameras

static void default_config(CAMERA camera, CONFIG_SECTION *sect) {
    camera->enabled = 1;

    config_getCharParameter(sect, "camera-command", &camera->cmd);
    fatalIfNull(camera->cmd, "No camera command provided\n");

    config_getCharParameter(sect, "camera-image", &camera->image);
    fatalIfNull(camera->image, "No camera image provided\n");
}

static void schedule_config(CAMERA camera, CONFIG_SECTION *sect) {
    // The schedule, defaults to every minute
    char *spec = "* *", *filter = NULL;
    config_getCharParameter(sect, "schedule", &spec);
    config_getCharParameter(sect, "schedule-filter", &filter);

    SCHEDULE_ENTRY *e = scheduler_new(handler, (void *) camera);

    // Lower the priority so sensors run first
    e->node.pri = -50;

    if (scheduler_parse(e, spec))
        fprintf(stderr, "Unsupported schedule \"%s\" in %s\n", spec, sect->node.name);

    if (filter)
        if (scheduler_parse_filter(e, filter))
            fprintf(stderr, "Unsupported filter \"%s\" in %s\n", filter, sect->node.name);

    scheduler_add(e);
}

void camera_init() {
    list_init(&cameras);


    struct Node *n = list_getHead(&config->sections);
    while (list_isNode(n)) {
        CONFIG_SECTION *sect = (CONFIG_SECTION *) n;
        n = n->n_succ;

        CONFIG_PARAM *p = config_getParameter(sect, "camera-command");
        if (p) {

            if (verbose > 1)
                fprintf(stderr, "Configuring camera %s\n", sect->node.name);

            CAMERA camera = (CAMERA) malloc(sizeof (struct camera_config));
            memset(camera, 0, sizeof (CAMERA));
            bytebuffer_init(&camera->imagedata);

            default_config(camera, sect);

            list_addTail(&cameras, &camera->node);

            schedule_config(camera, sect);

            if (verbose > 1)
                fprintf(stderr, "Configured camera %s\n", sect->node.name);
        }
    }
}