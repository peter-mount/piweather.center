/**
 * Renderer which takes the captured image and makes it available in the webserver as-is as /raw.jpg
 * 
 */
#include "weatherstation/main.h"
#include "camera/camera.h"
#include "webserver/webserver.h"
#include "lib/charbuffer.h"
#include "sensors/sensors.h"
#include <stdio.h>
#include <stdlib.h>
#include <gd.h>
#include "annotation.h"
#include "lib/config.h"
#include "renderers/imagerenderer.h"
#include "renderers/image/image.h"
#include "piweather_build.h"

#define CAMERA_TYPE_COUNT 4

static const char *camera_types[] = {
    "", "standard", "noir", "biology"
};
static const char *camera_abbrev[] = {
    "", "STD", "NOIR", "BIO"
};
static const char *months[] = {
    "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
};

struct state {
    // This must be first
    struct image_renderer meta;
    // Enable this image? implicitly set if any option is defined
    int enabled;
    // Camera title, goes on top of image
    const char *title;
    const char *site;
    // Font & colours
    uint32_t bg_colour;
    uint32_t bl_colour;
    uint32_t fg_colour;
    char *font;
    int fontsize;
    int shadow;
    // Primary & Secondary image dimensions
    struct image primary;
    struct image secondary;
    struct image tertiary;
    // =================
    // Internal use only
    // =================
    time_t time;
    // Buffer to hold the image date
    char date[32];
    // The camera type string
    const char *camera_type;
    // used to build the sensor text
    struct charbuffer sensors;
};

static void init(struct image_renderer *r, CAMERA camera) {
    struct state *state = (struct state *) r;

    CONFIG_SECTION *sect = config_getSection("annotation");
    if (!sect)
        return;

    state->enabled = 1;

    config_getCharParameter(sect, "title", (char **) &state->title);
    config_getCharParameter(sect, "site", (char **) &state->site);

    config_getHexParameter(sect, "background", &state->bg_colour);
    if (!state->bg_colour)
        state->bg_colour = 0x40263A93;

    config_getHexParameter(sect, "line-colour", &state->bl_colour);
    if (!state->bl_colour)
        state->bl_colour = 0x00FF0000;

    config_getHexParameter(sect, "foreground", &state->fg_colour);
    if (!state->fg_colour)
        state->fg_colour = 0x00FFFFFF;

    // Hopefully this default exists ;-)
    config_getCharParameter(sect, "font", &state->font);
    if (!state->font)
        state->font = "/usr/share/fonts/truetype/droid/DroidSans.ttf";

    config_getIntParameter(sect, "font-size", &state->fontsize);
    if (!state->fontsize)
        state->fontsize = 10;

    config_getBooleanParameter(sect, "font-shadow", &state->shadow);

    char *s = NULL;
    config_getCharParameter(sect, "camera-type", &s);
    if (s) {
        int i = 1;
        int used = 0;
        for (; i < CAMERA_TYPE_COUNT && !used; i++)
            if (strcmp(camera_types[i], s) == 0) {
                state->camera_type = camera_abbrev[i];
                used = 2;
            }
        if (!used) {
            fprintf(stderr, "Unsupported camera type: %s\n", s);
            exit(1);
        }
    }

    init_image(sect, &state->primary, "primary");
    init_image(sect, &state->secondary, "secondary");
    init_image(sect, &state->tertiary, "tertiary");
}

static void postinit(struct image_renderer *r, CAMERA camera) {

    // No camera then do nothing
    if (!camera->enabled)
        return;

    struct state *state = (struct state *) r;

    // Only run if configured
    if (!state->enabled)
        return;

    // Default annotated image to the raw image size
    if (!state->primary.width || !state->primary.height) {
        state->primary.width = camera->width;
        state->primary.height = camera->height;
    }

    state->primary.url = genurl((const char*) webserver.contextPath, "image.jpg");
    webserver_add_response_handler(state->primary.url);

    if (state->secondary.enabled) {
        state->secondary.url = genurl((const char*) webserver.contextPath, "image2.jpg");
        webserver_add_response_handler(state->secondary.url);
    }

    if (state->tertiary.enabled) {
        state->tertiary.url = genurl((const char*) webserver.contextPath, "image3.jpg");
        webserver_add_response_handler(state->tertiary.url);
    }
}

static void setDate(struct state *state) {
    struct tm timeinfo;

    time(&state->time);
    localtime_r(&state->time, &timeinfo);
    snprintf(state->date, sizeof (state->date),
            "%04d %3s %02d %02d:%02d %s",
            timeinfo.tm_year + 1900,
            months[timeinfo.tm_mon],
            timeinfo.tm_mday,
            timeinfo.tm_hour,
            timeinfo.tm_min,
            timeinfo.tm_zone);
}

static void annotate(struct image_renderers *ir, struct state *state, struct image *img) {
    gdImagePtr image = imagerenderer_getImageResize(ir, img->prefix, img->width, img->height);

    int w = image->sx;

    char tmp[256];
    int spacing = 4;
    int fontsize = state->fontsize;
    int height = fontsize + (spacing * 2);

    // subtitle
    //if(subtitle || info )
    //    height += fontsize*0.8+spacing;

    // Look at the sensors
    charbuffer_reset(&state->sensors);
    if (sensors) {
        int first = 1;
        struct sensor *s = sensors->sensors;
        while (s) {
            if (s->enabled & s->annotate) {
                if (first)
                    first = 0;
                else
                    charbuffer_append(&state->sensors, " ");
                charbuffer_append(&state->sensors, s->text);
            }
            s = s->next;
        }
    }

    // If present then the camera type/id under the date
    if (img->ident)
        snprintf(tmp, sizeof (tmp), "%s %s %s", state->camera_type, img->ident, PIWEATHER_VERSION);
    else
        snprintf(tmp, sizeof (tmp), "%s %s", state->camera_type, PIWEATHER_VERSION);

    // If we have sensors or type info then add room in the banner
    if (state->sensors.pos || *tmp)
        height += fontsize * 0.8 + spacing;

    // Top or bottom of image
    int top = 0;
    // top=h-height;

    // draw banner line
    gdImageFilledRectangle(image, 0, height + 1, w, height + 2, state->bl_colour);
    //gdImageFilledRectangle( image, 0, top-2, w, top-1, bl_colour );

    // Background box
    gdImageFilledRectangle(image, 0, top, w, top + height, state->bg_colour);

    int y = top + spacing + fontsize;

    // Site on top left
    draw_text(image, state->font, fontsize, spacing, y, ALIGN_LEFT, state->fg_colour, state->shadow, (char *) state->site);

    // title in middle
    draw_text(image, state->font, fontsize, w >> 1, y, ALIGN_CENTER, state->fg_colour, state->shadow, (char *) state->title);

    // Timestamp top right
    draw_text(image, state->font, fontsize, w - spacing, y, ALIGN_RIGHT, state->fg_colour, state->shadow, state->date);

    // Move to the second line
    y += spacing + fontsize * 0.8;

    // If present add the camera type under the date
    if (*tmp)
        draw_text(image, state->font, fontsize * 0.7, w - spacing, y, ALIGN_RIGHT, state->fg_colour, state->shadow, tmp);

    // Only include sensors if we have anything to annotate
    if (state->sensors.pos) {
        draw_text(image, state->font, fontsize * 0.7, spacing, y, ALIGN_LEFT, state->fg_colour, state->shadow, state->sensors.buffer);
    }

    // Generate the final image response
    imagerenderer_createJpegResponse(ir, image, 90, img->url);

    if (img->path)
        write_image(image, &state->time, img);
}

static void render(struct image_renderer *r, CAMERA camera) {

    // No camera then do nothing
    if (!camera->enabled)
        return;

    struct state *state = (struct state *) r;

    // Only run if configured
    if (!state->enabled)
        return;

    // Set the date string
    setDate(state);

    annotate(&camera->renderers, state, &state->primary);

    if (state->secondary.enabled)
        annotate(&camera->renderers, state, &state->secondary);

    if (state->tertiary.enabled)
        annotate(&camera->renderers, state, &state->tertiary);
}

struct image_renderer *create_annotatedrenderer() {

    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    // Register the functions
    state->meta.init = init;
    state->meta.postinit = postinit;
    state->meta.render = render;

    // TODO Reinstate later, for now remove (null) from image.
    state->camera_type = "";

    return (struct image_renderer *) state;
}
