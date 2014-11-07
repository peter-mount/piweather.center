/**
 * The sensor framework
 */

#include <stdlib.h>
#include "sensors/sensors.h"

time_t sensor_next_trigger(struct sensor *sensor) {
    return sensor->last_update + sensor->frequency;
}
