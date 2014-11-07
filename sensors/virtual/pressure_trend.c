/*
 * This virtual sensor monitors a pressure sensor and records the last
 * 3 hours (configurable) readings.
 * 
 * The logged output is the change in pressure in that time. So here we record the difference of the most recent
 * reading and the oldest.
 * 
 * So, if the oldest reading is 1011hPa and the most recent reading is 1009hPa then the recorded trend is -2hPa.
 * 
 * What do these values mean? Well the international standard says: Trend is the direction of change of the
 * barometric pressure over the last three hours.
 * 
 * Rising Rapidly - if the pressure increases by > 2mb
 * Rising Slowly - if the pressure increases by >1mb but <2mb
 * Steady - pressure changes by <1mb
 * Falling Slowly - pressure falls by >1mb but <2mb
 * Falling Rapidly - pressure falls by >2mb.
 * 
 * Note: 1mb (mbar, millibars) is equivalent to hPa (hecto-pascal), so 1001mb == 1001hPa. mb is not an SI unit whilst
 * hPa is derived from the Pa (Pascal), 1hPa = 100Pa
 */


#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <pthread.h>
#include <time.h>
#include "lib/config.h"
#include "lib/history.h"
#include "lib/string.h"
#include "sensors/sensors.h"

#include <stdio.h>

struct state {
    // This must be the first entry, it's what the state engine will see this struct as
    struct sensor sensor;
    // The pressure history
    struct History history;
    // The change in pressure since readings had begun
    int trend;
    // Unit to convert pressure raw value into a formatable value
    double unit;
    // Format to use when logging the sensor
    char *fmt;
    // mutex as updates can happen on different threads
    pthread_mutex_t mutex;
};

struct reading {
    struct HistoryNode node;
    int value;
};

static void add_reading(struct sensor *sensor, struct sensor *listener) {

    // Don't add a 0 pressure, can happen during startup
    if (sensor->value == 0)
        return;

    struct state *state = (struct state *) listener;

    // The reading, using the time & value of the sensor
    struct reading *r = (struct reading *) malloc(sizeof (struct reading));
    memset(r, 0, sizeof (struct reading));
    r->node.time = sensor->last_update;
    r->value = sensor->value;

    // Mutex as this data is shared across threads
    pthread_mutex_lock(&state->mutex);
    history_add(&state->history, &r->node);
    pthread_mutex_unlock(&state->mutex);
}

static void update(struct sensor *sensor) {
    struct state *state = (struct state *) sensor;

    // The change in pressure over time
    int trend = 0;

    // Mutex whilst we play with shared data
    pthread_mutex_lock(&state->mutex);

    // Remove any readings that are too old
    history_expire(&state->history);

    // Now calculate the trend - account for no data in the set
    // NB: if head is a node then tail will always be one so only test the head
    struct reading *oldest;
    struct reading *newest;
    if (history_get_old_new(&state->history, (struct HistoryNode *) &oldest, (struct HistoryNode *) &newest))
        trend = newest->value - oldest->value;

    // We are done with shared data from this point
    pthread_mutex_unlock(&state->mutex);

    // Log the numerical trend
    sensor_log(sensor, trend, state->fmt, (double) trend * state->unit);
}

void register_virtual_pressure_trend(CONFIG_SECTION *sect) {

    char *name = NULL;
    config_getCharParameter(sect, "barometer", &name);
    fatalIfNull(name, "barometer sensor name is mandatory for %s", sect->node.name);
    struct sensor *barometer = sensors_get(name);
    fatalIfNull(barometer, "barometer sensor \"%s\" is not present for %s", name, sect->node.name);

    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    state->sensor.name = sect->node.name;
    state->sensor.title = "Barometer Trend";
    state->sensor.update = update;

    // default is 3 hours
    int max_age = 3 * 60;
    config_getIntParameter(sect, "max-age", &max_age);
    // convert into seconds
    time_t t_max_age = max_age * 60;

    // Initialise the history
    history_init(&state->history, t_max_age);

    config_getCharParameter(sect, "format", &state->fmt);
    if (!state->fmt)
        state->fmt = "%0.3f";

    state->unit = 1.0;
    config_getDoubleParameter(sect, "unit", &state->unit);

    sensor_configure(sect, &state->sensor);

    pthread_mutex_init(&state->mutex, NULL);

    // Listen to the barometer
    sensor_add_listener(barometer, &state->sensor, add_reading);

    // Finally register it
    sensor_register(&state->sensor);
}

