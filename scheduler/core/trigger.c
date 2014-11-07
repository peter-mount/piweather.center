#include <stdlib.h>
#include <stdint.h>
#include "scheduler/scheduler.h"

/**
 * Tests to see if an entry is due to be triggered for a specific minute
 * @param s SCHEDULE
 * @param e SCHEDULE_ENTRY
 * @param m minute
 * @return 1 if it should trigger, 0 if not
 */
int scheduler_trigger(SCHEDULE_ENTRY *e, int m) {
    int t = scheduler_getBit(e, m);

    // Pass to optional filter
    if (t && e->filter)
        t = e->filter(e, m);

    return t;
}
