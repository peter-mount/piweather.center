/*
 * Dewpoint sensor.
 * 
 * Now this is not a real sensor as dewpoint is calculated from the
 * temperature and relative humidity so what we do here is define a
 * sensor which links to existing temperature and humidity sensors
 * within the config.
 * 
 * Then, when those sensors update then so does this one.
 * 
 * For the formula used see:
 * 
 * http://andrew.rsmas.miami.edu/bmcnoldy/Humidity.html
 *
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
    // The humidity sensor
    double humidity;
    // Unit to convert humidities raw value into percentage
    double humidity_unit;
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

// updates humidity when that sensor updates

static void update_humidity(struct sensor *sensor, struct sensor *listener) {
    struct state *state = (struct state *) listener;

    // Update state from inside mutex
    pthread_mutex_lock(&state->mutex);
    state->humidity = (double) sensor->value * state->humidity_unit;
    pthread_mutex_unlock(&state->mutex);
}

// For our sample frequency update using the received sensor values

static void update(struct sensor *sensor) {
    struct state *state = (struct state *) sensor;

    // Get data from inside mutex
    pthread_mutex_lock(&state->mutex);
    double RH = state->humidity;
    double T = state->temp;
    pthread_mutex_unlock(&state->mutex);

    double TD = 243.04 * (log(RH / 100)+((17.625 * T) / (243.04 + T))) / (17.625 - log(RH / 100)-((17.625 * T) / (243.04 + T)));

    sensor_log(sensor, (int) (TD * 10), "DP %0.1fC", TD);
}

void register_virtual_dewpoint(CONFIG_SECTION *sect) {

    char *name = NULL;
    config_getCharParameter(sect, "temperature", &name);
    fatalIfNull(name, "temperature sensor name is mandatory for %s", sect->node.name);
    struct sensor *temp = sensors_get(name);
    fatalIfNull(temp, "temperature sensor \"%s\" is not present for %s", name, sect->node.name);

    double tempUnit = 1.0;
    config_getDoubleParameter(sect, "temperature-unit", &tempUnit);

    name = NULL;
    config_getCharParameter(sect, "humidity", &name);
    fatalIfNull(name, "humidity sensor name is mandatory for %s", sect->node.name);
    struct sensor *humidity = sensors_get(name);
    fatalIfNull(humidity, "humidity sensor \"%s\" is not present for %s", name, sect->node.name);

    double humidityUnit = 1.0;
    config_getDoubleParameter(sect, "humidity-unit", &humidityUnit);

    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    state->sensor.name = sect->node.name;
    state->sensor.title = "Dewpoint";
    state->sensor.update = update;
    state->temp_unit = tempUnit;
    state->humidity_unit = humidityUnit;
    sensor_configure(sect, &state->sensor);

    pthread_mutex_init(&state->mutex, NULL);

    // Listen to those sensors
    sensor_add_listener(temp, &state->sensor, update_temp);
    sensor_add_listener(humidity, &state->sensor, update_humidity);

    // Finally register it
    sensor_register(&state->sensor);
}
