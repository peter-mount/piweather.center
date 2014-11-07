/* 
 * File:   scheduler.h
 * Author: Peter T Mount
 *
 * Created on April 29, 2014, 9:39 PM
 */

#ifndef SCHEDULER_H
#define	SCHEDULER_H

#include <time.h>
#include <stdint.h>
#include <pthread.h>
#include "astro/sun.h"
#include "astro/time.h"
#include "lib/list.h"
#include "lib/threadpool.h"

// Number of minutes in a day
#define SCHEDULE_MINS 1440

// Number of int's to hold the schedule, SCHEDULE_MINS/32 = 45
#define SCHEDULE_SIZE 45

typedef struct {
    struct List entries;
    // Sun rise/set for today
    SOLAR_EPHEMERIS today;
    // Mutex used to lock the schedule
    pthread_mutex_t mutex;
    // The scheduler thread
    pthread_t thread;
    // The worker thread pool
    struct thread_pool threads;
} SCHEDULE;

typedef struct schedule_entry SCHEDULE_ENTRY;

struct schedule_entry {
    struct Node node;
    // Bitmask of when schedule is to triggerminute schedule is to trigger
    // One bit for each minute of the day in 2 uint32_t for each hour
    uint32_t schedule[48];
    // Userdata for the scheduled job
    void *userdata;
    // Function to call when schedule is triggered
    void *(*handler)(void *userdata);
    // Filter used to determine if the entry should trigger
    int (*filter)(SCHEDULE_ENTRY *e, int m);
    // Userdata for the filter - unused by in future will allow us to
    // create filters based on sensor readings
    void *filter_userdata;
};

struct schedule_filter {
    const char *name;
    int (*filter)(SCHEDULE_ENTRY *e, int m);
};

extern SCHEDULE schedule;

void scheduler_add(SCHEDULE_ENTRY *e);
extern void scheduler_init();
extern SCHEDULE_ENTRY *scheduler_new(void *(*handler)(void *userdata), void *userdata);
extern int scheduler_nextMinute(SCHEDULE_ENTRY *e, int minute);
extern int scheduler_trigger(SCHEDULE_ENTRY *e, int m);

extern int scheduler_filter_above_horizon(RISE_SET *rs, int m);
#define scheduler_filter_below_horizon(e,m) (!scheduler_filter_above_horizon((e),(m)))
extern int scheduler_parse_filter(SCHEDULE_ENTRY *e, char *name);
extern struct schedule_filter *scheduler_getFilter(char *name);

extern int scheduler_getIndex(SCHEDULE_ENTRY *e, int minute);
extern int scheduler_getBit(SCHEDULE_ENTRY *e, int minute);
extern void scheduler_setBit(SCHEDULE_ENTRY *e, int minute);

#endif	/* SCHEDULER_H */

