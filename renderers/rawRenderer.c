/**
 * Renderer which takes the captured image and makes it available in the webserver as-is as /raw.jpg
 * 
 */
#include "weatherstation/main.h"
#include "camera/camera.h"
#include "webserver/webserver.h"
#include "renderers/imagerenderer.h"
#include "renderers/annotation.h"
#include "piweather_build.h"

struct state {
    struct image_renderer meta;
    struct MHD_Response *response;
    const char *url;
};

static void postinit(struct image_renderer *r, CAMERA camera) {

    // No camera then do nothing
    if (!camera->enabled)
        return;

    struct state *state = (struct state *) r;

    state->url = genurl((const char*) webserver.contextPath, "raw.jpg");
    webserver_add_response_handler(state->url);
}

/**
 * Renders the test card
 * 
 * @param r
 * @param camera
 */
static void renderTestCard(struct state *state, CAMERA camera) {
    // Create the initial info page, so until the camera first fires we have something to show
    gdImagePtr img = gdImageCreateTrueColor(PIWEATHER_WIDTH, PIWEATHER_HEIGHT);
    int y = PIWEATHER_TOP;

    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_CENTER, PIWEATHER_LABEL, 0, PIWEATHER_TITLE);
    y += PIWEATHER_LINE_HEIGHT;
    y += PIWEATHER_LINE_HEIGHT;

    char hostname[1024];
    gethostname(hostname, 1024);
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_RIGHT, PIWEATHER_LABEL, 0, "Station Host ");
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_LEFT, PIWEATHER_COLOUR, 0, hostname);
    y += PIWEATHER_LINE_HEIGHT;
    y += PIWEATHER_LINE_HEIGHT;

    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_RIGHT, PIWEATHER_LABEL, 0, "Version ");
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_LEFT, PIWEATHER_COLOUR, 0, PIWEATHER_VERSION);
    y += PIWEATHER_LINE_HEIGHT;
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_RIGHT, PIWEATHER_LABEL, 0, "Build Number ");
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_LEFT, PIWEATHER_COLOUR, 0, PIWEATHER_BUILD_NUMBER);
    y += PIWEATHER_LINE_HEIGHT;
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_RIGHT, PIWEATHER_LABEL, 0, "Build Host ");
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_LEFT, PIWEATHER_COLOUR, 0, PIWEATHER_BUILD_HOST);
    y += PIWEATHER_LINE_HEIGHT;
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_RIGHT, PIWEATHER_LABEL, 0, "Build Date ");
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_LEFT, PIWEATHER_COLOUR, 0, PIWEATHER_BUILD_TIME);
    y += PIWEATHER_LINE_HEIGHT;
    y += PIWEATHER_LINE_HEIGHT;
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_RIGHT, PIWEATHER_LABEL, 0, "Support ");
    draw_text(img, PIWEATHER_FONT, PIWEATHER_FONT_SIZE, PIWEATHER_LEFT, y, ALIGN_LEFT, PIWEATHER_COLOUR, 0, PIWEATHER_HOME);
    y += PIWEATHER_LINE_HEIGHT;

    imagerenderer_createJpegResponse(&camera->renderers, img, 90, state->url);
    imagerenderer_putImage(&camera->renderers, "raw", img);
    //imagerenderer_render(camera);
}

static void render(struct image_renderer *r, CAMERA camera) {

    // No camera then do nothing
    if (!camera->enabled)
        return;

    struct state *state = (struct state *) r;

    if (camera->imagedata.pos) {
        // The buffer is already formatted as a jpeg
        replaceResponseByteBuffer(state->url, &camera->imagedata, "image/jpeg");

        // Get the raw image version for the other renderers to use
        gdImagePtr image = jpg_image_from_bytebuffer(&camera->imagedata);
        imagerenderer_putImage(&camera->renderers, "raw", image);
    } else {
        // No image data so render the test card
        renderTestCard(state, camera);
    }
}

struct image_renderer *create_rawrenderer() {

    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));
    state->meta.postinit = postinit;
    state->meta.render = render;
    return (struct image_renderer *) state;
}