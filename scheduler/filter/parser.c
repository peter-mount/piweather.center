#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>
#include "scheduler/scheduler.h"
#include "lib/config.h"
#include "astro/location.h"

/*
 * A set of filters which enable schedules to be manipulated so that
 * they only apply under dynamic circumstances, for example only
 * trigger during daylight hours is a common one used by the PI based
 * skycamera
 */

static int normal_daylight(SCHEDULE_ENTRY *e, int m) {
    return scheduler_filter_above_horizon(&schedule.today.day, m);
}

static int civil_daylight(SCHEDULE_ENTRY *e, int m) {
    return scheduler_filter_above_horizon(&schedule.today.civil, m);
}

static int nautical_daylight(SCHEDULE_ENTRY *e, int m) {
    return scheduler_filter_above_horizon(&schedule.today.nautical, m);
}

static int astronomical_daylight(SCHEDULE_ENTRY *e, int m) {
    return scheduler_filter_above_horizon(&schedule.today.astronomical, m);
}

static int normal_night(SCHEDULE_ENTRY *e, int m) {
    return scheduler_filter_below_horizon(&schedule.today.day, m);
}

static int civil_night(SCHEDULE_ENTRY *e, int m) {
    return scheduler_filter_below_horizon(&schedule.today.civil, m);
}

static int nautical_night(SCHEDULE_ENTRY *e, int m) {
    return scheduler_filter_below_horizon(&schedule.today.nautical, m);
}

static int astronomical_night(SCHEDULE_ENTRY *e, int m) {
    return scheduler_filter_below_horizon(&schedule.today.astronomical, m);
}

static struct schedule_filter FILTERS[] = {
    // Daylight - when the sun is above some limit by the horizon
    { "daylight", normal_daylight},
    { "civil daylight", civil_daylight},
    { "nautical daylight", nautical_daylight},
    { "astronomical daylight", astronomical_daylight},
    // Night - when the sun is below some limit by the horizon
    { "night", normal_night},
    { "civil night", civil_night},
    { "nautical night", nautical_night},
    { "astronomical night", astronomical_night},
    // Terminate the list
    { NULL, NULL}
};

struct schedule_filter *scheduler_getFilter(char *name) {
    struct schedule_filter *f = FILTERS;
    while (f->name) {
        if (strcmp(name, f->name) == 0) {
            return f;
        }
        f++;
    }
    return NULL;
}

/**
 * Parse name and set the SCHEDULE_ENTRY to use the named filter
 * @param e SCHEDULE_ENTRY
 * @param name Name of filter
 * @return 1 on error, 0 if name matches a filter
 */
int scheduler_parse_filter(SCHEDULE_ENTRY *e, char *name) {
    struct schedule_filter *f = scheduler_getFilter(name);
    if (f) {
        e->filter = f->filter;
        return 0;
    }
    return 1;
}
