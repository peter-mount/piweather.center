
#include <stdlib.h>
#include <time.h>
#include "lib/list.h"
#include "lib/history.h"

/**
 * Adds a new HistoryNode to the tail of this History instance
 * 
 * Note: If the HistoryNode's time is 0 then it will be set to the current
 * time.
 * 
 * @param h History
 * @param n HistoryNode to add
 */
void history_add(struct History *h, struct HistoryNode *n) {
    // If not set then set the time to now
    if (!n->time)
        time(&n->time);
    
    list_addTail(&h->list, &n->node);
}
