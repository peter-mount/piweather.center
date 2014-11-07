/*
 * An example temperature sensor.
 * 
 * This sensor will log the Raspberry PI's CPU temperature
 */

#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include "lib/config.h"
#include "sensors/sensors.h"
#include "weatherstation/main.h"

struct state {
    // This must be the first entry, it's what the state engine will see this struct as
    struct sensor sensor;
    // Internal buffer used by update
    char *buffer;
    // The size of the buffer
    ssize_t buffer_length;
};

/**
 * Updates the sensor by reading the current value into it
 */
static void update(struct sensor *sensor) {
    struct state *state = (struct state *) sensor;

    FILE *f = fopen("/sys/class/thermal/thermal_zone0/temp", "r");
    if (f) {
        getline(&state->buffer, &state->buffer_length, f);
        fclose(f);

        int value;
        // The file contains a single integer which is our value
        if (sscanf(state->buffer, "%d", &value)) {
            // Update the sensor's state
            sensor_log(sensor, value, "CPU %0.1fC", (double) value / 1000.0);
        }
    }
}

/**
 * Used to create the sensor struct defining this sensor
 * @return 
 */
void register_cpu_sensor(CONFIG_SECTION *sect) {
    // Create our state and ensure it's all clear
    struct state *state = (struct state *) malloc(sizeof (struct state));
    memset(state, 0, sizeof (struct state));

    // The sensor name, used in command prefixes & rest url's
    state->sensor.name = sect->node.name;

    // The title is used in debug & help
    state->sensor.title = "CPU Temperature Sensor";

    // update is the only mandatory hook
    state->sensor.update = update;

    sensor_configure(sect, &state->sensor);

    sensor_register((struct sensor *) state);
}