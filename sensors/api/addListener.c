
#include <stdlib.h>
#include <string.h>
#include "sensors/sensors.h"

/**
 * Adds a listener to a sensor
 * 
 * @param sensor Sensor to listen to
 * @param listener Sensor that's listening to sensor
 * @param update function to call when the sensor is updated
 */
void sensor_add_listener(
        struct sensor *sensor,
        struct sensor *listener,
        void (*update)(struct sensor *sensor, struct sensor *listener)
        ) {
    struct sensor_listener *l = (struct sensor_listener *) malloc(sizeof (struct sensor_listener));
    memset(l, 0, sizeof (struct sensor_listener));
    l->listener = listener;
    l->update = update;
    l->next = sensor->listeners;
    sensor->listeners = l;
}