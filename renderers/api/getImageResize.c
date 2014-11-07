
#include <stdlib.h>
#include "imagefilter/filter.h"
#include "renderers/imagerenderer.h"

gdImagePtr imagerenderer_getImageResizeDuplicate(struct image_renderers *ir, const char *name, const char *srcName, int w, int h) {
    gdImagePtr image = imagerenderer_getImage(ir, name);
    if (!image) {
        gdImagePtr raw = imagerenderer_getImage(ir, srcName);
        if (raw) {
            image = imagefilter_createThumbnail(raw,w,h);

            // Cache the result
            imagerenderer_putImage(ir, name, image);
        }
    }

    return image;
}

gdImagePtr imagerenderer_getImageResize(struct image_renderers *ir, const char *name, int w, int h) {
    return imagerenderer_getImageResizeDuplicate(ir, name, "raw", w, h);
}
