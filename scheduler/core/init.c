#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <stdio.h>
#include "scheduler/scheduler.h"
#include "astro/time.h"
#include "astro/location.h"

SCHEDULE schedule;

struct job {
    struct Node node;
    SCHEDULE_ENTRY *e;
};

/**
 * Schedule job which refreshes the schedule for a new day
 */
static void *scheduler_recalc(void *userdata) {
    time_t now;
    time(&now);
    double jd = astro_julian_0h(astro_julday_time(&now));
    astro_sunriseset(jd, &observatory, &schedule.today);
}

/**
 * Worker to invoke a job on a worker thread
 */
static void *worker(struct Node *n, void *d) {
    struct job *j = (struct job *) n;
    SCHEDULE_ENTRY *e = j->e;
    free(j);
    if (e && e->handler)
        e->handler(e->userdata);
}

/**
 * The scheduler thread
 */
static void *scheduler(void *arg) {
    struct tm tm;
    time_t now;
    int m;

    while (1) {
        // Now get the current time, time remainin in this minute is how
        // long we need to sleep
        time(&now);
        gmtime_r(&now, &tm);
        m = 60 - tm.tm_sec;
        if (m > 0)
            sleep(m);

        // Now refresh time & calculate the schedule bit number
        time(&now);
        gmtime_r(&now, &tm);
        m = (tm.tm_hour * 60) + tm.tm_min;

        // Run through the entries looking for anything to trigger
        pthread_mutex_lock(&schedule.mutex);
        struct Node *n = schedule.entries.l_head;
        while (list_isNode(n)) {
            SCHEDULE_ENTRY *e = (SCHEDULE_ENTRY *) n;
            n = n->n_succ;

            if (scheduler_trigger(e, m)) {
                struct job *j = (struct job *) malloc(sizeof (struct job));
                memset(j, 0, sizeof (struct job));
                j->e = e;
                threadpool_submit(&schedule.threads, &j->node);
            }
        }
        pthread_mutex_unlock(&schedule.mutex);
    }
}

/**
 * Initialises the schedule ready for use
 */
void scheduler_init() {
    memset(&schedule, 0, sizeof (schedule));
    list_init(&schedule.entries);

    pthread_mutex_init(&schedule.mutex, NULL);

    CONFIG_SECTION *sect = config_getSection("scheduler");

    // Initialise the thread pool
    int poolsize = 4;
    if (sect)
        config_getIntParameter(sect, "pool-size", &poolsize);
    if (poolsize < 1)
        poolsize = 1;
    threadpool_init(&schedule.threads, poolsize);
    threadpool_start(&schedule.threads, worker, NULL);

    // Initialise the schedule ephemeris for the first time
    scheduler_recalc(NULL);

    // create a schedule entry to fire at midnight to recalc the scheduler's
    // ephemeris
    SCHEDULE_ENTRY *e = scheduler_new(scheduler_recalc, NULL);
    scheduler_parse(e, "0 0");
    scheduler_add(e);

    // Start the scheduler
    pthread_create(&schedule.thread, NULL, scheduler, NULL);
}

