
#include <stdlib.h>
#include "renderers/imagerenderer.h"

/**
 * Get an image from the hashmap for the specified name.e
 * 
 * @param ir image_renderers
 * @param name name of image
 * @return gdImagePtr or NULL
 */
void imagerenderer_putImage(struct image_renderers *ir, const char *name, gdImagePtr img) {
    gdImagePtr existing = (gdImagePtr) hashmapPut(ir->images, (void *) name, (void *) img);
    if (existing)
        gdImageDestroy(existing);
}

