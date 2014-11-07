#include <stdlib.h>
#include <stdint.h>
#include "scheduler/scheduler.h"

// The index in the schedule for the given minute

#define INDEX(d) ((d.quot<<1) + (d.rem<30?0:1))
#define BIT(d) ((d.rem<30?d.rem:(d.rem-30))+1)

int scheduler_getIndex(SCHEDULE_ENTRY *e, int minute) {
    div_t d = div(minute, 60);
    return d.quot << 1;
}

int scheduler_getBit(SCHEDULE_ENTRY *e, int minute) {
    div_t d = div(minute, 60);
    return (e->schedule[INDEX(d)] & 1 << BIT(d)) != 0;
}

void scheduler_setBit(SCHEDULE_ENTRY *e, int minute) {
    div_t d = div(minute, 60);
    e->schedule[INDEX(d)] |= 1 << BIT(d);
}
