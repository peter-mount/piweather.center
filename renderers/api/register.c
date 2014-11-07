
#include <stdlib.h>
#include "camera/camera.h"
#include "renderers/imagerenderer.h"

void imagerenderer_register(CAMERA camera, struct image_renderer *renderer) {
    list_addHead(&camera->renderers.renderers, &renderer->node);
}

