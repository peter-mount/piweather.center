/*
 * Applies a mask to an image.
 * 
 */
#include <gd.h>
#include <stdlib.h>
#include <stdint.h>

static void getPixel(gdImagePtr img, int x, int y, int *r, int *g, int *b) {
    if (x >= img->sx || y >= img->sy)
        return;

    int c = gdImageGetPixel(img, x, y);
    *r = gdImageRed(img, c);
    *g = gdImageGreen(img, c);
    *b = gdImageBlue(img, c);
}

/**
 * Applies a mask to an image.
 * 
 * The mask should be a 2 colour image. Areas that are black will be set
 * to black in the new image, otherwise the pixel from source is used.
 * 
 * @param src source image
 * @param mask Mask to apply
 * @return new image
 */
gdImagePtr imagefilter_apply_mask(gdImagePtr src, gdImagePtr mask) {
    int x, y, c, r, g, b;
    gdImagePtr img = gdImageCreateTrueColor(src->sx, src->sy);
    for (y = 0; y < src->sy; y++)
        for (x = 0; x < src->sx; x++) {
            getPixel(mask, x, y, &r, &g, &b);
            if (r || g || b)
                getPixel(src, x, y, &r, &g, &b);

            c = gdImageColorAllocate(img, r, g, b);

            gdImageSetPixel(img, x, y, c);
        }
    return img;
}
