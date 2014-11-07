/*
 * Wind Chill sensor.
 * 
 * This takes readings from an anemometer and a temperature sensor anc
 * calculates the equivalent wind chill.
 * 
 * The calculation is only valid for air temperatures below 10C and wind speeds above 4.8km/h
 * 
 * For the formula used:
 * 
 * http://en.wikipedia.org/wiki/Wind_chill#North_American_and_United_Kingdom_wind_chill_index
 * https://www.freemathhelp.com/wind-chill.html
 */

#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <pthread.h>
#include "lib/config.h"
#include "lib/string.h"
#include "sensors/sensors.h"

struct state {
    // This must be the first entry, it's what the state engine will see this struct as
    struct sensor sensor;
    double temp;
    // Unit to convert temperature raw value into celsius
    double temp_unit;
    // The anemometer
    double wind;
    // Anemometer counts per second to km/h
    double wind_factor;
    // The time (s) that the anenometer count covers
    double wind_sample;
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

// updates wind when that sensor updates

static void update_wind(struct sensor *sensor, struct sensor *listener) {
    struct state *state = (struct state *) listener;

    // Update state from inside mutex
    pthread_mutex_lock(&state->mutex);
    state->wind = (double) sensor->value * state->wind_factor / state->wind_sample;
    pthread_mutex_unlock(&state->mutex);
}

// For our sample frequency update using the received sensor values

static void update(struct sensor *sensor) {
    struct state *state = (struct state *) sensor;

    // Get data from inside mutex
    pthread_mutex_lock(&state->mutex);
    double V = state->wind;
    double T = state->temp;
    pthread_mutex_unlock(&state->mutex);

    V = pow(V, 0.16);
    double TC = 13.12 + (0.6215 * T) - (11.37 * V) + (0.3965 * T * V);

    sensor_log(sensor, (int) (TC * 10), "WChill %0.1fC", TC);
}

void register_virtual_windchill(CONFIG_SECTION *sect) {

    char *name = NULL;
    config_getCharParameter(sect, "temperature", &name);
    fatalIfNull(name, "temperature sensor name is mandatory for %s", sect->node.name);
    struct sensor *temp = sensors_get(name);
    fatalIfNull(temp, "temperature sensor \"%s\" is not present for %s", name, sect->node.name);

    double tempUnit = 1.0;
    config_getDoubleParameter(sect, "temperature-unit", &tempUnit);

    name = NULL;
    config_getCharParameter(sect, "wind", &name);
    fatalIfNull(name, "wind sensor name is mandatory for %s", sect->node.name);
    struct sensor *wind = sensors_get(name);
    fatalIfNull(wind, "wind sensor \"%s\" is not present for %s", name, sect->node.name);

    // 2.4 is default based on specs of the Maplin/SparkFun anemometers
    double wind_factor = 2.4;
    // Default is 60 seconds
    double wind_sample = 60;
    config_getDoubleParameter(sect, "wind-factor", &wind_factor);
    config_getDoubleParameter(sect, "wind-sample", &wind_sample);

    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    state->sensor.name = sect->node.name;
    state->sensor.title = "Wind Chill";
    state->sensor.update = update;
    state->temp_unit = tempUnit;
    state->wind_factor = wind_factor;
    state->wind_sample = wind_sample;
    sensor_configure(sect, &state->sensor);

    pthread_mutex_init(&state->mutex, NULL);

    // Listen to those sensors
    sensor_add_listener(temp, &state->sensor, update_temp);
    sensor_add_listener(wind, &state->sensor, update_wind);

    // Finally register it
    sensor_register(&state->sensor);
}
