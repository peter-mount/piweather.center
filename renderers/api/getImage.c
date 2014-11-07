
#include <stdlib.h>
#include "renderers/imagerenderer.h"

/**
 * Get an image from the hashmap for the specified name.e
 * 
 * @param ir image_renderers
 * @param name name of image
 * @return gdImagePtr or NULL
 */
gdImagePtr imagerenderer_getImage(struct image_renderers *ir, const char *name) {
    return (gdImagePtr) hashmapGet(ir->images, (void *) name);
}

/**
 * Get the raw image
 * @param ir image renderers
 * @return gdImagePtr or NULL
 */
gdImagePtr imagerenderer_getRawImage(struct image_renderers *ir) {
    return imagerenderer_getImage(ir, IMAGE_RAW);
}

/**
 * Get the SD (640x480) image, creating it from raw.
 * 
 * Note: This will return raw if that is also 640x480, saving memory
 * 
 * @param ir image renderers
 * @return gdImagePtr or NULL
 */
gdImagePtr imagerenderer_getSDImage(struct image_renderers *ir) {
    gdImagePtr image = imagerenderer_getImage(ir, IMAGE_SD);
    if (!image) {
        gdImagePtr raw = imagerenderer_getImage(ir, IMAGE_RAW);
        if (raw) {
            // If RAW is also SD size then just return it, don't cache it
            // otherwise we'll segfault by trying to free it twice
            if (raw->sx == IMAGE_SD_WIDTH && raw->sy == IMAGE_SD_HEIGHT)
                return raw;

            image = gdImageCreateTrueColor(IMAGE_SD_WIDTH, IMAGE_SD_HEIGHT);
            gdImageCopyResized(image, raw,
                    0, 0,
                    0, 0,
                    IMAGE_SD_WIDTH, IMAGE_SD_HEIGHT,
                    raw->sx, raw->sy);

            // Cache the result
            imagerenderer_putImage(ir, IMAGE_SD, image);
        }
    }
    return image;
}
