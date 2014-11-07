/* 
 * File:   annotation.h
 * Author: peter
 *
 * Created on February 26, 2014, 5:48 PM
 */

#ifndef ANNOTATION_H
#define	ANNOTATION_H

#ifdef	__cplusplus
extern "C" {
#endif

#include <stdio.h>
#include <stdlib.h>
#include <memory.h>
#include <stdint.h>
#include "lib/bytebuffer.h"

    enum ALIGN_T {
        ALIGN_LEFT,
        ALIGN_RIGHT,
        ALIGN_CENTER
    };

    /**
     * Create an image from a bytebuffer (i.e. from the camera)
     * 
     * @param buffer
     * @return 
     */
    extern gdImagePtr jpg_image_from_bytebuffer(struct bytebuffer *buffer);

    /**
     * Create an image from a bytebuffer (i.e. from the camera)
     * 
     * @param buffer
     * @return 
     */
    extern gdImagePtr png_image_from_bytebuffer(struct bytebuffer *buffer);

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
    extern int jpg_image_to_bytebuffer(gdImagePtr image, int quality, struct bytebuffer *buffer);

    /**
     * Write a JPeg image into a bytebuffer.
     * 
     * Note the buffer should be reset before this call as this will append the image
     * 
     * @param image gdImagePtr to write
     * @param buffer bytebuffer
     * @return 0 on success, !0 on error
     */
    extern int png_image_to_bytebuffer(gdImagePtr image, struct bytebuffer *buffer);

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
    extern void draw_text(gdImagePtr im, char *font, double size, int x, int y, char align, uint32_t colour, char shadow, char *text);

    /**
     * Overlay an image
     * @param filename
     * @param image
     * @return 
     */
    extern int overlay_image(char *filename, gdImagePtr image);

    extern gdImagePtr duplicate_image(gdImagePtr src);

#ifdef	__cplusplus
}
#endif

#endif	/* ANNOTATION_H */

