/*
 * Cloud Base Sensor.
 * 
 * This is not a real sensor but calculated based on Temperature, Dew Point
 * and station altitude.
 * 
 * http://www.csgnetwork.com/estcloudbasecalc.html
 */

#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <pthread.h>
#include "lib/config.h"
#include "lib/string.h"
#include "sensors/sensors.h"
#include "astro/location.h"
#include "astro/observatory.h"

struct state {
    // This must be the first entry, it's what the state engine will see this struct as
    struct sensor sensor;
    double temp;
    // Unit to convert temperature raw value into celsius
    double temp_unit;
    // The dewpoint sensor - no unit as we know it's 10
    double dewpoint;
    // mutex as updates can happen on different threads
    pthread_mutex_t mutex;
};

// updates temp when that sensor updates

static void update_temp(struct sensor *sensor, struct sensor *listener) {
    struct state *state = (struct state *) listener;

    // Update state from inside mutex
    pthread_mutex_lock(&state->mutex);
    state->temp = (double) sensor->value * state->temp_unit;
    pthread_mutex_unlock(&state->mutex);
}

// updates dewpoint when that sensor updates

static void update_dewpoint(struct sensor *sensor, struct sensor *listener) {
    struct state *state = (struct state *) listener;

    // Update state from inside mutex
    pthread_mutex_lock(&state->mutex);
    state->dewpoint = (double) sensor->value / 10.0;
    pthread_mutex_unlock(&state->mutex);
}

// For our sample frequency update using the received sensor values

static void update(struct sensor *sensor) {
    struct state *state = (struct state *) sensor;

    // Get data from inside mutex
    pthread_mutex_lock(&state->mutex);
    double DP = state->dewpoint;
    double T = state->temp;
    pthread_mutex_unlock(&state->mutex);

    // This formula is for Fahrenheit
    //int alt = (int)ceil((((T-DP)/4.5)*1000)+state->altitude);

    // Correct one for Celsius - http://en.wikipedia.org/wiki/Cloud_base
    int alt = (int) ceil((((T - DP) / 2.5)*1000)+(observatory.altitude / 3.2808399));

    sensor_log(sensor, alt, "CB %dft", alt);
}

void register_virtual_cloudbase(CONFIG_SECTION *sect) {

    char *name = NULL;
    config_getCharParameter(sect, "temperature", &name);
    fatalIfNull(name, "temperature sensor name is mandatory for %s", sect->node.name);
    struct sensor *temp = sensors_get(name);
    fatalIfNull(temp, "temperature sensor \"%s\" is not present for %s", name, sect->node.name);

    double tempUnit = 1.0;
    config_getDoubleParameter(sect, "temperature-unit", &tempUnit);

    name = NULL;
    config_getCharParameter(sect, "dewpoint", &name);
    fatalIfNull(name, "dewpoint sensor name is mandatory for %s", sect->node.name);
    struct sensor *humidity = sensors_get(name);
    fatalIfNull(humidity, "dewpoint sensor \"%s\" is not present for %s", name, sect->node.name);

    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    state->sensor.name = sect->node.name;
    state->sensor.title = "Cloud Base";
    state->sensor.update = update;
    state->temp_unit = tempUnit;
    sensor_configure(sect, &state->sensor);

    pthread_mutex_init(&state->mutex, NULL);

    // Listen to those sensors
    sensor_add_listener(temp, &state->sensor, update_temp);
    sensor_add_listener(humidity, &state->sensor, update_dewpoint);

    // Finally register it
    sensor_register(&state->sensor);
}
