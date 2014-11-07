
#include <stdlib.h>
#include <string.h>
#include "sensors/sensors.h"

struct sensor *sensors_get(const char *name) {
    struct sensor *s = sensors->sensors;
    while (s) {
        if (strcmp(s->name, name) == 0)
            return s;
        s = s->next;
    }
    return NULL;
}
