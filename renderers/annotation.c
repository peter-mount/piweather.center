/**
 * Common code for the manipulation of images, specifically annotating them
 */
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <memory.h>
#include <gd.h>
#include "lib/bytebuffer.h"
#include "annotation.h"

/**
 * Create an image from a bytebuffer (i.e. from the camera)
 * 
 * @param buffer
 * @return 
 */
gdImagePtr jpg_image_from_bytebuffer(struct bytebuffer *buffer) {
    if (0 != pthread_mutex_lock(&buffer->mutex)) {
        return NULL;
    }

    gdImagePtr img = gdImageCreateFromJpegPtr(buffer->pos, buffer->buffer);
    pthread_mutex_unlock(&buffer->mutex);
    return img;
}

/**
 * Create an image from a bytebuffer (i.e. from the camera)
 * 
 * @param buffer
 * @return 
 */
gdImagePtr png_image_from_bytebuffer(struct bytebuffer *buffer) {
    if (0 != pthread_mutex_lock(&buffer->mutex)) {
        return NULL;
    }

    gdImagePtr img = gdImageCreateFromPngPtr(buffer->pos, buffer->buffer);
    pthread_mutex_unlock(&buffer->mutex);
    return img;
}

/**
 * Write a JPeg image into a bytebuffer.
 * 
 * Note the buffer should be reset before this call as this will append the image
 * 
 * @param image gdImagePtr to write
 * @param quality JPeg quality
 * @param buffer bytebuffer
 * @return 0 on success, !0 on error
 */
int jpg_image_to_bytebuffer(gdImagePtr image, int quality, struct bytebuffer *buffer) {
    int size;
    void *data = (void *) gdImageJpegPtr(image, &size, quality);
    if (data) {
        bytebuffer_put(buffer, data, size);
        gdFree(data);
        return 0;
    }
    return 1;
}

/**
 * Write a JPeg image into a bytebuffer.
 * 
 * Note the buffer should be reset before this call as this will append the image
 * 
 * @param image gdImagePtr to write
 * @param buffer bytebuffer
 * @return 0 on success, !0 on error
 */
int png_image_to_bytebuffer(gdImagePtr image, struct bytebuffer *buffer) {
    int size;
    void *data = (void *) gdImagePngPtr(image, &size);
    if (data) {
        bytebuffer_put(buffer, data, size);
        gdFree(data);
        return 0;
    }
    return 1;
}

/**
 * Draw text on an image
 * @param im
 * @param font
 * @param size
 * @param x
 * @param y
 * @param align
 * @param colour
 * @param shadow
 * @param text
 */
void draw_text(gdImagePtr im, char *font, double size, int x, int y, char align, uint32_t colour, char shadow, char *text) {
    int brect[8];
    char *err;

    if (!text) return;

    if (shadow) {
        uint32_t scolour = colour & 0xFF000000;
        draw_text(im, font, size, x + 1, y + 1, align, scolour, 0, text);
    }

    /* Correct alpha value for GD. */
    colour = (((colour & 0xFF000000) / 2) & 0xFF000000) + (colour & 0xFFFFFF);

    // Pre-render the text to get the text dimensions
    err = gdImageStringFT(NULL, &brect[0], colour, font, size, 0.0, 0, 0, text);
    if (err) {
        fprintf(stderr, "%s", err);
        return;
    }

    switch (align) {
        case ALIGN_CENTER: x -= brect[4] / 2;
            break;
        case ALIGN_RIGHT: x -= brect[4];
            break;
    }

    gdImageStringFT(im, NULL, colour, font, size, 0.0, x, y, text);
}

/**
 * Overlay an image
 * @param filename
 * @param image
 * @return 
 */
int overlay_image(char *filename, gdImagePtr image) {
    FILE *f;
    gdImagePtr overlay;

    if (!filename) return (-1);

    f = fopen(filename, "rb");
    if (!f) {
        fprintf(stderr, "Unable to open '%s'", filename);
        return (-1);
    }

    overlay = gdImageCreateFromPng(f);
    fclose(f);

    if (!overlay) {
        fprintf(stderr, "Unable to read '%s'. Not a PNG image?", filename);
        return (-1);
    }

    gdImageCopy(image, overlay, 0, 0, 0, 0, overlay->sx, overlay->sy);
    gdImageDestroy(overlay);

    return (0);
}

gdImagePtr duplicate_image(gdImagePtr src) {
    gdImagePtr dst;

    dst = gdImageCreateTrueColor(gdImageSX(src), gdImageSY(src));
    if (!dst) return (NULL);

    gdImageCopy(dst, src, 0, 0, 0, 0, gdImageSX(src), gdImageSY(src));
    return (dst);
}

