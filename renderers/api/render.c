
#include <stdlib.h>
#include "camera/camera.h"
#include "renderers/imagerenderer.h"

void imagerenderer_render(CAMERA camera) {
    struct Node *n = camera->renderers.renderers.l_head;
    while (list_isNode(n)) {
        struct image_renderer *r = (struct image_renderer *) n;
        n = n->n_succ;
        if (r->render)
            r->render(r, camera);
    }

    // Free up memory
    imagerenderer_freeImages(&camera->renderers);
    bytebuffer_free(&camera->renderers.buffer);
}

