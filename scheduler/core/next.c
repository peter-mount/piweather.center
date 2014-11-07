#include <stdlib.h>
#include <stdint.h>
#include "scheduler/scheduler.h"

/**
 * Returns the number of minutes until the next invocation of
 * a SCHEDULE_ENTRY
 * 
 * TODO add start/end handling
 * 
 * @param e SCHEDULE_ENTRY
 * @param minute The current minute in the day
 * @return Number of minutes until the next run
 */
int scheduler_nextMinute(SCHEDULE_ENTRY *e, int minute) {
    int i = minute, j = 0;
    
    // Find next invocation
    for (; i < SCHEDULE_MINS; j++)
        if (scheduler_getBit(e, i))
            return j;
    
    // Start from beginning of next day
    for (i=0; i < minute; j++)
        if (scheduler_getBit(e, i))
            return j;
    
    // Should never happen, no schedule
    return INT_FAST32_MAX;
}
