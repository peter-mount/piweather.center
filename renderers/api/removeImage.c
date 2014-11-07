

#include <stdlib.h>
#include "renderers/imagerenderer.h"

void imagerenderer_removeImage(struct image_renderers *ir, const char *name) {
    gdImagePtr existing = (gdImagePtr) hashmapRemove(ir->images, (void *) name);
    if (existing)
        gdImageDestroy(existing);
}
