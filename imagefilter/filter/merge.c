/*
 * Merges two images by taking the maximum value of an image
 */
#include <gd.h>
#include <stdlib.h>
#include <stdint.h>

static int max(int a, int b) {
    if (a > b)
        return a;
    else
        return b;
}

gdImagePtr imagefilter_merge(gdImagePtr img1, gdImagePtr img2) {
    int x, y, c1, c2, r, g, b, np;
    int sx = img1->sx < img2->sx ? img1->sx : img2->sx;
    int sy = img1->sy < img2->sy ? img1->sy : img2->sy;

    gdImagePtr img = gdImageCreateTrueColor(sx, sy);

    for (y = 0; y < sy; y++)
        for (x = 0; x < sx; x++) {
            c1 = gdImageGetPixel(img1, x, y);
            c2 = gdImageGetPixel(img2, x, y);

            r = max(gdImageRed(img1, c1), gdImageRed(img2, c2));
            g = max(gdImageGreen(img1, c1), gdImageGreen(img2, c2));
            b = max(gdImageBlue(img1, c1), gdImageBlue(img2, c2));

            np = gdImageColorAllocate(img, r, g, b);

            gdImageSetPixel(img, x, y, np);
        }

    return img;
}

