/* 
 * File:   camera.h
 * Author: peter
 *
 * Created on February 6, 2014, 2:28 PM
 */

#ifndef CAMERA_H
#define	CAMERA_H

#include <pthread.h>
#include "lib/bytebuffer.h"
#include "lib/config.h"

typedef struct camera_config * CAMERA;

#include "renderers/imagerenderer.h"

/// Frame advance method
#define FRAME_NEXT_SINGLE        0
#define FRAME_NEXT_TIMELAPSE     1
#define FRAME_NEXT_KEYPRESS      2
#define FRAME_NEXT_FOREVER       3
#define FRAME_NEXT_GPIO          4
#define FRAME_NEXT_SIGNAL        5
#define FRAME_NEXT_IMMEDIATELY   6

#define MAX_USER_EXIF_TAGS      32
#define MAX_EXIF_PAYLOAD_LENGTH 128

/**
 * Camera configuration for capture and the raw image
 */
struct camera_config {
    // Must be first
    struct Node node;
    // Is the camera enabled
    int enabled;
    // Delay between each picture
    int timelapse;
    // Width of the raw image
    int width;
    // height of the raw image
    int height;
    // JPEG quality (1-100)
    int quality;
    // Flag for whether the JPEG metadata also contains the RAW bayer image
    int wantRAW;
    // Renderers to use with this camera
    struct image_renderers renderers;
    // =================
    // Internal use only
    // =================
    // Registry entry for this camera
    struct camera_registry *registry;
    // buffer used for capturing image data. Renderers can use this as they
    // run in the camera thread
    struct bytebuffer imagedata;
    // Command to take a picture
    char *cmd;
    // Path to captured image
    char *image;
    // Number of frames to drop
    int dropFrameCount;
};

struct List cameras;

// Internal structure to handle the camera types

struct camera_registry {
    // value of camera-type parameter
    const char *type;
    // initialise a new camera instance
    CAMERA(*init)(CONFIG_SECTION *sect);
    // Start the camera
    void (*start)(CAMERA camera);
    // Stop the camera
    void (*stop)(CAMERA camera);
    // Capture a frame
    void (*capture)(CAMERA camera);
};


extern void camera_init();
extern void camera_loop();
extern void camera_start();
extern void camera_stop();

// Default loop implementation
extern void *camera_timedloop(void *arg);

#endif	/* CAMERA_H */

