#include <stdlib.h>
#include <stdio.h>
#include "astro/sun.h"
#include "scheduler/scheduler.h"

/**
 * Utility to check if m is between the times an object is above the horizon
 * @param rs RISE_SET
 * @param m minute of day
 * @return 1 if m is while rs is above the horizon
 */
int scheduler_filter_above_horizon(RISE_SET *rs, int m) {
    if (rs->type == RISE_SET_NEVER_RISES)
        return 0;

    if (rs->type == RISE_SET_NEVER_SETS)
        return 1;

    int r = (int) round(rs->rise * 60.0);
    int s = (int) round(rs->set * 60.0);
    return r <= m && m <= s;
}
