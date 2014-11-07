
#include <stdlib.h>
#include "camera/camera.h"
#include "renderers/imagerenderer.h"

void imagerenderer_stop(CAMERA camera) {
    struct Node *n = camera->renderers.renderers.l_head;
    while (list_isNode(n)) {
        struct image_renderer *r = (struct image_renderer *) n;
        n = n->n_succ;
        if (r->stop)
            r->stop(r, camera);
    }
}

