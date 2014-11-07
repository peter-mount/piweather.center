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
#include "scheduler/scheduler.h"

static void *handler(void *arg) {
    struct sensor *s = (struct sensor *) arg;
    time(&s->last_update);
    if (s->enabled && s->update) {
        s->update(s);

        // Once updated, notify any listeners - safe as it's in the same thread
        if (s->listeners) {
            struct sensor_listener *l = s->listeners;
            while (l) {
                l->update(s, l->listener);
                l = l->next;
            }
        }
    }
}

void sensor_configure(CONFIG_SECTION *sect, struct sensor *sensor) {
    if (sect) {
        // The schedule
        char *spec = NULL, *filter = NULL;
        config_getCharParameter(sect, "schedule", &spec);
        config_getCharParameter(sect, "schedule-filter", &filter);
        if (spec) {
            SCHEDULE_ENTRY *e = scheduler_new(handler, (void *) sensor);
            sensor->enabled = !scheduler_parse(e, spec);

            if (!sensor->enabled)
                fprintf(stderr, "Unsupported filter \"%s\" in %s\n", filter, sect->node.name);

            if (filter)
                if (scheduler_parse_filter(e, filter))
                    fprintf(stderr, "Unsupported filter \"%s\" in %s\n", filter, sect->node.name);

            scheduler_add(e);
        }

        config_getBooleanParameter(sect, "annotate", &sensor->annotate);
    }
}
