/**
 * Renderer which takes the captured image and applies a R/B filter to identify clouds from the sky
 * 
 */

#include "camera/camera.h"
#include "webserver/webserver.h"
#include "clouds/clouds.h"

void clouds_render_postinit(struct image_renderer *r, CAMERA camera) {

    // No camera then do nothing
    if (!camera->enabled)
        return;

    struct cloud_renderer *state = (struct cloud_renderer *) r;

    state->response = genurl((const char*) webserver.contextPath, "cloud.jpg");
    webserver_add_response_handler(state->response);

    state->thumb = genurl((const char*) webserver.contextPath, "cloud_thumb.jpg");
    webserver_add_response_handler(state->thumb);
}

