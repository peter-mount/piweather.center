/*
 * Generate a histogram for an image
 */
#include <gd.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include "imagefilter/filter.h"

static void set(int *h, gdImagePtr img, int mask, int off, int *channel) {
    int x, y, c;
    if (img->trueColor) {
        for (y = 0; y < img->sy; y++)
            for (x = 0; x < img->sx; x++) {
                c = (gdImageGetPixel(img, x, y) & mask) >> off;
                if (c < 0)
                    c = 0;
                else if (c > 255)
                    c = 255;
                h[c]++;
            }
    } else {
        for (y = 0; y < img->sy; y++)
            for (x = 0; x < img->sx; x++) {
                c = channel[gdImageGetPixel(img, x, y)];
                if (c < 0)
                    c = 0;
                else if (c > 255)
                    c = 255;
                h[c]++;
            }
    }
}

int *imagefilter_histogram(gdImagePtr img, IMAGE_CHANNEL_T channel) {
    int *h = (int *) calloc(256,sizeof (int));

    if (channel == IMAGE_CHANNEL_RED)
        set(h, img, 0x00ff0000, 16, img->red);
    else if (channel == IMAGE_CHANNEL_GREEN)
        set(h, img, 0x0000ff00, 8, img->green);
    else if (channel == IMAGE_CHANNEL_BLUE)
        set(h, img, 0x000000ff, 0, img->blue);

    return h;
}