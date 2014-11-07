
#include <stdlib.h>
#include "camera/camera.h"
#include "renderers/imagerenderer.h"

void imagerenderer_initialise() {
    struct Node *n = cameras.l_head;
    while (list_isNode(n)) {
        CAMERA camera = (CAMERA) n;
        n = n->n_succ;

        bytebuffer_init(&camera->renderers.buffer);
        list_init(&camera->renderers.renderers);
        camera->renderers.images = hashmapCreate(10, hashmapStringHash, hashmapStringEquals);
    }
}

void imagerenderer_init(CAMERA camera) {
    struct Node *n = camera->renderers.renderers.l_head;
    while (list_isNode(n)) {
        struct image_renderer *r = (struct image_renderer *) n;
        n = n->n_succ;
        if (r->init)
            r->init(r, camera);
    }
}
