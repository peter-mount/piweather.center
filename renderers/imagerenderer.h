/* 
 * File:   imagerenderer.h
 * Author: peter
 *
 * Created on March 4, 2014, 5:11 PM
 */

#ifndef CAMERA_H
#include "camera/camera.h"
#else

#ifndef IMAGERENDERER_H
#define	IMAGERENDERER_H

#include <gd.h>
#include <microhttpd.h> 
#include "lib/bytebuffer.h"
#include "lib/list.h"
#include "lib/hashmap.h"

// Raw image name
#define IMAGE_RAW "raw"

// SD image name & dimensions
#define IMAGE_SD "sd"
#define IMAGE_SD_WIDTH 640
#define IMAGE_SD_HEIGHT 480

// Thumbnail size
#define IMAGE_THUMB_WIDTH 100
#define IMAGE_THUMB_HEIGHT 75

struct image_renderers {
    // List of renderers to run
    struct List renderers;
    // Map used to share data between renderers
    Hashmap *images;
    // Shared bytebuffer used in generating .jpg images
    struct bytebuffer buffer;
};

/**
 * A definition of an image renderer.
 * 
 * Note: a hook is called if it's not NULL, NULL means no operation
 */
struct image_renderer {
    struct Node node;
    // initialise the renderer
    void (*init)(struct image_renderer *r, CAMERA c);
    // postinitialise the renderer
    void (*postinit)(struct image_renderer *r, CAMERA c);
    // start the renderer
    void (*start)(struct image_renderer *r, CAMERA c);
    // Render the image
    void (*render)(struct image_renderer *r, CAMERA c);
    // shutdown the renderer
    void (*stop)(struct image_renderer *r, CAMERA c);
};

extern void imagerenderer_initialise();
extern void imagerenderer_register(CAMERA camera, struct image_renderer *renderer);
extern void imagerenderer_init(CAMERA camera);
extern void imagerenderer_postinit(CAMERA camera);
extern void imagerenderer_render(CAMERA camera);
extern void imagerenderer_start(CAMERA camera);
extern void imagerenderer_stop(CAMERA camera);

extern gdImagePtr imagerenderer_createImage(struct image_renderers *ir, const char *name, int w, int h);
extern void imagerenderer_freeImages(struct image_renderers *ir);
extern gdImagePtr imagerenderer_getImage(struct image_renderers *ir, const char *name);
extern gdImagePtr imagerenderer_getImageResize(struct image_renderers *ir, const char *name, int w, int h);
extern gdImagePtr imagerenderer_getImageResizeDuplicate(struct image_renderers *ir, const char *name, const char *srcName, int w, int h);
extern void imagerenderer_putImage(struct image_renderers *ir, const char *name, gdImagePtr img);
extern void imagerenderer_removeImage(struct image_renderers *ir, const char *name);

extern gdImagePtr imagerenderer_getRawImage(struct image_renderers *ir);
extern gdImagePtr imagerenderer_getSDImage(struct image_renderers *ir);

extern void imagerenderer_createJpeg(struct image_renderers *ir, gdImagePtr image, int quality);
extern void imagerenderer_createJpegResponse(struct image_renderers *ir, gdImagePtr image, int quality, const char *url);

#endif	/* IMAGERENDERER_H */

#endif 