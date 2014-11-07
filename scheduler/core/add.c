#include <stdlib.h>
#include "lib/threadpool.h"
#include "scheduler/scheduler.h"

void scheduler_add(SCHEDULE_ENTRY *e) {
    pthread_mutex_lock( &schedule.mutex);
    list_enqueue(&schedule.entries, &e->node);
    pthread_mutex_unlock( &schedule.mutex);
}
