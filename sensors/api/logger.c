/**
 * Used by a sensor to log a result
 */

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include "logger/logger.h"
#include "sensors/sensors.h"

void sensor_log(struct sensor *sensor, int value, char *fmt, ...) {
    va_list argp;

    // Set if we change. Note, this must not clear this flag here
    int updated = sensor->value != value ? 1 : 0;
    sensor->value = value;

    va_start(argp, fmt);
    vsnprintf(sensor->text, SENSOR_TEXT_SIZE, fmt, argp);
    va_end(argp);

    // Ensure we are terminated
    sensor->text[SENSOR_TEXT_SIZE - 1] = '\0';

    logger_log(sensor->last_update, (char *) sensor->name, sensor->text, value, updated);
}
