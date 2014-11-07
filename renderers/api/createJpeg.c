#include <stdlib.h>
#include "renderers/imagerenderer.h"

void imagerenderer_createJpeg(struct image_renderers *ir, gdImagePtr image, int quality) {
    bytebuffer_reset(&ir->buffer);
    jpg_image_to_bytebuffer(image, quality, &ir->buffer);
}

void imagerenderer_createJpegResponse(struct image_renderers *ir, gdImagePtr image, int quality, const char *url) {
    imagerenderer_createJpeg(ir, image, quality);
    replaceResponseByteBuffer(url, &ir->buffer, "image/jpeg");
}
