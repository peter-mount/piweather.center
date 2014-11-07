#include <stdlib.h>
#include <string.h>
#include "scheduler/scheduler.h"

SCHEDULE_ENTRY *scheduler_new(void *(*handler)(void *userdata), void *userdata) {
    SCHEDULE_ENTRY *e = (SCHEDULE_ENTRY *) malloc(sizeof (SCHEDULE_ENTRY));
    memset(e, 0, sizeof (SCHEDULE_ENTRY));
    e->handler = handler;
    e->userdata = userdata;
    return e;
}

