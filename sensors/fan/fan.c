/*
 * This virtual sensor acts as a sensor which listens to a temperature sensor
 * and if the temperature goes above a value it turns on a fan.
 * 
 * If the temperature drops below another value the fan is turned off.
 * 
 * This sensor does not do the actual switching - thats going to be specific
 * to an individual installation, but it handles the underlying logic.
 */

#include <stdlib.h>
#include <pthread.h>
#include <string.h>
#include "lib/config.h"
#include "sensors/sensors.h"
#include "lib/string.h"

/* Registry of supported fans - add your own here*/
struct fan_registry {
    char *name;
    void *(*init)(CONFIG_SECTION *sect);
    void (*handler)(void *arg, int state);
};

// PI Weather Board High Power Outputs
extern void *fan_piweather_init(CONFIG_SECTION *sect);
extern void fan_piweather_control(void *arg, int state);

struct fan_registry registry[] = {

    // PI Weather Board High Power Outputs
    {"piweather", fan_piweather_init, fan_piweather_control},

    // Terminate list
    {NULL, NULL}
};

struct temp {
    // 1 if this is enabled
    int enabled;
    // The last recorded temperature
    double temp;
    // Unit to convert raw temperature into celsius
    double temp_unit;
    // Temperature to turn fan off
    double off_temp;
    // Temperature to turn fan on
    double on_temp;
    // The sensor this is read from
    struct sensor *sensor;
};

struct state {
    struct sensor sensor;
    // Temperature, up to 3 sensors to watch
    struct temp temp1;
    struct temp temp2;
    struct temp temp3;
    // The fan state, 0=off, 1=on
    int fan_state;
    // Hook to code to control the fan
    void (*fan_control)(void *arg, int state);
    // Data for the fan
    void *fan_data;
    // mutex as updates happen on different threads
    pthread_mutex_t mutex;
};

static void update_temp1(struct sensor *sensor, struct sensor *listener) {
    struct state *state = (struct state *) listener;

    // Update state from inside mutex
    pthread_mutex_lock(&state->mutex);
    state->temp1.temp = (double) sensor->value * state->temp1.temp_unit;
    pthread_mutex_unlock(&state->mutex);
}

static void update_temp2(struct sensor *sensor, struct sensor *listener) {
    struct state *state = (struct state *) listener;

    // Update state from inside mutex
    pthread_mutex_lock(&state->mutex);
    state->temp2.temp = (double) sensor->value * state->temp2.temp_unit;
    pthread_mutex_unlock(&state->mutex);
}

static void update_temp3(struct sensor *sensor, struct sensor *listener) {
    struct state *state = (struct state *) listener;

    // Update state from inside mutex
    pthread_mutex_lock(&state->mutex);
    state->temp3.temp = (double) sensor->value * state->temp3.temp_unit;
    pthread_mutex_unlock(&state->mutex);
}

static int checkTemp(struct temp *t, int current) {
    if (!t->enabled)
        return current;

    if (t->temp <= t->off_temp)
        return 0;

    if (t->temp >= t->on_temp)
        return 1;

    return current;
}

static void update(struct sensor *sensor) {
    struct state *state = (struct state *) sensor;

    // Get data from inside mutex
    pthread_mutex_lock(&state->mutex);

    // Get current state
    int fan_state_current = state->fan_state;

    // Work out new state
    int fan_state = checkTemp(&state->temp1, fan_state_current);
    fan_state = checkTemp(&state->temp2, fan_state);
    fan_state = checkTemp(&state->temp3, fan_state);

    state->fan_state = fan_state;

    // We are now done with the shared data
    pthread_mutex_unlock(&state->mutex);

    // Log the state
    sensor_log(sensor, fan_state, "Fan %s", fan_state ? "on" : "off");

    // Now change the fan if state has changed
    if (fan_state != fan_state_current && state->fan_control)
        state->fan_control(state->fan_data, fan_state);
}

struct fan_registry *getFan(char *name) {
    struct fan_registry *r = registry;
    for (; r->name != NULL; r++)
        if (strcmp(name, r->name) == 0)
            return r;
    return NULL;
}

static void configTemp(CONFIG_SECTION *sect, int id, struct temp *t) {
    double tempUnit = 1.0;
    double on_temp = 0.0, off_temp = 0.0;
    char temp[128];

    char *name = NULL;
    snprintf(temp, sizeof (temp), "temperature%d", id);
    config_getCharParameter(sect, temp, &name);
    
    // Fatal if first one is missing, optional for rest
    if (id == 1)
        fatalIfNull(name, "temperature1 sensor name is mandatory for %s", sect->node.name);
    else if (!name) {
        t->enabled = 0;
        return;
    }

    struct sensor *sensor = sensors_get(name);
    fatalIfNull(temp, "temperature sensor \"%s\" is not present for %s", name, sect->node.name);

    snprintf(temp, sizeof (temp), "temperature%d.unit", id);
    config_getDoubleParameter(sect, temp, &tempUnit);

    snprintf(temp, sizeof (temp), "temperature%d.max", id);
    config_getDoubleParameter(sect, temp, &on_temp);

    snprintf(temp, sizeof (temp), "temperature%d.min", id);
    config_getDoubleParameter(sect, temp, &off_temp);

    if (on_temp < off_temp)
        fatalError("Max fan temperature must be higher than Min for %s", sect->node.name);

    t->enabled = 1;
    t->sensor = sensor;
    t->temp_unit = tempUnit;
    t->off_temp = off_temp;
    t->on_temp = on_temp;
}

void register_fan_sensor(CONFIG_SECTION *sect) {
    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    char *name = NULL;
    config_getCharParameter(sect, "fan-type", &name);
    fatalIfNull(name, "fan-type is mandatory for %s", sect->node.name);
    struct fan_registry *r = getFan(name);
    fatalIfNull(r, "fan-type %s is not supported for %s", name, sect->node.name);

    configTemp(sect, 1, &state->temp1);
    configTemp(sect, 2, &state->temp2);
    configTemp(sect, 3, &state->temp3);

    state->sensor.name = sect->node.name;
    state->sensor.title = "Fan control";
    state->sensor.update = update;

    // Add the fan control
    state->fan_control = r->handler;
    if (r->init)
        state->fan_data = r->init(sect);

    // Reset the fan to off initially
    if (r->handler)
        r->handler(state->fan_data, 0);

    sensor_configure(sect, &state->sensor);

    pthread_mutex_init(&state->mutex, NULL);

    // Listen to those sensors
    if (state->temp1.enabled)
        sensor_add_listener(state->temp1.sensor, &state->sensor, update_temp1);

    if (state->temp2.enabled)
        sensor_add_listener(state->temp2.sensor, &state->sensor, update_temp2);

    if (state->temp3.enabled)
        sensor_add_listener(state->temp3.sensor, &state->sensor, update_temp3);

    // Finally register it
    sensor_register(&state->sensor);
}
