/*
 * Heat Index sensor.
 * 
 * For the formula used see:
 * 
 * http://en.wikipedia.org/wiki/Heat_index
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

#define c1 16.923
#define c2 0.185212
#define c3 5.37941
#define c4 -0.100254
#define c5 9.41695e-3
#define c6 7.28898e-3
#define c7 3.45372e-4
#define c8 -8.14971e-4
#define c9 1.02102e-5
#define c10 -3.8646e-5
#define c11 2.91583e-5
#define c12 1.42721e-6
#define c13 1.97483e-7
#define c14 -2.18429e-8
#define c15 8.43296e-10
#define c16 -4.81975e-11

static void update(struct sensor *sensor) {
    struct state *state = (struct state *) sensor;

    // Get data from inside mutex
    pthread_mutex_lock(&state->mutex);
    double R = state->humidity;
    double T = state->temp;
    pthread_mutex_unlock(&state->mutex);

    // Formula uses Fahrenheit
    T = (T * 9.0 / 5.0) + 32.0;

    double HI = c1 + (c2 * T)+(c3 * R)+(c4 * T * R);

    double T2 = T*T;
    double R2 = R*R;
    HI += (c5 * T2)+(c6 * R2)+(c7 * T2 * R)+(c8 * T * R2)+(c9 * T2 * R2);

    double T3 = T*T2;
    double R3 = R*R2;
    HI += (c10 * T3)+(c11 * R3)+(c12 * T3 * R)+(c13 * T * R3)+(c14 * T3 * R2)+(c15 * T2 * R3)+(c16 * T3 * R3);

    // Back to celsius
    HI = (HI - 32.0)*5.0 / 9.0;

    sensor_log(sensor, (int) (HI * 10), "HInd %0.1fC", HI);
}

void register_virtual_heatindex(CONFIG_SECTION *sect) {

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
    state->sensor.title = "Heat Index";
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
