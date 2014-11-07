/**
 * The sensor framework
 */

#include <time.h>
#include "sensors/sensors.h"

time_t sensor_getAge(struct sensor *sensor, time_t now) {
    return now - sensor->last_update;
}
