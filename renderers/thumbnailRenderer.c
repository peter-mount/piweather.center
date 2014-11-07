/**
 * Renderer which generates a thumbnail image from the raw one
 * 
 */
#include "weatherstation/main.h"
#include "camera/camera.h"
#include "webserver/webserver.h"
#include "lib/charbuffer.h"
#include "sensors/sensors.h"
#include <stdio.h>
#include <stdlib.h>
#include <gd.h>
#include "annotation.h"
#include "lib/config.h"
#include "renderers/imagerenderer.h"
#include "renderers/image/image.h"
#include "imagefilter/filter.h"

struct state {
    // This must be first
    struct image_renderer meta;
    // Enable this image? implicitly set if any option is defined
    int enabled;
    // image dimensions
    struct image thumbnail;
};

static void init(struct image_renderer *r, CAMERA camera) {
    struct state *state = (struct state *) r;

    // Historically it's in with annotation
    CONFIG_SECTION *sect = config_getSection("annotation");
    if (!sect)
        return;

    init_image(sect, &state->thumbnail, "thumbnail");

    // If thumbnail is defined then enable it
    state->enabled = state->thumbnail.enabled;
}

static void postinit(struct image_renderer *r, CAMERA camera) {

    // No camera then do nothing
    if (!camera->enabled)
        return;

    struct state *state = (struct state *) r;

    // Only run if configured
    if (!state->enabled)
        return;

    // Default thumbnail to 100x75
    if (!state->thumbnail.width || !state->thumbnail.height) {
        state->thumbnail.width = IMAGE_THUMB_WIDTH;
        state->thumbnail.height = IMAGE_THUMB_HEIGHT;
    }

    state->thumbnail.url = genurl((const char*) webserver.contextPath, "imageThumb.jpg");
    webserver_add_response_handler(state->thumbnail.url);
}

static void render(struct image_renderer *r, CAMERA camera) {

    // No camera then do nothing
    if (!camera->enabled)
        return;

    struct state *state = (struct state *) r;

    // Only run if configured
    if (!state->enabled)
        return;

    gdImagePtr raw = imagerenderer_getRawImage(&camera->renderers);
    if (raw) {
        gdImagePtr thumb = imagefilter_createThumbnail(raw, state->thumbnail.width, state->thumbnail.height);
        imagerenderer_createJpegResponse(&camera->renderers, thumb, 90, state->thumbnail.url);
        gdImageDestroy(thumb);
    }
}

struct image_renderer *create_thumbnailrenderer() {

    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    // Register the functions
    state->meta.init = init;
    state->meta.postinit = postinit;
    state->meta.render = render;

    return (struct image_renderer *) state;
}
