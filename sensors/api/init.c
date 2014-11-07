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

void sensor_init() {
    struct sensor *s = sensors->sensors;
    while (s) {
        if (s->init)
            s->init(s);
        s = s->next;
    }
}

void sensor_postinit() {
    time_t t;
    time(&t);

    struct sensor *s = sensors->sensors;
    while (s) {
        if (s->postinit)
            s->postinit(s);
        s = s->next;
    }
}
