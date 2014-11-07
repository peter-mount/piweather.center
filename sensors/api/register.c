/**
 * The sensor framework
 */

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include <time.h>
#include "camera/camera.h"
#include "lib/blockingQueue.h"
#include "lib/config.h"
#include "weatherstation/main.h"
#include "webserver/webserver.h"
#include "sensors/sensors.h"

/**
 * Standard sensor registration
 * 
 * @param camera CAMERA_STATE
 * @param sensor sensor
 */
void sensor_register(struct sensor *sensor) {
    sensor->next = NULL;
    if (sensors->sensors) {
        struct sensor *s = sensors->sensors;
        while (s->next)
            s = s->next;
        s->next = sensor;
    } else {
        sensors->sensors = sensor;
    }
}
