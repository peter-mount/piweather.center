/* 
 * File:   image.h
 * Author: Peter T Mount
 *
 * Created on March 28, 2014, 10:36 AM
 */

#ifndef IMAGE_H
#define	IMAGE_H

#include <time.h>
#include <gd.h>

struct camera_type_t {
    char *mode;
    int id;
    char *abbrev;
};

struct image {
    // Dimension of image
    int width;
    int height;
    // Image enabled
    int enabled;
    // Ident
    char *ident;
    // Prefix
    const char *prefix;
    // The request holding the current image
    const char *url;
    // Path, if present to record the image
    char *path;
};

extern void init_image(CONFIG_SECTION *sect, struct image *image, const char *prefix);
extern void write_image(gdImagePtr img, time_t *date, struct image *image);

#endif	/* IMAGE_H */

