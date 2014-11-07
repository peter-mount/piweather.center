

#include <stdlib.h>
#include "renderers/imagerenderer.h"

/**
 * Get an image from the hashmap for the specified size.
 * If it's not present then take it from the raw image
 * 
 * @param ir image_renderers
 * @param name name of image
 * @param w width
 * @param h height
 * @return 
 */
gdImagePtr imagerenderer_createImage(struct image_renderers *ir, const char *name, int w, int h) {
    gdImagePtr image = gdImageCreateTrueColor(w, h);
    imagerenderer_putImage(ir, name, image);
    return image;
}

