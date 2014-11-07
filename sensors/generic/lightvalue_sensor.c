/**
 * lightvalue sensor used to monitor the light in the captured image
 * 
 * This was originally part of the auto exposure mode but it's value doesn't do anything.
 * It's kept as an example of using EXIF
 */

#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <math.h>
#include "main.h"
#include "commands.h"
#include "sensors/sensors.h"
#include "bytebuffer.h"
#include "RaspiCamControl.h"
#include <libexif/exif-loader.h>
#include <libexif/exif-utils.h>
#include "camera.h"

struct state {
    struct sensor sensor;
    // Shutter speed, i.e. 1/4
    double shutter_speed;
    // aperture, e.g. 2.9
    double aperture;
    // iso
    double iso;
    // light level
    double light_level;
    // Internals
    ExifLoader *loader;
};

static int getDouble(ExifData *ed, ExifByteOrder bo, ExifTag t, double *d) {
    ExifEntry * e = exif_data_get_entry(ed, t);
    if (!e)
        return 0;

    char *b = e->data;

    switch (e->format) {
        case EXIF_FORMAT_SHORT:
            *d = (double) exif_get_short(b, bo);
            return 1;

        case EXIF_FORMAT_SSHORT:
            *d = (double) exif_get_sshort(b, bo);
            return 1;

        case EXIF_FORMAT_LONG:
            *d = (double) exif_get_long(b, bo);
            return 1;

        case EXIF_FORMAT_SLONG:
            *d = (double) exif_get_slong(b, bo);
            return 1;

        case EXIF_FORMAT_RATIONAL:
        {
            ExifRational r = exif_get_rational(b, bo);
            *d = (double) r.numerator / (double) r.denominator;
            return 1;
        }

        case EXIF_FORMAT_SRATIONAL:
        {
            ExifSRational r = exif_get_srational(b, bo);
            *d = (double) r.numerator / (double) r.denominator;
            return 1;
        }

        default:
            return 0;
    }
}

static int get_shutter_speed(ExifData *d, ExifByteOrder bo, double *v) {
    // BulbDuration >0 then take that

    if (getDouble(d, bo, EXIF_TAG_EXPOSURE_TIME, v))
        return 1;

    if (getDouble(d, bo, EXIF_TAG_SHUTTER_SPEED_VALUE, v))
        return 1;

    return 0;
}

static int get_aperture(ExifData *d, ExifByteOrder bo, double *v) {
    if (getDouble(d, bo, EXIF_TAG_FNUMBER, v))
        return 1;

    if (getDouble(d, bo, EXIF_TAG_APERTURE_VALUE, v))
        return 1;

    return 0;
}

/**
 * Updates the sensor by reading the current value into it
 */
static void update(CAMERA_STATE *camera, struct sensor *sensor) {
    if (!sensor->enabled)
        return;

    struct state *state = (struct state *) sensor;

    if (!state->loader)
        state->loader = exif_loader_new();

    exif_loader_reset(state->loader);
    exif_loader_write(state->loader, camera->imagedata.buffer, camera->imagedata.pos);
    ExifData *data = exif_loader_get_data(state->loader);
    ExifByteOrder bo = exif_data_get_byte_order(data);

    int valid = get_shutter_speed(data, bo, &state->shutter_speed);

    if (valid)
        valid = get_aperture(data, bo, &state->aperture);

    if (valid)
        valid = getDouble(data, bo, EXIF_TAG_ISO_SPEED_RATINGS, &state->iso);

    exif_data_free(data);

    if (valid) {
        // Calculate the light level
        state->light_level = (2.0 * log(state->aperture) - log(state->shutter_speed) - log(state->iso / 100.0)) / log(2.0);

        // Update the sensor
        sensor_log(camera, &state->sensor, (int) (state->light_level * 1000), "Lvl %.1f", state->light_level);
    }
}

/**
 * Used to create the sensor struct defining this sensor
 * @return 
 */
void register_lightvalue_sensor(CAMERA_STATE *camera) {
    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    // The sensor name, used in command prefixes & rest url's
    state->sensor.name = "lightvalue";

    // The title is used in debug & help
    state->sensor.title = "Image Light Level";

    // update is the only mandatory hook
    state->sensor.update = update;

    _sensor_register(camera, (struct sensor *) state, SensorCommandEnable + SensorCommandAnnotate);
}