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

struct sensors *sensors;

void sensor_registerAll(struct sensor_registry *registries) {

    sensors = (struct sensors *) malloc(sizeof (struct sensors));
    memset(sensors, 0, sizeof (struct sensors));

    sensors->sensors = NULL;

    struct Node *sn = list_getHead(&config->sections);
    while (list_isNode(sn)) {
        CONFIG_SECTION *sect = (CONFIG_SECTION *) sn;
        sn = sn->n_succ;

        // We are only interested in those with a sensor type
        CONFIG_PARAM *p = config_getParameter(sect, "sensor-type");
        if (p) {
            int i = 0;
            void (*registry)(CONFIG_SECTION * sect) = NULL;
            while (!registry && registries[i].type) {
                if (strcmp(registries[i].type, p->value) == 0) {
                    registry = registries[i].registry;
                }
                i++;
            }
            if (registry) {
                registry(sect);
            } else {
                fprintf(stderr, "Unsupported sensor-type %s in section %s\n", p->value, sect->node.name);
                exit(1);
            }
        }
    }
}
