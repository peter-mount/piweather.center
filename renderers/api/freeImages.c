
#include <stdlib.h>
#include "renderers/imagerenderer.h"

static bool callback(void *key, void *value, void *context) {
    struct image_renderers *ir = (struct image_renderers *) context;
    hashmapRemove(ir->images, key);
    if (value)
        gdImageDestroy((gdImagePtr) value);
    return 1;
}

void imagerenderer_freeImages(struct image_renderers *ir) {
    hashmapForEach(ir->images, callback, (void *) ir);
}
