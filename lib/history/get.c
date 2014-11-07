
#include <stdlib.h>
#include <time.h>
#include "lib/list.h"
#include "lib/history.h"

/**
 * Utility to get the oldest and newest entry in a History.
 * 
 * If the History is empty then NULL is placed in the two variables
 * and 0 returned.
 * 
 * Otherwise the two variables are set to point to the oldest and newest
 * entries and 1 returned.
 * 
 * The History is not updated in either case.
 * 
 * Note: If the history contains just one entry then both variables will
 * point to the same entry as it's both the oldest and newest.
 * 
 * An example of this is pressure_trend which looks at the pressure
 * difference over a three hour period, so this method is used and we then
 * look at the pressures to determine the trend.
 * 
 * @param h History to check
 * @param old Pointer to variable to hold oldest entry
 * @param new Pointer to variable to hold newest entry
 * @return 1 if data exists, 0 of not.
 */
int history_get_old_new(struct History *h, struct HistoryNode **old, struct HistoryNode **new) {
    if (list_isEmpty(&h->list)) {
        *old = NULL;
        *new = NULL;
        return 0;
    } else {
        *old = (struct HistoryNode *) list_getHead(&h->list);
        *new = (struct HistoryNode *) list_getTail(&h->list);
        return 1;
    }
}
